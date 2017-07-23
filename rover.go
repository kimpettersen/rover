package rover

import (
	"errors"
	"fmt"
	"strings"
)

type Rover struct {
	x, y      int
	heading   string
	planetMap PlanetMap
}

type PlanetMap struct {
	xMax, yMax int
	obstacles  []Obstacle
}

type Obstacle struct {
	x, y int
}

func (r *Rover) safeToMove(x, y int) bool {
	for _, obs := range r.planetMap.obstacles {
		if obs.x == x && obs.y == y {
			return false
		}
	}
	return true
}

func (r *Rover) setY(n int) error {
	if r.safeToMove(r.x, n) {
		r.y = n
		return nil
	}
	return fmt.Errorf(
		"Obstacle found. Rover will stay at (%v,%v,%v)",
		r.x,
		r.y,
		r.heading,
	)
}

func (r *Rover) setX(n int) error {
	if r.safeToMove(n, r.y) {
		r.x = n
		return nil
	}
	return fmt.Errorf(
		"Obstacle found. Rover will stay at (%v,%v,%v)",
		r.x,
		r.y,
		r.heading,
	)
}

func (r *Rover) moveBackwards() error {
	var err error

	if r.heading == "N" {
		if r.y <= 0 {
			err = r.setY(r.planetMap.yMax)
		} else {
			err = r.setY(r.y - 1)
		}
	} else if r.heading == "E" {
		if r.x <= 0 {
			err = r.setX(r.planetMap.xMax)
		} else {
			err = r.setX(r.x - 1)
		}
	} else if r.heading == "S" {
		if r.y >= r.planetMap.yMax {
			err = r.setY(0)
		} else {
			err = r.setY(r.y + 1)
		}
	} else if r.heading == "W" {
		if r.x >= r.planetMap.xMax {
			err = r.setX(0)
		} else {
			err = r.setX(r.x + 1)
		}
	}
	return err // Default is nil
}

func (r *Rover) moveForwards() error {
	var err error
	if r.heading == "N" {
		if r.y >= r.planetMap.yMax {
			err = r.setY(0)
		} else {
			err = r.setY(r.y + 1)
		}
	} else if r.heading == "E" {
		if r.x >= r.planetMap.xMax {
			err = r.setX(0)
		} else {
			err = r.setX(r.x + 1)
		}
	} else if r.heading == "S" {
		if r.y <= 0 {
			err = r.setY(r.planetMap.yMax)
		} else {
			err = r.setY(r.y - 1)
		}
	} else if r.heading == "W" {
		if r.x <= 0 {
			err = r.setX(r.planetMap.xMax)
		} else {
			err = r.setX(r.x - 1)
		}
	}
	return err // Default is nil
}

func (r *Rover) moveRight() {
	if r.heading == "N" {
		r.heading = "E"
	} else if r.heading == "E" {
		r.heading = "S"
	} else if r.heading == "S" {
		r.heading = "W"
	} else if r.heading == "W" {
		r.heading = "N"
	}
}

func (r *Rover) moveLeft() {
	if r.heading == "N" {
		r.heading = "W"
	} else if r.heading == "W" {
		r.heading = "S"
	} else if r.heading == "S" {
		r.heading = "E"
	} else if r.heading == "E" {
		r.heading = "N"
	}
}

func parseCmds(cmds string) ([]string, error) {
	validCmds := "FBLR"

	cmds = strings.Replace(cmds, " ", "", -1)
	cmds = strings.ToUpper(cmds)
	cmdsSLice := strings.Split(cmds, "")

	for _, cmd := range cmdsSLice {
		if strings.Contains(validCmds, cmd) != true {
			return nil, errors.New("Invalid Command")
		}
	}

	return cmdsSLice, nil
}

func (r *Rover) Move(cmds string) error {
	cmdsSlice, err := parseCmds(cmds)
	if err != nil {
		return err
	}

	for _, cmd := range cmdsSlice {
		if cmd == "F" {
			if err := r.moveForwards(); err != nil {
				return err
			}
		} else if cmd == "B" {
			if err := r.moveBackwards(); err != nil {
				return err
			}
		} else if cmd == "R" {
			r.moveRight()
		} else if cmd == "L" {
			r.moveLeft()
		}
	}
	return nil
}
