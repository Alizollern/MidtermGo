package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	// "strconv"
	"github.com/gorilla/mux"
)
var(
	mutex sync.Mutex
)

var store = map[string]string{
	"1": "To Kill a Mockingbird",
	"2": "Pride and Prejudice",
	"3": "The Diary of a Young Girl",
	"4": "1984",
	"5": "The Little Prince",
	"6": "The Great Gatsby",
}
func getAll(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	mutex.Lock()
	for _, data := range store{
		fmt.Fprintf(w,"Data is: "+ data + " \n")
	}
	mutex.Unlock()
}
func getStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mutex.Lock()
	getParams := mux.Vars(r)
	values := store[getParams["key"]]
	fmt.Println(values)
	fmt.Fprintf(w, "Values is: "+values)
	mutex.Unlock()
	return 

}
func changeValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mutex.Lock()
	getParams := mux.Vars(r)

	store[getParams["key"]] = getParams["value"]
	fmt.Fprintf(w,"You change the value : ")
	fmt.Fprintf(w, getParams["key"] + " to " + getParams["value"])
	fmt.Println("Done")
	mutex.Unlock()
	return
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/store", getAll).Methods("GET")
	r.HandleFunc("/store/{key}", getStore).Methods("GET")
	r.HandleFunc("/store/{key}/{value}", changeValue)
	log.Fatal(http.ListenAndServe(":8000", r))
}
