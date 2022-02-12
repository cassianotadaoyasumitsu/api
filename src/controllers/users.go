package controllers

import (
	"api/src/db"
	"api/src/model"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// User Create
func UserCreate(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user model.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repo := repositories.NewUsersRepo(db)
	userID, err := repo.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Inserted ID: %d", userID)))
}

// Users Search
func UsersSearch(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users search!"))
}

// User Search
func UserSearch(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User search!"))
}

// User Update
func UserUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User update!"))
}

// User Delete
func UserDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User delete"))
}
