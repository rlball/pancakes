package main

import (
	"flag"
	"os"
	"log"
	"bufio"
	"io"

	"github.com/rlball/pancakes/flip"
)

var fPth = flag.String("file", "/dev/stdin", "input file name")

func main() {
	if err := run(); err != nil {
		log.Fatalln(err.Error())
	}
}

func run() error {
	flag.Parse()

	stacks, err := readStacks(*fPth)
	if err != nil {
		return err
	}

	flip.FlipAll(stacks)

	return nil
}

// readStacks will read from pth and return
func readStacks(pth string) ([]string, error) {
	// open
	f, err := os.Open(pth)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// read
	r := bufio.NewReader(f)
	stacks := make([]string, 0)
	for {
		ln, _, err := r.ReadLine()
		if len(ln) > 0 {
			stacks = append(stacks, string(ln))
		}

		if err != nil && err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}
	}

	if len(stacks) == 0 {
		return stacks, nil
	}

	return stacks[1:], nil
}