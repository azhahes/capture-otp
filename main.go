package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/otiai10/gosseract/v2"
	"gocv.io/x/gocv"
)

func main() {
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		log.Fatalf("Error opening capture device: %v\n", err)
	}
	defer webcam.Close()

	imagePath := "captured_image.jpg"
	if err := captureImage(webcam, imagePath); err != nil {
		log.Fatal("Failed to capture image:", err)
	}

	// Read text from the captured image using Tesseract OCR
	text, err := readTextFromImage(imagePath)
	if err != nil {
		log.Fatal("Failed to read text from image:", err)
	}

	fmt.Println("Text from the image:", text)
}

// Capture image from webcam and save to file
func captureImage(webcam *gocv.VideoCapture, imagePath string) error {
	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		return fmt.Errorf("cannot read from device")
	}

	// Add a text overlay to the image
	text := "Press any key to capture"
	gocv.PutText(&img, text, image.Point{X: 10, Y: 30}, gocv.FontHersheyPlain, 1.2, color.RGBA{255, 255, 255, 0}, 2)

	window := gocv.NewWindow("Press any key to capture")
	defer window.Close()

	for {
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}

	// Save the image
	if ok := gocv.IMWrite(imagePath, img); !ok {
		return fmt.Errorf("failed to write image to disk")
	}

	return nil
}

// Read text from the captured image using Tesseract OCR
func readTextFromImage(imagePath string) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	// Set the language for OCR (default is "eng")
	client.SetLanguage("eng")

	// Set the path to the image file
	client.SetImage(imagePath)

	// Perform OCR on the image
	text, err := client.Text()
	if err != nil {
		return "", err
	}

	return text, nil
}
