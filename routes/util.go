package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// Tmpl exports the compiled templates
var Tmpl *template.Template

func init() {
	templateFuncMap := template.FuncMap{
		"despace": func(s string) string {
			return strings.Replace(s, " ", "_", -1)
		},
		"isoTime": func(t time.Time) string {
			return t.Format(time.RFC3339)
		},
	}

	Tmpl = template.Must(template.New("main").Funcs(templateFuncMap).ParseGlob(getPath() + "/*"))
}

func getPath() (p string) {
	_, err := os.Open("./views")
	if err == nil {
		p = "./views/"
		return
	}
	_, err = os.Open("../views")
	if err == nil {
		p = "../views/"
		return
	}
	return
}

type errorPageData struct {
	Message string
}

func errRes(w http.ResponseWriter, r *http.Request, code int, message string, err error) {
	log.Printf("WARN Sending Error Response: %+v, %+v, %+v", code, message, err)
	log.Println(string(debug.Stack()))

	w.WriteHeader(code)
	if r.Header.Get("Accept") == "application/json" {
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, message)))
		return
	}

	if err := Tmpl.ExecuteTemplate(w, "error.html", errorPageData{
		Message: message,
	}); err != nil {
		log.Printf("ERROR error rendering the error template: %+v", err)
		w.Write([]byte("Error rendering the error template. Oh dear."))
		return
	}
}
