package game

import (
	"fmt"
	"go-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image     *ebiten.Image
	position  Vector
	game      *Game
	laserLoad *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite
	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2

	position := Vector{
		x: (screenWidthX / 2) - halfW,
		y: 500,
	}
	return &Player{
		image:     image,
		game:      game,
		position:  position,
		laserLoad: NewTimer(12),
	}
}

func (p *Player) Update() {
	speed := 10.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) && p.position.x > 0 {
		p.position.x -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) && p.position.x < 700 {
		fmt.Println(p.position.x)
		p.position.x += speed
	}

	p.laserLoad.Update()
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserLoad.IsReady() {
		p.laserLoad.Reset()
		bounds := p.image.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		spawnPos := Vector{
			p.position.x + halfW,
			p.position.y - halfH/2,
		}

		laser := NewLaser(spawnPos)
		p.game.AddLasers(laser)
	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//Posição x e y da imagem
	op.GeoM.Translate(p.position.x, p.position.y)

	//Desnha a imagem
	screen.DrawImage(p.image, op)

}
