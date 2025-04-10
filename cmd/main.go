package main

import (
	"encoding/json"
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/hexley21/noises/internal/noise"
	"github.com/hexley21/noises/internal/noise/perlin"
)

func main() {
	// Default parameters
	width := 500
	height := 500

	// Noise configuration
	scale := 0.01      // Scale factor (lower = smoother)
	octaves := 6       // Number of octaves (detail layers)
	persistence := 0.5 // How much each octave contributes

	// Image parameters
	useColor := false

	f, err := os.Create("perlin_noise.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	perlin := perlin.NewPerlinNoise(octaves, persistence, scale)
	writeResult(width, height, useColor, "perlin", perlin)
}

func writeResult(width, height int, useColor bool, name string, n noise.Noise) {
	f, err := os.Create(fmt.Sprint(name, "_noise.png"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	png.Encode(f, n.GenerateImage2D(width, height, useColor))
	res, _ := json.MarshalIndent(n.Generate2D(width, height), "", "\t")

	os.WriteFile(fmt.Sprint(name, "_noise.json"), res, 0644)
}
