package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	Anwesen  bool
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
	Bezeich          string
}

type Strafen_neu struct {
	Id                    int
	Id_strafe_typ         int
	Preis_strafe_typ      float32
	Bezeich_strafe_typ    string
	Id_mitglied           int
	Vname                 string
	Name                  string
	Nickname              string
	Datum_strafe          string
	Preis_strafe          float32
	Aktiv                 bool
	Anzahl                float32
	Bezeich_strafe        string
	Id_veranstaltung      int
	Datum_veranstaltung   string
	Bezeich_veranstaltung string
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

func main() {
	router := httprouter.New()
	router.GET("/", home)
	router.GET("/mitglieder", mitglieder)
	router.POST("/mitglieder/erstellen", mitglieder_erstellen)
	router.GET("/strafen", strafen)
	router.GET("/veranstaltungen", veranstaltungen)
	router.GET("/veranstaltungen/anwesenheit/:id", veranstaltungen_anwesenheit)
	router.POST("/veranstaltungen/zeitraum", veranstaltungen_zeitraum_post)
	router.GET("/strafen/erstellen_typ", strafen_erstellen_typ)
	router.POST("/strafen/erstellen_typ", strafen_erstellen_typ_post)
	router.GET("/strafen/erstellen_mitglied", strafen_erstellen_mitglied)
	router.POST("/strafen/erstellen_mitglied", strafen_erstellen_mitglied_post)
	router.POST("/strafen/zeitraum", strafenzeitraum)
	router.GET("/strafen/veranstaltung", strafen_veranstaltung)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func veranstaltungen_anwesenheit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	temp := params.ByName("id")
	fmt.Fprintf(w, temp)
}

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl := template.Must(template.ParseFiles("html/home.html"))
	tmpl.Execute(w, nil)
}

func veranstaltungen(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	veranstaltungen := get_veranstaltungen()
	tmpl := template.Must(template.ParseFiles("html/veranstaltungen.html"))
	tmpl.Execute(w, veranstaltungen)
}

func veranstaltungen_zeitraum_post(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	von_datum := r.PostFormValue("von_datum")
	bis_datum := r.PostFormValue("bis_datum")
	data := get_veranstaltungen_zeitraum(von_datum, bis_datum)
	for _, veranstaltung := range data {
		fmt.Fprintf(w, " %s | %s<br>", veranstaltung.Datum, veranstaltung.Bezeichnung)
	}
}

func strafen_veranstaltung(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	temp := create_strafen_grid(1)
	fmt.Fprintf(w, temp)
}

func mitglieder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	mitglieder := get_mitglieder()
	tmpl := template.Must(template.ParseFiles("html/mitglieder.html"))
	tmpl.Execute(w, mitglieder)
}

