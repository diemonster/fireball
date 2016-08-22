package main

import (
	"fmt"
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/blog/handlers"
	"net/http"
)

func main() {
	rootHandler := handlers.NewRootHandler()

	app := fireball.NewApp()
	app.StaticRoute("/static", "/static")
	app.Static
	router.Handle("/static", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// todo: app.Before = HTTPAuth()

	app.Routes = append(app.Routes, rootHandler.Routes()...)

	fmt.Println("Running on port 8000")
	http.ListenAndServe(":8000", app)
}
