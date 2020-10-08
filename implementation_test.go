package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToPrefix(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{ "2 2 +", "+ 2 2" },
		{ "4 2 - 3 * 5 +", "+ * - 4 2 3 5" },
	}

	for _, c := range cases {
		res, err := PostfixToPrefix(c.in)
		if assert.Nil(t, err) {
			assert.Equal(t, c.want, res)
		}
	}
}

func ExamplePostfixToPrefix() {
	res, _ := PostfixToPrefix("2 2 +")
	fmt.Println(res)

	// Output:
	// + 2 2
}
