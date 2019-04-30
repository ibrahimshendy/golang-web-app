package main

import (
	"net/http"
	"site/models"
	"site/routes"
	"site/utils"
)

func main() {
	models.Init()
	utils.LoadTemplate("templates/*.gohtml")
	r := routes.InitRoutes()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
