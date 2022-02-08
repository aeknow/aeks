<!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>{{.PageTitle}} - AEKs</title> 
    <meta name="description" content="{{.PageDescription}}">  
    <link rel="stylesheet" href="/views/editor.md/examples/css/style.css" />
    <link rel="stylesheet" href="/views/editor.md/css/editormd.css" />
    <link rel="stylesheet" href="/themes/mainroad/css/style.css">
    <link rel="shortcut icon" href="/favicon.ico">

</head>

<body class="body" style="text-align:left">
    <div class="container container--outer">

        {{ template "header" .}}

        <div class="wrapper flex">
            <div class="primary">

                <main class="main" role="main">
                    <article class="post">
                        <header class="post__header">

                            <h1 class="post__title">{{.PageTitle}} {{if eq .PageAuthor .Account}}
                                <small><a href="editpage?pubkey={{.Account}}&aid={{.PageAid}}" target=_blank>Edit</a></small> {{else}}
                                <small><a href="editpage?pubkey={{.PageAuthor}}&aid={{.PageAid}}&action=fork" target=_blank>Fork</a></small> {{end}}
                            </h1>
                            <div class="post__meta meta">
                                <div class="meta__item-datetime meta__item">
                                    <svg class="meta__icon icon icon-time" width="16" height="14" viewBox="0 0 30 28"><path d="M15 0C7 0 1 6 1 14s6 14 14 14 14-6 14-14S23 0 15 0zm0 25C9 25 4 20 4 14S9 3 15 3s11 5 11 11-5 11-11 11zm1-18h-2v8.4l6.8 4.4L22 18l-6-3.8V7z"/></svg>
                                    <time class="meta__text" datetime="{{.PubTime}}">Created at {{.PubTime}}</time></div>
                                <div class="meta__item-categories meta__item"><svg class="meta__icon icon icon-category" width="16" height="16" viewBox="0 0 16 16">
		<path d="m7 2l1 2h8v11h-16v-13z"/></svg><span class="meta__text">
			{{.CatgoriesLink}} 
	<div style="float:right">		
	
	{{$signed :="OK"}}
		{{if eq .SigStatus $signed}}		
		<b style="color:green">✔<a href="#" title="Page IPFS hash {{.PageHash}} is signed by {{.PageAuthor}} with Aeternity's Curve 25519 crypto signature {{.PageSignature}} "> Digital Signature Verified</a></b>
		
		{{else}}
		<b style="color:red">? Failed to verify Digital Signature </b>
		{{end}}
	
	(Last modified at {{.LastModTime}})
	</div>
	</span>
                                </div>
                            </div>
                        </header>
                        <div class="content post__content clearfix">
                            <center>
                                <h1>{{.PageTitle}}<small> by  {{.PageAuthorname}}</small></h1>
                            </center>

                            <pre>{{.PageDescription}}</pre>

                            <div id="test-markdown-view">
                                <!-- Server-side output Markdown text -->
                                <textarea style="display:none;">{{.PageContent}}</textarea>
                            </div>
                            <script src="/views/editor.md/examples/js/jquery.min.js"></script>
                            <script src="/views/editor.md/editormd.js"></script>
                            <script src="/views/editor.md/lib/marked.min.js"></script>
                            <script src="/views/editor.md/lib/prettify.min.js"></script>
                            <script type="text/javascript">
                                editormd.katexURL = {
                                    js: "/views/katex/katex.min", // default: //cdnjs.cloudflare.com/ajax/libs/KaTeX/0.3.0/katex.min
                                    css: "/views/katex/katex.min" // default: //cdnjs.cloudflare.com/ajax/libs/KaTeX/0.3.0/katex.min
                                };
                                $(function() {


                                    var testView = editormd.markdownToHTML("test-markdown-view", {

                                        //  htmlDecode : true,  // Enable / disable HTML tag encode.
                                        tocContainer: "#custom-toc-container", // 自定义 ToC 容器层
                                        tex: true,
                                        htmlDecode : "style,script,iframe",  // Note: If enabled, you should filter some dangerous HTML tags for website security.
                                    });


                                });
                            </script>
                        </div>

                        <footer class="post__footer">

                            <div class="post__tags tags clearfix">
                                <svg class="tags__badge icon icon-tag" width="16" height="16" viewBox="0 0 32 32"><path d="M32 19c0 1-1 2-1 2L21 31s-1 1-2 1-2-1-2-1L2 16c-1-1-1.4-2-1.4-2S0 12.5 0 11V3C0 1.5.8.8.8.8S1.5 0 3 0h8c1.5 0 3 .6 3 .6S15 1 16 2l15 15s1 1 1 2zM7 10a3 3 0 1 0 0-6 3 3 0 0 0 0 6z"/></svg>
                                <ul class="tags__list">


                                    {{.TagsLink}}

                                </ul>
                            </div>

                        </footer>
                    </article>
                </main>

                <div class="authorbox clearfix">
                    <figure class="authorbox__avatar">
                        <img alt="John Doe avatar" src="/themes/mainroad/img/avatar.jpg" class="avatar" height="90" width="90">
                    </figure>
                    <div class="authorbox__header">
                        <span class="authorbox__name">{{.AuthorLink}} 
		
		</span>
                    </div>
                    <div class="authorbox__description">
                        {{.Site.AuthorDescription}}
                    </div>
                </div>

                <nav class="pager flex">
                    <div class="pager__item pager__item--prev">
                        << {{.PreLink}} </div>

                            <div class="pager__item pager__item--next">
                                {{.NextLink}} >>
                            </div>


                </nav>


                </div>

                <!-- aside template removed-->

            </div>

            {{ template "footer" .}}

        </div>

</body>

</html>
