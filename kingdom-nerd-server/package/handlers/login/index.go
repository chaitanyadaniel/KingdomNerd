package login

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	authutils "server/utils/auth-utils"

	_ "github.com/go-sql-driver/mysql"
)

type Credentials struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func LoginCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"success": "ok",
	}
	jsonResponse, err5 := json.Marshal(response)
	if err5 != nil {
		http.Error(w, err5.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	var hashedPassword string
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/kingdom-db")
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer db.Close()

	err1 := r.ParseForm()
	if err1 != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	creds.Username = r.PostFormValue("username")
	creds.Password = r.PostFormValue("password")

	err2 := db.QueryRow("SELECT password,firstName,lastName FROM users1 WHERE email=?", creds.Username).Scan(&hashedPassword, &creds.Firstname, &creds.Lastname)

	if err2 != nil {
		http.Error(w, "Invalid username or password -1", http.StatusUnauthorized)
		return
	}

	if hashedPassword != creds.Password {
		http.Error(w, "Invalid username or password -2", http.StatusUnauthorized)
		return
	}

	tokenString, err4 := authutils.CreateToken(creds.Username, creds.Firstname, creds.Lastname)
	if err4 != nil {
		http.Error(w, "Error creating token: "+err4.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{
		"token": tokenString,
	}
	jsonResponse, err5 := json.Marshal(response)
	if err5 != nil {
		http.Error(w, err5.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
