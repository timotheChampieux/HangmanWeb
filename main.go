package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var (
	gameStarted bool   = false
	pseudo      string = ""
	difficulty  string = ""
)

func main() {
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	tmpl, tempErr := template.ParseGlob("./templates/*.html")
	if tempErr != nil {
		fmt.Printf("erreur avec le temp : %v", tempErr.Error())
		os.Exit(02)
	}

	http.HandleFunc("/lancement", func(w http.ResponseWriter, r *http.Request) {
		if gameStarted {
			http.Redirect(w, r, "/game", http.StatusSeeOther)
		}
		tmpl.ExecuteTemplate(w, "lancement", nil)
	})

	http.HandleFunc("/lancement/treatment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/lancement", http.StatusSeeOther)
			return
		}

		pseudo = r.FormValue("pseudo")
		difficulty = r.FormValue("difficulty")

		validPseudo := len(pseudo) >= 1 && len(pseudo) <= 32

		validDifficulty := difficulty == "1" || difficulty == "2"

		if !validPseudo || !validDifficulty {
			http.Redirect(w, r, "/lancement/error", http.StatusSeeOther)
			print("err")
			return
		}
		gameStarted = true
	})

	// Gestion des erreurs
	http.HandleFunc("/lancement/error", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "error", nil)
	})

	fmt.Println("Serveur démarré sur http://localhost:8080/lancement")
	http.ListenAndServe(":8080", nil)
}
