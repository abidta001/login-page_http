package main

import (
	"fmt"
	"login-http/Functions"
	"net/http"
)

func main() {
	http.HandleFunc("/", Function.LoginPage)
	http.HandleFunc("/login", Function.LoginPage)
	http.HandleFunc("/welcome", Function.WelcomePage)
	http.HandleFunc("/logout", Function.LogoutPage)
	
	
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
