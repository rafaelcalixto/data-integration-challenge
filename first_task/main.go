package main

import (
    "fmt"
    "encoding/csv"
    "os"
    "io"
    "log"
    "strings"
    "strconv"
    dt "neoway_db_conn"
)

type row struct {
    comp  string
    zip_int  int64
    zip      string
}

var (
    err      error
    line     row
    row_list []string
)

func main() {
    csvfile, errn := os.Open("../data/q1_catalog.csv")
    if errn != nil { log.Fatalln("Could not open csv file", err) }

    csvData := csv.NewReader(csvfile)
    _, _ = csvData.Read()

    db := dt.OpenConn()
    defer dt.CloseConn(db)

    for {
        row_list, err = csvData.Read()
        if err == io.EOF { break }
        if err != nil { log.Fatal(err) }

        row_list = strings.Split(row_list[0], ";")

        line.comp = strings.Replace(row_list[0], "'", "", -1)
        line.comp = strings.Replace(line.comp, "&", "", -1)
        line.comp = strings.ToUpper(line.comp)

        line.zip_int, err = strconv.ParseInt(row_list[1], 10, 32)
        if err != nil {
            line.zip = ""
        } else {
            line.zip = fmt.Sprintf("%05d", line.zip_int)
        }
        dt.CSVLoad(line.comp, line.zip, db)
    }
}
