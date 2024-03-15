package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image/color"
	"log"
	"os"
)

func detectFaceAndEyes(input gocv.Mat, faceCascade, eyeCascade gocv.CascadeClassifier) gocv.Mat {
	result := input.Clone()

	faces := faceCascade.DetectMultiScale(input)
	for _, face := range faces {
		gocv.Rectangle(&result, face, color.RGBA{G: 255}, 3)

		roi := input.Region(face)

		eyes := eyeCascade.DetectMultiScale(roi)
		for _, eye := range eyes {
			eye.Min.X += face.Min.X
			eye.Min.Y += face.Min.Y
			eye.Max.X += face.Min.X
			eye.Max.Y += face.Min.Y

			gocv.Rectangle(&result, eye, color.RGBA{B: 255}, 3)
		}

		roi.Close()
	}

	return result
}

func detectFromCamera() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatalf("Couldn't open webcam: %v", err)
	}
	defer webcam.Close()

	window := gocv.NewWindow("webcamwindow")
	defer window.Close()

	faceCascade := gocv.NewCascadeClassifier()
	if !faceCascade.Load("haarcascade_frontalface_default.xml") {
		log.Fatalf("Error loading face cascade")
	}
	defer faceCascade.Close()

	eyeCascade := gocv.NewCascadeClassifier()
	if !eyeCascade.Load("haarcascade_eye.xml") {
		log.Fatalf("Error loading eye cascade")
	}
	defer eyeCascade.Close()

	for {
		img := gocv.NewMat()
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Println("Unable to read from the webcam")
			continue
		}

		result := detectFaceAndEyes(img, faceCascade, eyeCascade)

		window.IMShow(result)
		window.WaitKey(20)
	}
}

func detectFromImage(imagePath string) {
	img := gocv.IMRead(imagePath, gocv.IMReadColor)
	if img.Empty() {
		log.Fatalf("Error loading image: %s", imagePath)
	}

	window := gocv.NewWindow("imageWindow")
	defer window.Close()

	faceCascade := gocv.NewCascadeClassifier()
	if !faceCascade.Load("haarcascade_frontalface_default.xml") {
		log.Fatalf("Error loading face cascade")
	}
	defer faceCascade.Close()

	eyeCascade := gocv.NewCascadeClassifier()
	if !eyeCascade.Load("haarcascade_eye.xml") {
		log.Fatalf("Error loading eye cascade")
	}
	defer eyeCascade.Close()

	result := detectFaceAndEyes(img, faceCascade, eyeCascade)

	window.IMShow(result)
	window.WaitKey(0)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <arg1>")
		os.Exit(1)
	}

	args := os.Args[1]
	if args == "camera" {
		detectFromCamera()
	} else {
		detectFromImage(args)
	}
}
