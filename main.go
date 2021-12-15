package main

import (
	_ "embed"
	"encoding/hex"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
)

//go:embed index.html
var indextpl string

type User struct {
	Username string
	Password string
}

var Users = []User{}

func randString(l int) string {
	rand.Seed(time.Now().UnixNano() + int64(rand.Int()))
	b := make([]byte, l)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

func validateAuth(username string, password string) bool {
	for _, u := range Users {
		if u.Username == username {
			if u.Password == password {
				return true
			}
		}
	}
	return false
}

func BasicAuth(w http.ResponseWriter, r *http.Request) {
	type data struct {
		Msg string
	}
	var tplData data
	tpl, _ := template.New("Index").Parse(indextpl)
	user, pass, ok := r.BasicAuth()
	if !ok || !validateAuth(user, pass) {
		w.Header().Set("WWW-Authenticate", `Basic realm="AuthDemo"`)
		w.WriteHeader(401)
		tplData.Msg = "Unauthorized"
	} else {
		tplData.Msg = fmt.Sprintf("Welcome %s", user)
	}
	tpl.Execute(w, tplData)
}

func FormsAuth(w http.ResponseWriter, r *http.Request) {
	type data struct {
		Msg string
	}
	var tplData data
	tpl, _ := template.New("Index").Parse(indextpl)
	if r.Method == "GET" {
		tpl.Execute(w, nil)
		return
	}

	r.ParseForm()
	user := r.FormValue("username")
	pass := r.FormValue("password")
	if validateAuth(user, pass) {
		tplData.Msg = fmt.Sprintf("Welcome %s", user)
	} else {
		tplData.Msg = "Invalid auth"
	}
	tpl.Execute(w, tplData)
}

func main() {
	var port *int = flag.Int("Port", 8091, "Listening port")
	flag.Parse()
	for i := 1; i < 5; i++ {
		Users = append(Users, User{Username: fmt.Sprintf("Account%d", i), Password: randString(16)})
	}

	fmt.Println("Created the following demo users:")
	b, _ := yaml.Marshal(Users)
	fmt.Println(string(b))

	fmt.Println("created endpoints /basic and /forms")
	fmt.Printf("Listening on port %d\n", *port)

	root := mux.NewRouter()
	root.HandleFunc("/basic", BasicAuth).Methods("GET", "POST")
	root.HandleFunc("/forms", FormsAuth).Methods("GET", "POST")
	root.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/forms", http.StatusFound)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", *port), root)

}
