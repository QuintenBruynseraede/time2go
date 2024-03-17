{{ define "index" }}
<!doctype html>
<html lang="en">
{{template "head" .}}
<body class="bg-sky-50">
    {{template "form" .}}
    {{template "spacer" .}}
    {{template "response" .}}
    {{template "footer" .}}
  </body>
</html>
{{end}}
