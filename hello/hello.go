package hello

import (
	"image"
	"image/color"
	"image/png"
	"ftree"
	"fmt"
    "net/http"
	"perlin"
	"time"
)

var (
	f = ftree.NewFTree("add", ftree.Var("x"), ftree.Var("y"), ftree.NewFTree("mul", ftree.Var("x"), ftree.Var("y"), ftree.Const(1000)))
)

func init() {
    http.HandleFunc("/", root)
}

const (
	W = 16
	H = 16
)

func root(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, W, H))
	b := m.Bounds()
//	fmt.Printf("Bounds: %v", b)
	nRed := perlin.NewPerlinNoise(time.Now().UnixNano())
	nGreen := perlin.NewPerlinNoise(time.Now().UnixNano())
	nBlue := perlin.NewPerlinNoise(time.Now().UnixNano())

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			ftree.SetVar("x", float64(x) / W)
			ftree.SetVar("y", float64(y) / H)
			v, err := f.EvalSigmoid()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
//			fmt.Printf("%v\t%v\t%v\n", v, v * 256.0, uint8(v * 256.0))
			fmt.Printf("%v\n", v)
			red := uint8(nRed.At2d(float64(x) / W, float64(y) / H) * 256)
			red = v
			green := uint8(nGreen.At2d(float64(x) / W, float64(y) / H) * 256)
			green = v
			blue := uint8(nBlue.At2d(float64(x) / W, float64(y) / H) * 256)
			blue = v
			alpha := uint8(255) //uint8(nAlpha.At2d(float64(x) / W, float64(y) / H) * 256)
	    	m.Set(x, y, color.RGBA{red, green, blue, alpha})
		}
	}
    w.Header().Set("Content-Type", "image/png")
	if err := png.Encode(w, m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
