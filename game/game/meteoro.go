package game

import (
	"go-game/assets"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Meteoro struct {
	image    *ebiten.Image
	speed    float64
	position Vector
}

func NewMetoro() *Meteoro {
	image := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]

	speed := (rand.Float64() * 13)
	position := Vector{
		x: rand.Float64() * screenWidthX,
		y: -100,
	}
	return &Meteoro{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func (m *Meteoro) Update() {
	m.position.y += m.speed
}

func (m *Meteoro) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(m.position.x, m.position.y)
	screen.DrawImage(m.image, op)
}
