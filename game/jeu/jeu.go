package jeu

import (
	"fmt"
	"hangmanWeb/game/affichage"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func ClearScreen() {
	switch runtime.GOOS {
	case "windows": // Commande pour effacer ecran Windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default: //ca efface sur linux
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func demanderElement(slice *[]string) string {
	fmt.Print("\n\nElement déja choisi :\n ")
	for _, char := range *slice {
		fmt.Print(char, "  ")
	}
	for {
		fmt.Print("\n\n\033[32mEntrez une lettre ou un mot:\033[0m \n")
		var lettre string
		fmt.Scanln(&lettre)

		if !elementDansSlice(lettre, *slice) {
			*slice = append(*slice, lettre)
			return lettre
		} else {
			time.Sleep(1 * time.Second)
			fmt.Println("\n\033[31mVous avez déja choisi cette lettre ou ce mot veuillez ressayer.\033[0m")
		}
	}
}

func elementDansSlice(element string, slice []string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func Jeu(mot string, motMasque []string) {
	essaie := 8
	lettreDejaPropose := []string{}
	var reussitte int

	if len(motMasque) <= 5 {
		reussitte = len(motMasque) - 1
	} else {
		reussitte = len(motMasque) - 2
	}

	for essaie > 1 && reussitte > 0 {

		//time.Sleep(2 * time.Second)
		ClearScreen()
		affichage.AfficherPendu(essaie)
		fmt.Printf("\033[31mNombre de vie : %v\n\n\033[0m", essaie-1)
		for i := 0; i < len(motMasque); i++ {
			fmt.Printf("%v ", motMasque[i])
		}
		lettre := demanderElement(&lettreDejaPropose)

		if len(lettre) == 1 {
			var count int
			for i := 0; i < len(mot); i++ {
				if lettre == string(mot[i]) && lettre != motMasque[i] {
					motMasque[i] = lettre
					reussitte--
					count++
				}
			}
			if count == 0 {
				essaie--
				fmt.Printf("\n\033[31mLa lettre que vous avez entré n'est pas contenue dans le mot.\033[0m \n")
			}

		} else {
			if lettre == mot {
				for i := 0; i < len(mot); i++ {
					motMasque[i] = string(lettre[i])
				}
				reussitte = 0
			} else {
				essaie -= 2
				fmt.Printf("\n\033[31mCe n'est pas le bon mot.\033[0m \n")

			}

		}
	}
	if reussitte <= 0 {

		affichage.AfficherPendu(-1)
		time.Sleep(1 * time.Second)
		fmt.Printf("\n\033[32mVous avez gagné !! Le mot etait bien %v.\033[0m", mot)

	} else {

		affichage.AfficherPendu(1)
		time.Sleep(1 * time.Second)
		fmt.Printf("\n\033[31mVous avez perdu... Le mot etait %v.\033[0m", mot)
	}
}
