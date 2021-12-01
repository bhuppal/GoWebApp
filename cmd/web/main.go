package main

import (
	"fmt"
	"github.com/bhuppal/go/goweb/pkg/config"
	"github.com/bhuppal/go/goweb/pkg/handlers"
	"github.com/bhuppal/go/goweb/pkg/render"
	"log"
	"net/http"
)

const PORT_NUMBER = ":8080"

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

	render.NewTemplate(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", PORT_NUMBER))

	//http.ListenAndServe(PORT_NUMBER, nil)

	srv := &http.Server{
		Addr:    PORT_NUMBER,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

/*
package main

import (
	"errors"
	"fmt"
	"net/http"
)

const PORT_NUMBER = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_,_ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

// Divide is the divide calculation page handler
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(8, 0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf( "%f divided by %f is %f", 100.0, 10.0, f))
}

func divideValues(x, y float64) (float64, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// addValues add two integers and returns the sum
func addValues(x, y int) int {
	sum := x + y
	return sum
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World!")
		if err != nil {
			fmt.Println(err, " - Error happened")
		}
		fmt.Println(fmt.Sprintf("Number of bytes writtern: %d", n))
	})

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	fmt.Println(fmt.Sprintf("Starting application on port %s", PORT_NUMBER))
	_ = http.ListenAndServe(PORT_NUMBER, nil)
}
*/
