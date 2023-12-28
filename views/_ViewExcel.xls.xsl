<?xml version="1.0" encoding="UTF-8"?>

<xsl:stylesheet version="1.0"
xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
	 xmlns="urn:schemas-microsoft-com:office:spreadsheet"
	 xmlns:o="urn:schemas-microsoft-com:office:office"
	 xmlns:x="urn:schemas-microsoft-com:office:excel"
	 xmlns:ss="urn:schemas-microsoft-com:office:spreadsheet"
	 xmlns:html="http://www.w3.org/TR/REC-html40">

<xsl:output method="xml" indent="yes"/> 

<xsl:variable name="DT_BOOL" select="'0'"/>
<xsl:variable name="DT_CHAR" select="'1'"/>
<xsl:variable name="DT_STRING" select="'2'"/>
<xsl:variable name="DT_INT" select="'3'"/>
<xsl:variable name="DT_DATE" select="'4'"/>
<xsl:variable name="DT_TIME" select="'5'"/>
<xsl:variable name="DT_TIMETZ" select="'6'"/>
<xsl:variable name="DT_DATETIME" select="'7'"/>
<xsl:variable name="DT_DATETIMETZ" select="'8'"/>
<xsl:variable name="DT_FLOAT" select="'9'"/>
<xsl:variable name="DT_TEXT" select="'10'"/>
<xsl:variable name="DT_ENUM" select="'11'"/>
<xsl:variable name="DT_PASSWORD" select="'12'"/>
<xsl:variable name="DT_INTERVAL" select="'13'"/>
<xsl:variable name="DT_JSON" select="'14'"/>
<xsl:variable name="DT_JSONB" select="'15'"/>
<xsl:variable name="DT_ARRAY" select="'16'"/>
<xsl:variable name="DT_BYTEA" select="'17'"/>
<xsl:variable name="DT_XML" select="'18'"/>
<xsl:variable name="DT_GEOMPOLYGON" select="'19'"/>
<xsl:variable name="DT_GEOMPOINT" select="'20'"/>

<!-- default widths for data types in px-->
<xsl:variable name="DEF_WIDTH_DATE" select="100"/>
<xsl:variable name="DEF_WIDTH_TIME" select="100"/>
<xsl:variable name="DEF_WIDTH_DATETIME" select="125"/>
<xsl:variable name="DEF_WIDTH_INT" select="65"/>
<xsl:variable name="DEF_WIDTH_DOUBLE" select="50"/>
<xsl:variable name="DEF_WIDTH_FILE" select="230"/>
<xsl:variable name="DEF_COL_WIDTH" select="100"/>

<xsl:variable name="EXCEl_STYLE_ID_STRING" select="'s21'"/>
<xsl:variable name="EXCEl_STYLE_ID_INT" select="'s26'"/>
<xsl:variable name="EXCEl_STYLE_ID_MONEY" select="'s23'"/>
<xsl:variable name="EXCEl_STYLE_ID_FLOAT" select="'s27'"/>
<xsl:variable name="EXCEl_STYLE_ID_DATETIME" select="'s24'"/>
<xsl:variable name="EXCEl_STYLE_ID_DATE" select="'s25'"/>

<xsl:variable name="EXCEl_DT_INT" select="'Number'"/>
<xsl:variable name="EXCEl_DT_FLOAT" select="'Number'"/>
<xsl:variable name="EXCEl_DT_STRING" select="'String'"/>
<xsl:variable name="EXCEl_DT_DATETIME" select="'DateTime'"/>
<xsl:variable name="EXCEl_DT_DATE" select="'DateTime'"/>

<xsl:template name="string-replace-all">
  <xsl:param name="text" />
  <xsl:param name="replace" />
  <xsl:param name="by" />
  <xsl:choose>
    <xsl:when test="contains($text, $replace)">
      <xsl:value-of select="substring-before($text,$replace)" />
      <xsl:value-of select="$by" />
      <xsl:call-template name="string-replace-all">
        <xsl:with-param name="text"
        select="substring-after($text,$replace)" />
        <xsl:with-param name="replace" select="$replace" />
        <xsl:with-param name="by" select="$by" />
      </xsl:call-template>
    </xsl:when>
    <xsl:otherwise>
      <xsl:value-of select="$text" />
    </xsl:otherwise>
  </xsl:choose>
