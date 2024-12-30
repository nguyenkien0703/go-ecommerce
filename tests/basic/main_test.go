package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input = 1 
	// 	output = 2 
	// )
	// actual := AddOne(input)
	// if actual != output {	
	// 	t.Errorf("AddOne(%d) = %d; want %d", input, actual, output)
	// }

	assert.Equal(t, AddOne(1), 2, "AddOne(1) should return 2")

	assert.NotEqual(t, AddOne(1), 3, "AddOne(1) should not return 3")
	assert.Nil(t, nil, "nil should be nil")



}

func TestAddOne2(t *testing.T) {
	var (
		input = 1 
		output = 2 
	)
	actual := AddOne(input)
	if actual != output {	
		t.Errorf("AddOne(%d) = %d; want %d", input, actual, output)
	}

	



}


func TestRequire(t *testing.T) {
	require.Equal(t, 1,2)
	fmt.Println("not executed")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 1,2)
	fmt.Println("executed")

}




