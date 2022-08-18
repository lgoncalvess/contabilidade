package controller

import (
	"contabilidade/database"
	"contabilidade/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func InserirConta(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var conta models.Conta
	if err = json.Unmarshal(corpoRequest, &conta); err != nil {
		panic(err)
	}

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into conta(name) values($1)", conta.Name)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("criado!")
}

func EditarConta(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var conta models.Conta
	if err = json.Unmarshal(corpoRequest, &conta); err != nil {
		panic(err)
	}

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("UPDATE conta SET Name = $1 WHERE id = $2", conta.Name, conta.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("editado!")
}

func ExcluirConta(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var conta models.Conta
	if err = json.Unmarshal(corpoRequest, &conta); err != nil {
		panic(err)
	}

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM conta WHERE id = $1", conta.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("deletado!")
}

func BuscarContas(w http.ResponseWriter, r *http.Request) {
	fmt.Println("comecou")
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from conta")

	if err != nil {
		log.Fatal(err)
	}
	var contas []models.Conta
	for rows.Next() {
		var conta models.Conta
		if err = rows.Scan(&conta.ID, &conta.Name); err != nil {
			panic(err)
		}
		contas = append(contas, conta)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contas)
}
