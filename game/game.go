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
	zoom         = 2
)

var (
	running = true

	tileMap     [][]int
	grassSprite rl.Texture2D

	Character   *ent.Player
	PauseButton *ui.Button

	mousePos rl.Vector2

	playerUp, playerDown, playerRight, playerLeft bool
	playerMoving                                  bool

	state      int
	framecount int

	cam rl.Camera2D
)

func init() {
	state = 0

	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
	rl.InitWindow(screenWidth, screenHeight, "CozyTown")
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
	cam = rl.NewCamera2D(rl.NewVector2(screenWidth/2, screenHeight/2), rl.NewVector2(float32(Character.X)+Character.Width/2, float32(Character.Y)+Character.Height/2), 0, zoom)

}

func input() {

	PauseButton.Bounds = rl.NewRectangle(0, 0, float32(PauseButton.Sprite.Width)*zoom, float32(PauseButton.Sprite.Height)*zoom)

	mousePos = rl.GetMousePosition()

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
		}
		if playerDown {
			Character.Move(0, Character.Speed)
		}
		if playerLeft {
			Character.Move(-Character.Speed, 0)
		}
		if playerRight {
			Character.Move(Character.Speed, 0)
		}
	}
	framecount++

	playerMoving = false
	playerUp, playerDown, playerRight, playerLeft = false, false, false, false

	cam.Target = rl.NewVector2(float32(Character.X)+Character.Width/2, float32(Character.Y)+Character.Height/2)
	PauseButton.SetPos(cam.Target.X-(screenWidth/(2*zoom)-1), cam.Target.Y-(screenHeight/(2*zoom))+1)
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
	Character.DrawPlayer(framecount, state)
	PauseButton.DrawButton(framecount)

}

func main() {

	for running {
		input()
		update()
		render()
	}
	quit()

}
