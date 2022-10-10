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

var app config.AppConfig

// main is the main app function
func main() {

	//change this in true when in production
	app.InProduction = false

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
