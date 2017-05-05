package main

import (
	"image/color"
	// "log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/kolasss/tic-tac-go/systems"
)

type GameScene struct{}

func (*GameScene) Preload() {
	// err := engo.Files.Load("Roboto-Regular.ttf")
	// if err != nil {
	// 	log.Println(err)
	// }
}
func (g *GameScene) Setup(w *ecs.World) {
	common.SetBackground(color.White)
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.MouseSystem{})
	w.AddSystem(&systems.ZoneControlSystem{})

	game := systems.NewGame()

	// systems.DrawHUD(w)
	systems.DrawBoard(w, &game)
}

func (*GameScene) Type() string { return "GameScene" }

func main() {
	opts := engo.RunOptions{
		Title:  "Tic-tac-go",
		Width:  800,
		Height: 800,
	}
	engo.Run(opts, &GameScene{})
}
