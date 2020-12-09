package main

import ebiten "github.com/hajimehoshi/ebiten/v2"
import "log"

const (
	screenWidth = 320
	screenHeight= 240
)

type Game struct{
	pixels []byte
}

func (g *Game) Update() error {
	return nil
}

type ColorRGBA struct {
	r,g,b,a int
}

type Vector2d struct {
	x, y int
}

func Clamp(i int) int {
	if i > 255 {
		return 255
	}
	if i < 0 {
		return 0
	}
	return i
}

func NewColorRGBA(r, g, b, a int) ColorRGBA {
	return ColorRGBA{r:Clamp(r), g:Clamp(g), b:Clamp(b), a:Clamp(a)}
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func drawLine(image []byte, start, end Vector2d, color ColorRGBA) {
	dx := abs(end.x - start.x)
	sx := -1
	if start.x < end.x {
		sx = 1
	}
	dy := -abs(end.y - start.y)
	sy := -1
	if start.y < end.y {
		sy = 1
	}
	err := dx + dy
	x := start.x
	y := start.y
	for {
		drawPixel(image, Vector2d{x, y}, color)
		if x == end.x && y == end.y {
			break
		}
		e2 := 2 * err
		if e2 >= dy {
			err += dy
			x += sx
		}
		if e2 <= dx {
			err += dx
			y += sy
		}
	}
}

func drawPixel(image []byte, position Vector2d, color ColorRGBA) {
	if position.y >= screenHeight || position.x >= screenWidth || position.x < 0 || position.y < 0{
		return
	}
	pos := (position.y * screenWidth + position.x) * 4
	image[pos+0] = byte(color.r)
	image[pos+1] = byte(color.g)
	image[pos+2] = byte(color.b)
	image[pos+3] = byte(color.a)
}

func (g *Game) TestDraw() {
	white := NewColorRGBA(255, 255, 255, 255)
	drawPixel(g.pixels, Vector2d{0, 0}, white)
	drawLine(g.pixels, Vector2d{40, 50}, Vector2d{200, 300}, white)
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.pixels == nil {
		g.pixels = make([]byte, screenWidth*screenHeight*4)
	}
	g.TestDraw()
	screen.ReplacePixels(g.pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Derp")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
