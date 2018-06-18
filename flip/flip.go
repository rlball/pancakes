package flip

import "fmt"

const (
	minus = '-'
)

// FlipAll will show the flips of each stack
// of pancakes by printing the number of flips
// to stdout.
func FlipAll(stacks []string) {
	for c, stack := range stacks {
		f := Flips(stack)
		fmt.Printf("Case #%d: %d\n", c+1, f)
	}
}

// Flips calculates the number of pancake flips
// needed for all smiles for a particular pancake
// stack.
//
// If len(stack) == 0 then flips will panic.
func Flips(stack string) int {
	changes := 0
	current := stack[0]

	// the number of flips is mostly a measure
	// of the number of times the stacks changes
	// from a '+' to a '-'.
	for i := 0; i < len(stack[1:]); i++ {
		if current != stack[i+1] {
			current = stack[i+1]
			changes++
		}
	}

	// a '-' at the last value means an extra flip.
	if stack[len(stack)-1] == minus {
		changes++
	}

	return changes
}