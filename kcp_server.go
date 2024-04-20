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

type Veranstaltung struct {
	Id          int
	Bezeichnung string
	Datum       string
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
	strafen := get_veranstaltungen()
	tmpl := template.Must(template.ParseFiles("html/strafen.html"))
	tmpl.Execute(w, strafen)
}

func strafenzeitraum(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	von_datum := r.PostFormValue("von_datum")
	bis_datum := r.PostFormValue("bis_datum")
	fmt.Fprintf(w, "Von: %s, Bis: %s", von_datum, bis_datum)
}

func main() {
	router := httprouter.New()
	router.GET("/", home)
	router.GET("/mitglieder", mitglieder)
	router.GET("/strafen", strafen)
	router.POST("/strafenzeitraum", strafenzeitraum)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func connect_to_db() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "kcp",
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
}

func get_mitglieder() []Mitglied {
	connect_to_db()

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
	db.Close()
	return mg
}
func get_veranstaltungen() []Veranstaltung {
	connect_to_db()
	rows, err := db.Query("SELECT * FROM veranstaltungen")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	vers := []Veranstaltung{}
	for rows.Next() {
		var id int
		var bezeichnung string
		var datum string
		err := rows.Scan(&id, &bezeichnung, &datum)
		if err != nil {
			log.Fatal(err)
		}
		vers = append(vers, Veranstaltung{id, bezeichnung, datum})
	}
	db.Close()
	return vers
}
