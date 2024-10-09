package main

import (
	"flag"
	"fmt"
	"os"

	hg "hangman/functions"
)

func main() {
	// Définir les flags pour hardmode et help
	hardmode := flag.Bool("hardmode", false, "Activer le mode difficile avec des mots plus complexes.")
	help := flag.Bool("help", false, "Afficher l'aide.")
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

	// Appeler la logique du jeu en passant le bon dictionnaire
	hg.Logic(dictionaryPath)
}

// Fonction pour afficher l'aide
func displayHelp() {
	fmt.Println("Usage:")
	fmt.Println("  go run . [options]")
	fmt.Println("Options:")
	fmt.Println("  -hardmode    Activer le mode difficile avec des mots plus complexes.")
	fmt.Println("  -h, -help    Afficher cette aide.")
	os.Exit(0)
}