</xsl:template>

<xsl:template name="format_date">
	<xsl:param name="val"/>
	<xsl:param name="formatStr"/>
	<xsl:choose>
		<xsl:when test="string-length($val)=10">
			<xsl:variable name="val_year" select="substring-before($val,'-')"/>
			<xsl:variable name="part_month" select="substring-after($val,'-')"/>
			<xsl:variable name="val_month" select="substring-before($part_month,'-')"/>
			<xsl:variable name="part_date" select="substring-after($part_month,'-')"/>
			<xsl:variable name="val_date" select="$part_date"/>
			<xsl:value-of select="concat($val_date,'/',$val_month,'/',$val_year)" />
		</xsl:when>
		<xsl:otherwise>
			<xsl:value-of select="$val" />
		</xsl:otherwise>
	</xsl:choose>
</xsl:template>


<xsl:template name="format_json">
	<xsl:param name="val"/>	
	<xsl:variable name="mark_beg" select="'&quot;descr&quot; : &quot;'"/>
	<xsl:choose>
		<xsl:when test="contains($val, $mark_beg)">
			<xsl:variable name="mark_end" select="'&quot;, &quot;dataType'"/>
			<xsl:variable name="descr_to_end" select="substring-after($val,$mark_beg)"/>		
			<xsl:value-of select="substring-before($descr_to_end,$mark_end)" />
		</xsl:when>
		<xsl:otherwise>
			<xsl:value-of select="$val" />
		</xsl:otherwise>		
	</xsl:choose>
</xsl:template>

<!-- Main template-->
<xsl:template match="/">
	<xsl:processing-instruction
		name="mso-application">progid="Excel.Sheet"
	</xsl:processing-instruction>

	<Workbook 
	 xmlns="urn:schemas-microsoft-com:office:spreadsheet"
	 xmlns:o="urn:schemas-microsoft-com:office:office"
	 xmlns:x="urn:schemas-microsoft-com:office:excel"
	 xmlns:ss="urn:schemas-microsoft-com:office:spreadsheet"
	 xmlns:html="http://www.w3.org/TR/REC-html40">

		<DocumentProperties xmlns="urn:schemas-microsoft-com:office:office">
			<Author><xsl:value-of select="page/user_details/@descr"/></Author>
			<LastAuthor></LastAuthor>
			<Created><xsl:value-of select="page/@created"/></Created>
			<Company><xsl:value-of select="page/@firm"/></Company>
			<Version>10.2625</Version>
		</DocumentProperties>
		<OfficeDocumentSettings xmlns="urn:schemas-microsoft-com:office:office">
			<DownloadComponents/>
			<LocationOfComponents HRef=""/>
		</OfficeDocumentSettings>
	 
		<ExcelWorkbook xmlns="urn:schemas-microsoft-com:office:excel">
			<WindowHeight>10485</WindowHeight>
			<WindowWidth>20955</WindowWidth>
			<WindowTopX>240</WindowTopX>
			<WindowTopY>15</WindowTopY>
			<RefModeR1C1/>
			<ProtectStructure>False</ProtectStructure>
			<ProtectWindows>False</ProtectWindows>
		</ExcelWorkbook>
		<Styles>
			<Style ss:ID="Default" ss:Name="Normal">
				<Alignment ss:Vertical="Bottom"/>
				<Borders/>
				<Font ss:FontName="Arial Cyr" x:CharSet="204"/>
				<Interior/>
				<NumberFormat/>
				<Protection/>
			</Style>
			<Style ss:ID="s21">
				<Font ss:FontName="Arial Cyr" x:CharSet="204"/>
			</Style>
			<Style ss:ID="s22">
				<Font ss:FontName="Arial Cyr" x:CharSet="204" ss:Bold="1"/>
				<Interior ss:Color="#C0C0C0" ss:Pattern="Solid"/>
			</Style>
			<Style ss:ID="s23">
				<NumberFormat ss:Format="Currency"/>
			</Style>
			<Style ss:ID="s27">
				<NumberFormat/>
			</Style>			
			<Style ss:ID="s24">
				<NumberFormat ss:Format="dd/mm/yy\ h:mm;@"/>
			</Style>			
			<Style ss:ID="s25">
				<!-- dd/mm/yy;@ -0 -->
				<NumberFormat ss:Format="Short Date"/>
			</Style>			
			<Style ss:ID="s26">
				<NumberFormat ss:Format="0"/>
			</Style>			
			
		</Styles>
		
		<!-- sheets -->
		<xsl:apply-templates select="document/model[@id != 'Response']"/>
		
	</Workbook>
