package Function

import (
	"html/template"
	"net/http"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

var predefinedUsername = "Abid"  //Username
var predefinedPassword = "12345" //Password

func LoginPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	if session.Values["authenticated"] == true {
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
		return
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" || password == "" {
			errorMessage := ""
			if username == "" {
				errorMessage += "Username is required. "
			}
			if password == "" {
				errorMessage += "Password is required."
			}
			renderTemplate(w, "login.html", map[string]string{"Error": errorMessage})
			return
		}

		if username == predefinedUsername && password == predefinedPassword {
			session.Values["authenticated"] = true
			session.Values["username"] = username
			session.Save(r, w)
			http.Redirect(w, r, "/welcome", http.StatusSeeOther)
			return
		}

		renderTemplate(w, "login.html", map[string]string{"Error": "Invalid credentials. Please try again."})
		return
	}

	renderTemplate(w, "login.html", nil)
}

func WelcomePage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	if session.Values["authenticated"] != true {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	username := session.Values["username"].(string)
	data := map[string]string{
		"Username": username,
	}
	renderTemplate(w, "welcome.html", data)
}

func LogoutPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Options.MaxAge = -1
	session.Values["authenticated"] = false
	session.Values["username"] = nil
	session.Save(r, w)

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}
