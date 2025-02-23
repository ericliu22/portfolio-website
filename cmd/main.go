package main

import (
    "html/template"
    "log"
    "net/http"
    "path/filepath"
)

// Global template cache (simple approach)
var tmpl *template.Template

func main() {
    // Parse all templates on start
    tmpl = parseTemplates()

    // Serve static files (CSS, JS, images) from the "static" directory
    fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static", fileServer))

    // Handle root route
    http.HandleFunc("/", indexHandler)

    // Start server
    log.Println("Listening on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}

// indexHandler is an example page handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
    // Render the index template (assumes "index" is defined in templates/pages/index.gohtml)
    // You could also serve different pages by splitting request paths, etc.
    err := tmpl.ExecuteTemplate(w, "index", nil)
    if err != nil {
        log.Println(err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

// parseTemplates loads all .gohtml files in the "templates" directory
func parseTemplates() *template.Template {
    pattern := filepath.Join("templates", "**", "*.gohtml")
    t, err := template.ParseGlob(pattern)
    if err != nil {
        log.Fatalf("Error parsing templates: %v", err)
    }
    return t
}

