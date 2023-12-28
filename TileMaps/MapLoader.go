package TileMaps

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func LoadMap(mapPath string) ([][]int, int32, int32) {
	var data map[string]interface{}

	file, err := os.Open(mapPath)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Unable to read file content: %v", err)
	}

	if err := json.Unmarshal(content, &data); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	width := int(data["width"].(float64))
	height := int(data["height"].(float64))
	layers := data["layers"].([]interface{})
	mapData := make([][]int, height)

	for i := 0; i < height; i++ {
		mapData[i] = make([]int, width)
		for _, layer := range layers {
			layerData := layer.(map[string]interface{})["data"].([]interface{})
			for j := 0; j < width; j++ {
				mapData[i][j] += int(layerData[i*width+j].(float64))
			}
		}
	}

	return mapData, int32(width), int32(height)
}