</xsl:template>

<!-- table -->
<xsl:template match="model">
	<xsl:variable name="model_id" select="@id"/>
	<!--<xsl:variable name="col_count" select="count(row[1]/*)"/>-->
	<xsl:variable name="col_count">
		<xsl:choose>
		<xsl:when test="count(/document/metadata[@modelId=$model_id]/field)>0">
			<xsl:value-of select="count(/document/metadata[@modelId=$model_id]/field[not(@sysCol='TRUE')])"/>
		</xsl:when>
		<xsl:otherwise>
			<xsl:value-of select="count(row[1]/*)"/>
		</xsl:otherwise>
		</xsl:choose>
	</xsl:variable>
	<xsl:variable name="row_count" select="count(row)"/>
	<Worksheet ss:Name="Rep{position()}">
		<Table ss:ExpandedColumnCount="{$col_count}" ss:ExpandedRowCount="{$row_count+1}" x:FullColumns="1"
		   x:FullRows="1">

			<!-- header
			Если есть метаданные и колич-во колонок совпадает,
			то берем их
			иначе формируем из данных
			-->
			<xsl:choose>
			<xsl:when test="count(/document/metadata[@modelId=$model_id]/field) = $col_count">
				<!-- Есть метаданные и количество колонок совпадает -->
				<!-- Заголовки -->
				<xsl:apply-templates select="/document/metadata[@modelId=$model_id]"/>
				
				<!-- all data rows -->
				<xsl:apply-templates select="row"/>				
			</xsl:when>
			<xsl:otherwise>
				<!-- все child первого ряда
				установка ширины
				-->
				<xsl:for-each select="row[1]/*">					
					<!-- Ширина колонки или из метаданных или по дефолту -->
					<xsl:variable name="field_id" select="name()"/>
					<xsl:variable name="metadata" select="/document/metadata[@modelId=$model_id]/field[@id=$field_id]"/>
					<xsl:if test="not($metadata) or ($metadata and not($metadata/@sysCol='TRUE'))">
						<!-- Либо нет метеданных либо колонка не системная -->
						<xsl:variable name="len">
							<xsl:choose>
								<xsl:when test="$metadata and not($metadata/@sysCol='TRUE') and ($metadata/@dataType=$DT_INT)">
									<xsl:value-of select="$DEF_WIDTH_INT"/>
								</xsl:when>
								<xsl:when test="$metadata and not($metadata/@sysCol='TRUE') and $metadata/@dataType=$DT_DATE">
									<xsl:value-of select="$DEF_WIDTH_DATE"/>
								</xsl:when>
								<xsl:when test="$metadata and not($metadata/@sysCol='TRUE') and $metadata/@dataType=$DT_TIME">
									<xsl:value-of select="$DEF_WIDTH_TIME"/>
								</xsl:when>					
								<xsl:when test="$metadata and not($metadata/@sysCol='TRUE') and ($metadata/@dataType=$DT_DATETIME or $metadata/@dataType=$DT_DATETIMETZ)">
									<xsl:value-of select="$DEF_WIDTH_DATETIME"/>
								</xsl:when>
								<xsl:when test="$metadata and not($metadata/@sysCol='TRUE') and $metadata/@dataType=$DT_FLOAT">
									<xsl:value-of select="$DEF_WIDTH_DOUBLE"/>
								</xsl:when>
								<xsl:otherwise>
									<xsl:value-of select="$DEF_COL_WIDTH"/>
								</xsl:otherwise>
							</xsl:choose>
						</xsl:variable>	
	
						<Column ss:Width="{$len}"/>					
					</xsl:if>
				</xsl:for-each>
								
				<xsl:if test="$row_count &gt; 0">
					<!-- все child первого ряда заголовок-->
					<Row>					
					<xsl:for-each select="row[1]/*">						
						<xsl:variable name="field_id" select="name()"/>
						<xsl:variable name="metadata" select="/document/metadata[@modelId=$model_id]/field[@id=$field_id]"/>
						
						<xsl:if test="not($metadata) or ($metadata and not($metadata/@sysCol='TRUE'))">
							<!-- Либо нет метеданных либо колонка не системная -->
							<xsl:variable name="col_descr">
								<xsl:choose>
									<xsl:when test="$metadata/@alias">
										<xsl:value-of select="$metadata/@alias"/>
									</xsl:when>
									<xsl:when test="$metadata">
										<xsl:value-of select="$metadata/@id"/>
									</xsl:when>									
									<xsl:otherwise>
										<xsl:value-of select="concat('Колонка ',position())"/>
									</xsl:otherwise>
								</xsl:choose>
							</xsl:variable>
						
							<Cell ss:StyleID="s22">
								<Data ss:Type="String"><xsl:value-of select="$col_descr"/></Data>
							</Cell>
						</xsl:if>
					</xsl:for-each>
					</Row>
					
					<!-- Данные-->
					<xsl:for-each select="row[position() &gt; 0]">
						<Row>
						<xsl:for-each select="*">
							<xsl:variable name="field_id" select="name()"/>
							<xsl:variable name="metadata" select="/document/metadata[@modelId=$model_id]/field[@id=$field_id]"/>
						
							<xsl:if test="not($metadata) or ($metadata and not($metadata/@sysCol='TRUE'))">
								<!-- Либо нет метеданных либо колонка не системная -->
							
								<xsl:variable name="field_dt" select="$metadata/@dataType"/>
								<xsl:variable name="field_float_prec" select="$metadata/@precision"/>
								
								<xsl:variable name="excel_style_id">		
									<xsl:choose>
										<xsl:when test="$field_dt=$DT_INT"><xsl:value-of select="$EXCEl_STYLE_ID_INT"/></xsl:when>										
										<xsl:when test="$field_dt=$DT_FLOAT"><xsl:value-of select="$EXCEl_STYLE_ID_FLOAT"/></xsl:when>
										<xsl:when test="$field_dt=$DT_DATETIME or $field_dt=$DT_DATETIMETZ"><xsl:value-of select="$EXCEl_STYLE_ID_DATETIME"/></xsl:when>
										<xsl:when test="$field_dt=$DT_DATE"><xsl:value-of select="$EXCEl_STYLE_ID_DATE"/></xsl:when>
										<xsl:otherwise><xsl:value-of select="$EXCEl_STYLE_ID_STRING"/></xsl:otherwise>
									</xsl:choose>
								</xsl:variable>
								<xsl:variable name="excel_data_type">
									<xsl:choose>
										<xsl:when test="$field_dt=$DT_INT"><xsl:value-of select="$EXCEl_DT_INT"/></xsl:when>
										<xsl:when test="$field_dt=$DT_FLOAT"><xsl:value-of select="$EXCEl_DT_FLOAT"/></xsl:when>
										<xsl:when test="$field_dt=$DT_DATETIME or $field_dt=$DT_DATETIMETZ"><xsl:value-of select="$EXCEl_DT_DATETIME"/></xsl:when>
										<xsl:when test="$field_dt=$DT_DATE"><xsl:value-of select="$EXCEl_DT_DATE"/></xsl:when>
										<xsl:otherwise><xsl:value-of select="$EXCEl_DT_STRING"/></xsl:otherwise>
									</xsl:choose>
								</xsl:variable>
								
								<!-- Исправление определенных значений-->
								<xsl:variable name="excel_val">
									<xsl:choose>
										<xsl:when test="$field_dt=$DT_BOOL and node()='true'"><xsl:value-of select="'да'"/></xsl:when>
										<xsl:when test="$field_dt=$DT_BOOL and node()='false'"><xsl:value-of select="'нет'"/></xsl:when>
										<xsl:when test="$field_dt=$DT_JSON and node()='false'"><xsl:value-of select="'нет'"/></xsl:when>
										
										<xsl:when test="($field_dt=$DT_DATETIME or $field_dt=$DT_DATETIMETZ) and string-length(node())=0">
											<xsl:value-of select="'1970-01-01T00:00:00'"/>
										</xsl:when>
										<xsl:when test="($field_dt=$DT_DATETIME or $field_dt=$DT_DATETIMETZ) and string-length(substring-before(node(),'+'))&gt;0">
											<xsl:value-of select="substring-before(node(),'+')"/>
										</xsl:when>
										
										<xsl:when test="$field_dt=$DT_DATE">
											<xsl:value-of select="concat(node(),'T00:00:00.000')"/>
										</xsl:when>
										
										<!--
										<xsl:when test="$field_dt=$DT_DATE">
											<xsl:call-template name="format_date">
												<xsl:with-param name="val" select="node()" />
												<xsl:with-param name="formatStr" select="''" />
											</xsl:call-template>		
										</xsl:when>
										-->
										<xsl:when test="$field_dt=$DT_JSON or $field_dt=$DT_JSONB">
											<xsl:call-template name="format_json">
												<xsl:with-param name="val" select="node()" />
											</xsl:call-template>		
										</xsl:when>							
										<xsl:otherwise><xsl:value-of select="node()"/></xsl:otherwise>
									</xsl:choose>
								</xsl:variable>
								
						
								<Cell ss:StyleID="{$excel_style_id}">
									<Data ss:Type="{$excel_data_type}"><xsl:value-of select="$excel_val"/></Data>
								</Cell>
							</xsl:if>
						</xsl:for-each>
						</Row>
					</xsl:for-each>					
				</xsl:if>
			</xsl:otherwise>
			</xsl:choose>
		</Table>
		<WorksheetOptions xmlns="urn:schemas-microsoft-com:office:excel">
			<PageSetup>
				<PageMargins x:Bottom="0.984251969" x:Left="0.78740157499999996"
				x:Right="0.78740157499999996" x:Top="0.984251969"/>
			</PageSetup>
			<Print>
				<ValidPrinterInfo/>
				<PaperSizeIndex>9</PaperSizeIndex>
				<HorizontalResolution>600</HorizontalResolution>
				<VerticalResolution>0</VerticalResolution>
			</Print>
			<Selected/>
			<ProtectObjects>False</ProtectObjects>
			<ProtectScenarios>False</ProtectScenarios>
		</WorksheetOptions>
	</Worksheet>
