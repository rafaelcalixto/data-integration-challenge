package main

import (
    // Core libraries
    "fmt"
    "encoding/csv"
    "os"
    "io"
    "log"
    "net/http"
    "io/ioutil"
    // Proprietary libraries
    dp "neoway_data_process"
)

// This type reflect the data type readed from the csv
type row struct {
    comp  string
    zip_int  int64
    zip      string
    website  string
}

var (
    err      error
    line     row
    row_list []string
    url      string
    bytes    []byte
)

func main() {
    // Loading the data
    csvfile, err := os.Open("../data/q2_clientData.csv")
    if err != nil { log.Fatalln("Could not open csv file", err) }

    csvData := csv.NewReader(csvfile)
    //// This line is used just to skip the column names
    _, _ = csvData.Read()

    // This is the central loop where the business rules are applied
    for {
        // Reading the csv data role by role
        row_list, err = csvData.Read()
        if err == io.EOF { break }
        if err != nil { log.Fatal(err) }

        line.comp, line.zip, line.website = dp.Normalize_data(row_list[0], true)

        // The data is send to the below API to be inserted on the Database
        url = fmt.Sprintf("http://localhost:8000/api/clientdata?c=%s&z=%s&w=%s",
                          line.comp, line.zip, line.website)
        ans, err := http.Get(url)
        if err != nil { log.Fatalln("Error while calling API", err) }
        bytes, err := ioutil.ReadAll(ans.Body)
        if err != nil { log.Fatalln("Error while read API return", err) }

        fmt.Println(string(bytes))
        ans.Body.Close()
    }
}
