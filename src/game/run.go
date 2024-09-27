package game

import (
	//"main/src/entity"
	"flag"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib" 
)


func (game *Game) Run() {
	rl.SetTargetFPS(120)
    rl.ToggleFullscreen()
	showFPS := flag.Bool("f", false, "Affiche les FPS")
	flag.Parse()

	
	for game.IsRunning {
		
		switch game.StateMenu {
			case HOME: 
				//game.homerandering()
				//game.homelogic()
				
			case SETTINGS:
				//game.SettingsLogic() 

			case PLAY:
				switch game.StateGame {
				case INGAME:
					//game.InGameRendering()
					//game.InGameLogic()
					if game.Player.Lives <= 0 {
						game.StateGame = GAMEOVER
					}
				
				case GAMEOVER:
					//game.GameOverRendering()
					//enginegame.GameOverLogic()
		}

		
		if *showFPS {
            fps := rl.GetFPS() 
            rl.DrawText(fmt.Sprintf("FPS: %d", fps), 10, 10, 20, rl.DarkGray)
        }

		rl.EndDrawing()
	}
}
}