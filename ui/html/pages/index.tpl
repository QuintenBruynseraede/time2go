{{ define "index" }}
<!doctype html>
<html lang="en">
{{template "head" .}}
<body class="bg-sky-50">
    {{template "form" .}}
    {{template "spacer" .}}
    <div id="response"></div>
    {{template "footer" .}}
  </body>
</html>
{{end}}
