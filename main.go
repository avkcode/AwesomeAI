package main

import (
    "embed"
    "fmt"
    "io/fs"
    "net/http"
)

// Embed all files into the binary
//go:embed LICENSE README.md favicon.svg index.html models.json
var content embed.FS

func main() {
    // Create a sub-filesystem containing only the embedded files
    staticFS, err := fs.Sub(content, ".")
    if err != nil {
        panic(err)
    }

    // Serve the static files using http.FileServer
    http.Handle("/", http.FileServer(http.FS(staticFS)))

    // Start the HTTP server
    fmt.Println("Starting server on :8080...")
    err = http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Printf("Failed to start server: %v\n", err)
    }
}
