package server

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUsr(w http.ResponseWriter, r *http.Request) {

	bodyRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var user user

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		w.Write([]byte("Erro ao converter"))
	}

	dbconn, erro := db.Connect()
	if erro != nil {
		w.Write([]byte("Erro ao conectar"))
		return
	}
	defer dbconn.Close()

	statement, erro := dbconn.Prepare("insert into users (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	insertion, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar statement"))
		return
	}

	idInserted, err := insertion.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ap obter id"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserted)))
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Erro ao conectar ao db"))
	}
	defer db.Close()

	lines, err := db.Query("select * from users")
	if err != nil {
		w.Write([]byte("Erro ao buscar usuarios"))
	}
	defer lines.Close()

	var users []user
	for lines.Next() {
		var user user

		if err := lines.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Erro ao escanear usuário"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Erro ao converter em json"))
		return
	}
}

func SearchUser(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)

	ID, err := strconv.ParseUint(param["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter parametro"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Erro em conectar ao db"))
		return
	}

	line, err := db.Query("select * from users where id = ?", ID)
	if err != nil {
		w.Write([]byte("Erro ao buscar usuário"))
		return
	}

	var user user

	if line.Next() {
		if err := line.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Erro ao buscar usuário"))
			return
		}
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Erro ao converter para json"))
		return
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}
