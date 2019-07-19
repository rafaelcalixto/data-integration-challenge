package data_process

import (
    "fmt"
    "strings"
    "strconv"
)

type row struct {
    comp     string
    zip_int  int64
    zip      string
    website  string
}

var (
    row_list []string
    line     row
    err      error
)

func Normalize_data(row string, api bool) (string, string, string) {
    row_list = strings.Split(row, ";")

    // Here the data is normalized for the company name and zip number
    line.comp = strings.Replace(row_list[0], "'", "", -1)
    line.comp = strings.Replace(line.comp, "&", "", -1)
    if api { line.comp = strings.Replace(line.comp, " ", "_", -1) }
    line.comp = strings.ToUpper(line.comp)

    line.zip_int, err = strconv.ParseInt(row_list[1], 10, 32)
    if err != nil {
        line.zip = ""
    } else {
        line.zip = fmt.Sprintf("%05d", line.zip_int)
    }
    if api { line.website = strings.ToLower(row_list[2]) }
    return line.comp, line.zip, line.website
}
