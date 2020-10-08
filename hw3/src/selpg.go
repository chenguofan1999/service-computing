package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/jessevdk/go-flags"
)

// Option describes The flags of SelPg
type Option struct {
	StartPg  int    `short:"s" required:"true" description:"Start page"`
	EndPg    int    `short:"e" required:"true" description:"End page"`
	LineOfPg int    `short:"l" default:"72" description:"How many lines in a page"`
	Fmode    bool   `short:"f" description:"Using /f as page "`
	Dest     string `short:"d" default:"" description:"Send pages to ? instead of os.Stdout"`
}

var opt Option
var fileName string

func main() {
	processArgs()
	processInput()
}

func processArgs() {

	// remainingArgs are filenames
	remainingArgs, _ := flags.Parse(&opt)

	// When filename not specified
	if len(remainingArgs) == 0 {
		fileName = ""
	} else {
		fileName = remainingArgs[0]
	}

	lineCount := 0
	if fileName != "" {
		// open fileName -> f
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: '%s' is Not a fileName \n", fileName)
			os.Exit(1)
		}
		defer f.Close()

		// get lineCount
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			lineCount++
		}
	}

	// check arg negative
	if opt.StartPg <= 0 {
		fmt.Fprintf(os.Stderr, "Error: page number <= 0\n")
		os.Exit(2)
	}

	// check arg beyond lineCount
	if fileName != "" && !opt.Fmode && opt.EndPg > lineCount {
		fmt.Fprintf(os.Stderr, "Error: page number beyond file\n")
		os.Exit(3)
	}

	// check start > end
	if opt.StartPg > opt.EndPg {
		fmt.Fprintf(os.Stderr, "Error: startPage > endPage\n")
		os.Exit(4)
		//log.Fatal("Error: startPage > endPage")
	}

	// check mode conflict
	if opt.Fmode && opt.LineOfPg != 72 {
		fmt.Fprintf(os.Stderr, "Mode conflict, don't use -f with -l\n")
		os.Exit(5)
	}

}

func processInput() {

	// write to dest or Stdout
	var writer *bufio.Writer
	if opt.Dest != "" {
		// write to dest
		cmd := exec.Command("lp", fmt.Sprintf("-d%s", opt.Dest))
		stdinPipe, _ := cmd.StdinPipe()
		defer stdinPipe.Close()
		cmd.Stdout = os.Stdout
		writer = bufio.NewWriter(stdinPipe)
	} else {
		//write to Stdout
		writer = bufio.NewWriter(os.Stdout)
	}
	defer writer.Flush()

	// Which file to use:
	// 1. If filename given : fileName
	// 2. Else : create a tempFile from Stdin
	var file *os.File
	if fileName != "" {
		// use given filename
		file, _ = os.Open(fileName)
	} else {
		// create a temp file
		file, _ = ioutil.TempFile("./", "temp")
		defer os.Remove(file.Name())

		tempScanner := bufio.NewScanner(os.Stdin)
		for tempScanner.Scan() {
			file.WriteString(tempScanner.Text())
			file.WriteString("\n")
		}

	}
	defer file.Close()

	if opt.Fmode {
		// Mode 2 :
		// Paging with /f
		fReader := bufio.NewReader(file)

		for i := 0; i < opt.StartPg; i++ {
			_, err := fReader.ReadString('\f')

			if err != nil || err == io.EOF {
				log.Fatal(err)
			}
		}

		for i := opt.StartPg; i <= opt.EndPg; i++ {
			page, err := fReader.ReadString('\f')

			if err == nil || err == io.EOF {
				fmt.Printf("%#v\n", page)
				if err == io.EOF {
					break
				}
			}

			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		// Mode 1 :
		// Paging by calculate lines
		startLine := (opt.StartPg - 1) * opt.LineOfPg
		endLine := opt.EndPg*opt.LineOfPg - 1

		file, _ = os.Open(file.Name())
		scanner := bufio.NewScanner(file)

		for i := 0; i < startLine; i++ {
			scanner.Scan()
		}

		for i := startLine; i <= endLine; i++ {
			scanner.Scan()
			writer.WriteString(scanner.Text())
			writer.WriteByte('\n')
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	}

}
