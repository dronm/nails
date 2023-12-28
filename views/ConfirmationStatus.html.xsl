<?xml version="1.0" encoding="UTF-8"?>

<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">

<xsl:import href="ViewBase.html.xsl"/>

<xsl:template match="/document">
<html>
	<head>
		<meta http-equiv="content-type" content="text/html; charset=UTF-8"/>
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1" />
		
		<link rel="stylesheet" href="{concat('/js20/ext/Limitless v1.4/layout_1/LTR/default/assets/icons/icomoon/styles.css','?',$VERSION)}" type="text/css"/>
		<link rel="stylesheet" href="{concat('/js20/ext/Limitless v1.4/layout_1/LTR/default/assets/css/bootstrap.css','?',$VERSION)}" type="text/css"/>
		<link rel="stylesheet" href="{concat('/js20/ext/Limitless v1.4/layout_1/LTR/default/assets/css/core.min.css','?',$VERSION)}" type="text/css"/>
		<link rel="stylesheet" href="{concat('/js20/ext/Limitless v1.4/layout_1/LTR/default/assets/css/components.min.css','?',$VERSION)}" type="text/css"/>
		<link rel="stylesheet" href="{concat('/js20/ext/Limitless v1.4/layout_1/LTR/default/assets/css/colors.min.css','?',$VERSION)}" type="text/css"/>
		<link rel="stylesheet" href="{concat('/js20/ext/Limitless v1.4/layout_1/LTR/default/assets/css/icons/fontawesome/styles.min.css','?',$VERSION)}" type="text/css"/>
		<link rel="stylesheet" href="{concat('/js20/custom-css/style.css','?',$VERSION)}" type="text/css"/>
		<title>Ногти</title>	
	</head>
	<body class="login-container">

	<div class="navbar">
		<div class="navbar-header">
			<a class="navbar-brand1" href="/index.html">
				<img id="main-logo-img" src="/img/logo2.png"></img>
			</a>
		</div>
	</div>
	<!-- Page container -->
	<div class="page-container">

		<!-- Page content -->
		<div class="page-content">

			<!-- Main content -->
			<div class="content-wrapper">

				<!-- Content area -->
				<div class="content">
					<xsl:choose>
					<xsl:when test="model[@id='Response']/row[1]/result/node()!='0'">	
					<div class="alert alert-danger alert-styled-left">
						<p><xsl:value-of select="model[@id='Response']/row[1]/descr"/>
						</p>
					</div>
					</xsl:when>
					<xsl:otherwise>
					<div class="alert alert-success alert-styled-left">
						<p>
							<strong>
								<xsl:value-of select="model[@id='ConfirmationStatusResult']/row[1]/name_full"/>
							</strong>
						</p>
						<xsl:if test="model[@id='ConfirmationStatusResult']/row[1]/field='tel'">	
							<p>Телефон подтвержден</p>
						</xsl:if>
						<xsl:if test="model[@id='ConfirmationStatusResult']/row[1]/field='email'">	
							<p>Адрес электронной почты подтвержден</p>
						</xsl:if>
					</div>
					</xsl:otherwise>
					</xsl:choose>
				
					<!-- Footer -->
					<div class="footer text-muted text-center">
						2023. <a href="#"></a>
					</div>
					<!-- /footer -->

				</div>
				<!-- /content area -->

			</div>
			<!-- /main content -->

		</div>
		<!-- /page content -->

	</div>
	<!-- /page container -->
		
	</body>
</html>		
</xsl:template>

</xsl:stylesheet>
