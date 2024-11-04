package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/weldsh2535/students-api/internal/config"
)

func main() {
    // Load config
    cfg := config.MustLoad()

    // Setup router
    router := http.NewServeMux()

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("welcome to student api"))
    })

    // Setup server
    server := http.Server{
        Addr:    cfg.HTTPServer.Addr,  // Corrected to cfg.HTTPServer.Addr
        Handler: router,
    }

    fmt.Printf("Server started on %s\n", cfg.HTTPServer.Addr)
    err := server.ListenAndServe()
    if err != nil {
        log.Fatal("failed to start server:", err)
    }
}
