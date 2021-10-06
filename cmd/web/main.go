package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/prakharmaurya/go_bookings/internal/config"
	"github.com/prakharmaurya/go_bookings/internal/handlers"
	"github.com/prakharmaurya/go_bookings/internal/render"
)

const portNumber = 8080

var app config.AppConfig

func main() {

	app.InProduction = false
	app.UseCache = true // In dev mode

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatalln("Failed to create tempalate cache", err)
	}

	app.TemplateCache = tc

	render.NewTemplate(&app)
	repo := handlers.NewRepositor(&app)
	handlers.NewHandlers(repo)

	fmt.Println("Server starting at :", portNumber)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", portNumber),
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to start server", err)
		defer func() {
			srv.Close()
		}()
	}
}
