package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gocv.io/x/gocv"
)

//	"fmt"

//	"io/fs"

func main() {
	fmt.Println("started up")
	srcDir := "./files/input"
	dstDir := "./files/output"
	targetWidth, targetHeight := 640, 360
	srchStr := ".pptx"
	os.MkdirAll(dstDir, os.ModePerm)

	//loop thru the filess
	files, err := os.ReadDir(srcDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Filter for image files
		ext := strings.ToLower(filepath.Ext(file.Name()))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			continue
		}

		//see if it is contains .pptx
		index := strings.Index(file.Name(), srchStr)
		if index == -1 {
			continue
		}

		namePart1 := file.Name()[0:index]
		index2 := strings.Index(file.Name(), ".png")
		namePart2 := file.Name()[index:index2]
		valStr1 := strings.ReplaceAll(namePart2, ".pptx", "")
		valStr2 := strings.ReplaceAll(valStr1, "(", "")
		valStr3 := strings.ReplaceAll(valStr2, ")", "")
		numFile := 0
		if valStr3 == "" {
			numFile = 1
		} else {
			valStr4 := strings.ReplaceAll(valStr3, " ", "0")
			num, err1 := strconv.Atoi(valStr4)
			if err1 != nil {
				log.Fatal("Conv to int error:", err1)
				continue
			}
			numFile = num + 1
		}

		namePart3 := fmt.Sprintf("%02d", numFile)
		outFileName := namePart1 + "_" + namePart3 + ".png"

		inputPath := filepath.Join(srcDir, file.Name())
		outputPath := filepath.Join(dstDir, outFileName)

		// 2. Read the image using GoCV
		img := gocv.IMRead(inputPath, gocv.IMReadColor)
		if img.Empty() {
			fmt.Printf("Error reading image: %s\n", file.Name())
			continue
		}
		defer img.Close()

		// 3. Resize the image
		resized := gocv.NewMat()
		defer resized.Close()

		gocv.Resize(img, &resized, image.Point{X: targetWidth, Y: targetHeight}, 0, 0, gocv.InterpolationDefault)

		// 4. Save the resized image
		if ok := gocv.IMWrite(outputPath, resized); !ok {
			fmt.Printf("Error saving image: %s\n", file.Name())
		} else {
			fmt.Printf("Resized and saved: %s  as  %s\n", file.Name(), outFileName)
		}
	}
}
