package flip

import "testing"

func TestFlips(t *testing.T) {
	cases := []struct {
		name string
		stack string // pancake stack
		expected int // expected flips
	}{
		{
			"case 1",
			"-",
			1,
		},
		{
			"case 2",
			"-+",
			1,
		},
		{
			"case 3",
			"+-",
			2,
		},
		{
			"case 4",
			"+++",
			0,
		},
		{
			"case 5",
			"--+-",
			3,
		},
	}

	for _, c := range cases {
		actual := Flips(c.stack)

		if actual != c.expected {
			t.Errorf("for case '%v' expected %v but go %v", c.name, c.expected, actual)
		}
	}
}

func ExampleFlipAll() {
	stacks := []string {
		"-",
		"-+",
		"+-",
		"+++",
		"--+-",
	}

	FlipAll(stacks)

	// Output:
	// Case #1: 1
	// Case #2: 1
	// Case #3: 2
	// Case #4: 0
	// Case #5: 3
}

func BenchmarkFlips(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		// 10 flips
		Flips("-+-+-+-+-+-")
	}
}

func BenchmarkFlipsLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		// 100 flips
		Flips("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")
	}
}