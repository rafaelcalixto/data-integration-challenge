package main

import (
    "fmt"
    "net/http"
    api "neoway_api"
)

func main() {
    fmt.Println("starting GenericAPI...")
    http.HandleFunc("/", api.Index_handler)
    http.HandleFunc("/api/clientdata", api.AssociateLink)
    http.ListenAndServe(":8000", nil)
}
