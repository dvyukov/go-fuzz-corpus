// Copyright 2015 Dmitry Vyukov. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package ttf

import (
	"image"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

var img = image.NewNRGBA64(image.Rectangle{image.Point{0, 0}, image.Point{59, 39}})

func Fuzz(data []byte) int {
	f, err := freetype.ParseFont(data)
	if err != nil {
		if f != nil {
			panic("font is not nil on error")
		}
		return 0
	}
	ctx := freetype.NewContext()
	ctx.SetFont(f)
	ctx.SetSrc(image.Black)
	ctx.SetHinting(font.HintingFull)
	ctx.SetDst(img)
	ctx.SetDPI(51)
	ctx.SetFontSize(9)
	if _, err = ctx.DrawString("go-фузз", freetype.Pt(1, 3)); err != nil {
		panic(err)
	}
	return 1
}
