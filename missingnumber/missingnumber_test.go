package missingnumber_test

import (
	"main/missingnumber"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	assert := assert.New(t)

	actual := missingnumber.Solutions([]int{1, 2, 3, 5, 6})
	expect := 4

	assert.Equal(expect, actual, "The two words should be the same.")

}
