<?xml version="1.0" encoding="UTF-8"?>

<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">

<xsl:import href="ViewBase.html.xsl"/>

<xsl:template match="/document">
<html>
	<head>
		<xsl:call-template name="initHead"/>		
		
		<script>		
			function pageLoad(){
			
				<xsl:call-template name="initApp"/>
				
				var view = new Login_View("Login");
				view.toDOM();
			}
		</script>		
		
	</head>
	<body onload="pageLoad();" class="login-container">

	<!--<xsl:call-template name="page_header"/>-->

	<!-- Page container -->
	<div class="page-container container-nopadding">

		<!-- Page content -->
		<div class="page-content">

			<!-- Main content -->
			<div class="content-wrapper">

				<!-- Content area -->
				<div class="content">

					<!-- Advanced login -->
					<form id="Login">
						<div class="panel panel-body login-form">
							<div class="text-center">
								<a href="/index.html">
									<img id="logo-img" src="img/logo.png">
									</img>
								</a>
								<div id="Login:error"></div>
							</div>

							<div class="form-group has-feedback has-feedback-left">
								<input id="Login:user" type="text" class="form-control" placeholder="Имя или номер телефона"/>
								<div class="form-control-feedback">
									<i class="icon-user text-muted"></i>
								</div>
								<div id="Login:user:error"/>
							</div>

							<div class="form-group has-feedback has-feedback-left">
								<input id="Login:pwd" type="password" class="form-control" placeholder="Пароль"/>
								<div class="form-control-feedback">
									<i class="icon-lock2 text-muted"></i>
								</div>
								<div id="Login:pwd:error"/>
							</div>
							
							<div class="form-group login-options">
								<div class="row">
									<div class="col-sm-6 text-right">
										<a href="PasswordRecovery">Забыли пароль?</a>
									</div>
								</div>
							</div>
							
							<!-- IE8 condition -->
							<xsl:comment><![CDATA[[if IE 8]>
							<div class="alert alert-danger alert-styled-left alert-bordered">К сожалению, работа в личном кабинете с данным браузером невозможна!
							</div>
							<![endif]]]></xsl:comment>
							
							<div class="form-group">
								<div id="Login:submit_login" class="btn bg-{$COLOR_PALETTE} btn-block">Войти <i class="icon-arrow-right14 position-right"></i></div>
							</div>
							
							<div class="content-divider text-muted form-group join_contract_text"><span>Если Вы не зарегистрированы, то ознакомьтесь с договором присоединения</span></div>
							<div id="Login:join_contract" class="join_contract form-control">
							<xsl:value-of select="/document/model[@id='ServerVars']/row[1]/joinContract"/>
							</div>
							<a href="Gettoknow" class="btn bg-{$COLOR_PALETTE} btn-block content-group">Ознакомлен (-а)</a>
							
						</div>
					</form>
					<!-- /advanced login -->


					<!-- Footer -->
					<div class="footer text-muted text-center">
						2023 <a href="#"></a>
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
		
	<script>
			var n = document.getElementById("Login:user");
			if (n){
				if (document.activeElement &amp;&amp; document.activeElement.id!="Login:pwd"){
					n.focus();
				}
			}
	</script>
		
		<xsl:call-template name="initJS"/>
	</body>
</html>		
</xsl:template>

</xsl:stylesheet>
