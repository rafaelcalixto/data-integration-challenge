package main

import (
    "fmt"
    "encoding/csv"
    "os"
    "io"
    "log"
    "strings"
    "strconv"
    // "crawler"
    // "strconv"
    dt "database_tools"
    // df "data_formater"
)

// var (
//     text      string
//     var_json  map[string]interface{}
//     cols      []string
//     col_len   int
//     count_1   int
//     count_2   int
//     f_line    bool = true
//     value     string
//     brute_row map[string]interface{}
// )

var (
    err      error
    row      []string
    row_list []string
    company  string
    zip_int  int64
    zip      string
)

func main() {
    csvfile, errn := os.Open("data/q1_catalog.csv")
    if errn != nil { log.Fatalln("Could not open csv file", err) }

    csv_data := csv.NewReader(csvfile)
    _, _ = csv_data.Read()
    for {
        row, err = csv_data.Read()
        if err == io.EOF { break }
        if err != nil { log.Fatal(err) }

        row_list = strings.Split(row[0], ";")

        company = strings.ToUpper(row_list[0])

        zip_int, err = strconv.ParseInt(row_list[1], 10, 32)
        if err != nil {
            zip = ""
        } else {
            zip = fmt.Sprintf("%05d", zip_int)
        }

        fmt.Println(company)
        fmt.Println(zip)
    }
    dt.DB_register("select * from companies", true)
    // text = crawler.Get("https://api.kraken.com/0/public/Assets")
    // b_text := []byte(text)
    //
    // result := df.JSONFormater(b_text)
    //
    // tps := []string{" varchar(10) not null, ",
    //                 " char(5) constraint fk primary key, ",
    //                 " integer not null, ",
    //                 " integer not null"}
    // ord := [4]int{1, 0, 2, 3}
    // rows := make([][]string, len(result))
    // for _, v1 := range result {
    //     brute_row = v1.(map[string]interface{})
    //     rows[count_1] = make([]string, len(brute_row))
    //     for k2, v2 := range brute_row {
    //         if f_line { cols = append(cols, k2) }
    //         final_str, err := v2.(string)
    //         if !err { final_str = strconv.FormatFloat(v2.(float64),'f', 0, 64) }
    //         rows[count_1][count_2 % 4] = final_str
    //         count_2 += 1
    //     }
    //     count_1 += 1
    //     f_line = false
    // }
    // col_len = len(cols)
    //
    // query_create := dt.FormatQuery("create", "kraken_test", cols, tps, rows, ord)
    // query_insert := dt.FormatQuery("insert", "kraken_test", cols, tps, rows, ord)
    //
    // tps = []string{"varchar", "char", "integer", "integer"}
    // query_select := dt.FormatQuery("select", "kraken_test", cols, tps, rows, ord)
    //
    // fmt.Println(query_create)
    // fmt.Println(query_insert)
    // fmt.Println(query_select)

    /*postgres_con.DB_register(query_create, false)
    postgres_con.DB_register(query_insert, false)
    postgres_con.DB_register(query_select, true)*/
}
