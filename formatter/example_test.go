package formatter

import (
	"fmt"
)

func ExampleFormat() {
	input := map[string]string{
		"Hello": "world!",
	}

	formatter, _ := Get("json")

	output, _ := formatter(input)

	fmt.Println(output)
	// Output:
	// {
	//   "Hello": "world!"
	// }
}
