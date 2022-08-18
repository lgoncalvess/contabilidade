package routers

import (
	"contabilidade/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	r := mux.NewRouter()

	r.HandleFunc("/", controller.RenderTemplate).Methods(http.MethodGet)
	r.HandleFunc("/contas", controller.BuscarContas).Methods(http.MethodGet)
	r.HandleFunc("/contas", controller.InserirConta).Methods(http.MethodPost)
	r.HandleFunc("/contas", controller.EditarConta).Methods(http.MethodPut)
	r.HandleFunc("/contas", controller.ExcluirConta).Methods(http.MethodDelete)

	r.HandleFunc("/transacoes", controller.BuscarTransacoes).Methods(http.MethodGet)
	r.HandleFunc("/transacoes", controller.InserirTransacao).Methods(http.MethodPost)
	r.HandleFunc("/transacoes/data", controller.BuscarTransacoesFromDate).Methods(http.MethodGet)
	r.HandleFunc("/transacoes", controller.DeleteTransaction).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":9090", r))
}
