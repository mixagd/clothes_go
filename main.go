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

func updateClothByIDOLD(w http.ResponseWriter, r *http.Request) {
	u := make(map[string]interface{})
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &u)
	if err != nil {
		panic(err)
	}
	vars := mux.Vars(r)
	strTmpl := `UPDATE clothes SET %s WHERE id = $1;`
	vals := make([]interface{}, 0)
	vals = append(vals, vars["id"])
	strstr := ""
	strcomma := ""
	cntVal := 2
	for k, v := range u {
		fmt.Printf("%v => %v\n", k, v)
		strstr = strstr + strcomma + fmt.Sprintf("%s=$%d", k, cntVal)
		//fmt.Sprintf("%s %s %s = $%d", strcomma, strstr, k, cntVal)
		cntVal = cntVal + 1
		strcomma = ","
		vals = append(vals, v)
	}
	fmt.Println(strstr)
	strQuery := fmt.Sprintf(strTmpl, strstr)
	fmt.Printf("%s\n%v\n", strQuery, vals)
	w.WriteHeader(http.StatusOK)
}

func updateClothByID(w http.ResponseWriter, r *http.Request) {
	var myCloth = Cloth{}

	vars := mux.Vars(r)
	clothID := vars["id"]

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	u := make(map[string]interface{})
	err = json.Unmarshal(b, &u)
	var requiredFields = []string{"colour", "fit", "type", "owner"}
	var missingFields = make([]string, 0)
	allFound := true
	for _, k := range requiredFields {
		if _, ok := u[k]; !ok {
			allFound = false
			missingFields = append(missingFields, k)
		}
	}

	if !allFound {
		w.WriteHeader(http.StatusBadRequest)
		bd := fmt.Sprintf("Missing fields: %v", missingFields)
		w.Write([]byte(bd))
		return
	}

	err = json.Unmarshal(b, &myCloth)
	if err != nil {
		panic(err)
	}

	strSQL := `UPDATE clothes SET type=$1, colour=$2,
		fit=$3, owner=$4 WHERE id=$5;
	`
	_, err = db.Exec(strSQL, myCloth.Type, myCloth.Colour,
		myCloth.Fit,
		myCloth.Owner, clothID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	cloclo, err := findCloth(clothID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		bdExplain := fmt.Sprintf("Error while retrieving cloth with id %s: %v",
			clothID, err.Error())
		w.Write([]byte(bdExplain))
		return
	}
	bd, err := json.Marshal(cloclo)
	w.WriteHeader(http.StatusOK)
	w.Write(bd)
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
	r.HandleFunc("/clothes/{id}", updateClothByID).Methods("PUT")
	http.ListenAndServe("0.0.0.0:8080", r)
}