func mitglieder_erstellen(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	name := r.PostFormValue("name")
	vname := r.PostFormValue("vname")
	nickname := r.PostFormValue("nickname")

	connect_to_db()
	stmt, err := db.Prepare("INSERT INTO mitglieder (name, vname, nickname) VALUES (?, ?, ?)")
	if err != nil {
		log.Printf("Error: %s", err)
	}
	_, err = stmt.Exec(name, vname, nickname)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	db.Close()
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
	data := get_strafen(0, von_datum, bis_datum)
	for _, strafe := range data {
		fmt.Fprintf(w, " %s | %f %f<br>", strafe.Datum_veranstaltung, strafe.Preis_strafe, strafe.Preis_strafe_typ)
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
		log.Printf("Error: %s", err)
	}
	_, err = stmt.Exec(bezeichnung, preis)
	if err != nil {
		log.Printf("Error: %s", err)
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
	bezeich := r.PostFormValue("bezeich")
	id_veranstaltung := r.PostFormValue("veranstaltungen")

	//Abbrechen wenn keine Veranstaltung und kein Datum gesetzt ist
	if id_veranstaltung == "0" && datum == "" {
		fmt.Fprintf(w, "FEHLER: Keine Veranstaltung und kein Datum gesetzt<br>")
		return
	}
	//Abbrechen wenn keine Strafe und Preis	0 ist
	if id_strafe_typ == "0" && preis == "0" {
		fmt.Fprintf(w, "FEHLER: Keine Strafe und Preis 0 gesetzt<br>")
		return
	}
	//Abbrechen wenn keine Strafe keine Bezeichnung hat
	if id_strafe_typ == "0" && bezeich == "" {
		fmt.Fprintf(w, "FEHLER: Keine Strafe und Bezeichnung leer<br>")
		return
	}
	//Abbrechen wenn Anzahl 0 ist
	if anzahl == "0" {
		fmt.Fprintf(w, "FEHLER: Anzahl 0 gesetzt<br>")
		return
	}
	// Abbrechen wenn kein Mitglied gesetzt ist
	if id_mitglied == "0" {
		fmt.Fprintf(w, "FEHLER: Kein Mitglied gesetzt<br>")
		return
	}

	// Wenn eine Strafe gesetzt ist, dann brauchen wir keinen Preis und keine Bezeichnung
	temp, err := strconv.Atoi(id_strafe_typ)
	if err == nil {
		if temp > 0 {
			preis = "0"
			bezeich = ""
		}
	}

	fmt.Fprintf(w, "Mitglied_ID: %s, Strafe_ID: %s, Preis: %s, Datum: %s, Anzahl: %s, Veranstaltung_ID: %s<br>", id_mitglied, id_strafe_typ, preis, datum, anzahl, id_veranstaltung)

	connect_to_db()
	stmt, err := db.Prepare("INSERT INTO strafen (id_strafe_typ, id_mitglied, preis, datum, anzahl, id_veranstaltung, bezeich) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Printf("Error: %s", err)
	}
	_, err = stmt.Exec(id_strafe_typ, id_mitglied, preis, NewNullString(datum), anzahl, id_veranstaltung, bezeich)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	db.Close()
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func get_strafen_typ_fuer_veranstaltung(id_veranstaltung int) []Strafe_typ {
	connect_to_db()
	var strafen []Strafe_typ
	rows, err := db.Query("SELECT * FROM strafen_typ WHERE id IN (SELECT id_strafe_typ FROM strafen WHERE id_veranstaltung = ?)", id_veranstaltung)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var bezeichnung string
		var preis float32
		var aktiv bool
		err := rows.Scan(&id, &bezeichnung, &preis, &aktiv)
		if err != nil {
			log.Printf("Error: %s", err)
		}
		strafen = append(strafen, Strafe_typ{id, bezeichnung, preis, aktiv})
	}
	db.Close()
	return strafen
}

func get_mitglieder_fuer_veranstaltung(id_veranstaltung int) []Mitglied {
	connect_to_db()
	var mitglieder []Mitglied
	rows, err := db.Query("SELECT * FROM mitglieder WHERE id IN (SELECT id_mitglied FROM strafen WHERE id_veranstaltung = ?)", id_veranstaltung)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var vname string
		var nickname string
		err := rows.Scan(&id, &name, &vname, &nickname)
		if err != nil {
			log.Printf("Error: %s", err)
		}
		mitglieder = append(mitglieder, Mitglied{id, name, vname, nickname, false})
	}
	db.Close()
	return mitglieder
}

func get_freie_strafen_typen_fuer_veranstaltung(id_veranstaltung int) []string {
	connect_to_db()
	var strafen []string
	rows, err := db.Query("select distinct bezeich from strafen where id_strafe_typ = 0 and id_veranstaltung = ?", id_veranstaltung)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	defer rows.Close()
	for rows.Next() {
		var bezeichnung string
		err := rows.Scan(&bezeichnung)
		if err != nil {
			log.Printf("Error: %s", err)
		}
		strafen = append(strafen, bezeichnung)
	}
	db.Close()
	return strafen
}

func get_freie_strafen_fuer_veranstaltung(id_veranstaltung int) []Strafe {
	connect_to_db()
	strafen := []Strafe{}
	rows, err := db.Query("select id, id_mitglied, preis, bezeich, anzahl from strafen where id_strafe_typ = 0 and id_veranstaltung = ?", id_veranstaltung)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var id_mitglied int
		var preis float32
		var bezeich string
		var anzahl float32
		err := rows.Scan(&id, &id_mitglied, &preis, &bezeich, &anzahl)
		if err != nil {
			log.Printf("Error: %s", err)
		}
		strafen = append(strafen, Strafe{id, 0, id_mitglied, preis, "", anzahl, id_veranstaltung, bezeich})
	}

	db.Close()
	return strafen
}

