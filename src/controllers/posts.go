package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/model"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Create a new post
func PostCreate(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.TokenUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post model.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorID = userID

	if err = post.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewPostRepo(db)
	post.ID, err = repo.Create(post)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}

// Search for posts
func PostsSearch(w http.ResponseWriter, r *http.Request) {

}

// Search for a post
func PostSearch(w http.ResponseWriter, r *http.Request) {

}

// Update a post
func PostUpdate(w http.ResponseWriter, r *http.Request) {

}

// Delete a post
func PostDelete(w http.ResponseWriter, r *http.Request) {

}
