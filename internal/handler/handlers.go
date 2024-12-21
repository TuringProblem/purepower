package handler

import (
    "html/template"
    "net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("web/templates/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}


func ProductHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("web/templates/product.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

func ContentHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<p>Ne wproduct details loaded via HTMX! :) </p>"))
}
