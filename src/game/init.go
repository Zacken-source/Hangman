package game

import (
	"main/src/entity"

	//rl "github.com/gen2brain/raylib-go/raylib"
)



func (g *Game) InitEntities(){
	g.Player = entity.Player{
//		Lives:	2*len(g.mot),
		IsAlive: true,
		Score:    0,
}}

/*func (g *Game) InitListes() {

	g.Slice.Slice = append(e.Seller.Inventory, item.Item{
		Name:         "Potion",
		Price:        5,
		IsConsumable: true,
		IsEquippable: false,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png"),
	})
	e.Seller.Inventory = append(e.Seller.Inventory, item.Item{
		Name:         "Epée",
		Price:        15,
		IsConsumable: false,
		IsEquippable: true,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png"),
	})
	e.Seller.Inventory = append(e.Seller.Inventory, item.Item{
		Name:         "Bouclier",
		Price:        25,
		IsConsumable: false,
		IsEquippable: true,
		Sprite:       rl.LoadTexture("textures/items/itemschelou.png"),
	})

}*/