func get_strafen_fuer_veranstaltung(id_veranstaltung int) []Strafe {
	connect_to_db()
	strafen := []Strafe{}
	rows, err := db.Query("SELECT * FROM strafen WHERE id_strafe_typ != 0 AND id_veranstaltung = ?", id_veranstaltung)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var id_strafe_typ int
		var id_mitglied int
		var id_veranstaltung int
		var datum sql.NullString
		var preis float32
		var anzahl float32
		var bezeich string
		err := rows.Scan(&id, &id_strafe_typ, &id_mitglied, &preis, &datum, &anzahl, &id_veranstaltung, &bezeich)
		if err != nil {
			log.Printf("Error: %s", err)
		}
		if datum.Valid {
			strafen = append(strafen, Strafe{id, id_strafe_typ, id_mitglied, preis, datum.String, anzahl, id_veranstaltung, bezeich})
		} else {
			strafen = append(strafen, Strafe{id, id_strafe_typ, id_mitglied, preis, "", anzahl, id_veranstaltung, bezeich})
		}
	}
	db.Close()
	return strafen
}

func create_strafen_grid(id_veranstaltung int) string {
	strafen := get_strafen_fuer_veranstaltung(id_veranstaltung)
	strafen_typen := get_strafen_typ_fuer_veranstaltung(id_veranstaltung)
	mitglieder := get_mitglieder_fuer_veranstaltung(id_veranstaltung)
	freie_strafen_typen := get_freie_strafen_typen_fuer_veranstaltung(id_veranstaltung)
	freie_strafen := get_freie_strafen_fuer_veranstaltung(id_veranstaltung)

	// Überschriften erstellen
	grid := "<table border='1'><tr><th>Mitglied</th>"
	for _, strafe := range strafen_typen {
		temp := strconv.FormatFloat(float64(strafe.Preis), 'f', 2, 32)
		grid += "<th> " + strafe.Bezeichnung + " " + temp + "€ </th>"
	}
	for _, strafe := range freie_strafen_typen {
		grid += "<th> " + strafe + " </th>"
	}
	grid += "<th>Summe</th>"
	grid += "</tr>"
	// Zeilen erstellen
	for _, mitglied := range mitglieder {
		summe := 0.0
		grid += "<tr><td>" + mitglied.Name + "</td>"
		for _, strafe_typ := range strafen_typen {
			found := false
			for _, strafe := range strafen {
				if strafe.Id_mitglied == mitglied.Id && strafe.Id_strafe_typ == strafe_typ.Id {
					grid += "<td>" + strconv.FormatFloat(float64(strafe.Anzahl), 'f', 2, 32) + "</td>"
					found = true
					summe += float64(strafe.Anzahl) * float64(strafe_typ.Preis)
				}
			}
			if !found {
				grid += "<td>0</td>"
			}
		}
		for _, strafe_typ := range freie_strafen_typen {
			found := false
			for _, strafe := range freie_strafen {
				if strafe.Id_mitglied == mitglied.Id && strafe.Bezeich == strafe_typ {
					grid += "<td>" + strconv.FormatFloat(float64(strafe.Preis), 'f', 2, 32) + "€</td>"
					found = true
					summe += float64(strafe.Anzahl) * float64(strafe.Preis)
				}
			}
			if !found {
				grid += "<td>0€</td>"
			}
		}
		grid += "<td>" + strconv.FormatFloat(summe, 'f', 2, 32) + "</td>"
		grid += "</tr>"
	}

	grid += "</table>"
	return grid
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
		log.Printf("Error: %s", err)
	}
}

func get_mitglieder() []Mitglied {
	connect_to_db()

	// Query the database.
	rows, err := db.Query("SELECT * FROM mitglieder")
	if err != nil {
		log.Printf("Error: %s", err)
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
			log.Printf("Error: %s", err)
		}
		mg = append(mg, Mitglied{id, name, vname, nickname, false})
	}
	db.Close()
	return mg
}

func get_veranstaltungen() []Veranstaltung {
	connect_to_db()
	rows, err := db.Query("SELECT * FROM veranstaltungen")
	if err != nil {
		log.Printf("Error: %s", err)
	}
	defer rows.Close()

	vers := []Veranstaltung{}
	for rows.Next() {
		var id int
		var bezeichnung string
		var datum string
		err := rows.Scan(&id, &bezeichnung, &datum)
		if err != nil {
			log.Printf("Error: %s", err)
		}
		vers = append(vers, Veranstaltung{id, bezeichnung, datum})
	}
	db.Close()
	return vers
}

func get_veranstaltungen_zeitraum(von_datum string, bis_datum string) []Veranstaltung {
	connect_to_db()
	rows, err := db.Query("SELECT * FROM veranstaltungen WHERE datum BETWEEN ? AND ? ORDER BY datum DESC", von_datum, bis_datum)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	defer rows.Close()

	vers := []Veranstaltung{}
	for rows.Next() {
		var id int
		var bezeichnung string
		var datum string
		err := rows.Scan(&id, &bezeichnung, &datum)
		if err != nil {
			log.Printf("Error: %s", err)
		}
		vers = append(vers, Veranstaltung{id, bezeichnung, datum})
	}
	db.Close()
	return vers
}