</xsl:template>

<!-- table header -->
<xsl:template match="metadata">
	<xsl:for-each select="field[not(@sysCol='TRUE')]">
		<xsl:variable name="len">
			<xsl:choose>
				<xsl:when test="@dataType=$DT_INT">
					<xsl:value-of select="$DEF_WIDTH_INT"/>
				</xsl:when>
				<xsl:when test="@dataType=$DT_DATE">
					<xsl:value-of select="$DEF_WIDTH_DATE"/>
				</xsl:when>
				<xsl:when test="@dataType=$DT_TIME">
					<xsl:value-of select="$DEF_WIDTH_TIME"/>
				</xsl:when>					
				<xsl:when test="@dataType=$DT_DATETIME or @dataType=$DT_DATETIMETZ">
					<xsl:value-of select="$DEF_WIDTH_DATETIME"/>
				</xsl:when>
				<xsl:when test="@dataType=$DT_FLOAT">
					<xsl:value-of select="$DEF_WIDTH_DOUBLE"/>
				</xsl:when>
				<xsl:otherwise>
					<xsl:value-of select="$DEF_COL_WIDTH"/>
				</xsl:otherwise>
			</xsl:choose>
		</xsl:variable>	
	
		<Column ss:Width="{$len}"/>
	</xsl:for-each>
	<Row>
		<xsl:apply-templates/>
	</Row>
