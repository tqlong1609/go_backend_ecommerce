package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	var (
		input  = 1
		output = 2
	)

	// result := AddOne(input)
	// if result != output {
	// 	t.Errorf("AddOne(1) = %d; expected %d; actual %d", result, output, result)
	// }

	assert.Equal(t, output, AddOne(input))

}

func TestRequire(t *testing.T) {
	require.Equal(t, 1, 2) // if fail will not execute code below
	fmt.Println("Not executing")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 1, 2) // if fail also execute code below
	fmt.Println("Executing")
}

func TestBasicCoverage(t *testing.T) {
	BasicCoverage()
}
