package game

import (
	"go-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Laser struct {
	image    *ebiten.Image
	position Vector
}

func NewLaser(position Vector) *Laser {
	image := assets.LaserSprite
	bound := image.Bounds()

	halfW := float64(bound.Dx()) / 2
	halfH := float64(bound.Dy()) / 2

	position.x -= halfW
	position.y -= halfH

	return &Laser{
		image:    image,
		position: (position),
	}
}

func (l *Laser) Update() {
	speed := 7.0

	l.position.y += -speed
}

func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(l.position.x, l.position.y)
	screen.DrawImage(l.image, op)

}
