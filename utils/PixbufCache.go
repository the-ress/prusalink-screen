package utils

import (
	"sync"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type PixbufCache struct {
	sync.Mutex

	cache map[pixbufCacheKey]*gdk.Pixbuf
}

type pixbufCacheKey struct {
	imageFileName string
	width         int
	height        int
}

func CreatePixbufCache() *PixbufCache {
	thisInstance := &PixbufCache{
		cache: map[pixbufCacheKey]*gdk.Pixbuf{},
	}

	return thisInstance
}

func (this *PixbufCache) getPixbuf(imageFileName string, width, height int) *gdk.Pixbuf {
	this.Lock()
	defer this.Unlock()

	key := pixbufCacheKey{
		imageFileName: imageFileName,
		width:         width,
		height:        height,
	}

	pixbuf, found := this.cache[key]
	if found {
		return pixbuf
	}

	pixbuf = MustPixbufFromFileWithSize(imageFileName, width, height)
	this.cache[key] = pixbuf

	return pixbuf
}

func (this *PixbufCache) MustImageFromFileWithSize(imageFileName string, width, height int) *gtk.Image {
	return MustImageFromPixbuf(this.getPixbuf(imageFileName, width, height))
}
