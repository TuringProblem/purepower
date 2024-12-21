package main

import (
    "log"
    "net/http"
    "github.com/TuringProblem/purepower/internal/handler"
)
func main() {
    //init handlers
    http.HandleFunc("/", handler.HomeHandler)
    http.HandleFunc("/athlete", handler.ProductHandler)
    http.HandleFunc("/explore", handler.ContentHandler)

    // Static files (i.e., CSS -> and JavaScript :/)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))


    // starting the server
    log.Println("Starter your server on port :42069")
    if err := http.ListenAndServe(":42069", nil); err != nil {
        log.Fatal(err)
    }
}
