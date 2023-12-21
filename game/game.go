package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity struct {
	Y      int32
	X      int32
	sprite rl.Texture2D
}

var Character Entity

func initialize() {
	rl.InitWindow(528, 528, "CozyTown")
	Character.X = 100
	Character.Y = 100
	Character.sprite = rl.LoadTexture("Assets/PlayerFemale.png")
}

func update() {

}

func input() {
	if rl.IsKeyPressed(rl.KeyD) {
		Character.X++
	}
}
func render() {
	rl.DrawTexture(Character.sprite, Character.X, Character.Y, rl.White)
}

func main() {

	initialize()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		
		input()
		update()
		render()

		rl.EndDrawing()
	}
	//quit()
	rl.CloseWindow()
}
