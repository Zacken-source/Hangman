package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	listeLVL1 := []string{"abcdefghijklmnopqrstuvwxyzéèàùâ", "test"}
	word := listeLVL1[rand.Intn(len(listeLVL1))]

	// Convertir le mot et les blanks en slices de runes pour gérer les caractères accentués
	wordRunes := []rune(word)
	blanks := make([]rune, len(wordRunes))
	for i := range blanks {
		blanks[i] = '_'
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
				}
			}

			if !correctGuess {
				lives-- // Diminuer les vies si la lettre n'est pas correcte
				fmt.Println("Lettre incorrecte!")
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
