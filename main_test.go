package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingForward(t *testing.T) {
	t.Run("When Facing North", func(t *testing.T) {
		rover := Rover{0, 0, "N"}
		rover.Move("F")
		assert.Equal(t, rover.y, 1, "it should move one step on the  y-axis")
		assert.Equal(t, rover.x, 0, "it should stay in the same position on x-axis")
	})
	t.Run("When Facing East", func(t *testing.T) {
		rover := Rover{0, 0, "E"}
		rover.Move("F")
		assert.Equal(t, rover.y, 0, "it should stay in the same position on y-axis")
		assert.Equal(t, rover.x, 1, "it should move one step on the  x-axis")
	})
}
