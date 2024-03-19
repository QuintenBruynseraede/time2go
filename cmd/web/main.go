package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"text/template"
)

type application struct {
	logger    *slog.Logger
	templates *template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network address to serve")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	files := []string{
		"./ui/html/form.tpl",
		"./ui/html/head.tpl",
		"./ui/html/response.tpl",
		"./ui/html/spacer.tpl",
		"./ui/html/footer.tpl",
		"./ui/html/pages/index.tpl",
	}
	templates, err := template.ParseFiles(files...)
	if err != nil {
		logger.Error("Unable to parse templates!")
		os.Exit(1)
	}

	app := &application{logger: logger, templates: templates}
	mux := app.routes()

	logger.Info("Starting server", "addr", *addr)
	err = http.ListenAndServe(*addr, mux)
	if err != nil {
		logger.Error("Exception when running server", "error", err)
		os.Exit(1)
	}
}
