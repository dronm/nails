package controllers

import(
	"os"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os/exec"
	"errors"
	"context"
	"strings"
	"bufio"
	
	"nails/nails_config"
	
	"osbe/view"
	
	"osbe/repo/docAttachment"
	
	"github.com/jackc/pgx/v4"	
)

const CONVERT_TO_PDF = "--headless --convert-to pdf:writer_pdf_Export --outdir"

type TemplateDocument struct {
	AttID int
	ContentID string
	Ref docAttachment.Ref_Type
	NeedSigning bool
	TemplateName string
	TemplateID int
	Values []byte//map[string]string	
}


//Applies parameters from TemplateDocument.Values to docx template
//Returns
//	path to template file
//	path to outfile.docx
//	path to outfile.pdf (if needPDF=true)
func ConvertDocFromTemplate(baseDir string, conf *nails_config.NailsConfig, doc *TemplateDocument, conn *pgx.Conn, needPDF bool) (string, string, string, error) {
	//---------------------------- Данные шаблона ----------------------------------------------	
	var err error
	template_fn := docAttachment.GetAttachmentCacheFileName(baseDir, doc.Ref.DataType, doc.Ref.Keys.Id, doc.ContentID) //имя файла в кэша
	if !view.FileExists(template_fn) {
		var template_f *os.File
		//no cache, read from db && save
		var f_cont []byte//&fields.ValBytea{}
		if err := conn.QueryRow(context.Background(),
			`SELECT
				content_data
			FROM attachments
			WHERE id = $1`,
			doc.AttID,
		).Scan(&f_cont); err != nil {
			return "", "", "", errors.New(fmt.Sprintf("conn.QueryRow() attachments failed: %v", err))
		}
		
		template_f, err = os.Create(template_fn)
		if err != nil {
			return "", "", "", errors.New(fmt.Sprintf("os.Create() on file from db failed: %v", err))
		}
		defer template_f.Close()
		if _, err := template_f.Write(f_cont); err != nil {
			return "", "", "", errors.New(fmt.Sprintf("Write() on file from db failed: %v", err))
		}
	}
	
	
	//------------------------------- Параметры ---------------------------------------------------
	params_doc, err := os.CreateTemp("", "nails_tmpl_params") //Файл со значениями для замены в шаблоне
	if err != nil {
		return "", "", "", errors.New(fmt.Sprintf("os.CreateTemp() on param file failed: %v", err))
	}
	defer os.Remove(params_doc.Name())
	
	//doc.Values["flSign"] = "   " //hide signature
	//doc_values, err := json.Marshal(&doc.Values)
	//if err != nil {
	//	return "", "", "", errors.New(fmt.Sprintf("json.Marshal() on param file failed: %v", err))
	//}
	if err := ioutil.WriteFile(params_doc.Name(), doc.Values, 0777); err != nil {
		return "", "", "", errors.New(fmt.Sprintf("ioutil.WriteFile() on param file failed: %v", err))
	}
	
	//Готовый документ docx
	out_doc, err := os.CreateTemp("", "nails_tmpl_out")
	if err != nil {
		return "", "", "", errors.New(fmt.Sprintf("os.CreateTemp() on out_doc file failed: %v", err))
	}


	//----------------------------- выполняем замены -------------------------------------------
//fmt.Println(conf.PHP, conf.PHPDocConverter, template_fn, params_doc.Name(), out_doc.Name())	
	cmd := exec.Command(conf.PHP, conf.PHPDocConverter, template_fn, params_doc.Name(), out_doc.Name())
	out_b, err := cmd.StdoutPipe()		
	if err != nil {
		return "", "", "", errors.New(fmt.Sprintf("cmd.StdoutPipe() on convert command failed: %v", err))
	}
	cmd.Stderr = cmd.Stdout	
	err = cmd.Run()
	if err != nil {
		return "", "", "", errors.New(fmt.Sprintf("cmd.Run() on convert command failed: %v", err))
	}	
	if !view.FileExists(out_doc.Name()) {
		var out_s strings.Builder
		scanner := bufio.NewScanner(out_b)
		for scanner.Scan() {
			if out_s.Len() > 0 {
				out_s.WriteString(" ")
			}
			out_s.WriteString(scanner.Text())
		}
		return "", "", "", errors.New(fmt.Sprintf("php error generating docx output: %s", out_s.String()))
	}
	
	if !needPDF {
		return template_fn, out_doc.Name(), "", nil
	}
	
	//------------------- Конветрация документа с заменами в PDF	
	//Файл pdf с заменами поместим в папку CACHE
	out_pdf_fn := filepath.Join(baseDir, docAttachment.CACHE_DIR, filepath.Base(out_doc.Name()) + ".pdf")
	//Это для отладки
	//fmt.Println("soffice --headless --convert-to pdf:writer_pdf_Export --outdir", filepath.Join(baseDir, docAttachment.CACHE_DIR), out_doc.Name())												
	cmd_args := strings.Split(CONVERT_TO_PDF, " ")
	cmd_args = append(cmd_args, filepath.Join(baseDir, docAttachment.CACHE_DIR))
	cmd_args = append(cmd_args, out_doc.Name())
	cmd = exec.Command("soffice", cmd_args...)
	err = cmd.Run()
	if err != nil {
		os.Remove(out_doc.Name())
		return "", "", "", errors.New(fmt.Sprintf("cmd.Run() on convert to PDF command failed: %v", err))
	}
	if _, err := os.Stat(out_pdf_fn); os.IsNotExist(err) {
		os.Remove(out_doc.Name())
		return "", "", "", errors.New("converted PDF file not found")
	}
	//Переименуем pdf в "правильное" имя для кэша	
	out_pdf_cor_for_cache := docAttachment.GetAttachmentCacheFileName(baseDir, doc.Ref.DataType, doc.Ref.Keys.Id, doc.TemplateName + ".pdf")
	if err := os.Rename(out_pdf_fn, out_pdf_cor_for_cache); err != nil {
		os.Remove(out_doc.Name())
		return "", "", "", errors.New(fmt.Sprintf("os.Rename() on out_pdf failed: %v", err))
	}
	
	return template_fn, out_doc.Name(), out_pdf_cor_for_cache, nil
}
