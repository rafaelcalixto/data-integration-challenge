package neoway_db_conn

import (
    "database/sql"
    _ "github.com/lib/pq"
    "fmt"
    "log"
    "strconv"
)

const (
    host    string = "wt-postgres-test.c0fua9u7ixva.us-east-2.rds.amazonaws.com"
    port    string = "5432"
    user    string = "neoway"
    dbname  string = "wt_postgres_test "
    pw      string = "flavorstone"
    sslmode string = "disable"
)

type connect struct {
    conn  *sql.DB
}

type entity struct {
    id      int64
    name    string
    zip     string
    website string
}

var (
    db        connect
    cs        string
    test      string
    query     string
    list      []string
    feedback  sql.Result
    work      int64
    err       error
    r_row     entity
    comp_data map[string]string
    str_test  sql.NullString
)

func OpenConn() (connect) {
    cs = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
                     host, port, user, dbname, pw, sslmode)

    db.conn, err = sql.Open("postgres", cs)
    if err != nil { fmt.Println("Conldn't connect to the database", err) }

    return db
}

func CloseConn(db connect) {
    db.conn.Close()
}

func CSVLoad(company string, zip string, db connect) (int64) {
    query = fmt.Sprintf("insert into companies (name, zip) values ('%s', '%s')",
                        company, zip)
    feedback, err = db.conn.Exec(query)
    if err != nil { fmt.Println(err) }
    rows, err := feedback.RowsAffected()
    if err != nil { fmt.Println(err) }
    return rows
}

func APILoad(company string, zip string, website string, db connect) (int64) {
    query = fmt.Sprintf("update companies set website = '%s' " +
                "where name = '%s' and zip = '%s'", website, company, zip)

    feedback, err = db.conn.Exec(query)
    if err != nil { fmt.Println(err) }
    work, err = feedback.RowsAffected()
    if err != nil { fmt.Println(err) }
    return work
}

func APIQuery(company string, zip string, db connect) ( map[string]string ) {
    query = fmt.Sprintf("select id, name, zip, website from companies where " +
                        "name like upper('%s') and zip = '%s'", company, zip)
    err := db.conn.QueryRow(query).Scan(&r_row.id, &r_row.name,
                                        &r_row.zip, &str_test)
    comp_data = make(map[string]string)

    switch {
    case err == sql.ErrNoRows:
        comp_data["msg"] = "No match results"
    case err != nil:
        log.Fatalf("Error while querying: %v", err)
    default:
        comp_data["id"] = strconv.FormatInt(r_row.id, 10)
        comp_data["name"] = r_row.name
        comp_data["zip"] = r_row.zip
        if str_test.Valid {
            comp_data["website"] = str_test.String
        } else {
            comp_data["website"] = " "
        }
    }
    return comp_data
}
