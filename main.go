package main

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type DbConfig struct {
	Development struct {
		Dialect    string
		Datasource string
	}
}

type Env struct {
	db *dbx.DB
}

var dbConfigFile = "dbconfig.yml"

type User struct {
	Id int
	Fam string
	Name string
	Otch string
	Birthday string
	Description string
	Avatar string
}

func main() {
	dbconf := DbConfig{}
	err := dbconf.getDbParamsFromYaml()
	if err != nil {
		panic(err)
	}

	dialect := &dbconf.Development.Dialect
	datasource := &dbconf.Development.Datasource

	var db, _ = dbx.Open(*dialect, *datasource)
	env := Env{db: db}
	
	router := mux.NewRouter()
	router.HandleFunc("/{user}{id:[0-9]+}", getUserById).Methods("GET")
	http.ListenAndServe(":8000", router)
}

func (env *Env) getUserById(w http.ResponseWriter, *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	user := User{}
	
	_ := env.db.Select("*").From("user").Where(dbx.HashExp{"id": id}).One(&user)
	
	if user.Id == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User not found"
		})
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}