package game

import (
	"fmt"
	"math/rand"
	"strings"
	//"main/entity"
)

func (g *Game) InGameLogic() {
	listeLVL1 := []string{"comme", "disais", "voir", "ensemble", "mots", "locutions", "qui", "sont", "vraiment", "par", "quotidien", "peu", "plus", "que", "habitude", "hein", "panique", "prends", "ton", "temps", "regarde", "plusieurs", "fois", "utilise", "fiche", "PDF", "avec", "cette", "pour", "bien", "assimiler", "bref", "stress"}
	//listeLVL2 := []string{"comme", "disais", "voir","ensemble","mots","locutions","qui","sont","vraiment","par","quotidien","peu","plus","que","habitude","hein","panique","prends","ton","temps","regarde","plusieurs","fois","utilise","fiche","PDF","avec","cette","pour","bien","assimiler","bref","stress"}
	word := listeLVL1[rand.Intn(len(listeLVL1))]
	//word2 := listeLVL2 [rand.Intn(len(listeLVL2))]
	//count := len(word)
	blanks := []string{}
	lives := 2 * len(word)
	for range word {
		blanks = append(blanks, "_")
	}
	for {
		fmt.Printf("word: %s Letter: ", strings.Join(blanks, " "))

		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		fmt.Println(input)
		fmt.Printf("lives = %s", lives)
		for _, inputLetter := range input {
			correctGuess := false

			for i, wordLetter := range word {
				if inputLetter == wordLetter {
					blanks[i] = string(inputLetter)
					correctGuess = true

				}
			}
			if correctGuess == false {
				lives--
			}
		}
		if lives <= 0 {
			fmt.Printf("shesh you lose\n")
			break
		}
		if word == strings.Join(blanks, "") {
			fmt.Printf("shit you won\n")
			break
		}
	}
}
