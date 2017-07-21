package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenMovingForward(t *testing.T) {
	rover := Rover{0, 0}
	rover.Move("F")
	assert.Equal(t, rover.y, 1, "it should move one step on the  y-axis")
	assert.Equal(t, rover.x, 0, "it should stay in the same position on x-axis")
}
