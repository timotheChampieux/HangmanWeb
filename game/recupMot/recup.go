package recupmot

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

func Recup(fichier string) string {

	file, err := os.Open(fichier)
	if err != nil {
		log.Fatalf("Erreur : Impossible d'ouvrir le fichier mot.txt. Vérifiez que le fichier existe et est lisible. Détails : %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var mots []string

	for scanner.Scan() {
		ligne := scanner.Text()
		mots = append(mots, ligne)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier mot.txt : %v", err)
	}

	if len(mots) == 0 {
		log.Fatalf("Erreur : Le fichier mot.txt est vide.")
	}

	rand.Seed(time.Now().UnixNano())
	indexMotAleatoire := rand.Intn(len(mots))
	return mots[indexMotAleatoire]
}
