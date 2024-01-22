package contollers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repos"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating User"))

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Conn()
	if err != nil {
		log.Fatal(err)
	}

	repos := repos.NewUserRepo(db)
	userID, err := repos.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("ID inserted %d", userID)))

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
