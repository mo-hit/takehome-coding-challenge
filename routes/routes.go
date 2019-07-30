package routes

import (
	"gopher/models"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Init() (h http.Handler) {
	r := mux.NewRouter()

	h = recoverWrap(r)
	h = handlers.LoggingHandler(os.Stdout, h)

	r.Path("/").
		Methods("GET").
		HandlerFunc(serveIndex)

	r.Path("/submit").
		Methods("GET").
		HandlerFunc(serveSubmit)

	r.Path("/submit").
		Methods("POST").
		HandlerFunc(jokeCreateHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))

	return
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	joke := models.Joke{}
	if err := joke.Random(); err != nil {
		errRes(w, r, 500, "Error looking up joke", err)
		return
	}
	if err := Tmpl.ExecuteTemplate(w, "index.html", indexPageData{
		Joke: joke,
	}); err != nil {
		errRes(w, r, 500, "Problem with template", err)
		return
	}
}

func serveSubmit(w http.ResponseWriter, r *http.Request) {
	if err := Tmpl.ExecuteTemplate(w, "submit.html", nil); err != nil {
		errRes(w, r, 500, "Problem with template", err)
		return
	}
}

type indexPageData struct {
	Joke models.Joke
}

func jokeCreateHandler(w http.ResponseWriter, r *http.Request) {
	errRes(w, r, 501, "Not Implemented", nil)
	return
}
