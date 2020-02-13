package ebitenui

import (
	"github.com/hajimehoshi/ebiten"
)

type ButtonTrigger int

const (
	// ButtonPress - OnClick fires as soon as it is clicked
	ButtonPress ButtonTrigger = iota
	// ButtonRelease - OnClick fires only when the mouse button is release
	// (cursor must be over button on both click and release)
	ButtonRelease
)

type Button struct {
	x            float64
	y            float64
	DefaultImage *ebiten.Image
	HoverImage   *ebiten.Image
	PressedImage *ebiten.Image
	OnClick      OnClick
	triggersOn   ButtonTrigger
	pressed      bool
}

func NewButton(image *ebiten.Image, onClick OnClick) *Button {
	newButton := Button{}
	newButton.DefaultImage = image
	newButton.HoverImage = image
	newButton.PressedImage = image
	newButton.OnClick = onClick

	newButton.pressed = false
	newButton.triggersOn = ButtonRelease

	return &newButton
}

func (b *Button) SetTriggersOn(trigger ButtonTrigger) {
	b.triggersOn = trigger
	b.pressed = false
}

func (b *Button) cursorIsOver(offsetX, offsetY float64) bool {
	cursorX, cursorY := ebiten.CursorPosition()
	x1 := int(offsetX + b.x)
	x2 := x1 + b.DefaultImage.Bounds().Dx()
	y1 := int(offsetY + b.y)
	y2 := y1 + b.DefaultImage.Bounds().Dy()

	if (cursorX >= x1 && cursorX <= x2) && (cursorY >= y1 && cursorY <= y2) {
		_, _, _, a := b.DefaultImage.At(cursorX-int(offsetX+b.x), cursorY-int(offsetY+b.y)).RGBA()
		if a != 0.0 {
			return true
		}
	}
	return false
}

func (b *Button) Update(screen *ebiten.Image, offsetX float64, offsetY float64) {
	if b.cursorIsOver(offsetX, offsetY) {
		if b.triggersOn == ButtonPress {
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				if !b.pressed {
					b.pressed = true
					b.OnClick()
				}
			} else if b.pressed {
				b.pressed = false
			}
		} else if b.triggersOn == ButtonRelease {
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				if !b.pressed {
					b.pressed = true
				}
			} else if b.pressed {
				b.OnClick()
				b.pressed = false
			}
		}
	}
	screen.DrawImage(b.DefaultImage, &ebiten.DrawImageOptions{})
}
