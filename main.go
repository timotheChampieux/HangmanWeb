package main

import (
	"fmt"
	"hangmanWeb/game/affichage"
	"hangmanWeb/game/jeu"
	recupmot "hangmanWeb/game/recupMot"
	"math/rand"
	"net/http"
	"os"
	"text/template"
)

var (
	gameStarted       bool = false
	essaie            int  = 7
	reussitte         int
	motCacher         string
	motMasque         []string
	motAleatoire      string
	count             int = 0
	lettreDejaPropose []string
	lettrePropose     string
	win               bool
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
		User2             user
		Mot               string
		MotCacher         string
		Essaie            int // nombre de vie
		Reussitte         int // nombre de lettre a trouver
		LettreDejaPropose string
		Message           string
	}
	motAleatoire = recupmot.Recup("./game/recupMot/mot.txt")
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {

		if count == 0 {
			motMasque = affichage.Debut(motAleatoire)
			for i := 0; i < len(motMasque); i++ {
				motMasque[i] = "_ "
			}
		}

		if user1.Difficulty == "difficile" && count == 0 {
			count++
			reussitte = len(motMasque) - 1
			motMasque[0] = string(motAleatoire[0])
		} else if user1.Difficulty == "facile" && count == 0 {
			reussitte = len(motMasque) - 2
			count++
			motMasque[0] = string(motAleatoire[0])
			max := len(motMasque)
			indexLettreAletoire := rand.Intn(max - 1)
			motMasque[indexLettreAletoire+1] = string(motAleatoire[indexLettreAletoire+1])
		}

		motCacher = ""
		for i := 0; i < len(motMasque); i++ {
			motCacher += motMasque[i]
		}

		data := data{user1, motAleatoire, motCacher, essaie, reussitte, lettrePropose, r.FormValue("message")}

		if essaie <= 0 {
			win = false
			http.Redirect(w, r, "/fin", http.StatusSeeOther)
		} else if reussitte == 0 {
			win = true
			http.Redirect(w, r, "/fin", http.StatusSeeOther)
		}
		fmt.Println(r.FormValue("message"))
		tmpl.ExecuteTemplate(w, "game", data)

	})

	//-------------------------------------------------------  traitement page de jeu  ----------------------------------------------

	http.HandleFunc("/game/treatment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/game", http.StatusSeeOther)
			return
		}
		message := ""
		lettre := r.FormValue("lettre")

		if jeu.ElementDansSlice(lettre, lettreDejaPropose) {
			message = "Vous avez déja entré cette lettre !"
		}

		if len(lettre) == 1 && !jeu.ElementDansSlice(lettre, lettreDejaPropose) {
			var count int
			for i := 0; i < len(motAleatoire); i++ {
				if lettre == string(motAleatoire[i]) && lettre != motMasque[i] {
					motMasque[i] = lettre
					reussitte--
					message = "Bien vu"
					count++
				}
			}
			if count == 0 {
				message = "Raté"
				essaie--
			}
		} else if !jeu.ElementDansSlice(lettre, lettreDejaPropose) {
			if lettre == motAleatoire {
				for i := 0; i < len(motAleatoire); i++ {
					motMasque[i] = string(lettre[i])
				}
				message = "GGGGGGGGG"
				reussitte = 0
			} else {
				message = "Bien tenté"
				essaie -= 2
			}

		}

		if !jeu.ElementDansSlice(lettre, lettreDejaPropose) {
			lettreDejaPropose = append(lettreDejaPropose, lettre)
			lettrePropose = ""
			for i := 0; i < len(lettreDejaPropose); i++ {
				lettrePropose += " "
				lettrePropose += lettreDejaPropose[i]
			}
		}

		motCacher = ""
		for i := 0; i < len(motMasque); i++ {
			motCacher += motMasque[i]
		}
		http.Redirect(w, r, "/game?message="+message, http.StatusSeeOther)
	})
	//-------------------------------------------------------  traitement page de fin ----------------------------------------------
	type dataEnd struct {
		Win bool
		Mot string
	}
	http.HandleFunc("/fin", func(w http.ResponseWriter, r *http.Request) {
		/*if  {
			http.Redirect(w, r, "/game", http.StatusSeeOther)
			return
		}*/
		data := dataEnd{Win: win, Mot: motAleatoire}
		tmpl.ExecuteTemplate(w, "fin", data)
	})

	//------------------------------------------------------------------------------------------------------------------------------
	fmt.Println("Serveur démarré sur http://localhost:8000/lancement")
	http.ListenAndServe(":8000", nil)
}
