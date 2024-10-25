package main

import (
	"hangman/affichage"
	"hangman/jeu"
	recupmot "hangman/recupMot"
)

func main() {
	motAleatoire := recupmot.Recup(".\\recupMot\\mot.txt")
	motMasqer := affichage.Debut(motAleatoire)
	jeu.Jeu(motAleatoire, motMasqer)
}
