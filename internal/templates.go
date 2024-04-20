package internal

import (
	"html/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func LoadTemplates() (*template.Template, error) {
	files := []string{
		"./ui/html/form.tpl",
		"./ui/html/head.tpl",
		"./ui/html/response.tpl",
		"./ui/html/search_results.tpl",
		"./ui/html/spacer.tpl",
		"./ui/html/footer.tpl",
		"./ui/html/pages/index.tpl",
	}
	funcMap := template.FuncMap{
		"title": title,
	}

	return template.New("all").Funcs(funcMap).ParseFiles(files...)
}

func title(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}
