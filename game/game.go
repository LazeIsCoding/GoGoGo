package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)
import ent "GoGoGo/Entities"

const (
	screenWidth  = 512
	screenHeight = 374
)

var (
	running   = true
	Character *ent.Player

	playerUp, playerDown, playerRight, playerLeft bool
	playerMoving                                  bool

	cam rl.Camera2D
)

func initialize() {
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
	rl.InitWindow(screenWidth, screenHeight, "CozyTown")
	Character = ent.NewPlayer(100, 100, 100, 1, "Assets/PlayerFemale.png")
	cam = rl.NewCamera2D(rl.NewVector2(screenWidth/2, screenHeight/2), rl.NewVector2(float32(Character.X)+Character.Width/2, float32(Character.Y)+Character.Height/2), 0, 3)
	fmt.Print(Character.Width, " H:", Character.Height)
}

func input() {
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

	playerMoving = false
	playerUp, playerDown, playerRight, playerLeft = false, false, false, false

	cam.Target = rl.NewVector2(float32(Character.X)+Character.Width/2, float32(Character.Y)+Character.Height/2)
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

	rl.DrawTexture(Character.Sprite, Character.X, Character.Y, rl.White)

}
func main() {

	initialize()

	for running {
		input()
		update()
		render()
	}
	quit()

}
