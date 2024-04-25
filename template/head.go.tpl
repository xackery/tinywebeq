{{ define "head" }}
<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
<meta name="robots" content="index, follow">
<title>{{ .Site.Title }}</title>

<meta name="description" content="thj-wiki.web.app">
<meta name="author" content="Xackery">
<link rel="canonical" href="{{ .Site.BaseURL }}">
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.15.0/cdn/themes/dark.css" />
<link rel="icon" href="{{ .Site.BaseURL }}/favicon.ico">
<noscript>
    <style>
        #theme-toggle,
        .top-link {
            display: none;
        }

    </style>
</noscript>
{{ if .Site.GoogleTag }}
<script async src="https://www.googletagmanager.com/gtag/js?id={{ .Site.GoogleTag }}"></script>
<script>
var doNotTrack = false;
if (!doNotTrack) {
	window.dataLayer = window.dataLayer || [];
	function gtag(){dataLayer.push(arguments);}
	gtag('js', new Date());
	gtag('config', '{{ .Site.GoogleTag }}', { 'anonymize_ip': false });
}
</script>
{{ end }}
<meta property="og:title" content="{{ .Site.Title }}" />
<meta property="og:description" content="{{ .Site.Description }}" />
<meta property="og:type" content="website" />
<meta property="og:url" content="{{ .Site.BaseURL }}" />
{{ if .Site.ImageURL }}
    <meta property="og:image" content="{{ .Site.BaseURL }}/{{ .Site.ImageURL }}"/>
    <meta name="twitter:card" content="summary_large_image"/>
    <meta name="twitter:image" content="{{ .Site.BaseURL }}{{ .Site.ImageURL }}"/>
    <meta name="twitter:title" content="{{ .Site.Title }}"/>
    <meta name="twitter:description" content="{{ .Site.Description }}"/>
{{ end }}
<script type="application/ld+json">
{
  "@context": "https://schema.org",
  "@type": "Organization",
  "name": "{{ .Site.Title }}",
  "url": "{{ .Site.BaseURL }}",
  "description": "{{ .Site.Description }}",
  "thumbnailUrl": "{{ .Site.BaseURL }}/favicon.ico",  
}
</script>
{{end}}