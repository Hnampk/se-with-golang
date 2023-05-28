package solid

import "fmt"

type Adder interface {
	Add(int, int) int
}

func PrintSum(a, b int, adder Adder) {
	fmt.Printf("%d + %d = %d", a, b, adder.Add(a, b))
}

// Int adds two integer values.
type Int struct{}

// Add returns the sum a+b.
func (Int) Add(a, b int) int { return a + b }

// Double adds two double values.
type Double struct{}

// Add returns the sum a+b.
func (Double) Add(a, b float64) float64 { return a + b }
