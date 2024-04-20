package main

import (
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/QuintenBruynseraede/time2go/internal"
	"github.com/QuintenBruynseraede/time2go/internal/cities"
	"github.com/QuintenBruynseraede/time2go/internal/trie"
)

type application struct {
	logger      *slog.Logger
	templates   *template.Template
	trie        *trie.Trie
	coordinates map[string]cities.City
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network address to serve")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	templates, err := internal.LoadTemplates()
	if err != nil {
		logger.Error("Unable to parse templates!")
		os.Exit(1)
	}

	cities, coordinates, err := cities.LoadCitiesFromFile(logger)
	if err != nil {
		logger.Error("Error reading cities", "error", err)
	}

	// Load all cities into trie
	trie := trie.NewTrie()
	for _, city := range cities {
		trie.Insert(city.Name)
	}

	app := &application{logger: logger, templates: templates, trie: trie, coordinates: coordinates}
	mux := app.routes()

	logger.Info("Starting server", "addr", *addr)
	err = http.ListenAndServe(*addr, mux)
	if err != nil {
		logger.Error("Exception when running server", "error", err)
		os.Exit(1)
	}
}
