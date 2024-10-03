package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)



func main() {
	rand.Seed(time.Now().UnixNano()) // Initialise le générateur de nombres aléatoires
	word, err := ReadLineFromWords(rand.Intn(2))
	if err != nil {
		fmt.Println("Erreur lors de la lecture du mot :", err)
		return
	}

	wordRunes := []rune(word)
	blanks := make([]rune, len(wordRunes))
	for i := range blanks {
		blanks[i] = '_'
	}

		// Calculer le nombre de lettres à révéler
		numLettersToReveal := (len(wordRunes) / 2) - 1
		if numLettersToReveal < 0 {
			numLettersToReveal = 0 // Assurez-vous qu'il n'y ait pas de nombre négatif
		}
	
		// Choisir des indices aléatoires pour révéler les lettres
		revealedIndices := make(map[int]struct{}) // Utiliser un map pour éviter les doublons
		for len(revealedIndices) < numLettersToReveal {
			index := rand.Intn(len(wordRunes))
			revealedIndices[index] = struct{}{}
		}
	
		// Remplacer les underscores par les lettres révélées
		for index := range revealedIndices {
			blanks[index] = wordRunes[index]
		}

	lives := 10

	// Jeu
	for {
		// Afficher les blanks avec les lettres devinées
		fmt.Printf("word: %s Letter: ", string(blanks))

		// Lire l'entrée utilisateur
		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)

		// Vérification si l'utilisateur n'entre rien
		if len(input) == 0 {
			// Réduire de 3 vies si aucune entrée n'est faite
			lives -= 3
			PrintHangman(lives)
			fmt.Println("Sérieusement !?")
		// Vérification si l'utilisateur essaie un mot entier
		} else if len(input) > 1 {
			// Si le mot entré est correct
			if input == word {
				fmt.Printf("Gagné! Le mot était : %s\n", word)
				break
			} else {
				// Si le mot est incorrect, perdre 2 vies
				lives -= 2
				PrintHangman(lives)
				fmt.Println("Mot incorrect!")
			}
		} else {
			// Si l'utilisateur entre une seule lettre, traiter chaque lettre
			correctGuess := false

			// Comparaison des lettres devinées avec le mot
			for i, wordLetter := range wordRunes {
				if rune(input[0]) == wordLetter {
					blanks[i] = wordLetter // Remplacer le "_" par la lettre correcte
					correctGuess = true
					fmt.Println("bon choix")
					PrintHangman(lives)
				}
			}

			if !correctGuess {
				lives-- // Diminuer les vies si la lettre n'est pas correcte
				fmt.Println("Lettre incorrecte!")
				PrintHangman(lives)
			}
		}

		// Vérifier si le joueur a perdu
		if lives <= 0 {
			fmt.Printf("Perdu! Le mot était : %s\n", word)
			break
		}

		// Vérifier si le joueur a gagné (toutes les lettres sont découvertes)
		if string(wordRunes) == string(blanks) {
			fmt.Printf("Gagné! Le mot était : %s\n", word)
			break
		}

		// Afficher le nombre d'essais restants
		fmt.Printf("%d attempts remaining.\n", lives)
	}

}
func PrintHangman(n int){
	// Ouvrir le fichier
	file, err := os.Open("hangman.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer file.Close()

	// Créer un scanner pour lire le fichier ligne par ligne
	scanner := bufio.NewScanner(file)
	var startLine, endLine int
	if n == 9 {
		startLine, endLine = 0, 7
	} else if n == 8 {
		startLine, endLine = 7, 14
	} else if n == 7 {
		startLine, endLine = 14, 21
	} else if n == 6 {
		startLine, endLine = 21, 28
	} else if n == 5 {
		startLine, endLine = 28, 35
	} else if n == 4 {
		startLine, endLine = 35, 42
	} else if n == 3 {
		startLine, endLine = 42, 49
	} else if n == 2 {
		startLine, endLine = 49, 56
	} else if n == 1 {
		startLine, endLine = 56, 63
	} else {
		fmt.Println("Valeur de n non valide.")
		return
	}

	// Lire et afficher les lignes appropriées
	lineCount := 0
	for scanner.Scan() {
		if lineCount >= startLine && lineCount < endLine {
			fmt.Println(scanner.Text())
		}
		lineCount++
		if lineCount >= endLine {
			break
		}
	}
}

func ReadLineFromWords(n int) (string, error) {
	var filename string
	if n == 0 {
		filename = "words.txt"
	} else {
		filename = "words2.txt"
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Compter le nombre de lignes
	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}
	
	// Vérifiez s'il y a eu des erreurs lors de la lecture du fichier
	if err := scanner.Err(); err != nil {
		return "", err
	}

	if lineCount == 0 {
		return "", fmt.Errorf("le fichier est vide")
	}

	// Réinitialiser le scanner pour relire le fichier
	file.Seek(0, 0) // Revenir au début du fichier
	scanner = bufio.NewScanner(file)

	// Générer un nombre aléatoire valide
	lineNumber := rand.Intn(lineCount) + 1 // +1 pour que ça commence à 1

	currentLine := 1
	for scanner.Scan() {
		if currentLine == lineNumber {
			return scanner.Text(), nil
		}
		currentLine++
	}

	return "", fmt.Errorf("erreur inconnue")
}
