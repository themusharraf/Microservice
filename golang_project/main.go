package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:1@tcp(mysql:3306)/mysql")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM users;")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Iterate over the rows returned by the query and write them to the response
		for rows.Next() {
			var id int
			var name string
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintf(w, "ID: %d, Name: %s\n", id, name)
		}
	})

    http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")
        query := `
        CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL
        );
        `
        _, err = db.Exec(query)
        if err != nil {
            fmt.Println(err)
            return
        }

        stmt, err := db.Prepare("INSERT INTO users (name) VALUES (?)")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer stmt.Close()

        _, err = stmt.Exec(name)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        fmt.Fprintln(w, "Data saved successfully")
    })

    log.Println("Server listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
