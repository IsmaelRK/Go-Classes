package contollers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Creating User"))

}

func SearchUsers(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Searching Users"))

}

func SearchUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Searching User"))

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Updating User"))

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Deleting User"))

}