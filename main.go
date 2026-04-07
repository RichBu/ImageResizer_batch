package main

import "fmt"
import (
	"gocv.io/x/gocv"
)
import (
	"fmt"
	"image"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"gocv.io/x/gocv"
)


func main() {
	fmt.Println("started up")
	srcDir := "./files/input"
	dstDir := "./files/output"
	targetWidth, targetHeight :=  640,360
	os.MkdirAll(dstDir, os.ModePerm)


	//loop thru the filess
	files, err := os.ReadDir(srcDir)
	if err != nil {
		log.fatal(err)
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

        inputPath := filepath.Join(srcDir, file.Name())
        outputPath := filepath.Join(dstDir, file.Name())

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
            fmt.Printf("Resized and saved: %s\n", file.Name())
        }
    }
}