</xsl:template>

<!-- header field -->
<xsl:template match="metadata/*">
	<xsl:if test="not(@sysCol='TRUE')">
		<xsl:variable name="col">
			<xsl:choose>
				<xsl:when test="@alias">
					<xsl:value-of select="@alias"/>
				</xsl:when>
				<xsl:otherwise>
					<xsl:value-of select="@id"/>
				</xsl:otherwise>			
			</xsl:choose>
		</xsl:variable>
		<Cell ss:StyleID="s22">
			<Data ss:Type="String"><xsl:value-of select="$col"/></Data>
		</Cell>
	</xsl:if>
</xsl:template>

<!-- table row
сохраним порядок из заголовка модели
пройдем циклом по заголовку
хотя данные могут быть выбраны в другом порядке
-->
<xsl:template match="row">
	<Row>
		<xsl:apply-templates/>
	</Row>
</xsl:template>

<!-- table cell все кроме системных колонок -->
<xsl:template match="row/*">
	<xsl:variable name="field_id" select="name()"/>
	<xsl:variable name="model_id" select="../../@id"/>	
	<xsl:if test="/document/metadata[@modelId=$model_id]/field[@id=$field_id and not(@sysCol='TRUE')]">	
		<xsl:variable name="field_dt" select="/document/metadata[@modelId=$model_id]/field[@id=$field_id]/@dataType"/>
		<xsl:variable name="excel_style_id">		
			<xsl:choose>
				<xsl:when test="$field_dt=$DT_INT"><xsl:value-of select="$EXCEl_STYLE_ID_INT"/></xsl:when>
				<xsl:when test="$field_dt=$DT_FLOAT"><xsl:value-of select="$EXCEl_STYLE_ID_FLOAT"/></xsl:when>
				<xsl:when test="$field_dt=$DT_DATETIME or $field_dt=$DT_DATETIMETZ"><xsl:value-of select="$EXCEl_STYLE_ID_DATETIME"/></xsl:when>
				<xsl:when test="$field_dt=$DT_DATE"><xsl:value-of select="$EXCEl_STYLE_ID_DATE"/></xsl:when>
				<xsl:otherwise><xsl:value-of select="$EXCEl_STYLE_ID_STRING"/></xsl:otherwise>
			</xsl:choose>
		</xsl:variable>
		<xsl:variable name="excel_data_type">
			<xsl:choose>
				<xsl:when test="$field_dt=$DT_INT"><xsl:value-of select="$EXCEl_DT_INT"/></xsl:when>
				<xsl:when test="$field_dt=$DT_FLOAT"><xsl:value-of select="$EXCEl_DT_FLOAT"/></xsl:when>
				<xsl:when test="$field_dt=$DT_DATETIME or $field_dt=$DT_DATETIMETZ"><xsl:value-of select="$EXCEl_DT_DATETIME"/></xsl:when>
				<xsl:when test="$field_dt=$DT_DATE"><xsl:value-of select="$EXCEl_DT_DATE"/></xsl:when>
				<xsl:otherwise><xsl:value-of select="$EXCEl_DT_STRING"/></xsl:otherwise>
			</xsl:choose>
		</xsl:variable>
		<xsl:variable name="excel_val">
			<xsl:choose>
				<xsl:when test="$field_dt=$DT_BOOL and node()='true'"><xsl:value-of select="'да'"/></xsl:when>
				<xsl:when test="$field_dt=$DT_BOOL and node()='false'"><xsl:value-of select="'нет'"/></xsl:when>
				<xsl:when test="($field_dt=$DT_DATETIME or $field_dt=$DT_DATETIMETZ) and string-length(node())=0">
					<xsl:value-of select="'1970-01-01T00:00:00'"/>
				</xsl:when>
				<xsl:when test="($field_dt=$DT_DATETIME or $field_dt=$DT_DATETIMETZ) and string-length(substring-before(node(),'+'))&gt;0">
					<xsl:value-of select="substring-before(node(),'+')"/>
				</xsl:when>
				
				<!--
				<xsl:when test="$field_dt=$DT_DATE">
					<xsl:call-template name="format_date">
						<xsl:with-param name="val" select="node()" />
						<xsl:with-param name="formatStr" select="''" />
					</xsl:call-template>		
				</xsl:when>
				-->
				<xsl:when test="$field_dt=$DT_DATE">
					<xsl:value-of select="concat(node(),'T00:00:00.000')"/>
				</xsl:when>
				
				
				<xsl:when test="$field_dt=$DT_JSON or $field_dt=$DT_JSONB">
					<xsl:call-template name="format_json">
						<xsl:with-param name="val" select="node()" />
					</xsl:call-template>		
				</xsl:when>														
				<xsl:otherwise><xsl:value-of select="node()"/></xsl:otherwise>
			</xsl:choose>
		</xsl:variable>
		
		<Cell ss:StyleID="{$excel_style_id}">
			<Data ss:Type="{$excel_data_type}"><xsl:value-of select="$excel_val"/></Data>
		</Cell>
	</xsl:if>
</xsl:template>

</xsl:stylesheet>
