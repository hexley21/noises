package noise

import "image"

type Noise interface {
	Generate2D(width, height int) [][]float64
	GenerateImage2D(width, height int, useColor bool) image.Image
}
