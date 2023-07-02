package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book
var id int = 1

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(id)
	id++
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	for index, item := range books {
		if item.ID == params["id"] {
			if book.ID != "" {
				books[index].ID = book.ID
			}
			if book.Title != "" {
				books[index].Title = book.Title
			}
			if book.Author != "" {
				books[index].Author = book.Author
			}
			if book.Year != "" {
				books[index].Year = book.Year
			}
			fmt.Fprintf(w, "book id: %s, updated", params["id"])
		} else {
			fmt.Fprintf(w, "book id: %s, Not Found", params["id"])
			return
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode(&item)
			fmt.Fprintf(w, "book id: %s deleted", params["id"])
			break
		}
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", getBooks).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", getBook).Methods(http.MethodGet)
	r.HandleFunc("/books", createBook).Methods(http.MethodPost)
	r.HandleFunc("/books/{id}", updateBook).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", deleteBook).Methods(http.MethodDelete)

	fmt.Println("server starting...")
	http.ListenAndServe(":8080", r)
}

/*
pelajaran yang bisa diambil dari crud di atas:
1. Decoder digunakan untuk mengambil json raw dan di proses ke backend (dari json ke backend)
2. Encoder digunakan untuk menampilkan hasil proses backedn ke bentuk json raw (dari backend ke json)
*/
