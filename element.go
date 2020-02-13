package ebitenui

import "github.com/hajimehoshi/ebiten"

type OnClick func()

type Element interface {
	Update(*ebiten.Image, float64, float64)
}
