package main

import (
	"crypto/subtle"
	"fmt"
	"net/http"
)

const username = "admin"
const password = "password123"

var realm = "login"

func login(w http.ResponseWriter, req *http.Request) {

	user, pass, ok := req.BasicAuth()
	cookieUser, _ := req.Cookie("username")

	if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
		w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
		w.WriteHeader(401)
		w.Write([]byte("Unauthorised.\n"))

		if cookieUser == nil {
			fmt.Println("cookie not found")
		}

		return
	}

	fmt.Println("Successfully logged in!")
	w.WriteHeader(200)

}

func main() {

	http.HandleFunc("/login", login)
	http.ListenAndServe(":8090", nil)
}
