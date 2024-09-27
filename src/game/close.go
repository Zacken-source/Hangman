package game

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib" // Import du module raylib pour gérer les fenêtres et le son
)

// Close : Fonction pour fermer proprement le moteur et quitter le jeu
func (g *Game) Close() {
	// Ferme le dispositif audio (libère les ressources audio)
	rl.CloseAudioDevice()

	// Ferme la fenêtre du jeu créée par Raylib
	rl.CloseWindow()

	// Quitte le programme avec un code de sortie 0 (indiquant que tout s'est bien passé)
	os.Exit(0)
}
