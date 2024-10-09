package hangman

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Dictionnaire de correspondance entre les lettres accentuées et non accentuées
var accentMap = map[rune][]rune{
	'e': {'e', 'é', 'è', 'ê', 'ë'},
	'a': {'a', 'à', 'â', 'ä'},
	'u': {'u', 'ù', 'û', 'ü'},
	'o': {'o', 'ô', 'ö'},
	'i': {'i', 'î', 'ï'},
	'c': {'c', 'ç'},
}


func Logic(dictionaryPath string) {
	// Lire le fichier hangman.txt
	hangmanStages, err := readHangmanStages("ASKII Display/hangman.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	// Lire le fichier victory.txt
	victoryMessage, err := readVictoryMessage("ASKII Display/victory.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier de victoire :", err)
		return
	}

	// Lire les mots du dictionnaire en utilisant le chemin passé en argument
	dictionaryWords, err := readDictionaryWords(dictionaryPath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier dictionnaire :", err)
		return
	}

	// Choisir un mot aléatoire à partir du dictionnaire
	rand.Seed(time.Now().UnixNano())
	word := dictionaryWords[rand.Intn(len(dictionaryWords))]

	// Convertir le mot et les blanks en slices de runes pour gérer les caractères accentués
	wordRunes := []rune(word)
	blanks := make([]rune, len(wordRunes))
	for i := range blanks {
		blanks[i] = '_'
	}

	guessedLetters := make(map[rune]struct{})

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
		fmt.Printf("Word: [%s], Lettres déjà proposées: %s\n", string(blanks), getGuessedLetters(guessedLetters))
		fmt.Print("Entrez une lettre: ")

		// Lire l'entrée utilisateur
		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)

		// Traiter l'entrée utilisateur pour retirer les accents
		input = removeAccents(input)

		// Vérification si l'utilisateur n'entre rien
		if len(input) == 0 {
			// Réduire de 3 vies si aucune entrée n'est faite
			lives -= 3
			fmt.Println("Sérieusement !?")
		} else if len(input) > 1 {
			// Si le mot entré est correct
			// Si le mot entré est correct (comparaison sans accents)
			if removeAccents(input) == removeAccents(word) {
				fmt.Println(victoryMessage) // Afficher le message de victoire
				fmt.Printf("Le mot était : %s\n", word)
				break
			} else {
				// Si le mot est incorrect, perdre 2 vies
				lives -= 2
				fmt.Println("Mot incorrect!")
			}
		} else {
			// Vérification si la lettre a déjà été proposée
			if _, exists := guessedLetters[rune(input[0])]; exists {
				fmt.Println("Vous avez déjà proposé cette lettre. Essayez une autre.")
				continue
			}

			// Ajouter la lettre à guessedLetters
			guessedLetters[rune(input[0])] = struct{}{}
			correctGuess := false

			// Comparaison des lettres devinées avec le mot
			for i, wordLetter := range wordRunes {
				if containsAccentMatch(rune(input[0]), wordLetter) {
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

		// Assurer que le nombre de vies ne tombe pas en dessous de zéro
		if lives < 0 {
			lives = 0
		}

		// Afficher l'état du Hangman en fonction des vies restantes
		hangmanIndex := len(hangmanStages) - lives - 1
		if hangmanIndex >= 0 && hangmanIndex < len(hangmanStages) {
			fmt.Println(hangmanStages[hangmanIndex])
		}

		// Vérifier si le joueur a perdu
		if lives <= 0 {
			fmt.Printf("Perdu! Le mot était : %s\n", word)
			break
		}

		// Vérifier si le joueur a gagné (toutes les lettres sont découvertes)
		if string(wordRunes) == string(blanks) {
			fmt.Println(victoryMessage) // Afficher le message de victoire
			fmt.Printf("Le mot était : %s\n", word)
			break
		}

		// Afficher le nombre d'essais restants
		fmt.Printf("%d attempts remaining.\n", lives)
	}
}

// Fonction pour lire les étapes du Hangman à partir d'un fichier
func readHangmanStages(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Diviser le fichier en étapes, chaque étape est séparée par une ligne vide
	stages := strings.Split(string(content), "\n\n")
	return stages, nil
}

// Fonction pour lire le message de victoire à partir d'un fichier
func readVictoryMessage(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// Fonction pour lire les mots du dictionnaire à partir d'un fichier
func readDictionaryWords(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Diviser le fichier en mots, chaque mot étant sur une nouvelle ligne
	words := strings.Split(string(content), "\n")
	return words, nil
}

// Fonction pour retirer les accents des lettres
func removeAccents(input string) string {
	replacer := strings.NewReplacer(
		"é", "e", "è", "e", "ê", "e", "ë", "e",
		"à", "a", "â", "a", "ä", "a",
		"ù", "u", "û", "u", "ü", "u",
		"ô", "o", "ö", "o",
		"î", "i", "ï", "i",
		"ç", "c",
	)
	return replacer.Replace(input)
}

// Fonction pour extraire les lettres déjà proposées et les afficher sous forme de chaîne
func getGuessedLetters(m map[rune]struct{}) string {
	letters := make([]string, 0, len(m))
	for k := range m {
		letters = append(letters, string(k))
	}
	return strings.Join(letters, ", ")
}

// Fonction pour vérifier si une lettre non accentuée correspond à une lettre accentuée
func containsAccentMatch(input, wordLetter rune) bool {
	// Si la lettre sans accent existe dans le dictionnaire des accents
	if accentedLetters, exists := accentMap[input]; exists {
		// Vérifier si la lettre du mot fait partie des lettres accentuées correspondantes
		for _, accentedLetter := range accentedLetters {
			if wordLetter == accentedLetter {
				return true
			}
		}
	}
	// Comparer directement si ce n'est pas une lettre accentuée
	return input == wordLetter
}
