package main

import (
    // Core libraries
    "fmt"
    "net/http"
    // Proprietary libraries
    api "neoway_api"
)

// This is a simple package used just to call the API
func main() {
    fmt.Println("starting the API...")
    http.HandleFunc("/", api.Index_handler)
    http.HandleFunc("/api/clientdata", api.AssociateLink)
    http.ListenAndServe(":8000", nil)
}
