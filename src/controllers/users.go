package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/model"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err = user.Preparar("register"); err != nil {
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
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := db.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUsersRepo(db)
	users, err := repo.Search(nameOrNick)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)

}

// User Search
func UserSearch(w http.ResponseWriter, r *http.Request) {
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

	repo := repositories.NewUsersRepo(db)
	user, err := repo.SearchByID(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// User Update
func UserUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := auth.TokenUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDToken {
		responses.Err(w, http.StatusForbidden, errors.New("user id is not authorized to update user id"))
		return
	}

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

	if err = user.Preparar("edit"); err != nil {
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
	if err = repo.Update(userID, user); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// User Delete
func UserDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := auth.TokenUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDToken {
		responses.Err(w, http.StatusForbidden, errors.New("user id is not authorized to delete user id"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUsersRepo(db)
	if err = repo.Delete(userID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Follow User
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.TokenUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if followerId == userID {
		responses.Err(w, http.StatusForbidden, errors.New("you can't follow yourself"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepo(db)
	if err = repository.Follow(followerId, userID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

// Unfollow User
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.TokenUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if followerId == userID {
		responses.Err(w, http.StatusForbidden, errors.New("you can't unfollow yourself"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepo(db)
	if err = repository.Unfollow(followerId, userID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// GetFollowers
func UserFollowers(w http.ResponseWriter, r *http.Request) {
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

	repository := repositories.NewUsersRepo(db)
	followers, err := repository.GetFollowers(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

// GetFollowing
func UserFollowing(w http.ResponseWriter, r *http.Request) {
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

	repository := repositories.NewUsersRepo(db)
	following, err := repository.GetFollowing(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, following)
}

// Password Update
func PasswordUpdate(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := auth.TokenUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if userID != userIDToken {
		responses.Err(w, http.StatusForbidden, errors.New("user id is not authorized to update password"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	var password model.Password
	if err = json.Unmarshal(requestBody, &password); err != nil {
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
	passwordDB, err := repo.GetPassword(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.Verify(passwordDB, password.Old); err != nil {
		responses.Err(w, http.StatusUnauthorized, errors.New("old password is not correct"))
		return
	}

	passwordHash, err := security.Hash(password.New)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = repo.UpdatePassword(userID, string(passwordHash)); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
