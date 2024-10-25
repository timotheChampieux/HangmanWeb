package main

import (
	"fmt"
	recupmot "hangmanWeb/game/recupMot"
	"net/http"
	"os"
	"text/template"
)

var (
	gameStarted bool = false
)

type user struct {
	Pseudo     string
	Difficulty string
}

func main() {
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	tmpl, tempErr := template.ParseGlob("./templates/*.html")
	if tempErr != nil {
		fmt.Printf("erreur avec le temp : %v", tempErr.Error())
		os.Exit(02)
	}
	//------------------------------------------------------- 1ere page ----------------------------------------------
	http.HandleFunc("/lancement", func(w http.ResponseWriter, r *http.Request) {
		if gameStarted {
			http.Redirect(w, r, "/game", http.StatusSeeOther)
		}
		tmpl.ExecuteTemplate(w, "lancement", nil)
	})
	user1 := user{}
	http.HandleFunc("/lancement/treatment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/lancement", http.StatusSeeOther)
			return
		}

		user1.Pseudo = r.FormValue("pseudo")
		if r.FormValue("difficulty") == "1" {
			user1.Difficulty = "facile"
		} else {
			user1.Difficulty = "difficile"
		}

		validPseudo := len(user1.Pseudo) >= 1 && len(user1.Pseudo) <= 32

		validDifficulty := r.FormValue("difficulty") == "1" || r.FormValue("difficulty") == "2"

		if !validPseudo || !validDifficulty {
			http.Redirect(w, r, "/lancement/error", http.StatusSeeOther)
			print("err")
			return
		}
		gameStarted = true
		if gameStarted {
			http.Redirect(w, r, "/game", http.StatusSeeOther)
		}
	})
	http.HandleFunc("/lancement/error", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "error", nil)
	})
	//------------------------------------------------------- page de jeu  ----------------------------------------------
	type data struct {
		User2 user
		Mot   string
	}
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {

		motAleatoire := recupmot.Recup("./game/recupMot/mot.txt")
		data := data{user1, motAleatoire}
		tmpl.ExecuteTemplate(w, "game", data)

	})

	//------------------------------------------------------- page de jeu  ----------------------------------------------
	fmt.Println("Serveur démarré sur http://localhost:8080/lancement")
	http.ListenAndServe(":8080", nil)
}
