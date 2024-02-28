package main

import (
	ent "GoGoGo/Entities"
	"GoGoGo/TileMaps"
	ui "GoGoGo/UI"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
	"math/rand"
)

const (
	screenWidth  = 1600
	screenHeight = 900
	tileSize     = 16
	zoom         = 4
)

var (
	running    = true
	framecount int

	tileMap   [][]int
	mapWidth  int32
	mapHeight int32
	Obstacles []rl.Rectangle

	grassSprite  rl.Texture2D
	rockSprite   rl.Texture2D
	cursorSprite rl.Texture2D

	Character   *ent.Player
	PauseButton *ui.Button
	ItemBar     *ui.ItemBar
	Butterflies []*ent.Butterfly

	Rand *rand.Rand

	mouseWheelUsage float32

	mousePos rl.Vector2

	playerUp, playerDown, playerRight, playerLeft bool
	playerMoving                                  bool
	dir                                           rl.Vector2

	cam rl.Camera2D
)

func init() {
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
	rl.SetWindowMonitor(2)
	rl.InitWindow(screenWidth, screenHeight, "")
	initTextures()

	Obstacles = []rl.Rectangle{}
	Butterflies = []*ent.Butterfly{}

	tileMap, mapWidth, mapHeight = TileMaps.LoadMap("Assets/Maps/map3.tmj")

	//Add CollisionObjects to Obstacles array
	for i := range tileMap {
		for j := range tileMap[i] {
			if tileMap[i][j] == 2 {
				Obstacles = append(Obstacles, rl.NewRectangle(float32(j*tileSize), float32(i*tileSize), 16, 16))
			}
		}
	}

	initEntities()

	dir = rl.NewVector2(0, 0)
}

func input() {
	dir.X = 0
	dir.Y = 0

	mouseWheelUsage = rl.GetMouseWheelMove()
	mousePos = rl.GetMousePosition()

	//Change selected item based on mouseWheelUsage
	if mouseWheelUsage != 0 {
		ItemBar.Selected = (ItemBar.Selected - int32(mouseWheelUsage) + ItemBar.ItemCount) % ItemBar.ItemCount
	}

	//check Collision with mousePos
	if rl.CheckCollisionPointRec(mousePos, PauseButton.Bounds) {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			PauseButton.OnClick()
		}
	}

	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		dir.Y -= Character.Speed
		playerMoving = true
		playerUp = true
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		dir.Y += Character.Speed
		playerMoving = true
		playerDown = true
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		dir.X -= Character.Speed
		playerMoving = true
		playerLeft = true
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		dir.X += Character.Speed
		playerMoving = true
		playerRight = true
	}
	if rl.IsKeyPressed(rl.KeyF) {
		rl.ToggleFullscreen()
	}

	if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		PauseButton.Pressed = false
	}

}

func update() {
	running = !rl.WindowShouldClose()
	checkCollisions()
	if playerMoving {
		if playerDown {
			Character.Move(0, Character.Speed)
			Character.State = 1
		}
		if playerUp {
			Character.Move(0, -Character.Speed)
			Character.State = 2
		}
		if playerRight {
			Character.Move(Character.Speed, 0)
			Character.State = 3
		}
		if playerLeft {
			Character.Move(-Character.Speed, 0)
			Character.State = 4
		}

	}
	if !playerMoving {
		Character.State = 0
	}

	framecount++
	playerMoving = false
	playerUp, playerDown, playerRight, playerLeft = false, false, false, false
	for _, but := range Butterflies {
		but.Move()
	}
	cam.Target = rl.NewVector2(float32(math.Round(float64(Character.GetX()+Character.GetWidth()/2))), float32(math.Round(float64(Character.GetY()+Character.GetHeight()/2))))
	ItemBar.SetPos(cam.Target.X-float32(ItemBar.Sprite.Width/2), cam.Target.Y+(screenHeight/(2*zoom))-float32(ItemBar.Sprite.Height+2))
	PauseButton.SetPos(cam.Target.X-(screenWidth/(2*zoom)), cam.Target.Y-(screenHeight/(2*zoom)))
	spawnEntities(framecount)
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	rl.BeginMode2D(cam)
	drawScene()
	rl.EndMode2D()
	rl.EndDrawing()

}

func quit() {
	rl.UnloadTexture(Character.Sprite)
	rl.CloseWindow()
}

// initialize Sprites
func initTextures() {
	cursorSprite = rl.LoadTexture("Assets/Cursor/cursor_std.png")
	grassSprite = rl.LoadTexture("Assets/Tiles/grass_1.png")
	rockSprite = rl.LoadTexture("Assets/Tiles/rock_1.png")
}
func initEntities() {
	Character = ent.NewPlayer(200, 200, 100, 1, "Assets/PlayerFemaleAnim.png")
	PauseButton = ui.NewButton(0, 0, 0, "Assets/Buttons/pausebutton.png", "Assets/Buttons/pausebutton_pressed.png")
	PauseButton.Bounds = rl.NewRectangle(0, 0, float32(PauseButton.Sprite.Width)*zoom, float32(PauseButton.Sprite.Height)*zoom)

	ItemBar = ui.NewItemBar(screenWidth/2, screenHeight-18, "Assets/ItemBar/item_bar.png", "Assets/ItemBar/item_bar_sel.png", "")
	cam = rl.NewCamera2D(rl.NewVector2(screenWidth/2, screenHeight/2), rl.NewVector2(Character.GetX()+Character.GetWidth()/2, Character.GetWidth()+Character.GetHeight()/2), 0, zoom)
}

// draw components on screen
func drawScene() {

	//Loading only a slightly bigger chunk than the screen size
	for i := int(math.Max(0, float64(Character.GetY()/tileSize-40/zoom))); i < int(math.Min(float64(mapHeight), float64(Character.GetY()/tileSize+44/zoom))); i++ {
		for j := int(math.Max(0, float64(Character.GetX()/tileSize-52/zoom))); j < int(math.Min(float64(mapWidth), float64(Character.GetX()/tileSize+64/zoom))); j++ {
			switch tileMap[i][j] {
			case 1:
				rl.DrawTexture(grassSprite, int32(j*tileSize), int32(i*tileSize), rl.White)
			case 2:
				rl.DrawTexture(rockSprite, int32(j*tileSize), int32(i*tileSize), rl.White)
			}
		}
	}

	Character.DrawPlayer(framecount)
	for _, but := range Butterflies {
		but.Draw(framecount)
	}
	PauseButton.DrawButton(framecount)
	ItemBar.DrawItemBar()
}

func checkCollisions() {
	for _, rec := range Obstacles {
		if rl.CheckCollisionRecs(rl.NewRectangle(Character.Pos.X+dir.X, Character.Pos.Y+dir.Y, Character.GetWidth(), Character.GetHeight()), rec) {
			playerMoving = false
		}
	}
	if true {
		for i, rec := range Butterflies {
			if rl.CheckCollisionRecs(Character.Pos, rec.Pos) {
				Butterflies = append(Butterflies[:i], Butterflies[i+1:]...)
				break
			}
		}
	}
}

func spawnEntities(framecount int) {
	if framecount%5 == 0 {

		if rand.Float32() > 0.5 {
			Butterflies = append(Butterflies, ent.NewButterfly(200, 200, "Assets/Entities/Butterflies/butterfly_yellow_anim.png"))
		}
	}

}
func main() {

	for running {
		input()
		update()
		render()
	}
	quit()

}
