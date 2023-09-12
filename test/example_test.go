package UnitTests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	request_a := 1
	request_b := 2
	executed_result := addTwoNumbers(request_a, request_b)
	assert.Equal(t, executed_result, 3)
}

func addTwoNumbers(a int, b int) int {
	return a + b + 1
}

func addSubNumber(a int, b int) int {
	return a + b
}
