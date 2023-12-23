package main

import (
	ent "GoGoGo/Entities"
	ui "GoGoGo/UI"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 600
	tileSize     = 16
	mapWidth     = 100
	mapHeight    = 80
	zoom         = 4
)

var (
	running = true

	tileMap      [][]int
	grassSprite  rl.Texture2D
	cursorSprite rl.Texture2D

	mouseWheelUsage float32

	Character   *ent.Player
	PauseButton *ui.Button
	ItemBar     *ui.ItemBar

	mousePos rl.Vector2

	playerUp, playerDown, playerRight, playerLeft bool
	playerMoving                                  bool

	framecount int

	cam rl.Camera2D
)

func init() {
	
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
	rl.InitWindow(screenWidth, screenHeight, "CozyTown")
	cursorSprite = rl.LoadTexture("Assets/Cursor/cursor_std.png")
	grassSprite = rl.LoadTexture("Assets/Tiles/grass_1.png")

	tileMap = make([][]int, mapHeight)
	for i := range tileMap {
		tileMap[i] = make([]int, mapWidth)
		for j := range tileMap[i] {
			tileMap[i][j] = 0
		}
	}

	Character = ent.NewPlayer(200, 100, 100, 1, "Assets/PlayerFemaleAnim.png")
	PauseButton = ui.NewButton(0, 0, 0, "Assets/Buttons/pausebutton.png", "Assets/Buttons/pausebutton_pressed.png")
	ItemBar = ui.NewItemBar(screenWidth/2, screenHeight-18, "Assets/ItemBar/item_bar.png", "Assets/ItemBar/item_bar_sel.png")
	cam = rl.NewCamera2D(rl.NewVector2(screenWidth/2, screenHeight/2), rl.NewVector2(Character.GetX()+Character.GetWidth()/2, Character.GetWidth()+Character.GetHeight()/2), 0, zoom)

}

func input() {

	PauseButton.Bounds = rl.NewRectangle(0, 0, float32(PauseButton.Sprite.Width)*zoom, float32(PauseButton.Sprite.Height)*zoom)

	mouseWheelUsage = rl.GetMouseWheelMove()
	mousePos = rl.GetMousePosition()

	if mouseWheelUsage != 0 {
		ItemBar.Selected = (ItemBar.Selected + int32(mouseWheelUsage) + ItemBar.ItemCount) % ItemBar.ItemCount
		fmt.Println(ItemBar.Selected)
	}

	if rl.CheckCollisionPointRec(mousePos, PauseButton.Bounds) {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			PauseButton.OnClick()
		}
	}

	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerMoving = true
		playerUp = true
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerMoving = true
		playerDown = true
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerMoving = true
		playerLeft = true
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerMoving = true
		playerRight = true
	}
	if rl.IsKeyPressed(rl.KeyF) {
		rl.ToggleFullscreen()
	}
	if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		fmt.Print("hi")
		PauseButton.Pressed = false
	}
}

func update() {
	running = !rl.WindowShouldClose()
	if playerMoving {
		if playerUp {
			Character.Move(0, -Character.Speed)
			Character.State = 2
		}
		if playerDown {
			Character.Move(0, Character.Speed)
			Character.State = 1
		}
		if playerLeft {
			Character.Move(-Character.Speed, 0)
			Character.State = 4
		}
		if playerRight {
			Character.Move(Character.Speed, 0)
			Character.State = 3
		}
	}

	if !playerMoving {
		Character.State = 0
	}

	framecount++
	playerMoving = false
	playerUp, playerDown, playerRight, playerLeft = false, false, false, false

	cam.Target = rl.NewVector2(Character.GetX()-Character.GetWidth()/2, Character.GetY()-Character.GetHeight()/2)
	PauseButton.SetPos(cam.Target.X-(screenWidth/(2*zoom)-1), cam.Target.Y-(screenHeight/(2*zoom))+1)
	ItemBar.SetPos(cam.Target.X-float32(ItemBar.Sprite.Width/2), cam.Target.Y+(screenHeight/(2*zoom))-float32(ItemBar.Sprite.Height+2))
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

func drawScene() {
	for i := range tileMap {
		for j := range tileMap[i] {
			rl.DrawTexture(grassSprite, int32(j*tileSize), int32(i*tileSize), rl.White)
		}
	}
	Character.DrawPlayer(framecount)

	PauseButton.DrawButton(framecount)
	ItemBar.DrawItemBar()
}

func main() {

	for running {
		input()
		update()
		render()
	}
	quit()

}
