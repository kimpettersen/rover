package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingForward(t *testing.T) {
	t.Run("When Facing North", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "N", pluto}
		_ = rover.Move("F")
		assert.Equal(t, 1, rover.y, "it should move one step up the y-axis")
		assert.Equal(t, 0, rover.x, "it should stay in the same position on x-axis")
	})
	t.Run("When Facing East", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "E", pluto}
		_ = rover.Move("F")
		assert.Equal(t, 0, rover.y, "it should stay in the same position on y-axis")
		assert.Equal(t, 1, rover.x, "it should move one step up the x-axis")
	})
	t.Run("When Facing South", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{1, 1, "S", pluto}
		_ = rover.Move("F")
		assert.Equal(t, 0, rover.y, "it should move one step down the y-axis")
		assert.Equal(t, 1, rover.x, "it should move one step on the x-axis")
	})
	t.Run("When Facing West", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{1, 1, "W", pluto}
		_ = rover.Move("F")
		assert.Equal(t, 1, rover.y, "it should move one step on the y-axis")
		assert.Equal(t, 0, rover.x, "it should move one step down the x-axis")
	})
}

func TestMovingBackward(t *testing.T) {
	t.Run("When Facing North", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{1, 1, "N", pluto}
		_ = rover.Move("B")
		assert.Equal(t, 0, rover.y, "it should move one step down the y-axis")
		assert.Equal(t, 1, rover.x, "it should stay in the same position on x-axis")
	})
	t.Run("When Facing East", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{1, 1, "E", pluto}
		_ = rover.Move("B")
		assert.Equal(t, 1, rover.y, "it should stay in the same position on y-axis")
		assert.Equal(t, 0, rover.x, "it should move one step down the x-axis")
	})
	t.Run("When Facing South", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "S", pluto}
		_ = rover.Move("B")
		assert.Equal(t, 1, rover.y, "it should move one step up the y-axis")
		assert.Equal(t, 0, rover.x, "it should stay in the same position on y-axis")
	})
	t.Run("When Facing West", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "W", pluto}
		_ = rover.Move("B")
		assert.Equal(t, 0, rover.y, "it should stay in the same position on y-axis")
		assert.Equal(t, 1, rover.x, "it should move one step up the x-axis")
	})
}

func TestTurnRight(t *testing.T) {
	t.Run("When Facing North", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "N", pluto}
		_ = rover.Move("R")
		assert.Equal(t, "E", rover.heading, "It should head East")
	})
	t.Run("When Facing East", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "E", pluto}
		_ = rover.Move("R")
		assert.Equal(t, "S", rover.heading, "It should head South")
	})
	t.Run("When Facing South", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "S", pluto}
		_ = rover.Move("R")
		assert.Equal(t, "W", rover.heading, "It should head West")
	})
	t.Run("When Facing West", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "W", pluto}
		_ = rover.Move("R")
		assert.Equal(t, "N", rover.heading, "It should head North")
	})
}

func TestTurnLeft(t *testing.T) {
	t.Run("When Facing North", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "N", pluto}
		_ = rover.Move("L")
		assert.Equal(t, "W", rover.heading, "It should head West")
	})
	t.Run("When Facing West", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "W", pluto}
		_ = rover.Move("L")
		assert.Equal(t, "S", rover.heading, "It should head South")
	})
	t.Run("When Facing South", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "S", pluto}
		_ = rover.Move("L")
		assert.Equal(t, "E", rover.heading, "It should head East")
	})
	t.Run("When Facing East", func(t *testing.T) {
		pluto := PlanetMap{10, 5}
		rover := Rover{0, 0, "E", pluto}
		_ = rover.Move("L")
		assert.Equal(t, "N", rover.heading, "It should head North")
	})
}

