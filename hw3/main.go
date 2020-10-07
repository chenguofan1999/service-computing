package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

// Option describes The flags of SelPg
type Option struct {
	StartPg  int    `short:"s" required:"true" description:"Start page"`
	EndPg    int    `short:"e" required:"true" description:"End page"`
	LineOfPg int    `short:"l" default:"72" description:"How many lines in a page"`
	HasF     bool   `short:"f" description:"Using f or not"`
	Dest     string `short:"t" description:"Send pages to ? instead of os.Stdout"`
}

func main() {
	var opt Option
	remainingArgs, err := flags.Parse(&opt)

	if err != nil {
		panic(err)
	}

	fmt.Println(opt)
	fmt.Println(remainingArgs)
}
