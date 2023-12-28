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
				
				var view = new Registration_View("Registration");
				view.toDOM();
			}
		</script>		
		
	</head>
	<body onload="pageLoad();" class="login-container">

	<!-- Page container -->
	<div class="page-container container-nopadding">

		<!-- Page content -->
		<div class="page-content">

			<!-- Main content -->
			<div class="content-wrapper">

				<!-- Content area -->
				<div class="content">

					<div class="row">
						<div id="windowData" class="col-lg-12">

							<div id="Registration" class="panel panel-body login-form">
								<div class="text-center">
									<a href="/index.html">
										<img id="logo-img2" src="img/logo2.png">
										</img>
									</a>
								</div>
							
								<div id="Registration:inn"/>
								
								<div id="Registration:banks_ref"/>
								
								<div id="Registration:bank_acc"/>
								
								<div id="Registration:name_full" />
								<div id="Registration:tel" />
								<div id="Registration:email" />
								<div id="Registration:address_reg" />
								<div id="Registration:birthdate" />
								
								<div id="Registration:passport" />
								<div id="Registration:name" />

								<div class="form-group">
									<div class="input-group col-lg-12">
										<div id="Registration:captcha"/>
									</div>
								</div>
								
								<div id="Registration:info" class="alert alert-styled-left hidden">
								</div>								
								
								<a id="Registration:submit" href="#" class="btn bg-{$COLOR_PALETTE} hidden">Завершить</a>	
								
								<div id="Registration:registered" class="alert alert-info alert-styled-left hidden">
									<p>Пользователь зарегистрирован, ждите активации
									</p>
								</div>
								
							</div>
						</div>
						
						<div class="windowMessage hidden">
						</div>
						<!--waiting  -->
						<div id="waiting">
							<div>Обработка...</div>
							<img src="img/loading.gif?{$VERSION}"/>
						</div>
						
					</div>
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
		
	<script>
			var n = document.getElementById("Registration:user");
			if (n){
				if (document.activeElement &amp;&amp; document.activeElement.id!="Registration:user"){
					n.focus();
				}
			}
	</script>
		
		<xsl:call-template name="initJS"/>
	</body>
</html>		
</xsl:template>

</xsl:stylesheet>
