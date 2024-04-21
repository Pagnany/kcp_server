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

type Strafe struct {
	Id               int
	Id_strafe_typ    int
	Id_mitglied      int
	Preis            float32
	Datum            string
	Anzahl           float32
	Id_veranstaltung int
}

type Strafe_typ struct {
	Id          int
	Bezeichnung string
	Preis       float32
	Aktiv       bool
}

type Page_date_strafen struct {
	Veranstaltungen []Veranstaltung
	Strafen         []Strafe
}

type Page_data_mitglieder struct {
	Mitglieder      []Mitglied
	Strafen_typen   []Strafe_typ
	Veranstaltungen []Veranstaltung
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
	data := Page_date_strafen{get_veranstaltungen(), []Strafe{}}
	tmpl := template.Must(template.ParseFiles("html/strafen.html"))
	tmpl.Execute(w, data)
}

func strafenzeitraum(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	von_datum := r.PostFormValue("von_datum")
	bis_datum := r.PostFormValue("bis_datum")
	data := Page_date_strafen{get_veranstaltungen(), get_strafen(0, von_datum, bis_datum)}
	for _, strafe := range data.Strafen {
		fmt.Fprintf(w, "ID: %d, Mitglied: %d, Preis: %f, Datum: %s, Anzahl: %f, Veranstaltung: %d<br>", strafe.Id, strafe.Id_mitglied, strafe.Preis, strafe.Datum, strafe.Anzahl, strafe.Id_veranstaltung)
	}
}

func strafen_erstellen_typ(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := get_strafen_typen()
	tmpl := template.Must(template.ParseFiles("html/strafen_erstellen_typ.html"))
	tmpl.Execute(w, data)
}

func strafen_erstellen_typ_post(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	bezeichnung := r.PostFormValue("bezeichnung")
	preis := r.PostFormValue("preis")

	connect_to_db()
	stmt, err := db.Prepare("INSERT INTO strafen_typ (bezeichnung, preis, aktiv) VALUES (?, ?, 1)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(bezeichnung, preis)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}

func strafen_erstellen_mitglied(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := Page_data_mitglieder{get_mitglieder(), get_strafen_typen(), get_veranstaltungen()}
	tmpl := template.Must(template.ParseFiles("html/strafen_erstellen_mitglied.html"))
	tmpl.Execute(w, data)
}
func strafen_erstellen_mitglied_post(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	id_mitglied := r.PostFormValue("mitglieder")
	id_strafe_typ := r.PostFormValue("strafe")
	preis := r.PostFormValue("preis")
	datum := r.PostFormValue("datum")
	anzahl := r.PostFormValue("anzahl")
	id_veranstaltung := r.PostFormValue("veranstaltungen")
	fmt.Fprintf(w, "ID: %s, Strafe: %s, Preis: %s, Datum: %s, Anzahl: %s, Veranstaltung: %s<br>", id_mitglied, id_strafe_typ, preis, datum, anzahl, id_veranstaltung)
	// connect_to_db()
	// stmt, err := db.Prepare("INSERT INTO strafen (id_strafe_typ, id_mitglied, preis, datum, anzahl, id_veranstaltung) VALUES (?, ?, ?, ?, ?, ?)")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = stmt.Exec(id_strafe_typ, id_mitglied, preis, datum, anzahl, id_veranstaltung)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// db.Close()
}

func main() {
	router := httprouter.New()
	router.GET("/", home)
	router.GET("/mitglieder", mitglieder)
	router.GET("/strafen", strafen)
	router.GET("/strafen/erstellen_typ", strafen_erstellen_typ)
	router.POST("/strafen/erstellen_typ", strafen_erstellen_typ_post)
	router.GET("/strafen/erstellen_mitglied", strafen_erstellen_mitglied)
	router.POST("/strafen/erstellen_mitglied", strafen_erstellen_mitglied_post)
	router.POST("/strafen/zeitraum", strafenzeitraum)

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

func get_strafen(mitgliedid int, von_datum string, bis_datum string) []Strafe {
	connect_to_db()
	strafen := []Strafe{}
	// wenn mitgliedid == 0 dann alle mitglieder

	var rows *sql.Rows
	var err error
	if mitgliedid == 0 {
		rows, err = db.Query("SELECT * FROM strafen WHERE datum BETWEEN ? AND ?", von_datum, bis_datum)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		rows, err = db.Query("SELECT * FROM strafen WHERE id_mitglied = ? AND datum BETWEEN ? AND ?", mitgliedid, von_datum, bis_datum)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var id_strafe_typ int
		var id_mitglied int
		var id_veranstaltung int
		var datum string
		var preis float32
		var anzahl float32
		err := rows.Scan(&id, &id_strafe_typ, &id_mitglied, &preis, &datum, &anzahl, &id_veranstaltung)
		if err != nil {
			log.Fatal(err)
		}
		strafen = append(strafen, Strafe{id, id_strafe_typ, id_mitglied, preis, datum, anzahl, id_veranstaltung})
	}

	db.Close()
	return strafen
}

func get_strafen_typen() []Strafe_typ {
	connect_to_db()
	strafen_typen := []Strafe_typ{}
	rows, err := db.Query("SELECT * FROM strafen_typ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var bezeichnung string
		var preis float32
		var aktiv bool
		err := rows.Scan(&id, &bezeichnung, &preis, &aktiv)
		if err != nil {
			log.Fatal(err)
		}
		strafen_typen = append(strafen_typen, Strafe_typ{id, bezeichnung, preis, aktiv})
	}
	db.Close()
	return strafen_typen
}
