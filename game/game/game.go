package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player      *Player
	lasers      []*Laser
	meteoros    []*Meteoro
	meteorSpawn *Timer
}

func NewGame() *Game {
	g := &Game{
		meteorSpawn: NewTimer(24),
	}
	player := NewPlayer(g)
	g.player = player
	return g
}

func (g *Game) Update() error {

	g.player.Update()

	for _, l := range g.lasers {
		l.Update()
	}

	g.meteorSpawn.Update()
	if g.meteorSpawn.IsReady() {
		g.meteorSpawn.Reset()

		m := NewMetoro()
		g.meteoros = append(g.meteoros, m)
	}

	for _, m := range g.meteoros {
		m.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, l := range g.lasers {
		l.Draw(screen)
	}

	for _, m := range g.meteoros {
		m.Draw(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidthX, screenHeightY
}

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}
