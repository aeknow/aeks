{{ define "header" }}
<header class="header">
	<div class="container header__container">
		
	<div class="logo">
		<a class="logo__link" href="/view?pubkey={{.Site.Pubkey}}&viewtype=author" title="{{.Site.Title}}" rel="home">
			<div class="logo__item logo__text">
					<div class="logo__title">{{.Site.Title}}</div>
					<div class="logo__tagline">{{.Site.Subtitle}}</div>
				</div>
		</a>
	</div>
		
<nav class="menu">
	<button class="menu__btn" aria-haspopup="true" aria-expanded="false" tabindex="0">
		<span class="menu__btn-title" tabindex="-1">Menu</span>
	</button>
	<ul class="menu__list">
		<li class="menu__item">
			<a class="menu__link" href="/view?pubkey={{.Site.Pubkey}}&viewtype=author">
				
				<span class="menu__text">Haeme</span>
				
			</a>
		</li>
		
		<!-- 
		<li class="menu__item">
			<a class="menu__link" href="/post/goisforlovers/">
				
				<span class="menu__text">Article</span>
				
			</a>
		</li>
		<li class="menu__item">
			<a class="menu__link" href="/about/">
				
				<span class="menu__text">Wiki</span>
				
			</a>
		</li>
		<li class="menu__item">
			<a class="menu__link" href="/post/hugoisforlovers/">
				
				<span class="menu__text">Book</span>
				
			</a>
		</li>
		
		<li class="menu__item">
			<a class="menu__link" href="/post/hugoisforlovers/">
				
				<span class="menu__text">Video</span>
				
			</a>
		</li>
		
		-->
	</ul>
</nav>

	</div>
</header>

{{ end }}
