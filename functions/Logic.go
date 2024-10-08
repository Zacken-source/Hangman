package hangman

import (
	"encoding/json"
	"fmt"
	"bufio"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type GameState struct {
    Word           string
    Blanks         []rune
    GuessedLetters map[rune]struct{}
    Lives          int
}

var accentMap = map[rune][]rune{
    'a': {'à', 'â', 'ä'},
    'e': {'é', 'è', 'ê', 'ë'},
    'i': {'î', 'ï'},
    'o': {'ô', 'ö'},
    'u': {'ù', 'û', 'ü'},
    'c': {'ç'},
}

func readHangmanStages(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var stages []string
    scanner := bufio.NewScanner(file)
    var currentStage string

    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            stages = append(stages, currentStage)
            currentStage = ""
        } else {
            currentStage += line + "\n"
        }
    }
    if currentStage != "" {
        stages = append(stages, currentStage)
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return stages, nil
}

func readDictionaryWords(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var words []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        words = append(words, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return words, nil
}

func readVictoryMessage(filename string) (string, error) {
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        return "", err
    }
    return string(content), nil
}

// Fonction pour lire le message de défaite à partir d'un fichier
func readLoseMessage(filename string) (string, error) {
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        return "", err
    }
    return string(content), nil
}

func Logic(dictionaryPath string) {
	// Lire le fichier hangman.txt
	hangmanStages, err := readHangmanStages("ASKIIDisplay/hangman.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	// Lire le fichier victory.txt
	victoryMessage, err := readVictoryMessage("ASKIIDisplay/victory.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier de victoire :", err)
		return
	}

	// Lire le fichier lose.txt (message de défaite)
	loseMessage, err := readLoseMessage("ASKIIDisplay/lose.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier de défaite :", err)
		return
	}

	// Lire les mots du dictionnaire en utilisant le chemin passé en argument
	dictionaryWords, err := readDictionaryWords(dictionaryPath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier du dictionnaire :", err)
		return
	}

	// Choisir un mot aléatoire à partir du dictionnaire
	rand.Seed(time.Now().UnixNano())
	word := dictionaryWords[rand.Intn(len(dictionaryWords))]

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
			PrintHangman(lives, hangmanStages) // Passer hangmanStages ici
			fmt.Println("Sérieusement !?")
		} else if len(input) > 1 {
			// Si le mot entré est correct (comparaison sans accents)
			if removeAccents(input) == removeAccents(word) {
				fmt.Println(victoryMessage) // Afficher le message de victoire
				fmt.Printf("Le mot était : %s\n", word)
				break
			} else {
				// Si le mot est incorrect, perdre 2 vies
				lives -= 2
				PrintHangman(lives, hangmanStages) // Passer hangmanStages ici
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
					PrintHangman(lives, hangmanStages) // Passer hangmanStages ici
				}
			}

			if !correctGuess {
				lives-- // Diminuer les vies si la lettre n'est pas correcte
				fmt.Println("Lettre incorrecte!")
				PrintHangman(lives, hangmanStages) // Passer hangmanStages ici
			}
		}

		if lives <= 0 {
			fmt.Println(loseMessage)  // Utiliser le message de défaite lu depuis le fichier
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

func PrintHangman(lives int, hangmanStages []string) {
    // Calculer l'indice correspondant aux vies restantes
    index := len(hangmanStages) - lives - 1

    // S'assurer que l'indice est dans les limites
    if index >= 0 && index < len(hangmanStages) {
        fmt.Println(hangmanStages[index])
    } else {
        fmt.Println("Erreur: index des étapes du pendu hors limites.")
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

// Fonction pour sauvegarder l'état du jeu dans un fichier
func saveGame(gameState GameState) error {
	// Sérialiser l'état du jeu en JSON
	data, err := json.MarshalIndent(gameState, "", "  ")
	if err != nil {
		return err
	}

	// Écrire les données dans le fichier de sauvegarde
	return ioutil.WriteFile("save/save.txt", data, 0644)
}

// Fonction pour charger et démarrer le jeu depuis une sauvegarde
func StartFromSave(saveFile string, hangmanStages []string) error {
    // Vérifier si le fichier de sauvegarde existe
    if _, err := os.Stat(saveFile); os.IsNotExist(err) {
        return fmt.Errorf("le fichier de sauvegarde n'existe pas")
    }

    // Lire le fichier de sauvegarde
    data, err := ioutil.ReadFile(saveFile)
    if err != nil {
        return err
    }

    // Décoder le contenu du fichier en un objet GameState
    var gameState GameState
    err = json.Unmarshal(data, &gameState)
    if err != nil {
        return err
    }

    // Continuer le jeu à partir de cet état
    return continueGame(gameState, hangmanStages)
}

func continueGame(gameState GameState, hangmanStages []string) error {
    word := gameState.Word
    blanks := gameState.Blanks
    guessedLetters := gameState.GuessedLetters
    lives := gameState.Lives

    // Convertir le mot en runes pour gérer les caractères accentués
    wordRunes := []rune(word)

    // Reprendre la logique du jeu à partir de cet état
    for {
        // Afficher les blanks avec les lettres devinées
        fmt.Printf("Mot: [%s], Lettres déjà proposées: %s\n", string(blanks), getGuessedLetters(guessedLetters))
        fmt.Print("Entrez une lettre ou 'stop' pour sauvegarder: ")

        // Lire l'entrée utilisateur
        var input string
        fmt.Scanln(&input)
        input = strings.ToLower(strings.TrimSpace(input))

        // Gérer le cas où l'utilisateur entre "stop"
        if input == "stop" {
            // Sauvegarder l'état du jeu
            err := saveGame(GameState{
                Word:           word,
                Blanks:         blanks,
                GuessedLetters: guessedLetters,
                Lives:          lives,
            })
            if err != nil {
                return fmt.Errorf("erreur lors de la sauvegarde: %v", err)
            }
            fmt.Println("Jeu sauvegardé dans save/save.txt")
            break
        }

        // Vérifier la validité de l'entrée
        if len(input) == 0 {
            lives -= 3
            fmt.Println("Sérieusement !? Vous devez entrer une lettre.")
            continue
        } else if len(input) > 1 {
            // Vérification si l'utilisateur entre un mot
            if removeAccents(input) == removeAccents(word) {
                fmt.Println("Bravo! Vous avez gagné !")
                fmt.Printf("Le mot était : %s\n", word)
                break
            } else {
                lives -= 2
                fmt.Println("Mot incorrect!")
            }
            continue
        }

        // Traitement d'une seule lettre
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
                fmt.Println("Bon choix")
            }
        }

        if !correctGuess {
            lives-- // Diminuer les vies si la lettre n'est pas correcte
            fmt.Println("Lettre incorrecte!")
        }

        // Vérifier le nombre de vies restantes
        if lives <= 0 {
            fmt.Printf("Perdu! Le mot était : %s\n", word)
            break
        }

        // Vérifier si le joueur a gagné
        if string(wordRunes) == string(blanks) {
            fmt.Println("Vous avez gagné !")
            break
        }

        // Afficher l'état du Hangman en fonction des vies restantes
        hangmanIndex := len(hangmanStages) - lives - 1
        if hangmanIndex >= 0 && hangmanIndex < len(hangmanStages) {
            fmt.Println(hangmanStages[hangmanIndex])
        }

        // Afficher le nombre d'essais restants
        fmt.Printf("%d tentatives restantes.\n", lives)
    }
    return nil
}
