package main

import (
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"unicode/utf8"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

const (
	screenWidth     = 800
	screenHeight    = 600
	healthBarWidth  = 200
	healthBarHeight = 20
	healthLabelY    = healthBarHeight + 5 // Space between bar and label
)

type Game struct {
	background  *ebiten.Image
	characterX  int
	characterY  int
	isAttacking bool
}

func (g *Game) Update() error {
	//PLaceholder
	g.isAttacking = !g.isAttacking // Simulate attack movement for testing
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.background, nil)

	// Draw the health bar
	x := (screenWidth - healthBarWidth) / 2
	y := 10

	//health bar outlining
	for i := 0; i < healthBarWidth; i++ {
		for j := 0; j < healthBarHeight; j++ {
			screen.Set(x+i, y+j, colornames.Gray)
		}
	}

	// Draw the health label
	// TODO: Not displaying properly
	text := "Health"
	for i, c := range text {
		screen.Set(x+(healthBarWidth-utf8.RuneLen(c))/2+i*8, y+healthLabelY, colornames.Black) // Draw character
	}

	// Draw the warrior character
	g.drawWarrior(screen)
}

// Placeholder until functionality works for combat and movement
func (g *Game) drawWarrior(screen *ebiten.Image) {
	bodyColor := color.RGBA{0, 0, 255, 255}     // Blue for the body
	weaponColor := color.RGBA{139, 69, 19, 255} // Brown for the weapon

	// Draw the body
	for dx := -10; dx <= 10; dx++ {
		for dy := -20; dy <= 20; dy++ {
			screen.Set(g.characterX+dx, g.characterY+dy, bodyColor)
		}
	}

	// Draw the head
	for dx := -5; dx <= 5; dx++ {
		for dy := -5; dy <= 5; dy++ {
			if dx*dx+dy*dy <= 25 {
				screen.Set(g.characterX+dx, g.characterY-30+dy, bodyColor)
			}
		}
	}

	// Draw the weapon/ a stick
	if g.isAttacking {
		// Extended position for "attack"
		for i := 0; i < 15; i++ {
			screen.Set(g.characterX+15+i, g.characterY-10, weaponColor)
		}
	} else {
		// Rest position
		for i := 0; i < 10; i++ {
			screen.Set(g.characterX+10, g.characterY-10-i, weaponColor)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	bgImage, err := loadImage("background.jpg")
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Simple Game Template with Warrior Character and Health Bar")

	game := &Game{
		background: bgImage,
		characterX: screenWidth / 2,  // Fix the positioning
		characterY: screenHeight / 2, //create actual 'grounding'
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

// Load image from file (for now its only background)
func loadImage(file string) (*ebiten.Image, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()

	img, err := jpeg.Decode(imgFile)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}
