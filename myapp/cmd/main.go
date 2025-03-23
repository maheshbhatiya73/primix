package main

import (
    "net/http"   
    "primix/logger"
    "primix/server"
    "primix/utils"
)

func main() {
    logger.Init()
    utils.PrintBanner()

    s := server.NewServer()
    s.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, Primix!"))
    })
    s.Start(":8080")
}
