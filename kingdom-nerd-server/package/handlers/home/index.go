package home

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// HomeHandler is a basic handler for the home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// // Generate a UUIDv4 for the ID column
	// id := uuid.New().String()

	// // Create a new row in the table
	// stmt, err := db.Prepare("INSERT INTO users1 (id, firstName, lastName, email, password) VALUES (?, ?, ?, ?, ?)")
	// if err != nil {
	// 	fmt.Println("Failed to prepare the SQL statement:", err)
	// 	return
	// }
	// defer stmt.Close()

	// _, err = stmt.Exec(id, "chaitanya", "daniel", "pchaitanyadaniel@gmail.com", "password")
	// if err != nil {
	// 	fmt.Println("Failed to execute the SQL statement:", err)
	// 	return
	// }
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, Authorization")
	// w.Header().Set("Expose-Headers", "*")
	w.Write([]byte("Hello, world!"))
}