func get_strafen(mitgliedid int, von_datum string, bis_datum string) []Strafen_neu {
	connect_to_db()
	strafen := []Strafen_neu{}
	// wenn mitgliedid == 0 dann alle mitglieder
	var rows *sql.Rows
	var err error
	if mitgliedid == 0 {
		rows, err = db.Query("select strafen.id, strafen.id_strafe_typ, strafen_typ.preis, strafen_typ.bezeichnung, strafen.id_mitglied, mitglieder.vname, mitglieder.name, mitglieder.nickname, strafen.datum, strafen.preis, strafen_typ.aktiv, strafen.anzahl, strafen.bezeich, strafen.id_veranstaltung, veranstaltungen.datum, veranstaltungen.bezeichnung from strafen left join veranstaltungen on strafen.id_veranstaltung = veranstaltungen.id left join strafen_typ on strafen.id_strafe_typ = strafen_typ.id join mitglieder on strafen.id_mitglied = mitglieder.id WHERE (veranstaltungen.datum BETWEEN ? AND ?) OR (strafen.datum BETWEEN ? AND ?)", von_datum, bis_datum, von_datum, bis_datum)
		if err != nil {
			log.Printf("Error: %s", err)
		}
	} else {
		rows, err = db.Query("select strafen.id, strafen.id_strafe_typ, strafen_typ.preis, strafen_typ.bezeichnung, strafen.id_mitglied, mitglieder.vname, mitglieder.name, mitglieder.nickname, strafen.datum, strafen.preis, strafen_typ.aktiv, strafen.anzahl, strafen.bezeich, strafen.id_veranstaltung, veranstaltungen.datum, veranstaltungen.bezeichnung from strafen left join veranstaltungen on strafen.id_veranstaltung = veranstaltungen.id left join strafen_typ on strafen.id_strafe_typ = strafen_typ.id join mitglieder on strafen.id_mitglied = mitglieder.id WHERE (veranstaltungen.datum BETWEEN ? AND ?) OR (strafen.datum BETWEEN ? AND ?) AND strafen.id_mitglied = ?", von_datum, bis_datum, von_datum, bis_datum, mitgliedid)
		if err != nil {
			log.Printf("Error: %s", err)
		}
	}
	defer rows.Close()
	for rows.Next() {
		var Id int
		var id_strafe_typ int
		var preis_strafe_typ sql.NullFloat64
		var bezeich_strafe_typ sql.NullString
		var id_mitglied int
		var vname string
		var name string
		var nickname string
		var datum_strafe sql.NullString
		var preis_strafe float32
		var aktiv sql.NullBool
		var anzahl float32
		var bezeich_strafe sql.NullString
		var id_veranstaltung int
		var datum_veranstaltung sql.NullString
		var bezeich_veranstaltung sql.NullString
		err := rows.Scan(&Id, &id_strafe_typ, &preis_strafe_typ, &bezeich_strafe_typ, &id_mitglied, &vname, &name, &nickname, &datum_strafe, &preis_strafe, &aktiv, &anzahl, &bezeich_strafe, &id_veranstaltung, &datum_veranstaltung, &bezeich_veranstaltung)
		if err != nil {
			log.Printf("Error: %s", err)
		}
		strafen = append(strafen, Strafen_neu{Id, id_strafe_typ, float32(preis_strafe_typ.Float64), bezeich_strafe_typ.String, id_mitglied, vname, name, nickname, datum_strafe.String, preis_strafe, aktiv.Bool, anzahl, bezeich_strafe.String, id_veranstaltung, datum_veranstaltung.String, bezeich_veranstaltung.String})
	}

	db.Close()
	return strafen
}

func get_strafen_typen() []Strafe_typ {
	connect_to_db()
	strafen_typen := []Strafe_typ{}
	rows, err := db.Query("SELECT * FROM strafen_typ")
	if err != nil {
		log.Printf("Error: %s", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var bezeichnung string
		var preis float32
		var aktiv bool
		err := rows.Scan(&id, &bezeichnung, &preis, &aktiv)
		if err != nil {
			log.Printf("Error: %s", err)
		}
		strafen_typen = append(strafen_typen, Strafe_typ{id, bezeichnung, preis, aktiv})
	}
	db.Close()
	return strafen_typen
}
