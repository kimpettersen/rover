package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingForward(t *testing.T) {
	t.Run("When Facing North", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "N", pluto}
		rover.Move("F")
		assert.Equal(t, 1, rover.y, "it should move one step up the y-axis")
		assert.Equal(t, 0, rover.x, "it should stay in the same position on x-axis")
	})
	t.Run("When Facing East", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "E", pluto}
		rover.Move("F")
		assert.Equal(t, 0, rover.y, "it should stay in the same position on y-axis")
		assert.Equal(t, 1, rover.x, "it should move one step up the x-axis")
	})
	t.Run("When Facing South", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{1, 1, "S", pluto}
		rover.Move("F")
		assert.Equal(t, 0, rover.y, "it should move one step down the y-axis")
		assert.Equal(t, 1, rover.x, "it should move one step on the x-axis")
	})
	t.Run("When Facing West", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{1, 1, "W", pluto}
		rover.Move("F")
		assert.Equal(t, 1, rover.y, "it should move one step on the y-axis")
		assert.Equal(t, 0, rover.x, "it should move one step down the x-axis")
	})
}

func TestMovingBackward(t *testing.T) {
	t.Run("When Facing North", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{1, 1, "N", pluto}
		rover.Move("B")
		assert.Equal(t, 0, rover.y, "it should move one step down the y-axis")
		assert.Equal(t, 1, rover.x, "it should stay in the same position on x-axis")
	})
	t.Run("When Facing East", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{1, 1, "E", pluto}
		rover.Move("B")
		assert.Equal(t, 1, rover.y, "it should stay in the same position on y-axis")
		assert.Equal(t, 0, rover.x, "it should move one step down the x-axis")
	})
	t.Run("When Facing South", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "S", pluto}
		rover.Move("B")
		assert.Equal(t, 1, rover.y, "it should move one step up the y-axis")
		assert.Equal(t, 0, rover.x, "it should stay in the same position on y-axis")
	})
	t.Run("When Facing West", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "W", pluto}
		rover.Move("B")
		assert.Equal(t, 0, rover.y, "it should stay in the same position on y-axis")
		assert.Equal(t, 1, rover.x, "it should move one step up the x-axis")
	})
}

func TestTurnRight(t *testing.T) {
	t.Run("When Facing North", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "N", pluto}
		rover.Move("R")
		assert.Equal(t, "E", rover.heading, "It should head East")
	})
	t.Run("When Facing East", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "E", pluto}
		rover.Move("R")
		assert.Equal(t, "S", rover.heading, "It should head South")
	})
	t.Run("When Facing South", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "S", pluto}
		rover.Move("R")
		assert.Equal(t, "W", rover.heading, "It should head West")
	})
	t.Run("When Facing West", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "W", pluto}
		rover.Move("R")
		assert.Equal(t, "N", rover.heading, "It should head North")
	})
}

func TestTurnLeft(t *testing.T) {
	t.Run("When Facing North", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "N", pluto}
		rover.Move("L")
		assert.Equal(t, "W", rover.heading, "It should head West")
	})
	t.Run("When Facing West", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "W", pluto}
		rover.Move("L")
		assert.Equal(t, "S", rover.heading, "It should head South")
	})
	t.Run("When Facing South", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "S", pluto}
		rover.Move("L")
		assert.Equal(t, "E", rover.heading, "It should head East")
	})
	t.Run("When Facing East", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "E", pluto}
		rover.Move("L")
		assert.Equal(t, "N", rover.heading, "It should head North")
	})
}

func TestGridWrapping(t *testing.T) {
	pluto := PlanetMap{10, 5}
	t.Run("When Facing North And Moving Forwards", func(t *testing.T) {
		rover := Rover{0, 5, "N", pluto}
		rover.Move("F")
		assert.Equal(t, 0, rover.y, "It should wrap the grid")
	})
	t.Run("When Facing East And Moving Forwards", func(t *testing.T) {
		rover := Rover{10, 0, "E", pluto}
		rover.Move("F")
		assert.Equal(t, 0, rover.x, "It should wrap the grid")
	})
	t.Run("When Facing South And Moving Forwards", func(t *testing.T) {
		rover := Rover{0, 0, "S", pluto}
		rover.Move("F")
		assert.Equal(t, 5, rover.y, "It should wrap the grid")
	})
	t.Run("When Facing West And Moving Forwards", func(t *testing.T) {
		rover := Rover{0, 0, "W", pluto}
		rover.Move("F")
		assert.Equal(t, 10, rover.x, "It should wrap the grid")
	})
	t.Run("When Facing North And Moving Backwards", func(t *testing.T) {
		rover := Rover{0, 0, "N", pluto}
		rover.Move("B")
		assert.Equal(t, 5, rover.y, "It should wrap the grid")
	})
	t.Run("When Facing East And Moving Backwards", func(t *testing.T) {
		rover := Rover{0, 0, "E", pluto}
		rover.Move("B")
		assert.Equal(t, 10, rover.x, "It should wrap the grid")
	})
	t.Run("When Facing South And Moving Backwards", func(t *testing.T) {
		rover := Rover{0, 5, "S", pluto}
		rover.Move("B")
		assert.Equal(t, 0, rover.y, "It should wrap the grid")
	})
	t.Run("When Facing West And Moving Backwards", func(t *testing.T) {
		rover := Rover{10, 0, "W", pluto}
		rover.Move("B")
		assert.Equal(t, 0, rover.x, "It should wrap the grid")
	})
}
