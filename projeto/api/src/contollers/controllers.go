package contollers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repos"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating User"))

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {

		responses.Error(w, http.StatusUnprocessableEntity, err)
		return

	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {

		responses.Error(w, http.StatusBadRequest, err)
		return

	}

	if err = user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conn()
	if err != nil {

		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repos := repos.NewUserRepo(db)
	user.ID, err = repos.Create(user)
	if err != nil {

		responses.Error(w, http.StatusInternalServerError, err)
		return

	}

	responses.JSON(w, http.StatusCreated, user)

}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching Users"))

	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Conn()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repos.NewUserRepo(db)
	users, err := repo.Search(nameOrNick)

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
