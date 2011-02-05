// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package walk

import (
	"os"
)

type ImageView struct {
	CustomWidget
	image Image
}

func NewImageView(parent IContainer) (*ImageView, os.Error) {
	iv := &ImageView{}

	cw, err := NewCustomWidget(parent, 0, func(surface *Surface, updateBounds Rectangle) os.Error {
		return iv.drawImage(surface, updateBounds)
	})
	if err != nil {
		return nil, err
	}

	iv.CustomWidget = *cw

	iv.SetInvalidatesOnResize(true)

	widgetsByHWnd[iv.hWnd] = iv
	customWidgetsByHWND[iv.hWnd] = &iv.CustomWidget

	return iv, nil
}

func (iv *ImageView) Image() Image {
	return iv.image
}

func (iv *ImageView) SetImage(value Image) os.Error {
	iv.image = value

	_, isMetafile := value.(*Metafile)
	iv.SetClearsBackground(isMetafile)

	return iv.Invalidate()
}

func (iv *ImageView) drawImage(surface *Surface, updateBounds Rectangle) os.Error {
	if iv.image == nil {
		return nil
	}

	bounds := iv.ClientBounds()

	return surface.DrawImageStretched(iv.image, bounds)
}
