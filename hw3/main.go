package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

// Option describes The flags of SelPg
type Option struct {
	StartPg  int    `short:"s" required:"true" description:"Start page"`
	EndPg    int    `short:"e" required:"true" description:"End page"`
	LineOfPg int    `short:"l" default:"72" description:"How many lines in a page"`
	Fmode    bool   `short:"f" description:"Using f or not"`
	Dest     string `short:"d" description:"Send pages to ? instead of os.Stdout"`
}

var opt Option
var fileName string

func main() {
	process_args()
	process_input()

	// fmt.Println(opt)
	// fmt.Println("File name: ", remainingArgs)

	//io.Copy(os.Stdout, f)

	// // create a tmpFile to store all lines
	// tmpFile, err := ioutil.TempFile("/home/chen/service-computing/hw3", "tempFile")
	// if err != nil {
	// 	panic("1")
	// }
	// //defer os.Remove(tmpFile.Name())

	// _, err = io.Copy(tmpFile, f)
	// if err != nil {
	// 	panic("2")
	// }

	// // test!
	// f, err = os.Open(tmpFile.Name())
	// _, err = io.Copy(os.Stdout, f)
	// if err != nil {
	// 	panic("3")
	// }
}

func process_args() {

	remainingArgs, err := flags.Parse(&opt)
	// remainingArgs are filenames

	if err != nil {
		panic(err)
	}

	// When filename not specified
	if len(remainingArgs) == 0 {
		_, err := fmt.Scanf("%s", &fileName)
		if err != nil {
			panic(err)
		}
	} else {
		fileName = remainingArgs[0]
	}

	// open fileName -> f
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	// get lineCount

	// check arg range
	if opt.StartPg < 0 || opt.EndPg < 0 {
		panic("Error: page number < 0")
	}

	// check start > end
	if opt.StartPg > opt.EndPg {
		panic("Error: startPage > endPage")
	}

	// check mode conflict
	if opt.Fmode && opt.LineOfPg != 72 {
		panic("You can't use -l with -f")
	}
}

func process_input() {

}
