package main

import (
    "fmt"
    "os"         // Added for os.Args
    "net/http"   // Added for http.ResponseWriter and http.Request
    "primix/cmd"
    "primix/logger"
    "primix/server"
    "primix/utils"
)

func main() {
    logger.Init()
    utils.PrintBanner()

    // If command-line args are provided, run CLI
    if len(os.Args) > 1 {
        fmt.Println("🚀 Running Primix CLI...")
        cmd.Execute()
        return
    }

    // Otherwise, start the server
    s := server.NewServer()

    // Add routes
    s.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to Primix!"))
    })

    s.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, Primix World!"))
    })

    // Serve static files
    s.Static("/static", "./static")

    // Start server
    if err := s.Start(":8080"); err != nil {
        logger.Logger.Fatalf("Failed to start server: %v", err)
    }
}