package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TomislavGalic/bookings/pkg/config"
	"github.com/TomislavGalic/bookings/pkg/handlers"
	"github.com/TomislavGalic/bookings/pkg/render"
)

const portNumber = ":8080"

// main is the main app function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//	http.HandleFunc("/", handlers.Repo.Home)
	//	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting app on port: %s", portNumber))
	//	_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
