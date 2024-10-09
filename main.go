package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"io/ioutil"

	hg "hangman/functions"
)

// Fonction pour lire les étapes du Hangman à partir d'un fichier
func readHangmanStages(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// Diviser le fichier en étapes, chaque étape étant séparée par une ligne vide
	stages := strings.Split(string(content), "\n\n")
	return stages, nil
}

func main() {
	// Définir les flags pour hardmode, help et startWith
	hardmode := flag.Bool("hardmode", false, "Activer le mode difficile avec des mots plus complexes.")
	help := flag.Bool("help", false, "Afficher l'aide.")
	startWithSave := flag.Bool("startWith", false, "Démarrer avec la sauvegarde dans save/save.txt.")
	flag.Parse()

	// Si l'option help est activée, afficher l'aide et arrêter l'exécution
	if *help {
		displayHelp()
		return
	}

	// Dictionnaire par défaut
	dictionaryPath := "Word Selection/dictionnaire.txt"

	// Si l'option hardmode est activée, changer le dictionnaire
	if *hardmode {
		fmt.Println("Mode difficile activé !")
		dictionaryPath = "Word Selection/dictionnaire_hardmode.txt"
	} else {
		fmt.Println("Mode normal activé.")
	}

	// Lire les étapes du Hangman
	hangmanStages, err := readHangmanStages("ASKII Display/hangman.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture des étapes du Hangman :", err)
		return
	}

	// Si l'option startWithSave est activée, charger la sauvegarde
	if *startWithSave {
		fmt.Println("Démarrage avec la sauvegarde...")
		// Appeler la logique du jeu en utilisant la sauvegarde
		err := hg.StartFromSave("save/save.txt", hangmanStages)
		if err != nil {
			fmt.Println("Erreur lors du chargement de la sauvegarde:", err)
			return
		}
	} else {
		// Démarrer une nouvelle partie avec le bon dictionnaire
		hg.Logic(dictionaryPath)
	}
}

// Fonction pour afficher l'aide
func displayHelp() {
	fmt.Println("Usage:")
	fmt.Println("  go run . [options]")
	fmt.Println("Options:")
	fmt.Println("  -hardmode    Activer le mode difficile avec des mots plus complexes.")
	fmt.Println("  -startWith   Démarrer avec la sauvegarde dans save/save.txt.")
	fmt.Println("  -h, -help    Afficher cette aide.")
	os.Exit(0)
}
