package main

import (
    "log"
    "net/http"

    "url_processor/internal/service"
    "url_processor/internal/handler"
    "url_processor/internal/config"

)

func main() {
    config.Load()

    port := config.Get(config.KeyPort)
    if port == ""{
        port = "8080"
    }

    urlService := service.NewURLService()
    urlHandler := handler.NewURLHandler(urlService)

    http.HandleFunc("/process-url", urlHandler.ProcessURL)

    log.Println("Server starting on %s", port)
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
