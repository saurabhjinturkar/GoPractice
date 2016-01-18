package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

var store map[int]string
var server string

type Pair struct {
	Server string `json:"server,omitempty"`
	Key    int    `json:"key"`
	Value  string `json:"value"`
}

func init() {
	store = make(map[int]string)
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	key_id := vars["key_id"]
	value := vars["value"]
	key, err := strconv.Atoi(key_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	store[key] = value
	pair := Pair{}
	pair.Server = server
	pair.Key = key
	pair.Value = value
	json.NewEncoder(w).Encode(pair)
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	key_id := vars["key_id"]
	key, err := strconv.Atoi(key_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	if val, ok := store[key]; ok {
		pair := Pair{}
		pair.Key = key
		pair.Value = val
		json.NewEncoder(w).Encode(pair)
		return
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Error\":\"Key Not Found\"}"))
		return
	}
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	output := []Pair{}
	for key, _ := range store {
		pair := Pair{}
		pair.Key = key
		pair.Value = store[key]
		output = append(output, pair)
	}
	json.NewEncoder(w).Encode(output)
	return
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please pass port number")
		return
	}
	r := mux.NewRouter()
	r.HandleFunc("/keys/{key_id}/{value}", put).Methods("PUT", "POST")
	r.HandleFunc("/keys/{key_id}", get).Methods("GET")
	r.HandleFunc("/keys", getAll).Methods("GET")
	http.Handle("/", r)
	fmt.Println("Listening on port", os.Args[1])
	server = "localhost:" + os.Args[1]
	log.Fatal(http.ListenAndServe(":"+os.Args[1], r))
}
