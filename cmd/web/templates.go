package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/AliiAhmadi/ShareCode/pkg/models"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache.
	cache := map[string]*template.Template{}

	// Use filepath.Glob function to get a slice of all filepaths with the
	// extension '.page.tmpl'. This essentially gives us a slice of all the
	// 'page' templates for the application.
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))

	if err != nil {
		return nil, err
	}

	// Loop through the pages.
	for _, page := range pages {
		// Extract the file name.
		name := filepath.Base(page)

		// Parse the page template file.
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return nil, err
		}

		// Use the ParseGlob method to add any 'layout' templates.
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// Use the ParseGlob method to add any 'partial' templates.
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		// Add the template set to the cache.
		cache[name] = ts
	}

	// Return the map :)
	return cache, nil
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}