func TestGridWrapping(t *testing.T) {
	pluto := PlanetMap{10, 5}
	t.Run("When Facing North And Moving Forwards", func(t *testing.T) {
		rover := Rover{0, 5, "N", pluto}
		_ = rover.Move("F")
		assert.Equal(t, 0, rover.y, "It should wrap the grid")
	})
	t.Run("When Facing East And Moving Forwards", func(t *testing.T) {
		rover := Rover{10, 0, "E", pluto}
		_ = rover.Move("F")
		assert.Equal(t, 0, rover.x, "It should wrap the grid")
	})
	t.Run("When Facing South And Moving Forwards", func(t *testing.T) {
		rover := Rover{0, 0, "S", pluto}
		_ = rover.Move("F")
		assert.Equal(t, 5, rover.y, "It should wrap the grid")
	})
	t.Run("When Facing West And Moving Forwards", func(t *testing.T) {
		rover := Rover{0, 0, "W", pluto}
		_ = rover.Move("F")
		assert.Equal(t, 10, rover.x, "It should wrap the grid")
	})
	t.Run("When Facing North And Moving Backwards", func(t *testing.T) {
		rover := Rover{0, 0, "N", pluto}
		_ = rover.Move("B")
		assert.Equal(t, 5, rover.y, "It should wrap the grid")
	})
	t.Run("When Facing East And Moving Backwards", func(t *testing.T) {
		rover := Rover{0, 0, "E", pluto}
		_ = rover.Move("B")
		assert.Equal(t, 10, rover.x, "It should wrap the grid")
	})
	t.Run("When Facing South And Moving Backwards", func(t *testing.T) {
		rover := Rover{0, 5, "S", pluto}
		_ = rover.Move("B")
		assert.Equal(t, 0, rover.y, "It should wrap the grid")
	})
	t.Run("When Facing West And Moving Backwards", func(t *testing.T) {
		rover := Rover{10, 0, "W", pluto}
		_ = rover.Move("B")
		assert.Equal(t, 0, rover.x, "It should wrap the grid")
	})
}

func TestMoveRoverAround(t *testing.T) {
	pluto := PlanetMap{4, 3}
	t.Run("When moving rover from (1,1, N) to (2,2, E)", func(t *testing.T) {
		rover := Rover{1, 1, "N", pluto}
		_ = rover.Move("RFLFR")
		assert.Equal(t, 2, rover.x, "It should be at x=2")
		assert.Equal(t, 2, rover.y, "It should be at y=2")
		assert.Equal(t, "E", rover.heading, "It should be heading E")
	})
	t.Run("When moving rover from (1,1, N) to (4,2, N) when grid is wrapped x Direction", func(t *testing.T) {
		rover := Rover{1, 1, "N", pluto}
		_ = rover.Move("LFFRF")
		assert.Equal(t, 4, rover.x, "It should be at x=4")
		assert.Equal(t, 2, rover.y, "It should be at y=2")
		assert.Equal(t, "N", rover.heading, "It should be heading N")
	})
	t.Run("When moving rover from (1,1, N) to (3,1, W) when grid is wrapped y Direction", func(t *testing.T) {
		rover := Rover{1, 1, "N", pluto}
		_ = rover.Move("BBL")
		assert.Equal(t, 1, rover.x, "It should be at x=1")
		assert.Equal(t, 3, rover.y, "It should be at y=3")
		assert.Equal(t, "W", rover.heading, "It should be heading W")
	})
}

func TestCommandParsing(t *testing.T) {
	pluto := PlanetMap{4, 3}
	t.Run("When passing lower case commands", func(t *testing.T) {
		rover := Rover{1, 1, "E", pluto}
		err := rover.Move("f")
		assert.Equal(t, 2, rover.x, "It should be at x=2")
		assert.Nil(t, err, "It should not throw an error")
	})
	t.Run("When passing unexisting commands", func(t *testing.T) {
		rover := Rover{1, 1, "E", pluto}
		err := rover.Move("LBFLXS")
		assert.EqualError(t, err, "Invalid Command")
	})
	t.Run("When passing commands with spaces", func(t *testing.T) {
		rover := Rover{1, 1, "E", pluto}
		_ = rover.Move("F F")
		assert.Equal(t, 3, rover.x, "It should be x=3")
	})
}
