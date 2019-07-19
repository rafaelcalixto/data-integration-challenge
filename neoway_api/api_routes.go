package neoway_api

import (
    // Core libraries
    "fmt"
    "net/http"
    "strings"
    "strconv"
    // Proprietary libraries
    dt "neoway_db_conn"
)

// This type reflect the data type readed from the csv
type row struct {
    comp  string
    zip_int  int64
    zip      string
    website  string
}

var (
    line  row
    err   error
    work  int64
    msg   string
    query_return map[string]string
)

// This function returns for the Browsers some informations about the API
// This is mandary for some Browsers allows the access to the API
func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// This function returns a "Welcome message" to the API
func Index_handler(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    fmt.Fprintf(w, "This is the API for the Neoway test")
}

// This function take the parameters passed on the API and update the website on
// the database
func AssociateLink(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    //// Here the parameters are taken from the URL
    line.comp = r.URL.Query()["c"][0]
    line.zip = r.URL.Query()["z"][0]
    line.website = r.URL.Query()["w"][0]

    // Opening the Database Connection
    db := dt.OpenConn()
    defer dt.CloseConn(db)

    //// This line is necessary to revert the replace of the space for underline
    line.comp = strings.Replace(line.comp, "_", " ", -1)


    // This block returns a message if the database was successfull updated
    work = dt.APILoad(line.comp, line.zip, line.website, db)
    if work == 1 {
        msg = fmt.Sprintf( "The page %s was updated", line.website)
    } else {
        msg = fmt.Sprintf( "Parameter not found")
    }
    fmt.Fprintf(w, msg)
}

// This function take the parameters passed on the API and search for matchs on
// the database
func ConsultCompanies(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    //// Here the parameters are taken from the URL
    line.comp = r.URL.Query()["c"][0]
    line.zip = r.URL.Query()["z"][0]

    line.comp = strings.Replace(line.comp, "'", "", -1)
    line.comp = strings.Replace(line.comp, "_", " ", -1)
    line.zip_int, err = strconv.ParseInt(line.zip, 10, 32)
    if err != nil {
        line.zip = ""
    } else {
        line.zip = fmt.Sprintf("%05d", line.zip_int)
    }

    // Opening the Database Connection
    db := dt.OpenConn()
    defer dt.CloseConn(db)

    // The below function returns the datafrom the database (if matched) and
    // the below if steatment builds a string to be printed on the API
    query_return = dt.APIQuery(line.comp, line.zip, db)
    if _, ok := query_return["msg"]; ok {
        msg = fmt.Sprintf("{\n\t\"msg\": \"" + query_return["msg"] + "\"\n}")
    } else {
        msg = fmt.Sprintf( "{\n\t\"id\": \"" + query_return["id"] + "\",\n" +
                           "\t\"name\": \"" + query_return["name"] + "\",\n" +
                           "\t\"zip\": \"" + query_return["zip"] + "\",\n" +
                           "\t\"website\": \"" + query_return["website"] +
                           "\"\n}" )
    }
    fmt.Fprintf(w, msg)
}
