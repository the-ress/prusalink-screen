package uiUtils

import (
	"sync"

	"github.com/gotk3/gotk3/gdk"
)

type PixbufCache struct {
	sync.Mutex

	cache map[pixbufCacheKey]*gdk.Pixbuf
}

type pixbufCacheKey struct {
	fileName ImageFileName
	width    int
	height   int
}

func NewPixbufCache() *PixbufCache {
	thisInstance := &PixbufCache{
		cache: map[pixbufCacheKey]*gdk.Pixbuf{},
	}

	return thisInstance
}

func (this *PixbufCache) GetPixbuf(fileName ImageFileName) *gdk.Pixbuf {
	return this.GetPixbufWithSize(fileName, -1, -1)
}

func (this *PixbufCache) GetPixbufWithSize(fileName ImageFileName, width, height int) *gdk.Pixbuf {
	this.Lock()
	defer this.Unlock()

	key := pixbufCacheKey{
		fileName: fileName,
		width:    width,
		height:   height,
	}

	pixbuf, found := this.cache[key]
	if found {
		return pixbuf
	} else {
		return nil
	}
}

func (this *PixbufCache) SetPixbuf(fileName ImageFileName, pixbuf *gdk.Pixbuf) {
	this.SetPixbufWithSize(fileName, -1, -1, pixbuf)
}

func (this *PixbufCache) SetPixbufWithSize(fileName ImageFileName, width, height int, pixbuf *gdk.Pixbuf) {
	this.Lock()
	defer this.Unlock()

	key := pixbufCacheKey{
		fileName: fileName,
		width:    width,
		height:   height,
	}

	this.cache[key] = pixbuf
}
