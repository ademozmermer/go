package app

import (
	"fmt"
	"net/http"

	database "../../system/drivers"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	users := database.ConnectSQL()

	users.Query("SELECT * FROM users")
	fmt.Println("Selam dünya")

	defer users.Close()
}
