{{ define "aside" }}
  <aside class="sidebar">
				<div class="widget-search widget">
				
	<form class="widget-search__form" role="search" method="get" action="https://google.com/search">
		<label>
			<input class="widget-search__field" type="search" placeholder="SEARCH…" value="" name="q" aria-label="SEARCH…">
		</label>
		<input class="widget-search__submit" type="submit" value="Search">
		<input type="hidden" name="sitesearch" value="https://mainroad-demo.netlify.app/" />
	</form>
</div>

<div class="widget-categories widget">
	<h4 class="widget__title">Table of Content</h4>
	<div class="widget__content">
		<div class="markdown-body editormd-preview-container" id="custom-toc-container">#custom-toc-container</div>
	</div>
</div>


<div class="widget-recent widget">
	<h4 class="widget__title">Recent Posts</h4>
	<div class="widget__content">
		<ul class="widget__list">
			{{.LastTenLink}}			
		</ul>
	</div>
</div>
<div class="widget-categories widget">
	<h4 class="widget__title">Categories</h4>
	<div class="widget__content">
		<ul class="widget__list">
			{{.AllCatgoriesLink}}		
		</ul>
	</div>
</div>
<div class="widget-taglist widget">
	<h4 class="widget__title">Tags</h4>
	<div class="widget__content">
		{{.AllTagsLink}}		
	</div>
</div><div class="widget-text widget">
	<h4 class="widget__title">Support us</h4>
	<div class="widget__content">
		<p>If you like <strong>Mainroad</strong> theme, remember that you can show your support by starring it on Github. It only takes a couple of seconds!</p>
	</div>
</div>
</aside>

{{ end }}
