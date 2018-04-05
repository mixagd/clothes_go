package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "clothes"
	password = "clothes"
	dbname   = "clothes"
)

var db *sql.DB

type Cloth struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Colour string `json:"colour"`
	Fit    string `json:"fit"`
	Owner  string `json:"owner"`
}

// var data = []Cloth{
// 	Cloth{ID: "1", Type: "Blouse", Colour: "Blue", Fit: "Tight", Owner: "Michaila"},
// 	Cloth{ID: "2", Type: "Blouse", Colour: "White", Fit: "Loose", Owner: "Maria"},
// 	Cloth{ID: "3", Type: "Blouse", Colour: "Black", Fit: "Loose", Owner: "Maria"},
// 	Cloth{ID: "4", Type: "Jeans", Colour: "Blue", Fit: "Tight", Owner: "Michaila"},
// 	Cloth{ID: "5", Type: "Trousers", Colour: "Beige", Fit: "Tight", Owner: "Katerina"},
// }

func getAllClothes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.Query("SELECT * FROM clothes")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	myRetList := make([]Cloth, 0)
	for rows.Next() {
		var cloth = Cloth{}
		err = rows.Scan(&cloth.ID, &cloth.Type, &cloth.Colour, &cloth.Fit, &cloth.Owner)
		if err != nil {
			// handle this error
			panic(err)
		}
		myRetList = append(myRetList, cloth)
	}
	bd, _ := json.Marshal(myRetList)
	w.Write(bd)

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

}

func getClothByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	cloth, err := findCloth(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	bd, err := json.Marshal(cloth)
	w.Write(bd)
}

func findCloth(id string) (Cloth, error) {

	sqlStatement := `SELECT * FROM clothes WHERE id=$1;`
	var cloth = Cloth{}
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&cloth.ID, &cloth.Type, &cloth.Colour, &cloth.Fit, &cloth.Owner)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return Cloth{}, err
	case nil:
		return cloth, nil
	default:
		return Cloth{}, err
	}
}

func deleteCloth(id string) {
	sqlStatement := `DELETE FROM clothes WHERE id = $1;`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
}

func deleteClothByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	deleteCloth(vars["id"])
	w.WriteHeader(http.StatusNoContent)
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	r := mux.NewRouter()
	r.HandleFunc("/clothes", getAllClothes).Methods("GET")
	r.HandleFunc("/clothes/{id}", getClothByID).Methods("GET")
	r.HandleFunc("/clothes/{id}", deleteClothByID).Methods("DELETE")
	http.ListenAndServe("0.0.0.0:8080", r)
}
