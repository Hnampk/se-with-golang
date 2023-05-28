package solid

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	PrintSum(1, 2, Int{}) // prints: "1 + 2 = 3"
	// This line will trigger a compile-time error:
	// cannot use adder.Double literal (type adder.Double) as type Adder
	// in argument to PrintSum: adder.Double does not implement Adder
	// (wrong type for Add method)
	// have Add(float64, float64) float64
	// want Add(int, int) int
	// PrintSum(1, 2, Double{})
}

// In the cases where the type of the object to be substituted is not known at compile time, the
// compiler will automatically generate code to perform the check at runtime
func TestExplain(t *testing.T) {
	var placeholder interface{}
	// Cast to io.Reader works; os.Stdin implements io.Reader
	placeholder = os.Stdin
	_ = placeholder.(io.Reader)
	// Cast to io.Reader triggers a run-time panic:
	// "panic: interface conversion: string is not io.Reader: missing method Read"
	// placeholder = "cast check"
	// _ = placeholder.(io.Reader)

	/*TIP
	When you aren't sure whether a type instance or an interface{} can be
	cast to another type or interface at runtime, it is often good practice to use
	the dual return value variant of the cast operator
	to avoid potential panics while your program is executing.
	*/
	// Cast to io.Reader fails and isReader is set to false
	placeholder = "cast check"
	if _, isReader := placeholder.(io.Reader); !isReader {
		fmt.Printf("%T does not implement io.Reader\n", placeholder)
	}
}
