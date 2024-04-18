package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

var db *sql.DB

type Mitglied struct {
	Id       int
	Name     string
	Vname    string
	Nickname string
}

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl := template.Must(template.ParseFiles("html/home.html"))
	tmpl.Execute(w, nil)
}

func mitglieder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	mitglieder := get_mitglieder()
	tmpl := template.Must(template.ParseFiles("html/mitglieder.html"))
	tmpl.Execute(w, mitglieder)
}

func strafen(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Strafen")
}

func main() {
	router := httprouter.New()
	router.GET("/", home)
	router.GET("/mitglieder", mitglieder)
	router.GET("/strafen", strafen)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func get_mitglieder() []Mitglied {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "testo",
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query the database.
	rows, err := db.Query("SELECT * FROM mitglieder")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	mg := []Mitglied{}
	for rows.Next() {
		var id int
		var name string
		var vname string
		var nickname string
		err := rows.Scan(&id, &name, &vname, &nickname)
		if err != nil {
			log.Fatal(err)
		}
		mg = append(mg, Mitglied{id, name, vname, nickname})
	}
	return mg
}
