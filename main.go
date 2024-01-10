package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
)

func main() {

	// Create a new gin router
	router := gin.Default()

	//Handler for the root path
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello and welcome to the watermark app, hit the /watermark for more!"})
	})

	// Handler for the healthcheck
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	// Define a handler function to handle the watermark request
	router.POST("/watermark", func(c *gin.Context) {
		// Get the input parameters from the request
		url := c.PostForm("url")
		text := c.PostForm("text")

		// Download the image from the URL
		log.Println("Downloading the image from", url)
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Failed to download the image", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to download the image"})
			return
		}
		defer resp.Body.Close()

		// Decode the image
		img, err := imaging.Decode(resp.Body)
		if err != nil {
			log.Println("Failed to decode the image", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode the image"})
			return
		}
		// image size
		imgSize := img.Bounds().Size()

		// create a string watermark with a name
		log.Println("Creating the watermark on image with text", text)
		watermark := createWatermark(text, imgSize.X, imgSize.Y)

		// Add the overlay to the background image with 50% opacity
		result := imaging.Overlay(img, watermark, image.Point{0, 0}, 0.5)

		// Save the result image to file
		thirdImage, err := os.Create("image-with-overlay.jpg")
		if err != nil {
			log.Fatalf("Failed to create: %s", err)
		}
		jpeg.Encode(thirdImage, result, &jpeg.Options{Quality: 100})
		defer thirdImage.Close()

		// Send the result image to the client
		c.Header("Content-Type", "image/jpeg")
		c.File("image-with-overlay.jpg")
		defer os.Remove("image-with-overlay.jpg")
		return

	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func createWatermark(name string, bgWidth, bgHeight int) *image.RGBA {
	// Create a new RGBA image
	img := image.NewRGBA(image.Rect(0, 0, bgWidth, bgHeight))

	// Load the font and create a font face
	fnt, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}

	// Create a font face with a proportional font size
	face := truetype.NewFace(fnt, &truetype.Options{
		Size: float64(bgWidth / 10),
	})

	// Calculate the width of the text using the font face
	width := font.MeasureString(face, name).Round()

	// Calculate the starting point of the text to center it
	// This will render from right to left
	startX := (img.Bounds().Dx() / 2) - width/2

	// this will render from bottom to top
	startY := (bgHeight / 2) + width/14

	// Draw the text onto the image
	col := color.RGBA{0, 0, 0, 128}
	draw.Draw(img, img.Bounds(), image.NewUniform(col), image.Point{}, draw.Src)

	watermark := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.White),
		Face: face,
		Dot:  fixed.P(startX, startY),
	}

	// text shadow
	shadow := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.Black),
		Face: face,
		Dot:  fixed.P(startX+2, startY+2),
	}

	shadow.DrawString(name)
	watermark.DrawString(name)

	return img
}
