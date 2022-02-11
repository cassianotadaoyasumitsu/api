package controllers

import "net/http"

// User Create
func UserCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User create!"))
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
