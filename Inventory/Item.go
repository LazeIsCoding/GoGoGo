package Inventory

import rl "github.com/gen2brain/raylib-go/raylib"

type Item struct {
	Id     int32
	Count  int32
	Sprite rl.Texture2D
}

func CreateItem(id int32, sprite rl.Texture2D) *Item {
	return &Item{
		Id:     id,
		Count:  1,
		Sprite: sprite,
	}
}

func (i *Item) getID() int32 {
	return i.Id
}
