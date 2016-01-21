package gotee

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ExampleNewTee() {
	tee, err := NewTee("tee.out")
	if err != nil {
		fmt.Printf("ERROR: tee: %s\n", err)
		return
	}
	// you can write headers, dates etc just to the
	// tee file if necessary
	fmt.Fprintf(tee.OrigStdout, "write to stdout but not tee file\n")
	fmt.Fprintf(tee.TeeFile, "test to tee file\n")
	fmt.Fprintf(os.Stderr, "test to stderr\n")
	fmt.Printf("test to stdout\n")
	tee.Close()
	fmt.Printf("more to stdout\n")

	teed, _ := ioutil.ReadFile("tee.out")
	fmt.Printf("teed output:\n%s", teed)

	// Output:
	// write to stdout but not tee file
	// test to stdout
	// more to stdout
	// teed output:
	// test to tee file
	// test to stderr
	// test to stdout
}

func ExampleNewTee_CloseOnPanic() {
	tee, err := NewTee("tee.log")
	if err != nil {
		fmt.Printf("ERROR: tee: %s\n", err)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			tee.Close()
			// handle panic

			teed, _ := ioutil.ReadFile("tee.log")
			fmt.Printf("teed output:\n%s", teed)
		}
	}()

	fmt.Printf("stdout\n")
	panic(fmt.Errorf("whoops"))

	// Output:
	// stdout
	// teed output:
	// stdout
}
