package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingForward(t *testing.T) {
	t.Run("When Facing North", func(t *testing.T) {
		rover := Rover{0, 0, "N"}
		rover.Move("F")
		assert.Equal(t, rover.y, 1, "it should move one step up the  y-axis")
		assert.Equal(t, rover.x, 0, "it should stay in the same position on x-axis")
	})
	t.Run("When Facing East", func(t *testing.T) {
		rover := Rover{0, 0, "E"}
		rover.Move("F")
		assert.Equal(t, rover.y, 0, "it should stay in the same position on y-axis")
		assert.Equal(t, rover.x, 1, "it should move one step up the x-axis")
	})
	t.Run("When Facing South", func(t *testing.T) {
		rover := Rover{1, 1, "S"}
		rover.Move("F")
		assert.Equal(t, rover.y, 0, "it should move one step down the y-axis")
		assert.Equal(t, rover.x, 1, "it should move one step on the x-axis")
	})
	t.Run("When Facing West", func(t *testing.T) {
		rover := Rover{1, 1, "W"}
		rover.Move("F")
		assert.Equal(t, rover.y, 1, "it should move one step on the y-axis")
		assert.Equal(t, rover.x, 0, "it should move one step down the x-axis")
	})
}
