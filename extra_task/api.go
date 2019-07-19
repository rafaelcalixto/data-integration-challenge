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
    fmt.Println("starting API...")
    http.HandleFunc("/", api.Index_handler)
    http.HandleFunc("/api/companies", api.ConsultCompanies)
    http.ListenAndServe(":8080", nil)
}
