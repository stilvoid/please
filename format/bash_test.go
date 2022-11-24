package format

import "testing"

func TestBash(t *testing.T) {
	expecteds := []string{
		"123",
		"456.789",
		"abc",
		"true",
		"false",
		"",
		"([0]=\"123\" [1]=\"abc\")",
		"([foo]=\"bar\")",
		"([123]=\"([0]=\\\"baz\\\" [1]=\\\"quux\\\")\")",
		"([true]=\"([null]=\\\"\\\")\")",
		"([0]=\"456\" [1]=\"def\" [2]=\"([3]=\\\"4\\\")\" [3]=\"([0]=\\\"first\\\" [1]=\\\"second\\\" [2]=\\\"([0]=\\\\\\\"deeper\\\\\\\")\\\")\")",
	}

	if len(expecteds) != len(testCases) {
		t.Fatalf("insufficient test cases implemented")
	}

	for i, expected := range expecteds {
		testCase := testCases[i]

		actual, err := formatBash(testCase)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if actual != expected {
			t.Errorf("unexpected:\n'%v'\nvs\n'%v'", actual, expected)
		}
	}
}
