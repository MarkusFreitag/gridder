package gridder

import (
	"image"
	"image/gif"

	"github.com/andybons/gogif"
)

func (g *Gridder) scopedCtx(fn func()) {
	g.ctx.Push()
	defer g.ctx.Pop()

	fn()

	if g.animation != nil {
		bounds := g.ctx.Image().Bounds()
		palettedImage := image.NewPaletted(bounds, nil)
		quantizer := gogif.MedianCutQuantizer{NumColor: 64}
		quantizer.Quantize(palettedImage, bounds, g.ctx.Image(), image.ZP)

		g.animation.Image = append(g.animation.Image, palettedImage)
		g.animation.Delay = append(g.animation.Delay, 5)
		g.animation.Disposal = append(g.animation.Disposal, gif.DisposalBackground)
	}
}
