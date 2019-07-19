package main

import (
    // Core libraries
    "fmt"
    "encoding/csv"
    "os"
    "io"
    "log"
    // Proprietary libraries
    dt "neoway_db_conn"
    dp "neoway_data_process"
)

// This type reflect the data type readed from the csv
type row struct {
    comp     string
    zip      string
}

var (
    err      error
    line     row
    row_list []string
    rows_in  int64
)

func main() {
    // Loading the data
    csvfile, errn := os.Open("../data/q1_catalog.csv")
    if errn != nil { log.Fatalln("Could not open csv file", err) }

    csvData := csv.NewReader(csvfile)
    //// This line is used just to skip the column names
    _, _ = csvData.Read()

    // Opening the Database Connection
    db := dt.OpenConn()
    defer dt.CloseConn(db)

    // This is the central loop where the business rules are applied
    for {
        // Reading the csv data role by role
        row_list, err = csvData.Read()
        if err == io.EOF { break }
        if err != nil { log.Fatal(err) }

        line.comp, line.zip, _ = dp.Normalize_data(row_list[0], false)
        // This function inserts the data on the Database
        rows_in += dt.CSVLoad(line.comp, line.zip, db)
    }
    fmt.Println("Total of rows inserted on the database:", rows_in)
}
