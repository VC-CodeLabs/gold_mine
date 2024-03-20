package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	input := [][]int{
		{1, 3, 3},
		{2, 1, 4},
		{0, 6, 4}}

	assert.Equal(t, 12, goldMine(input))
}

func TestExample2(t *testing.T) {
	input := [][]int{
		{1, 3, 1, 5},
		{2, 2, 4, 1},
		{5, 0, 2, 3},
		{0, 6, 1, 2}}

	assert.Equal(t, 16, goldMine(input))
}

func TestExample3(t *testing.T) {
	input := [][]int{
		{10, 33, 13, 15},
		{22, 21, 4, 1},
		{5, 0, 2, 3},
		{0, 6, 14, 2}}

	assert.Equal(t, 83, goldMine(input))
}
