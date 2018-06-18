package main

import (
	"os"
)

func ExampleRun() {
	pth := "test-input.txt"
	writeExampleFile(pth)
	fPth = &pth

	run()

	os.Remove(pth)

	// Output:
	// Case #1: 1
	// Case #2: 1
	// Case #3: 2
	// Case #4: 0
	// Case #5: 3
}

func writeExampleFile(pth string) {
	f, _ := os.Create(pth)

	f.WriteString("5\n")
	f.WriteString("-\n")
	f.WriteString("-+\n")
	f.WriteString("+-\n")
	f.WriteString("+++\n")
	f.WriteString("--+-\n")
	f.Close()
}
