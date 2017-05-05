package systems

import (
	"image/color"
	// "log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type scoreText struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func DrawHUD(w *ecs.World) {
	basicFont := &common.Font{
		URL:  "Roboto-Regular.ttf",
		Size: 32,
		FG:   color.Black,
	}
	if err := basicFont.CreatePreloaded(); err != nil {
		// log.Println("Could not load font:", err)
	}

	score := scoreText{BasicEntity: ecs.NewBasic()}

	// score.RenderComponent = common.RenderComponent{Drawable: basicFont.Render(" ")}
	score.RenderComponent.Drawable = common.Text{
		Font: basicFont,
		Text: "Score players",
		// LetterSpacing: 0.15,
	}
	score.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{100, 5},
		// Width:    100,
		// Height:   100,
	}

	// Add our entity to the appropriate systems
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&score.BasicEntity, &score.RenderComponent, &score.SpaceComponent)
			// case *ScoreSystem:
			// 	sys.Add(&score.BasicEntity, &score.RenderComponent, &score.SpaceComponent)
		}
	}
}
