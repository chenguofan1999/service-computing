package main

import (
	"os"
	"testing"
)

func TestArgsProcess(t *testing.T) {
	args := [][]string{
		[]string{os.Args[0], "-s1", "-e3", "-l2"},
		[]string{os.Args[0], "-s2", "-e3", "-l2"},
		[]string{os.Args[0], "-s1", "-e5"}}

	expected := [3][3]int{
		{1, 3, 2},
		{2, 3, 2},
		{1, 5, 72}}

	for i, arg := range args {
		os.Args = arg
		processArgs()
		if opt.StartPg != expected[i][0] || opt.EndPg != expected[i][1] || opt.LineOfPg != expected[i][2] {
			t.Error("not right")
		}
	}
}
