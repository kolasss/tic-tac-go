package systems

import (
	"image/color"
	// "log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type myShape struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type boardZone struct {
	ecs.BasicEntity
	common.MouseComponent
	common.RenderComponent
	common.SpaceComponent
	gameBoardCoords [2]int
	game            *Game
}

type ZoneControlSystem struct {
	entities []*boardZone
	world    *ecs.World
}

func (c *ZoneControlSystem) New(w *ecs.World) {
	c.world = w
}

func (c *ZoneControlSystem) Add(zone *boardZone) {
	c.entities = append(c.entities, zone)
}

func (c *ZoneControlSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range c.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		c.entities = append(c.entities[:delete], c.entities[delete+1:]...)
	}
}

func (c *ZoneControlSystem) Update(float32) {
	for _, e := range c.entities {
		if e.MouseComponent.Enter {
			engo.SetCursor(engo.CursorHand)
		} else if e.MouseComponent.Leave {
			engo.SetCursor(engo.CursorNone)
		} else if e.MouseComponent.Clicked {
			err := e.game.MakeMove(e.gameBoardCoords[0], e.gameBoardCoords[1])
			if !err {
				switch e.game.Board[e.gameBoardCoords[0]][e.gameBoardCoords[1]] {
				case player1:
					drawOAtZone(e, c.world)
				case player2:
					drawXAtZone(e, c.world)
				}
			}
		}
	}
}

func DrawBoard(w *ecs.World, game *Game) {
	drawLines(w)

	// зоны для крестиков/ноликов
	zonesCoords := [9][2]float32{
		{110, 110}, {110, 310}, {110, 510},
		{310, 110}, {310, 310}, {310, 510},
		{510, 110}, {510, 310}, {510, 510},
	}

	var zones [9]*boardZone

	for i := range zones {
		zones[i] = &boardZone{BasicEntity: ecs.NewBasic()}
	}

	for i, zone := range zones {
		zone.RenderComponent = common.RenderComponent{
			Drawable: common.Rectangle{},
			Color:    color.RGBA{50, 50, 50, 255},
		}

		zone.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{zonesCoords[i][0], zonesCoords[i][1]},
			Width:    180,
			Height:   180,
		}
		zone.game = game
	}

	// n := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			zones[i*3+j].gameBoardCoords = [2]int{i, j}
			// n++
		}
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		// case *common.RenderSystem:
		// 	for _, zone := range zones {
		// 		sys.Add(
		// 			&zone.BasicEntity,
		// 			&zone.RenderComponent,
		// 			&zone.SpaceComponent,
		// 		)
		// 	}
		case *common.MouseSystem:
			for _, zone := range zones {
				sys.Add(
					&zone.BasicEntity,
					&zone.MouseComponent,
					&zone.SpaceComponent,
					&zone.RenderComponent,
				)
			}
		case *ZoneControlSystem:
			for _, zone := range zones {
				sys.Add(zone)
			}
		}
	}
}

func drawLines(w *ecs.World) {
	// линии для поля
	linesCoords := [4][2]float32{
		{300, 100},
		{500, 100},
		{100, 300},
		{100, 500},
	}

	var lines [4]*myShape

	for i := range lines {
		lines[i] = &myShape{BasicEntity: ecs.NewBasic()}
	}
	for i, line := range lines {
		line.RenderComponent = common.RenderComponent{
			Drawable: common.Rectangle{},
			Color:    color.Black,
		}

		var lineWidth float32
		var lineHeight float32
		if i < 2 {
			lineWidth = 5
			lineHeight = 600
		} else {
			lineWidth = 600
			lineHeight = 5
		}

		line.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{linesCoords[i][0], linesCoords[i][1]},
			Width:    lineWidth,
			Height:   lineHeight,
		}
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, line := range lines {
				sys.Add(&line.BasicEntity, &line.RenderComponent, &line.SpaceComponent)
			}
		}
	}
}

func drawXAtZone(zone *boardZone, w *ecs.World) {
	figure1 := myShape{BasicEntity: ecs.NewBasic()}

	figure1.RenderComponent = common.RenderComponent{
		Drawable: common.Rectangle{},
		Color:    color.RGBA{250, 50, 50, 255},
	}
	figure1.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{zone.SpaceComponent.Position.X + 150, zone.SpaceComponent.Position.Y + 30},
		Width:    7,
		Height:   180,
		Rotation: 45,
	}

	figure2 := myShape{BasicEntity: ecs.NewBasic()}

	figure2.RenderComponent = common.RenderComponent{
		Drawable: common.Rectangle{},
		Color:    color.RGBA{250, 50, 50, 255},
	}
	figure2.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{zone.SpaceComponent.Position.X + 25, zone.SpaceComponent.Position.Y + 35},
		Width:    7,
		Height:   180,
		Rotation: -45,
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&figure1.BasicEntity, &figure1.RenderComponent, &figure1.SpaceComponent)
			sys.Add(&figure2.BasicEntity, &figure2.RenderComponent, &figure2.SpaceComponent)
		}
	}
}

func drawOAtZone(zone *boardZone, w *ecs.World) {
	figure1 := myShape{BasicEntity: ecs.NewBasic()}

	figure1.RenderComponent = common.RenderComponent{
		Drawable: common.Circle{BorderWidth: 10, BorderColor: color.RGBA{50, 250, 50, 255}},
		Color:    color.White,
	}
	figure1.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{zone.SpaceComponent.Position.X + 15, zone.SpaceComponent.Position.Y + 15},
		Width:    150,
		Height:   150,
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&figure1.BasicEntity, &figure1.RenderComponent, &figure1.SpaceComponent)
		}
	}
}
