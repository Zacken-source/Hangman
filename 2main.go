package main

import (
	"strings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var screenWidth int
var screenHeight int
var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var guessedLetters = []rune{}
var secretWord = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var revealedWord = strings.Repeat("_", len(secretWord))  // Mot masqué au départ

func main() {
	rl.InitWindow(2880, 1800, "Hangman")
	rl.ToggleFullscreen()

	screenWidth = rl.GetScreenWidth()
	screenHeight = rl.GetScreenHeight()

	rl.SetTargetFPS(120)

	missed := 0

	for !rl.WindowShouldClose() {
		// Gestion des clics sur le clavier virtuel ou des touches physiques
		handleKeyboardClicks(&missed)

		// Commencer le dessin
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// 1. Afficher les emplacements de lettres (centré en haut)
		drawWordToGuess()

		// 2. Dessiner le Hangman (centré, légèrement plus bas que le milieu)
		drawHangmanStructure(missed)

		// 3. Dessiner le clavier virtuel (centré en bas)
		drawVirtualKeyboard()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

// Fonction pour dessiner les lettres du mot à deviner
func drawWordToGuess() {
	fontSize := int32(50)
	wordWidth := int32(len(revealedWord)) * (fontSize + 20) // Largeur totale du mot
	startX := int32(screenWidth)/2 - wordWidth/2

	for i, letter := range revealedWord {
		rl.DrawText(string(letter), startX+int32(i)*(fontSize+20), 50, fontSize, rl.White)
	}
}

// Fonction pour dessiner la structure du hangman, ajustée pour prendre plus de place à l'écran
func drawHangmanStructure(missed int) {
	centerX := screenWidth / 2
	baseY := screenHeight / 2 + 100 // Position centrale légèrement vers le bas

	// Facteur d'échelle pour agrandir les éléments du pendu
	scaleFactor := float32(screenHeight) / 800.0 // Ajuste l'échelle en fonction de la hauteur de l'écran

	// Dessiner les éléments en fonction des erreurs
	if missed >= 1 {
		rl.DrawRectangle(int32(float32(centerX)-150*scaleFactor), int32(float32(baseY)+100*scaleFactor), int32(300*scaleFactor), int32(20*scaleFactor), rl.White) // Base
	}
	if missed >= 2 {
		rl.DrawRectangle(int32(float32(centerX)-10*scaleFactor), int32(float32(baseY)-250*scaleFactor), int32(20*scaleFactor), int32(350*scaleFactor), rl.White) // Pilier vertical
	}
	if missed >= 3 {
		rl.DrawRectangle(int32(float32(centerX)-10*scaleFactor), int32(float32(baseY)-350*scaleFactor), int32(200*scaleFactor), int32(20*scaleFactor), rl.White) // Poutre horizontale
	}
	if missed >= 4 {
		rl.DrawLine(int32(float32(centerX)), int32(float32(baseY)-150*scaleFactor), int32(float32(centerX)-100*scaleFactor), int32(float32(baseY)+100*scaleFactor), rl.White) // Renfort oblique
	}
	if missed >= 5 {
		rl.DrawLine(int32(float32(centerX)+100*scaleFactor), int32(float32(baseY)-350*scaleFactor), int32(float32(centerX)+100*scaleFactor), int32(float32(baseY)-300*scaleFactor), rl.White) // Corde
	}
	if missed >= 6 {
		rl.DrawCircle(int32(float32(centerX)+100*scaleFactor), int32(float32(baseY)-275*scaleFactor), float32(50*scaleFactor), rl.White) // Tête
	}
	if missed >= 7 {
		rl.DrawLine(int32(float32(centerX)+100*scaleFactor), int32(float32(baseY)-225*scaleFactor), int32(float32(centerX)+100*scaleFactor), int32(float32(baseY)-125*scaleFactor), rl.White) // Corps
	}
	if missed >= 8 {
		rl.DrawLine(int32(float32(centerX)+100*scaleFactor), int32(float32(baseY)-200*scaleFactor), int32(float32(centerX)+50*scaleFactor), int32(float32(baseY)-150*scaleFactor), rl.White) // Bras gauche
	}
	if missed >= 9 {
		rl.DrawLine(int32(float32(centerX)+100*scaleFactor), int32(float32(baseY)-200*scaleFactor), int32(float32(centerX)+150*scaleFactor), int32(float32(baseY)-150*scaleFactor), rl.White) // Bras droit
	}
	if missed == 10 {
		rl.DrawLine(int32(float32(centerX)+100*scaleFactor), int32(float32(baseY)-125*scaleFactor), int32(float32(centerX)+50*scaleFactor), int32(float32(baseY)-50*scaleFactor), rl.White) // Jambe gauche
		rl.DrawLine(int32(float32(centerX)+100*scaleFactor), int32(float32(baseY)-125*scaleFactor), int32(float32(centerX)+150*scaleFactor), int32(float32(baseY)-50*scaleFactor), rl.White) // Jambe droite
	}
}

// Fonction pour dessiner un clavier virtuel
func drawVirtualKeyboard() {
	fontSize := int32(30)
	buttonSize := int32(50) // Agrandi les boutons pour correspondre à l'échelle de l'écran
	padding := int32(10)

	startX := int32(screenWidth)/2 - (13*(buttonSize+padding))/2
	startY := screenHeight - 200

	for i, letter := range letters {
		x := startX + int32(i%13)*(buttonSize+padding)
		y := int32(startY) + int32(i/13)*(buttonSize+padding)

		rect := rl.Rectangle{X: float32(x), Y: float32(y), Width: float32(buttonSize), Height: float32(buttonSize)}
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rect) {
			rl.DrawRectangleRec(rect, rl.DarkGray)
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && !contains(guessedLetters, letter) {
				guessedLetters = append(guessedLetters, letter)
			}
		} else {
			rl.DrawRectangleRec(rect, rl.Gray)
		}
		rl.DrawText(string(letter), x+10, y+5, fontSize, rl.White)
	}
}

// Fonction pour vérifier si une lettre a déjà été devinée
func contains(slice []rune, letter rune) bool {
	for _, l := range slice {
		if l == letter {
			return true
		}
	}
	return false
}

// Fonction pour gérer les clics sur le clavier (physique ou virtuel)
func handleKeyboardClicks(missed *int) {
	// Gestion des lettres tapées sur le clavier physique (AZERTY)
	for _, letter := range letters {
		var keyPressed int32 = -1 // Variable pour stocker la touche pressée

		// Vérifiez chaque lettre dans le mappage AZERTY
		switch letter {
		case 'A':
			if rl.IsKeyPressed(rl.KeyQ) {
				keyPressed = rl.KeyA
			}
		case 'Q':
			if rl.IsKeyPressed(rl.KeyA) {
				keyPressed = rl.KeyQ
			}
		case 'Z':
			if rl.IsKeyPressed(rl.KeyW) {
				keyPressed = rl.KeyZ
			}
		case 'W':
			if rl.IsKeyPressed(rl.KeyZ) {
				keyPressed = rl.KeyW
			}
		case'M':
			if rl.IsKeyPressed(rl.KeySemicolon) {
				keyPressed = rl.KeyComma
			}
		default:
			if rl.IsKeyPressed(int32(letter)) {
				keyPressed = int32(letter)
			}
		}

		// Si une touche valide a été pressée, ajoutez-la aux lettres devinées
		if keyPressed != -1 && !contains(guessedLetters, letter) {
			guessedLetters = append(guessedLetters, letter)
		}
	}

	// Mise à jour du mot et des erreurs
	for _, letter := range guessedLetters {
		if !strings.ContainsRune(secretWord, letter) {
			*missed = len(guessedLetters) - len(guessedCorrectLetters(secretWord)) // Met à jour en fonction des erreurs
		} else {
			for i, l := range secretWord {
				if l == letter {
					revealedWord = revealedWord[:i] + string(l) + revealedWord[i+1:]
				}
			}
		}
	}
}

// Fonction pour obtenir les lettres correctes déjà devinées
func guessedCorrectLetters(word string) []rune {
	correct := []rune{}
	for _, letter := range guessedLetters {
		if strings.ContainsRune(word, letter) {
			correct = append(correct, letter)
		}
	}
	return correct
}
