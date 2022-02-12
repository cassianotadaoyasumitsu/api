package controllers

import (
	"api/src/db"
	"api/src/model"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// User Create
func UserCreate(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUsersRepo(db)
	user.ID, err = repo.Create(user)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
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
