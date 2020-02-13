package ebitenui

import "github.com/hajimehoshi/ebiten"

import "log"

// TODO: consider replacing offsetX, offsetY with r2.Point

type Layer struct {
	x               float64 // position of this layer within its parent layer
	y               float64
	currentSubLayer int
	subLayers       []Layer
	elements        []Element
}

func NewLayer() *Layer {
	newLayer := Layer{}
	newLayer.currentSubLayer = -1
	return &newLayer
}

func (layer *Layer) getRenderOpts() *ebiten.DrawImageOptions {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(layer.x, layer.y)
	return &opts
}

// Update contains all logic for this Element
// Takes screen *ebiten.Image to which it will draw itself,
//       offsetX, offsetY float64 - total position of parent
func (layer *Layer) Update(screen *ebiten.Image, offsetX, offsetY float64) {
	layerImg, err := ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy(), ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	for _, elem := range layer.elements {
		elem.Update(layerImg, offsetX+layer.x, offsetY+layer.y)
	}

	if layer.currentSubLayer > 0 {
		layer.subLayers[layer.currentSubLayer].Update(layerImg, offsetX+layer.x, offsetY+layer.y)
	}
	screen.DrawImage(layerImg, layer.getRenderOpts())
}

func (layer *Layer) AddElement(e Element) {
	layer.elements = append(layer.elements, e)
}

func (layer *Layer) AddSubLayer(subLayer Layer) {
	layer.subLayers = append(layer.subLayers, subLayer)
}
