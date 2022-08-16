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

func BuscarTransacoes(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM transaction")
	if err != nil {
		log.Fatal(err)
	}
	var transactions []models.Transaction
	for rows.Next() {
		var trs models.Transaction
		if err = rows.Scan(&trs.ID, &trs.Name, &trs.Value, &trs.Date, &trs.Type, &trs.Conta_ID); err != nil {
			log.Fatal(err)
		}
		transactions = append(transactions, trs)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Println(transactions)
	json.NewEncoder(w).Encode(transactions)
}

func BuscarTransacoesFromDate(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]interface{})
	if err = json.Unmarshal(corpoRequest, &data); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM transaction WHERE date >= $1 AND date <= $2 AND conta_id = $3", data["dataInicio"], data["dataFim"], data["contaId"])
	if err != nil {
		log.Fatal(err)
	}
	var transactions []models.Transaction
	for rows.Next() {
		var trs models.Transaction
		if err = rows.Scan(&trs.ID, &trs.Name, &trs.Value, &trs.Date, &trs.Type, &trs.Conta_ID); err != nil {
			log.Fatal(err)
		}
		transactions = append(transactions, trs)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}

func InserirTransacao(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var transaction models.Transaction
	if err = json.Unmarshal(corpoRequest, &transaction); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into transaction(name,value,date,type,conta_id) values($1,$2,$3,$4,$5)", transaction.Name, transaction.Value, transaction.Date, transaction.Type, transaction.Conta_ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.RowsAffected())

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("criado!")
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var transaction models.Transaction
	if err = json.Unmarshal(corpoRequest, &transaction); err != nil {
		panic(err)
	}

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM transaction WHERE id = $1", transaction.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("deletado!")
}
