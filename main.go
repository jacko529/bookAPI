package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


// book struct

type Book struct {
	ID  string `json:"id"`
	Isbn  string `json:"isbn"`
	Title  string `json:"title"`
	Author  *Author `json:"author"`

}

type Author struct {
	Forename string `json:"forename"`
	Surname  string `json:"surname"`
}


var books []Book


func getBooks(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}


func getBook(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// loop through
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000))//not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)

}

func deleteBook(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}


// slice


func main(){
	// init router
	r := mux.NewRouter();

	// mock data - implement db
	books = append(books, Book{ID: "1", Isbn:"34334DF", Title:"Book One", Author:
		&Author{Forename: "Jack", Surname:"Churchill"}})
	books = append(books, Book{ID: "2", Isbn:"34dfdf4DF", Title:"Book Two", Author:
		&Author{Forename: "Magda", Surname:"Churchill"}})


	// route handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8000", r))
}