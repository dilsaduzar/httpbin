package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type City struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Region      string `json:"region,omitempty"`
	Country     string `json:"country,omitempty"`
	NumberPlate int64  `json:"number_plate,omitempty"`
	Population  int64  `json:"population,omitempty"`
}

// DB is a reference to our MySQL database
var DB *sql.DB

func CitiesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["city"]

	var city City
	if r.Method == http.MethodGet {
		result, err := fetchCity(DB, name)
		if err != nil {
			w.WriteHeader(500)
			io.WriteString(w, err.Error())
			return
		}
		if name != result.Name {
			io.WriteString(w, errMsg("the city is not found."))
			return
		}
		city = result
	} else if r.Method == http.MethodPost {
		out, err := io.ReadAll(r.Body)
		if err != nil {
			io.WriteString(w, errMsg("Error code: C02"))
			return
		}

		var cty City
		err = json.Unmarshal(out, &cty)
		if err != nil {
			io.WriteString(w, errMsg("You need to send a JSON string"))
			return
		}

		err = insertCity(DB, cty)
		if err != nil {
			w.WriteHeader(500)
			io.WriteString(w, err.Error())
			return
		}

		city = cty
	} else {
		w.WriteHeader(501)
		io.WriteString(w, errMsg("Supports only GET and POST methods. Please use GET or POST method."))
		return
	}

	json, err := json.Marshal(city)
	if err != nil {
		io.WriteString(w, errMsg("Error Code: C01"))
		return
	}

	fmt.Fprintln(w, string(json))

}

func fetchCity(db *sql.DB, name string) (City, error) {
	rows, err := db.Query("select * from cities where name=?", name)
	if err != nil {
		return City{}, err
	}
	defer rows.Close()

	var city City
	for rows.Next() {
		err = rows.Scan(&city.ID, &city.Name, &city.Region, &city.Country, &city.NumberPlate, &city.Population)
		if err != nil {
			return City{}, err
		}
	}

	err = rows.Err()
	if err != nil {
		return City{}, err
	}
	return city, nil
}

func insertCity(db *sql.DB, city City) error {
	_, err := db.Exec("insert into cities (name, region, country, number_plate, population) values (?, ?, ?, ?, ?)",
		city.Name, city.Region, city.Country, city.NumberPlate, city.Population)
	if err != nil {
		return err
	}

	return nil
}
