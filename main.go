package main

import (
	"flag"
	"fmt"
	"os"

	hg "hangman/functions"
)

func main() {
	// Définir les options pour activer le mode difficile ou afficher l'aide
	hardmode := flag.Bool("hardmode", false, "Activer le mode difficile avec des mots plus complexes.")
	help := flag.Bool("help", false, "Afficher l'aide.")
	flag.Parse()

	// Vérifier si l'option d'aide est activée, afficher l'aide et arrêter l'exécution si nécessaire
	if *help {
		displayHelp()
		return
	}

	// Chemin par défaut du dictionnaire
	dictionaryPath := "Word Selection/dictionnaire.txt"

	// Modifier le dictionnaire si le mode difficile est activé
	if *hardmode {
		fmt.Println("Mode difficile activé !")
		dictionaryPath = "Word Selection/dictionnaire_hardmode.txt"
	} else {
		fmt.Println("Mode normal activé.")
	}

	// Lancer la logique du jeu avec le dictionnaire sélectionné
	hg.Logic(dictionaryPath)
}

// Fonction pour afficher les instructions d'utilisation du programme
func displayHelp() {
	fmt.Println("Usage:")
	fmt.Println("  go run . [options]")
	fmt.Println("Options:")
	fmt.Println("  -hardmode    Activer le mode difficile avec des mots plus complexes.")
	fmt.Println("  -h, -help    Afficher cette aide.")
	os.Exit(0)
}
