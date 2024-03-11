package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

func count(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%d", 5)
}

func main() {

	cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: os.Getenv("DBNAME"),
    }

	// Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

	http.HandleFunc("/count", count)
	http.ListenAndServe(":8080", nil)
}


// getCount
func getCount() (int, error) {
	count := 0
    row := db.QueryRow("SELECT COUNT(*) FROM scan")
    if err := row.Scan(&count); err != nil {
        return 0, fmt.Errorf("error scanning count",err)
    }
    return count, nil
}