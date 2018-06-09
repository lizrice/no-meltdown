package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	var conninfo string = "host=database sslmode=disable user=postgres password=insecurepassword dbname=postgres"
	time.Sleep(4 * time.Second)
	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		panic(err)
	}

	// create table if not there
	ct := `
		CREATE TABLE IF NOT EXISTS Students (
			name varchar(40) NOT NULL,
			age integer NOT NULL
		)
	`

	_, err = db.Exec(ct)
	if err != nil {
		panic(err)
	}

	// add some data
	ad := `
		INSERT INTO Students (name, age) VALUES ('Robert', '23');
		INSERT INTO Students (name, age) VALUES ('Mary', '24');
	`

	_, err = db.Exec(ad)
	if err != nil {
		panic(err)
	}

	log.Println("Ready to accept connections")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("request received")
		qname := r.URL.Query().Get("name")
		// allows SQL injection!
		query := fmt.Sprintf("SELECT name, age FROM Students WHERE name='%s'", qname)
		log.Println(query)
		rows, err := db.Query(query)
		// safe if you use the language SQL construction tools!
		// rows, err := db.Query("SELECT name, age FROM Students WHERE name=?", qname)
		if err != nil {
			fmt.Fprintf(w, "error: %v\n", err)
			return
		}
		defer rows.Close()
		if rows.Next() {
			var name string
			var age int
			err = rows.Scan(&name, &age)
			if err == nil {
				fmt.Fprintf(w, "Hello, %s you are %d years old\n", name, age)
			}
		} else {
			fmt.Fprintf(w, "Sorry I do not know how old you are\n")
		}
	})
	err = http.ListenAndServe(":80", nil)
	panic(err)
}
