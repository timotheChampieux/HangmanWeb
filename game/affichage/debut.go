package affichage

import (
	"fmt"
	"math/rand"
	"time"
)

func Debut(mot string) []string {
	motMasque := []string{}
	for i := 0; i < len(mot); i++ {
		motMasque = append(motMasque, "_")
	}
	if len(mot) <= 5 {
		rand.Seed(time.Now().UnixNano())
		max := len(motMasque)
		indexLettreAletoire := rand.Intn(max)
		motMasque[indexLettreAletoire] = string(mot[indexLettreAletoire])
	} else if len(mot) > 5 {
		max := len(motMasque)
		indexLettreAletoire := rand.Intn(max - 1)
		motMasque[indexLettreAletoire+1] = string(mot[indexLettreAletoire+1])
		motMasque[0] = string(mot[0])
	}
	for i := 0; i < len(motMasque); i++ {
		fmt.Printf("%v ", motMasque[i])
	}
	return motMasque
}
