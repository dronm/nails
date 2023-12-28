func convert_doc(app osbe.Applicationer, doc *convertDocument, conn *pgx.Conn, specialist_id int) error {
	var template_f *os.File
	var err error
	template_fn := docAttachment.GetAttachmentCacheFileName(app.GetBaseDir(), doc.content_id) //имя файла в кэша
	if view.FileExists(template_fn) {
		//from cache		
		template_f, err = os.Open(template_fn)
		if err != nil {
			return errors.New(fmt.Sprintf("os.Open() from cache failed: %v", err))
		}
		defer template_f.Close()		
		
	}else{
		//no cache, read from db && save
		var f_cont []byte//&fields.ValBytea{}
		if err := conn.QueryRow(context.Background(),
			`SELECT
				content_data
			FROM attachments
			WHERE id = $1`,
			doc.att_id,
		).Scan(&f_cont); err != nil {
			return errors.New(fmt.Sprintf("conn.QueryRow() attachments failed: %v", err))
		}
		
		template_f, err = os.Create(template_fn)
		if err != nil {
			return errors.New(fmt.Sprintf("os.Create() on file from db failed: %v", err))
		}
		defer template_f.Close()
		if _, err := template_f.Write(f_cont); err != nil {
			return errors.New(fmt.Sprintf("Write() on file from db failed: %v", err))
		}
	}
		
	params_doc, err := os.CreateTemp("", "nails_reg_params") //Файл со значениями для замены в шаблоне
	if err != nil {
		return errors.New(fmt.Sprintf("os.CreateTemp() on param file failed: %v", err))
	}
	defer os.Remove(params_doc.Name())
	//fmt.Println("params_doc=", params_doc.Name())	
	doc.values["flSign"] = "   " //hide signature
	doc_values, err := json.Marshal(&doc.values)
	if err != nil {
		return errors.New(fmt.Sprintf("json.Marshal() on param file failed: %v", err))
	}
	if err := ioutil.WriteFile(params_doc.Name(), doc_values, 0777); err != nil {
		return errors.New(fmt.Sprintf("ioutil.WriteFile() on param file failed: %v", err))
	}
		
	out_doc, err := os.CreateTemp("", "nails_reg_out") //Готовый документ docx
	if err != nil {
		return errors.New(fmt.Sprintf("os.CreateTemp() on out_doc file failed: %v", err))
	}
	defer os.Remove(out_doc.Name())
	//fmt.Println("out_doc=", out_doc.Name())		
	
	//----------------------------- выполняем замены
	conf := app.GetConfig().(*nails_config.NailsConfig)
	//fmt.Println(conf.PHP, conf.PHPDocConverter, template_fn, params_doc.Name(), out_doc.Name())	
	cmd := exec.Command(conf.PHP, conf.PHPDocConverter, template_fn, params_doc.Name(), out_doc.Name())
	out_b, err := cmd.StdoutPipe()		
	if err != nil {
		return errors.New(fmt.Sprintf("cmd.StdoutPipe() on convert command failed: %v", err))
	}
	cmd.Stderr = cmd.Stdout	
	err = cmd.Run()
	if err != nil {
		return errors.New(fmt.Sprintf("cmd.Run() on convert command failed: %v", err))
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
		return errors.New(fmt.Sprintf("php error generating docx output: %s", out_s.String()))
	}
	
	//Документ docx с заменами в базу данных без preview - для дальнейших замен (подпись)	
	ref := &docAttachment.Ref_Type{Keys: docAttachment.AttachmentKey{Id: docAttachment.HttpInt(specialist_id)}, DataType: "specialists"}	
	out_doc_buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(out_doc_buf, out_doc); err != nil {
		return errors.New(fmt.Sprintf("io.Copy() on out_doc. file failed: %v", err))
	}		
	out_doc_fi := &docAttachment.Content_info_Type{Id: doc.template_name,
			Name: doc.template_name + ".docx",
			Size: int64(out_doc_buf.Len())}	
	
	
	//------------------- Конветрация документа с заменами в PDF
	//Файл pdf с заменами поместим в папку CACHE
	out_pdf_fn := filepath.Join(app.GetBaseDir(), docAttachment.CACHE_DIR, filepath.Base(out_doc.Name()) + ".pdf")
	//Это для отладки
	//fmt.Println("soffice --headless --convert-to pdf:writer_pdf_Export --outdir", filepath.Join(app.GetBaseDir(), docAttachment.CACHE_DIR), out_doc.Name())												
	cmd_args := strings.Split("--headless --convert-to pdf:writer_pdf_Export --outdir", " ")
	cmd_args = append(cmd_args, filepath.Join(app.GetBaseDir(), docAttachment.CACHE_DIR))
	cmd_args = append(cmd_args, out_doc.Name())
	cmd = exec.Command("soffice", cmd_args...)
	err = cmd.Run()
	if err != nil {
		return errors.New(fmt.Sprintf("cmd.Run() on convert to PDF command failed: %v", err))
	}
	if _, err := os.Stat(out_pdf_fn); os.IsNotExist(err) {
		return errors.New("converted PDF file not found")
	}
	//Переименуем pdf в "правильное" имя для кэша	
	out_pdf_fi := &docAttachment.Content_info_Type{Id: doc.template_name + ".pdf", Name: doc.template_name + ".pdf"}	
	out_pdf_cor_for_cache := docAttachment.GetAttachmentCacheFileName(app.GetBaseDir(), out_pdf_fi.Id)
	if err := os.Rename(out_pdf_fn, out_pdf_cor_for_cache); err != nil {
		return errors.New(fmt.Sprintf("os.Rename() on out_pdf failed: %v", err))
	}
	out_pdf_preview_fn := docAttachment.GetPreviewCacheFileName(app.GetBaseDir(), out_pdf_fi.Id)
	if err := docAttachment.GenThumbnail(out_pdf_cor_for_cache, out_pdf_preview_fn, out_pdf_fi.Name); err != nil {
		return errors.New(fmt.Sprintf("docAttachment.GenThumbnail() on out_pdf failed: %v", err))
	}
	out_pdf_bt, err := ioutil.ReadFile(out_pdf_cor_for_cache)
	if err != nil {
		return errors.New(fmt.Sprintf("ioutil.ReadFile() on out_pdf file failed: %v", err))
	}		
	out_pdf_prw_bt, err := ioutil.ReadFile(out_pdf_preview_fn)	
	if err != nil {
		return errors.New(fmt.Sprintf("ioutil.ReadFile() on out_pdf preview file failed: %v", err))
	}
	
	//----------------------------- Складываем все в базу данных
	var out_doc_att_id int
	var out_pdf_att_id int
	
	if _, err := conn.Exec(context.Background(), `BEGIN`); err != nil {
		return errors.New(fmt.Sprintf("conn.Exec() BEGIN failed: %v", err))
	}
	
	//delete all previously written documents + attachments	
	if _, err := conn.Exec(context.Background(),
		`DELETE FROM specialist_documents WHERE specialist_id = $1`,
		specialist_id,
	); err != nil {
		conn.Exec(context.Background(), `ROLLBACK`)
		return errors.New(fmt.Sprintf("conn.Exec() DELETE specialist_documents failed: %v", err))	
	}
	
	//out_doc
	if err := conn.QueryRow(context.Background(),
		`INSERT INTO attachments
		(ref, content_info, content_data)
		VALUES ($1, $2, $3)
		RETURNING id
		`,
			ref,
			out_doc_fi,	
			out_doc_buf.Bytes(),
	).Scan(&out_doc_att_id); err != nil {	
		conn.Exec(context.Background(), `ROLLBACK`)
		return errors.New(fmt.Sprintf("conn.QueryRow() on out_doc file failed: %v", err))	
	}

	//out_pdf + preview
	if err := conn.QueryRow(context.Background(),
		`INSERT INTO attachments
		(ref, content_info, content_data, content_preview)
		VALUES ($1, $2, $3, $4)
		RETURNING id
		`,
			ref,
			out_pdf_fi,	
			out_pdf_bt,
			out_pdf_prw_bt,
	).Scan(&out_pdf_att_id); err != nil {	
		conn.Exec(context.Background(), `ROLLBACK`)
		return errors.New(fmt.Sprintf("conn.QueryRow() on out_doc.pdf file failed: %v", err))	
	}
	
	if _, err := conn.Exec(context.Background(),
		`INSERT INTO specialist_documents
		(specialist_id, template_att_id, document_att_id, need_signing)
		VALUES ($1, $2, $3, $4)
		`,
		specialist_id,
		out_doc_att_id,
		out_pdf_att_id,
		doc.need_signing,
	); err != nil {
		conn.Exec(context.Background(), `ROLLBACK`)
		return errors.New(fmt.Sprintf("conn.Exec() INSERT specialist_documents failed: %v", err))	
	}

	if _, err := conn.Exec(context.Background(), `COMMIT`); err != nil {
		return errors.New(fmt.Sprintf("conn.Exec() COMMIT failed: %v", err))
	}
	return nil
}

