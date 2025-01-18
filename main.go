package main

import (
    "log"
    "net/http"
    "github.com/PraveenUpadhyay/techtalk_logging/handlers"
)

func main() {
    http.HandleFunc("/user", handlers.CreateUserHandler)
    log.Println("Starting server on :8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
