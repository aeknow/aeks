<!DOCTYPE html>
<html >
<head>
	<meta name="generator" content="Hugo 0.48" />
	 <meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<title>{{.Title}} - AEKs</title>
	
	<meta name="description" content="John Doe&#39;s Personal blog about everything">
		<meta property="og:title" content="Mainroad" />
<meta property="og:description" content="John Doe&#39;s Personal blog about everything" />
<meta property="og:type" content="website" />
<meta property="og:url" content="https://mainroad-demo.netlify.app/" />
<meta property="og:updated_time" content="2018-04-16T00:00:00&#43;00:00"/>
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
	<link rel="dns-prefetch" href="//fonts.googleapis.com">
	<link rel="dns-prefetch" href="//fonts.gstatic.com">
	<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Open+Sans:400,400i,700">

	<link rel="stylesheet" href="/themes/mainroad/css/style.css">
	
	<link rel="alternate" type="application/rss+xml" href="/index.xml" title="Mainroad">

	<link rel="shortcut icon" href="/favicon.ico">
		
</head>
<body class="body">
	<div class="container container--outer">
  {{ template "header" .}}
  
		<div class="wrapper flex">
			<div class="primary">

<main class="main list" role="main">
<article class="list__item post">
<h2><a  href="/goaens?gohome=gohome">Home </a> >> {{.Title}}</h2>
</article>	
	
{{range .Posts}}
{{if gt (len .Title) 0}}
<article class="list__item post">
	{{if gt (len .Remark) 0}}
	<figure class="list__thumbnail thumbnail">
		<a class="thumbnail__link" href="view?pubkey={{.Pubkey}}&hash={{.Hash}}&aid={{.Aid}}" >
		<img class="thumbnail__image" src="{{.Remark}}" alt="{{.Title}}">
		</a>
	</figure>
	{{end}}
	<header class="list__header">
		<h2 class="list__title post__title">
			<a href="view?pubkey={{.Pubkey}}&hash={{.Hash}}&aid={{.Aid}}" rel="bookmark">
			 {{.Title}}
			</a>
			{{if gt (len .IsOwner) 1}}
			<small><a href="editpage?pubkey={{.Pubkey}}&aid={{.Aid}}" target=_blank>Edit</a></small>
			{{else}}
			<small><a href="editpage?pubkey={{.Pubkey}}&aid={{.Aid}}&action=fork" target=_blank>Fork</a></small>
			{{end}}
			
		</h2>
		
		<div class="list__meta meta">
<div class="meta__item-datetime meta__item">
	<svg class="meta__icon icon icon-time" width="16" height="14" viewBox="0 0 30 28"><path d="M15 0C7 0 1 6 1 14s6 14 14 14 14-6 14-14S23 0 15 0zm0 25C9 25 4 20 4 14S9 3 15 3s11 5 11 11-5 11-11 11zm1-18h-2v8.4l6.8 4.4L22 18l-6-3.8V7z"/></svg>
	<time class="meta__text" datetime="2018-04-16T00:00:00Z">{{.LastModTime}}</time>
	</div>
	{{if gt (len .Tags) 90}}
	
	<div class="meta__item-categories meta__item"><svg class="meta__icon icon icon-category" width="16" height="16" viewBox="0 0 16 16"><path d="m7 2l1 2h8v11h-16v-13z"/></svg>
	
	<span class="meta__text">{{.Tags}}</span>
	
	</div>
	{{end}}
	
	
</div>
	</header>
<div class="content list__excerpt post__content clearfix">
		<p>{{.Abstract}}</p>
	</div>	
</article>
    {{end}}
{{end}}
     
</main>

{{- if or (.Paginator.HasPrev) (.Paginator.HasNext) }}
<div class="pagination">	
	{{- if .Paginator.HasPrev }}
	<a class="pagination__item pagination__item--prev btn" href="{{ .Paginator.PrevURL }}">«</a>
	{{- end }}
	
	<span class="pagination__item pagination__item--current">
		{{- .Paginator.PageNumber }}/{{ .Paginator.TotalPages -}}
	</span>
	{{- if .Paginator.HasNext }}
	<a class="pagination__item pagination__item--next btn" href="{{ .Paginator.NextURL }}">»</a>
	{{- end }}
</div>
{{- end }}


			</div>

<aside class="sidebar"><div class="widget-search widget">
	<form class="widget-search__form" role="search" method="get" action="https://google.com/search">
		<label>
			<input class="widget-search__field" type="search" placeholder="SEARCH…" value="" name="q" aria-label="SEARCH…">
		</label>
		<input class="widget-search__submit" type="submit" value="Search">
		<input type="hidden" name="sitesearch" value="https://mainroad-demo.netlify.app/" />
	</form>
</div>
<div class="widget-recent widget">
	<h4 class="widget__title">Recent Posts</h4>
	<div class="widget__content">
		<ul class="widget__list">
			
{{range .Posts}}
	{{if gt (len .Title) 0}}
<li class="widget__item"><a class="widget__link" href="/view?pubkey=ak_fCCw1JEkvXdztZxk8FRGNAkvmArhVeow89e64yX4AxbCPrVh5&hash={{.Hash}}&aid={{.Aid}}">{{.Title}}</a></li>
    {{end}}
{{end}}		
		</ul>
	</div>
</div>
<div class="widget-categories widget">
	<h4 class="widget__title">Related Categories</h4>
	<div class="widget__content">
		<ul class="widget__list">
			{{.Categories}}			
		</ul>
	</div>
</div>
<div class="widget-taglist widget">
	<h4 class="widget__title">Related Tags</h4>
	<div class="widget__content">
		{{.Tags}}		
	</div>
</div><div class="widget-text widget">
	<h4 class="widget__title">About</h4>
	<div class="widget__content">
		<p>{{.Site.Description}}</p>
	</div>
</div>
</aside>



		</div>
		{{ template "footer" .}}
	</div>
<script async defer src="/js/menu.js"></script>
</body>
</html>
