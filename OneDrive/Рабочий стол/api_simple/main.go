package main

import (
	helper "api_simple/helper"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	uName, email, pass, passConfirm := "", "", "", ""

	// registration
	mux.HandleFunc("/signUp", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		uName = r.FormValue("username")
		email = r.FormValue("email")
		pass = r.FormValue("password")
		passConfirm = r.FormValue("passwordConfirm")
		if helper.IsEmpty(uName) || helper.IsEmpty(email) || helper.IsEmpty(pass) || helper.IsEmpty(passConfirm) {
			w.Write([]byte("Hechqaysi filed bo'sh bo'lishi mumkin emas !!"))
			return
		}
		if pass == passConfirm {
			w.Write([]byte("Muofiqiyatli ro'yxatdan o'tingiz "))
			
		}else {
			w.Write([]byte("Parol Confirm bir xilda emas"))
		}

	})

	// login
	mux.HandleFunc("/signIn", func(w http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe(":8080", mux)
}
