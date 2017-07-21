package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingForward(t *testing.T) {
	t.Run("When Facing North", func(t *testing.T) {
		rover := Rover{0, 0, "N"}
		rover.Move("F")
		assert.Equal(t, rover.y, 1, "it should move one step up the y-axis")
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

func TestMovingBackward(t *testing.T) {
	t.Run("When Facing North", func(t *testing.T) {
		rover := Rover{1, 1, "N"}
		rover.Move("B")
		assert.Equal(t, rover.y, 0, "it should move one step down the y-axis")
		assert.Equal(t, rover.x, 1, "it should stay in the same position on x-axis")
	})
	t.Run("When Facing East", func(t *testing.T) {
		rover := Rover{1, 1, "E"}
		rover.Move("B")
		assert.Equal(t, rover.y, 1, "it should stay in the same position on y-axis")
		assert.Equal(t, rover.x, 0, "it should move one step down the x-axis")
	})
	t.Run("When Facing South", func(t *testing.T) {
		rover := Rover{0, 0, "S"}
		rover.Move("B")
		assert.Equal(t, rover.y, 1, "it should move one step up the y-axis")
		assert.Equal(t, rover.x, 0, "it should stay in the same position on y-axis")
	})
	t.Run("When Facing West", func(t *testing.T) {
		rover := Rover{0, 0, "W"}
		rover.Move("B")
		assert.Equal(t, rover.y, 0, "it should stay in the same position on y-axis")
		assert.Equal(t, rover.x, 1, "it should move one step up the x-axis")
	})
}

func TestTurnRight(t *testing.T) {
	t.Run("When Facing North", func(t *testing.T) {
		rover := Rover{0, 0, "N"}
		rover.Move("R")
		assert.Equal(t, rover.heading, "E", "It should head East")
	})
	t.Run("When Facing East", func(t *testing.T) {
		rover := Rover{0, 0, "E"}
		rover.Move("R")
		assert.Equal(t, rover.heading, "S", "It should head South")
	})
	t.Run("When Facing South", func(t *testing.T) {
		rover := Rover{0, 0, "S"}
		rover.Move("R")
		assert.Equal(t, rover.heading, "W", "It should head West")
	})
	t.Run("When Facing West", func(t *testing.T) {
		rover := Rover{0, 0, "W"}
		rover.Move("R")
		assert.Equal(t, rover.heading, "N", "It should head North")
	})
}

func TestTurnLeft(t *testing.T) {
	t.Run("When Facing North", func(t *testing.T) {
		rover := Rover{0, 0, "N"}
		rover.Move("L")
		assert.Equal(t, rover.heading, "W", "It should head West")
	})
	t.Run("When Facing West", func(t *testing.T) {
		rover := Rover{0, 0, "W"}
		rover.Move("L")
		assert.Equal(t, rover.heading, "S", "It should head South")
	})
	t.Run("When Facing South", func(t *testing.T) {
		rover := Rover{0, 0, "S"}
		rover.Move("L")
		assert.Equal(t, rover.heading, "E", "It should head East")
	})
	t.Run("When Facing East", func(t *testing.T) {
		rover := Rover{0, 0, "E"}
		rover.Move("L")
		assert.Equal(t, rover.heading, "N", "It should head North")
	})

}
