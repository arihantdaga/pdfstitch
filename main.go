package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	// "gopkg.in/gographics/imagick.v3/imagick"
)

const MEDIA_DIRECTORY = "./media/"

var counter = 0

func main() {
	fmt.Println("Starting...")
	// PDF_NAMES := []string{
	// 	"AmazonFile_0.pdf",
	// }
	// imagick.Initialize()
	// defer imagick.Terminate()
	// mw := imagick.NewMagickWand()
	// defer mw.Destroy()
	// LoadPDF(PDF_NAMES[0], mw)
	readMediaDirectory()
	// executeCommand("sleep", "3")
	// executeCommand("sleep", "2")
}

func readMediaDirectory() {
	var multiple = false
	files, err := ioutil.ReadDir(MEDIA_DIRECTORY)
	if err != nil {
		log.Fatal(err)
	}
	pdfFiles := []string{}
	imageFiles := []string{}
	for _, f := range files {
		// fmt.Println(f.Name())
		if filepath.Ext(f.Name()) == ".pdf" {
			pdfFiles = append(pdfFiles, f.Name())
		} else if filepath.Ext(f.Name()) == ".jpeg" {
			imageFiles = append(imageFiles, f.Name())
		}
	}
	if len(pdfFiles) != len(imageFiles) {
		log.Fatal("PDF and image files are not the same number")
	}
	if len(pdfFiles) > 1 {
		multiple = true
	}
	for i := 0; i < len(pdfFiles); i++ {
		var outLocation string
		if i%2 == 0 {
			outLocation = "top.jpg"
		} else {
			outLocation = "bottom.jpg"
		}
		step1(MEDIA_DIRECTORY + pdfFiles[i])
		step2(MEDIA_DIRECTORY + imageFiles[i])
		step3(pdfFiles[i], pdfFiles[i], outLocation)
		if i%2 != 0 || i == len(pdfFiles)-1 {
			if i%2 == 0 {
				multiple = false
			}
			step4(multiple)
			step5()
			step6()
			counter++
		}
	}
	if len(pdfFiles)%2 != 0 {
		step6()
	}
}
func step1(invoiceFileName string) {
	args := strings.Fields("-density 300 -resize 987x1480^ -background white -alpha remove ")
	args = append(args, invoiceFileName, "invoice.jpg")
	executeCommand("convert", args...)
}
func step2(labelFileName string) {
	args := strings.Fields("-density 300 -resize 987x1480^ -background white -alpha remove ")
	args = append(args, labelFileName, "label.jpg")
	executeCommand("convert", args...)
}
func step3(invoiceFileName string, labelFileName string, outLocation string) {
	args := strings.Fields("label.jpg invoice.jpg -gravity South +append " + outLocation)
	executeCommand("convert", args...)
}

func step4(multiple bool) {
	var args []string
	if multiple {
		args = strings.Fields("top.jpg bottom.jpg -append final.jpg")
	} else {
		args = strings.Fields("top.jpg -append final.jpg")
	}
	executeCommand("convert", args...)
}

func step5() {
	args := strings.Fields(fmt.Sprintf("final.jpg -quality 100 final%d.pdf", counter))
	executeCommand("convert", args...)
}

// Cleanup
func step6() {
	// clean media files
	args := strings.Fields("final.jpg top.jpg bottom.jpg")
	executeCommand("rm", args...)
}

func executeCommand(command string, args ...string) {
	fmt.Println(command, args)
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	// time.Sleep(3 * time.Second)
}

// func LoadPDF(f string, mw *imagick.MagickWand) {
// 	mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_DEACTIVATE)
// 	mw.SetImageAlpha(0)
// 	mw.ReadImage(MEDIA_DIRECTORY + f)
// 	mw.SetIteratorIndex(0)
// 	mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_DEACTIVATE)
// 	mw.SetImageAlpha(0)
// 	mw.SetImageFormat("jpg")
// 	mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_DEACTIVATE)
// 	mw.SetImageAlpha(0)
// 	mw.WriteImage("out.jpg")
// 	fmt.Println("Done")
// }

// convert  -density 300 -resize 1100x1649^ -background white -alpha remove ./media/AmazonFile_0.pdf invoice.jpg
// convert  -density 300 -resize 1100x1649^ -background white -alpha remove './media/WhatsApp Image 2022-07-25 at 1.21.29 AM.jpeg' label.jpg

// convert label.jpg invoice.jpg -gravity South +append top.jpg
// convert label.jpg out.jpg -gravity South +append bottom.jpg
// convert top.jpg bottom.jpg -append final.jpg
// convert final.jpg -quality 100  final.pdf

// # Cleanup
