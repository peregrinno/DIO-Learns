package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Endereco struct {
	Logradouro string
	Cidade     string
	Estado     string
}

type Cliente struct {
	ID       int
	Nome     string
	Idade    int
	Endereco Endereco
}

var clientes []Cliente
var nextID int = 1

func initializeData() {
	clientes = append(clientes, Cliente{
		ID:    nextID,
		Nome:  "John Doe",
		Idade: 30,
		Endereco: Endereco{
			Logradouro: "Rua das Flores",
			Cidade:     "São Paulo",
			Estado:     "SP",
		},
	})
	nextID++
}

func getClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

func getCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID inválido"})
		return
	}

	for _, cliente := range clientes {
		if cliente.ID == id {
			json.NewEncoder(w).Encode(cliente)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Cliente não encontrado"})
}

func createCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cliente Cliente
	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Dados inválidos"})
		return
	}

	cliente.ID = nextID
	nextID++
	clientes = append(clientes, cliente)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cliente)
}

func updateCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID inválido"})
		return
	}

	var clienteAtualizado Cliente
	err = json.NewDecoder(r.Body).Decode(&clienteAtualizado)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Dados inválidos"})
		return
	}

	for i, cliente := range clientes {
		if cliente.ID == id {
			clienteAtualizado.ID = id
			clientes[i] = clienteAtualizado
			json.NewEncoder(w).Encode(clienteAtualizado)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Cliente não encontrado"})
}

func deleteCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID inválido"})
		return
	}

	for i, cliente := range clientes {
		if cliente.ID == id {
			clientes = append(clientes[:i], clientes[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Cliente removido com sucesso"})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Cliente não encontrado"})
}

func main() {
	initializeData()

	router := mux.NewRouter()

	router.HandleFunc("/clientes", getClientes).Methods("GET")
	router.HandleFunc("/clientes/{id}", getCliente).Methods("GET")
	router.HandleFunc("/clientes", createCliente).Methods("POST")
	router.HandleFunc("/clientes/{id}", updateCliente).Methods("PUT")
	router.HandleFunc("/clientes/{id}", deleteCliente).Methods("DELETE")

	fmt.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
