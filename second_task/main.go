package main

import (
    "fmt"
    "encoding/csv"
    "os"
    "io"
    "log"
    "strings"
    "net/http"
    "io/ioutil"
)

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
    csvfile, err := os.Open("../data/q2_clientData.csv")
    if err != nil { log.Fatalln("Could not open csv file", err) }

    csvData := csv.NewReader(csvfile)
    _, _ = csvData.Read()

    for {
        row_list, err = csvData.Read()
        if err == io.EOF { break }
        if err != nil { log.Fatal(err) }

        row_list = strings.Split(row_list[0], ";")

        line.comp = strings.Replace(row_list[0], "'", "", -1)
        line.comp = strings.Replace(line.comp, " ", "_", -1)
        line.comp = strings.Replace(line.comp, "&", "", -1)
        line.comp = strings.ToUpper(line.comp)
        line.zip = row_list[1]
        line.website = row_list[2]

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
