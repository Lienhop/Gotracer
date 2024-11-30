package main

import (
	"bufio"
	"fmt"
	"os"
)

type Camera struct {
	center            Vec3
	aspectRatio       float64
	imageWidth        int
	imageHeight       int
	samplesPerPixel   int
	pixelSamplesScale float64

	focalLength    float64
	viewportHeight float64
	viewportWidth  float64

	viewportU   Vec3
	viewportV   Vec3
	pixelDeltaU Vec3
	pixelDeltaV Vec3

	viewportUpperLeft Vec3
	pixel00Loc        Vec3
}

func (c *Camera) init() {
	c.aspectRatio = 16.0 / 9.0
	c.imageWidth = 400
	c.imageHeight = int(float64(c.imageWidth) / c.aspectRatio)
	c.center = Vec3{0, 0, 0.5}
	c.samplesPerPixel = 100
	c.pixelSamplesScale = 1 / float64(c.samplesPerPixel)

	// Viewport Dimensions
	c.focalLength = 1.0
	c.viewportHeight = 2.0
	c.viewportWidth = c.aspectRatio * c.viewportHeight

	//Vectors across the viewport edges
	c.viewportU = Vec3{c.viewportWidth, 0.0, 0.0}
	c.viewportV = Vec3{0.0, -c.viewportHeight, 0.0}

	// Delta vectors from pixel to pixel
	c.pixelDeltaU = c.viewportU.divide(float64(c.imageWidth))
	c.pixelDeltaV = c.viewportV.divide(float64(c.imageHeight))

	// Location of upper left pixel and corner
	c.viewportUpperLeft = c.center.subtract(Vec3{0, 0, c.focalLength}).subtract(c.viewportU.divide(2)).subtract(c.viewportV.divide(2))
	c.pixel00Loc = c.viewportUpperLeft.add(c.pixelDeltaU.add(c.pixelDeltaV).scale(0.5))
}

func (c Camera) render(world hittable) {
	c.init()
	f, err := os.Create("image.ppm")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	writer := bufio.NewWriter(f)

	fmt.Fprintf(writer, "P3\n%d %d\n255\n", c.imageWidth, c.imageHeight)
	for i := 0; i < c.imageHeight; i++ {
		fmt.Printf("\rLine rendered: %d/%d", i+1, c.imageHeight)
		for j := 0; j < c.imageWidth; j++ {
			pixelColor := Color{r: 0, g: 0, b: 0}
			for s := 0; s < c.samplesPerPixel; s++ {
				ray := c.getRay(j, i)
				pixelColor = pixelColor.add(ray.rayColor(world))
			}
			pixelColor = pixelColor.scale(c.pixelSamplesScale)
			fmt.Fprintf(writer, pixelColor.writeColor())
		}
	}
	writer.Flush()
}

func (c Camera) getRay(x, y int) Ray {
	offset := sampleSquare()
	pixelSample := c.pixel00Loc.add(c.pixelDeltaU.scale(float64(x) + offset.x)).add(c.pixelDeltaV.scale(float64(y) + offset.y))
	rayOrigin := c.center
	rayDirection := pixelSample.subtract(rayOrigin)
	return Ray{origin: rayOrigin, direction: rayDirection}
}
