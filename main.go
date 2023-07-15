package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"log"
	"regexp"

	"github.com/otiai10/gosseract/v2"
	"gocv.io/x/gocv"
)

func main() {
	var showWindow bool
	parseCmdLineArgs(&showWindow)
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		log.Fatalf("Error opening capture device: %v\n", err)
	}
	defer webcam.Close()

	imagePath := "captured_image.jpg"
	if err := captureImage(webcam, imagePath, showWindow); err != nil {
		log.Fatal("Failed to capture image:", err)
	}

	text, err := readTextFromImage(imagePath)
	if err != nil {
		log.Fatal("Failed to read text from image:", err)
	}

	r := regexp.MustCompile(`\d{4,}`)
	match := r.Find([]byte(text))

	fmt.Println("Matched string:", text)
	fmt.Println("=====================")
	fmt.Println("OTP:", string(match))
}

func parseCmdLineArgs(showWindow *bool) {
	flag.BoolVar(showWindow, "show-window", false, "flag to show window")
	flag.BoolVar(showWindow, "sw", false, "flag to show window (shorthand)")
	flag.Parse()
}

// Capture image from webcam and save to file
func captureImage(webcam *gocv.VideoCapture, imagePath string, showWindow bool) error {
	var window *gocv.Window
	if showWindow {
		fmt.Println("show window set to true")
		window = gocv.NewWindow("Captura")
		defer window.Close()
	}

	for {
		img := gocv.NewMat()
		defer img.Close()

		if ok := webcam.Read(&img); !ok {
			return fmt.Errorf("cannot read from device")
		}

		var key int
		if showWindow {
			gocv.PutText(&img, "Press Space to capture", image.Point{X: 10, Y: 30}, gocv.FontHersheyPlain, 2, color.RGBA{255, 0, 0, 0}, 5)
			window.IMShow(img)
			key = window.WaitKey(1)
			if key == 27 { // Press 'Esc' key to exit the loop
				break
			}
		}

		if !showWindow || key == 32 {
			if ok := gocv.IMWrite(imagePath, img); !ok {
				return fmt.Errorf("failed to write image to disk")
			}
			break
		}
	}

	return nil
}

func readTextFromImage(imagePath string) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetLanguage("eng")

	client.SetImage(imagePath)

	text, err := client.Text()
	if err != nil {
		return "", err
	}

	return text, nil
}
