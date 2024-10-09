package hangman

import (
	"bufio"
	"io/ioutil"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)


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
		fmt.Println("Erreur lors de la lecture du fichier de victoire :", err)
		return
	}
	// Choisir un mot aléatoire à partir du dictionnaire
	rand.Seed(time.Now().UnixNano()) // Initialise le générateur de nombres aléatoires
	word, err := ReadLineFromWords(rand.Intn(2))

	// Fonction pour simplifier les caractères accentués
	word = removeAccents(strings.ToLower(word))

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
			PrintHangman(lives)
			fmt.Println("Sérieusement !?")
		} else if len(input) > 1 {
			// Si le mot entré est correct
			if input == word {
				fmt.Println(victoryMessage) // Afficher le message de victoire
				fmt.Printf("Le mot était : %s\n", word)
				break
			} else {
				// Si le mot est incorrect, perdre 2 vies
				lives -= 2
				PrintHangman(lives)
				fmt.Println("Mot incorrect!")
			}
		} else {

			// Vérification si la lettre a déjà été proposée (avant de l'ajouter à guessedLetters)
			if _, exists := guessedLetters[rune(input[0])]; exists {
				fmt.Println("Vous avez déjà proposé cette lettre. Essayez une autre.")
				continue
			}

			// Ajouter la lettre à guessedLetters après avoir vérifié si elle a déjà été proposée
			guessedLetters[rune(input[0])] = struct{}{}
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
			fmt.Println(LoseMessage)
			break
		}

		// Vérifier si le joueur a gagné (toutes les lettres sont découvertes)
		if string(wordRunes) == string(blanks) {
			fmt.Printf("Le mot était : %s\n", word)
			fmt.Println(victoryMessage) // Afficher le message de victoire
			fmt.Printf("Le mot était : %s\n", word)
			break
		}

		// Afficher le nombre d'essais restants
		fmt.Printf("%d attempts remaining.\n", lives)
	}
}

func PrintHangman(n int){
	// Ouvrir le fichier
	file, err := os.Open("ASKIIDisplay/hangman.txt")
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
		filename = "WordSelection/dictionnaire1.txt"
	} else {
		filename = "WordSelection/dictionnaire2.txt"
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
// Fonction pour lire le message de victoire à partir d'un fichier
func readMessage(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(content), nil
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
