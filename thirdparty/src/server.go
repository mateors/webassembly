package main

import (
    "log"
    "net/http"
)

func main() {
    log.Println("Server listening on port 3000")
    http.ListenAndServe(":3000", http.FileServer(http.Dir("public/")))
}
