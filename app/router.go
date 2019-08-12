package app

import (
	"fmt"
	"net/http"

	app "./controllers"
	"github.com/gorilla/mux"
)

func Init() {
	router := mux.NewRouter()

	router.HandleFunc("/", app.HomeController)
	router.HandleFunc("/home", app.HomeController)

	fmt.Println("Sunucu :8080 portu üzerinden dinleniyor")
	http.ListenAndServe(":8080", router)
}
