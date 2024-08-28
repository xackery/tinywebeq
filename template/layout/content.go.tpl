<!DOCTYPE html>
<html data-layout="default" data-shoelace-version="2.15.0" class="flavor-html sl-theme-dark" lang="en">
<head>
    {{template "head" . }}
</head>

<body class="list dark" id="top">
    {{template "header" .}}
    {{template "sidebar" .}}
    <main class="main">
        {{template "data" .}}
    </main>
    {{template "footer" .}}
</body>

</html>
