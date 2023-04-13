package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Block struct {
	Pos       int
	Data      BookCheckout
	TimeStamp string
	Hash      string
	PrevHash  string
}

type BookCheckout struct {
	BookID       string `json:"bookid"`
	User         string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis    string `json:"is_genesis"`
}

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishDate string `json:"publish_date"`
	ISBN        string `json:"isbn"`
}

type Blockchain struct {
	block []*Block
}

var blockchain *Blockchain

func newBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("couldnt not create: %v", err)
		w.Write([]byte("couldnt create new book"))
		return
	}

	h := md5.New()
	io.WriteString(h, book.ISBN+book.PublishDate)
	book.ID = fmt.Sprintf("%x", h.Sum(nil))

	resp, err := json.MarshalIndent(book, "", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("couldnt marshal payload: %v", err)
		w.Write([]byte("couldnt save the book data"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func main() {
	r := mux.NewRouter()
	// r.HandleFunc("/", getBlockchain).Methods("GET")
	// r.HandleFunc("/", writeBlockchain).Methods("POST")
	r.HandleFunc("/new", newBook).Methods("POST")

	log.Println("Listening to port 3000")

	log.Fatal(http.ListenAndServe(":3000", r))
}
