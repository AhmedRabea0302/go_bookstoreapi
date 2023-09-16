package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AhmedRabea0302/go_bookstoreapi/pkg/models"
	"github.com/gorilla/mux"
)

type NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error parsing bookId")
	}

	bookDetails, _ := models.GetBookById(ID)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
