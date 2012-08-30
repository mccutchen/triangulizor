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
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	log.Printf("Processing %dx%d tile", w, h)
	colors := make([]color.Color, w*h)
	for y := 0; y < h; y += 1 {
		for x := 0; x < w; x += 1 {
			colors[x*y] = img.At(x, y)
		}
	}
	if len(colors) > w*h {
		log.Fatalf("Too many colors in tile! Expected %d, got %d", w*h, len(colors))
	}
	log.Printf("Got %d colors for tile", len(colors))
	avgColor := getAverageColor(colors)
	log.Printf("Avg color for tile: %v", avgColor)
}

func getAverageColor(colors []color.Color) color.Color {
	var tr, tg, tb, ta uint8
	var colorCount = uint8(len(colors))
	colorCount = 1
	log.Printf("Averaging %d colors...", len(colors))
	for i, c := range colors {
		log.Printf("Color %04d: %v", i, c)
		r, g, b, a := c.RGBA()
		tr += uint8(r)
		tg += uint8(g)
		tb += uint8(b)
		ta += uint8(a)
	}
	return color.RGBA{tr / colorCount, tg / colorCount, tb / colorCount, ta / colorCount}
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
