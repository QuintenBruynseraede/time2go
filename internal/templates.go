package internal

import (
	"html/template"

	"github.com/QuintenBruynseraede/time2go/internal/utils"
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
		"title": utils.Title,
	}

	return template.New("all").Funcs(funcMap).ParseFiles(files...)
}
