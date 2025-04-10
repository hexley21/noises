package perlin

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"time"
)

type PerlinNoise struct {
	permutation []int
	octaves     int
	persistence float64
	scale       float64
}

func NewPerlinNoise(octoctaves int, persistence, scale float64) *PerlinNoise {
	perm := make([]int, 512)
	for i := range len(perm) / 2 {
		perm[i] = i
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(perm)/2, func(i, j int) {
		perm[i], perm[j] = perm[j], perm[i]
	})

	return &PerlinNoise{
		permutation: perm,
		octaves:     octoctaves,
		persistence: persistence,
		scale:       scale,
	}
}

func (p *PerlinNoise) noise2D(x, y float64) float64 {
	X := int(math.Floor(x)) & 255
	Y := int(math.Floor(y)) & 255

	x -= math.Floor(x)
	y -= math.Floor(y)

	u := fade(x)
	v := fade(y)

	A := p.permutation[X] + Y
	AA := p.permutation[A]
	AB := p.permutation[A+1]
	B := p.permutation[X+1] + Y
	BA := p.permutation[B]
	BB := p.permutation[B+1]

	return lerp(v,
		lerp(u, grad(AA, x, y, 0), grad(BA, x-1, y, 0)),
		lerp(u, grad(AB, x, y-1, 0), grad(BB, x-1, y-1, 0)))
}

// Fade function as defined by Ken Perlin
func fade(t float64) float64 {
	return t * t * t * (t*(t*6-15) + 10)
}

// Linear interpolation between a and b by t
func lerp(t, a, b float64) float64 {
	return a + t*(b-a)
}

// Convert hash code into gradient direction
func grad(hash int, x, y, z float64) float64 {
	switch hash & 15 {
	case 0:
		return x + y
	case 1:
		return -x + y
	case 2:
		return x - y
	case 3:
		return -x - y
	case 4:
		return x + z
	case 5:
		return -x + z
	case 6:
		return x - z
	case 7:
		return -x - z
	case 8:
		return y + z
	case 9:
		return -y + z
	case 10:
		return y - z
	case 11:
		return -y - z
	case 12:
		return y + x
	case 13:
		return -y + z
	case 14:
		return y - x
	case 15:
		return -y - z
	}

	return 0 // never happens
}

// Generate fractal noise by combining multiple octaves
func (p *PerlinNoise) octaveNoise2D(x, y float64) float64 {
	total := 0.0
	frequency := 1.0
	amplitude := 1.0
	maxValue := 0.0 // Used for normalizing result

	for range p.octaves {
		total += p.noise2D(x*frequency, y*frequency) * amplitude

		maxValue += amplitude
		amplitude *= p.persistence
		frequency *= 2
	}

	return total / maxValue
}

func (p *PerlinNoise) Generate2D(width, height int) [][]float64 {
	// Initialize 2D slice
	noise2D := make([][]float64, height)
	for i := range noise2D {
		noise2D[i] = make([]float64, width)
	}

	// Fill with noise values
	for y := range height {
		for x := range width {
			nx := float64(x) * p.scale
			ny := float64(y) * p.scale
			noise2D[y][x] = p.octaveNoise2D(nx, ny)
		}
	}

	return noise2D
}

func (p *PerlinNoise) GenerateImage2D(width, height int, useColor bool) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	noise2D := p.Generate2D(width, height)

	for x, n := range noise2D {
		for y, v := range n {

			gray := uint8((v + 1.0) * 127.5)

			if useColor {
				r := uint8(math.Min(255, float64(gray*2)))
				g := uint8(math.Max(0, math.Min(255, float64(gray-64)*2)))
				b := uint8(math.Max(0, math.Min(255, float64(gray-128)*2)))
				img.Set(x, y, color.RGBA{r, g, b, 255})
			} else {
				img.Set(x, y, color.Gray{gray})
			}
		}
	}
	return img
}
