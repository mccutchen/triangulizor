package main

import (
	//"flag"
	//"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func Triangulize(origImg image.Image, tileSize int) (image.Image, error) {
	log.Printf("Triangulizing image with size %dx%d",
		origImg.Bounds().Dx(), origImg.Bounds().Dy())
	img, err := prepImage(origImg, tileSize)
	if err != nil {
		log.Fatalf("Unable to prepare image for processing: %s", err)
	}
	tilePoints := getTilePoints(img, tileSize)
	log.Printf("Image pixels: %d", len(img.Pix))
	log.Printf("Tile count:   %v", len(tilePoints))
	log.Printf("Tile coords:  %v", tilePoints)
	for _, pt := range tilePoints {
		tileBounds := image.Rect(pt.X, pt.Y, pt.X+tileSize, pt.Y+tileSize)
		tile := imageToRGBA(img.SubImage(tileBounds))
		processTile(tile)
	}
	return img, nil
}

func imageToRGBA(src image.Image) (dst *image.RGBA) {
	bounds := src.Bounds()
	dst = image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(dst, dst.Bounds(), src, bounds.Min, draw.Src)
	return
}

func prepImage(img image.Image, tileSize int) (*image.RGBA, error) {
	newImg := imageToRGBA(img)
	return newImg, nil
}

func processTile(img *image.RGBA) {
	b := img.Bounds()
	w := b.Dx()
	h := b.Dy()
	log.Printf("Processing %dx%d tile", w, h)
	colors := make([]color.Color, w*h)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.Y; x++ {
			c := img.At(x, y)
			if c == nil {
				log.Fatalf("Nil color!")
			}
			colors[x*y] = c
		}
	}
	avgColor := getAverageColor(colors)
	log.Printf("Avg color for tile: %v", avgColor)
}

func getAverageColor(colors []color.Color) color.Color {
	var tr, tg, tb, ta uint8
	count := uint8(len(colors))
	log.Printf("Averaging %d colors...", len(colors))
	nils := 0
	for _, c := range colors {
		// log.Printf("Color %04d: %v", i, c)
		if c == nil {
			nils++
			// log.Printf("Color %04d is nil", i)
			continue
		}
		r, g, b, a := c.RGBA()
		tr += uint8(r)
		tg += uint8(g)
		tb += uint8(b)
		ta += uint8(a)
	}
	log.Printf("%d nil colors in tile", nils)
	return color.RGBA{tr / count, tg / count, tb / count, ta / count}
}

func getTilePoints(img *image.RGBA, tileSize int) (tilePts []image.Point) {
	inWidth := img.Bounds().Dx()
	inHeight := img.Bounds().Dy()
	xTileCount := inWidth / tileSize
	yTileCount := inHeight / tileSize
	outWidth := xTileCount * tileSize
	outHeight := yTileCount * tileSize
	xOff := (inWidth - outWidth) / 2
	yOff := (inHeight - outHeight) / 2

	for y := 0; y < outHeight; y += tileSize {
		for x := 0; x < outWidth; x += tileSize {
			tilePts = append(tilePts, image.Point{x + xOff, y + yOff})
		}
	}
	return
}

func main() {
	log.Println("Triangulizor!")

	imgPath := "examples/in.jpg"
	imgFile, err := os.Open(imgPath)
	if err != nil {
		log.Fatalf("Unable to open file at %s: %s", imgPath, err)
	}
	defer imgFile.Close()

	img, format, err := image.Decode(imgFile)
	if err != nil {
		log.Fatalf("Unable to decode image from %s: %s", imgPath, err)
	}
	log.Printf("Decoded %s image", format)

	img, err = Triangulize(img, 40)
	if err != nil {
		log.Fatalf("Error triangulizing image %s: %s", imgPath, err)
	}
}
