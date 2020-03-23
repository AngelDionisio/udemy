package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {

		// get form values
		userName := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		// username taken?
		if _, ok := dbUsers[userName]; ok {
			http.Error(w, "Username already exists", http.StatusForbidden)
			return
		}

		// create session
		sID, err := uuid.NewV4()
		if err != nil {
			log.Fatalf("error tryign to create UUID due to: %v", err)
		}

		cookie := http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		http.SetCookie(w, &cookie)
		dbSessions[cookie.Value] = userName

		// store user in dbUsers
		u := user{
			UserName: userName,
			Password: p,
			First:    f,
			Last:     l,
		}

		dbUsers[userName] = u

		fmt.Printf("db: %#v\n", dbUsers)

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
