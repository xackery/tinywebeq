{{ define "head" }}
<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
<meta name="robots" content="index, follow">
<title>{{ .Site.Title }}</title>

<meta name="description" content="thj-wiki.web.app">
<meta name="author" content="Xackery">
<link rel="canonical" href="https://thj-wiki.web.app/">
<link rel="icon" href="https://thj-wiki.web.app/favicon.ico">
<link rel="icon" type="image/png" sizes="16x16" href="https://thj-wiki.web.app/favicon-16x16.png">
<link rel="icon" type="image/png" sizes="32x32" href="https://thj-wiki.web.app/favicon-32x32.png">
<link rel="apple-touch-icon" href="https://thj-wiki.web.app/apple-touch-icon.png">
<link rel="mask-icon" href="https://thj-wiki.web.app/safari-pinned-tab.svg">
<meta name="theme-color" content="#2e2e33">
<meta name="msapplication-TileColor" content="#2e2e33">
<link rel="alternate" type="application/rss+xml" href="https://thj-wiki.web.app/index.xml">
<link rel="alternate" type="application/json" href="https://thj-wiki.web.app/index.json">
<noscript>
    <style>
        #theme-toggle,
        .top-link {
            display: none;
        }

    </style>
</noscript>
<script async src="https://www.googletagmanager.com/gtag/js?id=G-Y9XZCMEGZ7"></script>
<script>
var doNotTrack = false;
if (!doNotTrack) {
	window.dataLayer = window.dataLayer || [];
	function gtag(){dataLayer.push(arguments);}
	gtag('js', new Date());
	gtag('config', 'G-Y9XZCMEGZ7', { 'anonymize_ip': false });
}
</script>
<meta property="og:title" content="Xackery" />
<meta property="og:description" content="thj-wiki.web.app" />
<meta property="og:type" content="website" />
<meta property="og:url" content="https://thj-wiki.web.app/" /><meta property="og:image" content="https://thj-wiki.web.app/papermod-cover.png"/>

<meta name="twitter:card" content="summary_large_image"/>
<meta name="twitter:image" content="https://thj-wiki.web.app/papermod-cover.png"/>

<meta name="twitter:title" content="Xackery"/>
<meta name="twitter:description" content="thj-wiki.web.app"/>

<script type="application/ld+json">
{
  "@context": "https://schema.org",
  "@type": "Organization",
  "name": "THJ-Wiki",
  "url": "https://thj-wiki.web.app/",
  "description": "thj-wiki.web.app",
  "thumbnailUrl": "https://thj-wiki.web.app/favicon.ico",
  "sameAs": [
      "https://github.com/xackery/thj-wiki", "https://ko-fi.com/xackery", "index.xml"
  ]
}
</script>
{{end}}