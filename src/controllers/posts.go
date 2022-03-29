package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/model"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	userID, err := auth.TokenUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewPostRepo(db)
	posts, err := repo.Search(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

// Search for a post
func PostSearch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
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
	post, err := repo.SearchByID(postID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, post)
}

// Update a post
func PostUpdate(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.TokenUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
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
	postDB, err := repo.SearchByID(postID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if postDB.AuthorID != userID {
		responses.Err(w, http.StatusForbidden, errors.New("user not authorized"))
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

	if err = post.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = repo.Update(postID, post); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Delete a post
func PostDelete(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.TokenUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
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
	postDB, err := repo.SearchByID(postID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if postDB.AuthorID != userID {
		responses.Err(w, http.StatusForbidden, errors.New("user not authorized"))
		return
	}

	if err = repo.Delete(postID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Search posts by userID
func PostsByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
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
	posts, err := repo.SearchByUserID(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

// Like Post
func LikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
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
	if err = repo.Like(postID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
