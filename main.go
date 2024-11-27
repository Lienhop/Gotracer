package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	aspectRatio = 16.0 / 9.0
	imgWidth    = 400
	imgHeight   = int(float64(imgWidth) / aspectRatio)

	//Camera
	focalLength    = 1.0
	viewportHeight = 2.0
	viewportWidth  = aspectRatio * viewportHeight
	cameraCenter   = Vec3{0.0, 0.0, 0}

	// Vectors across the viewport edges
	viewportU = Vec3{viewportWidth, 0.0, 0.0}
	viewportV = Vec3{0.0, -viewportHeight, 0.0}

	// Delta vectors from pixel to pixel
	pixelDeltaU = viewportU.divide(float64(imgWidth))
	pixelDeltaV = viewportV.divide(float64(imgHeight))

	// Location of upper left pixel
	viewportUpperLeft = cameraCenter.subtract(Vec3{0, 0, focalLength}).subtract(viewportU.divide(2)).subtract(viewportV.divide(2))
	pixel00Loc        = viewportUpperLeft.add(pixelDeltaU.add(pixelDeltaV).scale(0.5))
)

func (r Ray) rayColor() Color {
	unitDirection := r.direction.unitVector()
	a := 0.5 * (unitDirection.y + 1.0)
	return Color{r: 1.0, g: 1.0, b: 1.0}.scale(1.0 - a).add(Color{r: 0.5, g: 0.7, b: 1.0}.scale(a))
}

func main() {
	f, err := os.Create("image.ppm")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	writer := bufio.NewWriter(f)

	fmt.Fprintf(writer, "P3\n%d %d\n255\n", imgWidth, imgHeight)
	for i := 0; i < imgHeight; i++ {
		fmt.Printf("\rLine rendered: %d/%d", i+1, imgHeight)
		for j := 0; j < imgWidth; j++ {
			pixelCenter := pixel00Loc.add(pixelDeltaU.scale(float64(j))).add(pixelDeltaV.scale(float64(i)))
			rayDirection := pixelCenter.subtract(cameraCenter)
			ray := Ray{origin: cameraCenter, direction: rayDirection}
			pixelColor := ray.rayColor()
			//color := Color{r: float64(i) / float64(imgWidth-1), g: float64(j) / float64(imgHeight-1), b: 0.5}
			fmt.Fprintf(writer, pixelColor.writeColor())
		}
	}
	writer.Flush()
}
