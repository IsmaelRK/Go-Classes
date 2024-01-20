package server

import (
	"crud/db"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

	idInserted := insertion.LastInsertId()

}
