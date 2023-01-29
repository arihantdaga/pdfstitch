package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const MEDIA_DIRECTORY = "./media/"

var counter = 0

func main() {
	fmt.Println("Starting...")
	readMediaDirectory()
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
	args := strings.Fields("-density 100 -resize 987x1480^ -background white -alpha remove ")
	args = append(args, invoiceFileName, "invoice.jpg")
	executeCommand("convert", args...)
}
func step2(labelFileName string) {
	args := strings.Fields("-density 100 -resize 987x1480^ -background white -alpha remove ")
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
}
