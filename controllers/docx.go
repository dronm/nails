package controllers

import(
	"archive/zip"
	"os"
	"io"
	"io/fs"	
	"fmt"
)

func ChangeImage(docxFileName, newImageFileName, zipImagePath string) (string, error) {
	r, err := zip.OpenReader(docxFileName) //reader
	if err != nil {
		return "", err
	}	
	defer r.Close()
	
	target_f, err := os.CreateTemp("", "nails_docx")
	if err != nil {
		return "", err
	}
	defer target_f.Close()	
	
	w := zip.NewWriter(target_f)
	defer w.Close()
	
	changed := false
	for _, f := range r.File {		
		var item_reader io.ReadCloser
		var err error
		var file_inf fs.FileInfo
		if len(f.Name) >= len(zipImagePath) && f.Name[0:len(zipImagePath)] == zipImagePath {
			//new file
			changed = true
			new_item_reader, err := os.Open(newImageFileName)
			if err != nil {
				return "", err
			}
			defer new_item_reader.Close()			
			
			file_inf, err = new_item_reader.Stat()
			if err != nil {
				return "", err
			}
			item_reader = new_item_reader			
			
		}else{
			//original file
			item_reader, err = f.Open()
			if err != nil {
				return "", err
			}
			defer item_reader.Close()
			file_inf = f.FileInfo()
		}
		
		header, err := zip.FileInfoHeader(file_inf)			
		if err != nil {
			return "", err
		}
		//header.Method = zip.Deflate		
		header.Name = f.Name						
		target_item, err := w.CreateHeader(header)
		if err != nil {
			return "", err
		}
		_, err = io.Copy(target_item, item_reader)
		if err != nil {
			return "", err
		}					
	}
	if !changed {
		return "", fmt.Errorf("image at path '%s' not found", zipImagePath)
	}
	
	return target_f.Name(), nil	
}
