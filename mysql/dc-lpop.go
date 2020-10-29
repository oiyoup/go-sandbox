package main

//http://golang.site/go/article/107-MySql-%EC%82%AC%EC%9A%A9---%EC%BF%BC%EB%A6%AC

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

//"https://github.com/go-sql-driver/mysql"

func main() {
    db, err := sql.Open("mysql", "lpop:lpop5913kc@tcp(127.0.0.1:3307)/dc_lpop")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    //singRow(*db)
    muitlRows(*db)
}

func singRow(db sql.DB) {
    var name string
    err := db.QueryRow("Select contact_name from tmerchant where merchant_id = 'clickbuy'").Scan(&name)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(name)
}

func muitlRows(db sql.DB) {
    var id int
    var name string
    rows, err := db.Query("Select id, contact_name from tmerchant where id >= ? and id <= 10", 1)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    
    for rows.Next() {
        err := rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(id, name)
    }
}
