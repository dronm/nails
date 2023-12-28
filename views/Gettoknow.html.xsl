<?xml version="1.0" encoding="UTF-8"?>

<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">

<xsl:import href="ViewBase.html.xsl"/>

<xsl:template match="/document">
<html>
	<head>
		<xsl:call-template name="initHead"/>		
		
		<script>		
			function pageLoad(){
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

							<div id="Gettoknow" class="panel panel-body login-form">
								<div class="text-center">
									<a href="/index.html">
										<img id="logo-img2" src="img/logo2.png">
										</img>
									</a>
								</div>
								
								<div class="text-center">
									<h5>знакомство</h5>
								</div>
								<p>Компания <span class="stressed-text">"PROНогти"</span> - региональная сеть студий красоты, своей миссией мы хотим сделать мир прекраснее, а особенно милых дам.
<span class="stressed-text">ООО "ПроНогти"</span> - действует только в соответствии с действующим законодательством, поэтому для того, чтобы начать с нами сотрудничество, необходимо открыть самозанятость!
								</p>
								
								<p>Преимущества получения официального дохода самозанятым при сотрудничестве с нами:
									<ul>
										<li>возможность легально работать;</li>
										<li>отличный кредитный рейтинг;</li>
										<li>не нужно отчитываться за крупные покупки перед налоговыми органами;</li>
										<li>прозрачное движение денежных средств;</li>
										<li>отсутствие отчетности и деклараций;</li>
										<li>самая низкая ставка налога в РФ и половину налога (3%) мы дополнительно компенсируем;</li>
										<li>регулярные поступления (каждые 2 недели).</li>
										<li>множество клиентов без лишних забот!</li>	
									</ul>
								</p>
								
								<p>Узнать больше о самозанятости, а так же о том, как зарегистрироваться и открыть расчетный счёт, Вы можете пройдя по ссылкам банков:
									<ul>
										<li> <a href="https://allo.tochka.com/samozanyatye">Точка Банк (Открытие);</a></li>
										<li><a href="https://www.tinkoff.ru/business/help/business-self-employment/self-employed/work/start">Тинькофф;</a></li>
										<li><a href="https://www.sberbank.com/ru/svoedelo/start">Сбер.</a></li>
									</ul>
								</p>
								
								<p>Если Вы зарегистрировались в качестве самозанятого и открыли расчетный счёт, то переходите к регистрации в системе <span class="stressed-text">PROНогти</span>:
								</p>
								<a id="Gettoknow:submit" href="Registration" class="btn bg-{$COLOR_PALETTE}">Начать регистрацию</a>									
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
		
		<xsl:call-template name="initJS"/>
	</body>
</html>		
</xsl:template>

</xsl:stylesheet>
