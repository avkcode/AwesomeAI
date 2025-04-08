package main

import (
    "embed"
    "fmt"
    "io/fs"
    "net/http"
    "os"
    "strings"
    "path/filepath"
)

//go:embed LICENSE README.md favicon.svg index.html models.json
var content embed.FS

func main() {
    modelDataURL := os.Getenv("MODEL_DATA_URL")
    if modelDataURL == "" {
        modelDataURL = "https://raw.githubusercontent.com/avkcode/awesomeai/refs/heads/main/models.json"
    } else if strings.HasPrefix(modelDataURL, "./") {
        // Handle local file paths - convert to absolute path
        absPath, err := filepath.Abs(modelDataURL)
        if err != nil {
            fmt.Printf("Error resolving local path: %v\n", err)
            absPath = modelDataURL
        }
        modelDataURL = absPath
    }

    staticFS, err := fs.Sub(content, ".")
    if err != nil {
        panic(err)
    }

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        indexHTML, err := fs.ReadFile(content, "index.html")
        if err != nil {
            http.Error(w, "Failed to read index.html", http.StatusInternalServerError)
            return
        }

        modifiedHTML := strings.Replace(
            string(indexHTML),
            "window.MODEL_DATA_URL = window.MODEL_DATA_URL ||",
            fmt.Sprintf("window.MODEL_DATA_URL = '%s' ||", modelDataURL),
            1,
        )

        w.Header().Set("Content-Type", "text/html")
        w.Write([]byte(modifiedHTML))
    })

    fmt.Printf("Starting server on :8080 with MODEL_DATA_URL=%s...\n", modelDataURL)
    err = http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Printf("Failed to start server: %v\n", err)
    }
}