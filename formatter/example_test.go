package formatter

import (
	"fmt"
)

func ExampleFormat() {
	input := map[string]string{
		"Hello": "world!",
	}

	output, _ := Format(input, "json")

	fmt.Println(output)
	// Output:
	// {
	//   "Hello": "world!"
	// }
}
