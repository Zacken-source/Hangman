package hangman

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Dictionnaire pour la correspondance entre les lettres accentuées et non accentuées
var accentMap = map[rune][]rune{
	'e': {'e', 'é', 'è', 'ê', 'ë'},
	'a': {'a', 'à', 'â', 'ä'},
	'u': {'u', 'ù', 'û', 'ü'},
	'o': {'o', 'ô', 'ö'},
	'i': {'i', 'î', 'ï'},
	'c': {'c', 'ç'},
}

func Logic(dictionaryPath string) {
	// Lire les étapes du pendu à partir du fichier
	hangmanStages, err := readHangmanStages("ASCIIDisplay/hangman.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	// Lire le message de défaite
	loseMessage, err := readLoseMessage("ASCIIDisplay/lose.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier de lose :", err)
		return
	}

	// Lire le message de victoire
	victoryMessage, err := readVictoryMessage("ASCIIDisplay/victory.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier de victoire :", err)
		return
	}

	// Lire les mots du dictionnaire en utilisant le chemin fourni
	dictionaryWords, err := readDictionaryWords(dictionaryPath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier dictionnaire :", err)
		return
	}

	// Sélectionner un mot aléatoire à partir du dictionnaire
	rand.Seed(time.Now().UnixNano())
	word := dictionaryWords[rand.Intn(len(dictionaryWords))]

	// Convertir le mot en une slice de runes pour gérer les caractères accentués
	wordRunes := []rune(word)
	blanks := make([]rune, len(wordRunes))
	for i := range blanks {
		blanks[i] = '_'
	}

	// Stocker les lettres devinées dans une map
	guessedLetters := make(map[rune]struct{})

	// Calculer le nombre de lettres à révéler au début du jeu
	numLettersToReveal := (len(wordRunes) / 2) - 1
	if numLettersToReveal < 0 {
		numLettersToReveal = 0 // Assurer qu'il n'y ait pas de nombre négatif
	}

	// Sélectionner des indices aléatoires pour révéler les lettres
	revealedIndices := make(map[int]struct{}) // Utiliser un map pour éviter les doublons
	for len(revealedIndices) < numLettersToReveal {
		index := rand.Intn(len(wordRunes))
		revealedIndices[index] = struct{}{}
	}

	// Révéler les lettres aux indices choisis
	for index := range revealedIndices {
		blanks[index] = wordRunes[index]
	}

	lives := 10 // Nombre de vies du joueur

	// Boucle de jeu principale
	for {
		// Afficher l'état actuel du mot et des lettres déjà devinées
		fmt.Printf("Word: [%s], Lettres déjà proposées: %s\n", string(blanks), getGuessedLetters(guessedLetters))
		fmt.Print("Entrez une lettre: ")

		// Lire l'entrée utilisateur
		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)

		// Retirer les accents de l'entrée utilisateur
		input = removeAccents(input)

		// Gérer le cas où aucune entrée n'est fournie
		if len(input) == 0 {
			lives -= 3 // Réduire de 3 vies si aucune lettre n'est saisie
			fmt.Println("Sérieusement !?")
		} else if len(input) > 1 {
			// Si un mot complet est saisi et correspond au mot à deviner
			if removeAccents(input) == removeAccents(word) {
				fmt.Println(victoryMessage) // Afficher le message de victoire
				fmt.Printf("Le mot était : %s\n", word)
				break
			} else {
				lives -= 2 // Réduire de 2 vies si le mot saisi est incorrect
				fmt.Println("Mot incorrect!")
			}
		} else {
			// Vérifier si la lettre a déjà été devinée
			if _, exists := guessedLetters[rune(input[0])]; exists {
				fmt.Println("Cette lettre a déjà été proposée.")
				continue
			}

			// Ajouter la lettre aux lettres devinées
			guessedLetters[rune(input[0])] = struct{}{}
			correctGuess := false

			// Vérifier si la lettre devinée correspond à une lettre du mot
			for i, wordLetter := range wordRunes {
				if containsAccentMatch(rune(input[0]), wordLetter) {
					blanks[i] = wordLetter // Révéler la lettre dans les blancs
					correctGuess = true
					fmt.Println("Bon choix")
				}
			}

			// Si la lettre est incorrecte, réduire les vies
			if !correctGuess {
				lives--
				fmt.Println("Lettre incorrecte!")
			}
		}

		// Empêcher le nombre de vies de descendre en dessous de zéro
		if lives < 0 {
			lives = 0
		}

		// Afficher l'état du pendu en fonction des vies restantes
		hangmanIndex := len(hangmanStages) - lives
		if hangmanIndex >= 0 && hangmanIndex < len(hangmanStages) {
			fmt.Println(hangmanStages[hangmanIndex])
		}

		// Vérifier si toutes les vies sont épuisées
		if lives <= 0 {
			fmt.Println(loseMessage) // Afficher le message de défaite
			fmt.Printf("Perdu! Le mot était : %s\n", word)
			break
		}

		// Vérifier si toutes les lettres ont été devinées
		if string(wordRunes) == string(blanks) {
			fmt.Println(victoryMessage) // Afficher le message de victoire
			fmt.Printf("Le mot était : %s\n", word)
			break
		}

		// Afficher le nombre d'essais restants
		fmt.Printf("%d essais restants.\n", lives)
	}
}

// Fonction pour lire les étapes du pendu depuis un fichier
func readHangmanStages(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Diviser le fichier en étapes, chaque étape est séparée par une ligne vide
	stages := strings.Split(string(content), "\n\n")
	return stages, nil
}

// Fonction pour lire le message de défaite depuis un fichier
func readLoseMessage(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// Fonction pour lire le message de victoire depuis un fichier
func readVictoryMessage(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// Fonction pour lire les mots du dictionnaire depuis un fichier
func readDictionaryWords(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Diviser le contenu du fichier en mots, chaque mot étant sur une nouvelle ligne
	words := strings.Split(string(content), "\n")
	return words, nil
}

// Fonction pour retirer les accents des lettres dans une chaîne
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

// Fonction pour afficher les lettres déjà devinées sous forme de chaîne
func getGuessedLetters(m map[rune]struct{}) string {
	letters := make([]string, 0, len(m))
	for k := range m {
		letters = append(letters, string(k))
	}
	return strings.Join(letters, ", ")
}

// Fonction pour vérifier si une lettre non accentuée correspond à une lettre accentuée
func containsAccentMatch(input, wordLetter rune) bool {
	// Vérifier si la lettre sans accent existe dans le dictionnaire des accents
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
