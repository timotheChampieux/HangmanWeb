package affichage

import "fmt"

func AfficherPendu(etape int) {
	switch etape {
	case 7:
		fmt.Println(`
  +---+
  |   |
      |
      |
      |
      |
=========
		`)
	case 6:
		fmt.Println(`
  +---+
  |   |
  O   |
      |
      |
      |
=========
		`)
	case 5:
		fmt.Println(`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========
		`)
	case 4:
		fmt.Println(`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========
		`)
	case 3:
		fmt.Println(`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========
		`)
	case 2:
		fmt.Println(`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========
		`)
	case 1:
		fmt.Println(`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========
		`)
	case -1:
		fmt.Println(`
  +---+
      |
      |
  O/  |
 /|   |
 / \  |
=========
		`)
	default:
		fmt.Println(`
  +---+
      |
      |
      |
      |
      |
=========
                      `)
	}
}
