package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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
	Id          int    `db:"id"`
	AuthId      int    `db:"auth_id"`
	Fam         string `db:"fam"`
	Name        string `db:"name"`
	Otch        string `db:"otch"`
	Birthday    string `db:"birthday"`
	Description string `db:"description"`
	Avatar      string `db:"avatar"`
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
	router.HandleFunc("/{user}/{id:[0-9]+}", env.getUserById).Methods("GET")
	router.HandleFunc("/{user}", env.createUser).Methods("POST")
	http.ListenAndServe(":8001", router)
}

func (env *Env) getUserById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	user := User{}

	_ = env.db.Select("*").From("user").Where(dbx.HashExp{"id": id}).One(&user)

	if user.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User not found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (env *Env) createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := User{}
	_ = json.NewDecoder(r.Body).Decode(&user)

	err := env.db.Model(&user).Insert()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Bad request",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (dbconf *DbConfig) getDbParamsFromYaml() error {
	fopen, err := ioutil.ReadFile(dbConfigFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fopen, &dbconf)
	if err != nil {
		return err
	}

	return nil
}
