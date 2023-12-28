<?xml version="1.0" encoding="UTF-8"?>

<xsl:stylesheet version="1.0" 
 xmlns:html="http://www.w3.org/TR/REC-html40"
 xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
 xmlns:fo="http://www.w3.org/1999/XSL/Format">

<xsl:output method="html"/> 

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

<!-- Main template-->
<xsl:template match="/">
	<xsl:apply-templates select="document/model"/>	
</xsl:template>

<!-- Error -->
<xsl:template match="model[@id='Response']">
	<div class="error">
		<xsl:value-of select="row[1]/descr"/>
	</div>
</xsl:template>

<!-- table -->
<xsl:template match="model">
	<xsl:variable name="model_id" select="@id"/>	
	<table id="{$model_id}" class="tabel table-bordered table-striped">
		<thead>
			<tr>
				<xsl:for-each select="./row[1]/*">
					<xsl:variable name="field_id" select="name()"/>
					<xsl:if test="$field_id != 'sys_level_val' and $field_id != 'sys_level_count' and $field_id != 'sys_level_col_count'">
					<xsl:variable name="label">
						<xsl:choose>
							<xsl:when test="/document/metadata[@modelId=$model_id]/field[@id=$field_id]/@alias">
								<xsl:value-of select="/document/metadata[@modelId=$model_id]/field[@id=$field_id]/@alias"/>
							</xsl:when>
							<xsl:when test="/document/metadata[@modelId=$model_id]/@id">
								<xsl:value-of select="/document/metadata[@modelId=$model_id]/@id"/>
							</xsl:when>
							<xsl:otherwise>
								<!-- <xsl:value-of select="$field_id"/>-->
								<xsl:call-template name="string-replace-all">
									<xsl:with-param name="text" select="$field_id"/>
									<xsl:with-param name="replace" select="'_x0020_'"/>
									<xsl:with-param name="by" select="' '"/>
								</xsl:call-template>																					
							</xsl:otherwise>
						</xsl:choose>
					</xsl:variable>
					<!--<th>&#160;&#160;&#160;&#160;&#160;<xsl:value-of select="$label"/>&#160;&#160;&#160;&#160;&#160;</th>-->
					<th><xsl:value-of select="$label"/></th>
					</xsl:if>
				</xsl:for-each>
			</tr>
		</thead>
	
		<tbody>
			<xsl:apply-templates/>
		</tbody>
	</table>
</xsl:template>

<!-- table header -->

<!-- header field -->

<!-- table row -->
<xsl:template match="row">
	<xsl:variable name="class_name">
		<xsl:choose>
		<xsl:when test="sys_level_val &gt;= 0">level_<xsl:value-of select="sys_level_val"/>
		</xsl:when>
		<xsl:otherwise></xsl:otherwise>
		</xsl:choose>	
	</xsl:variable>
	<tr class="{$class_name}">
		<xsl:apply-templates/>
	</tr>
</xsl:template>

<!-- table cell -->
<xsl:template match="row/*">
	<xsl:variable name="pos" select="position() div 2"/>
	<xsl:variable name="field_id" select="name()"/>
	<xsl:variable name="model_id" select="./../../@id"/>	
	<xsl:variable name="dataType" select="/document/metadata[@modelId=$model_id]/field[@id=$field_id]/@dataType"/>
	
	<xsl:variable name="td_align">
		<xsl:choose>
		<xsl:when test="$dataType='0' or $dataType='1' or $dataType='3' or $dataType='4' or $dataType='5' or $dataType='6'">right</xsl:when>
		<xsl:when test="$dataType='9' or $dataType='10' or $dataType='11' or $dataType='15' or $dataType='7'">center</xsl:when>
		<xsl:otherwise>left</xsl:otherwise>
		</xsl:choose>
	</xsl:variable>
	
	<xsl:choose>
	<xsl:when test="(number(../sys_level_count)=number(../sys_level_val))">
		<!-- ИТОГ
		<td align="{$td_align}" fieldId="{name()}" colspan="{greatest(../sys_level_col_count/element)}">Итого</td>
		-->
		<td align="{$td_align}" fieldId="{name()}"><xsl:value-of select="node()"/></td>
	</xsl:when>
	<!--ToDO
		- Сделать итоги нормально в гуи
		- передавать как то количество полей во всех группировках
		- убрать поля слева у группировок, которые не на максимальном своем уровне
	<xsl:when test="node() = (../[name()=$field_id]/node())">
		<td></td>
	</xsl:when>
	-->
	<xsl:otherwise>
		<!-- НЕТ span-->
		<td align="{$td_align}" fieldId="{name()}"><xsl:value-of select="node()"/></td>
	</xsl:otherwise>
	</xsl:choose>
</xsl:template>

<!-- заглушка -->
<xsl:template match="row/sys_level_val">
</xsl:template>

<xsl:template match="row/sys_level_count">
</xsl:template>

<xsl:template match="row/sys_level_col_count">
</xsl:template>

<!-- javascript -->
<xsl:template match="model[@id='ModelJavaScript']">
</xsl:template>
<!-- css -->
<xsl:template match="model[@id='ModelStyleSheet']">
</xsl:template>
<!-- menu-->
<xsl:template match="model[@id='MainMenu_Model']">
</xsl:template>
<!-- vars-->
<xsl:template match="model[@id='ModelVars']">
</xsl:template>

</xsl:stylesheet>
