package main

import (
	"errors"
	"strings"
)

type Rover struct {
	x, y      int
	heading   string
	planetMap PlanetMap
}

type PlanetMap struct {
	xMax, yMax int
}

func (r *Rover) moveBackwards() {
	if r.heading == "N" {
		if r.y <= 0 {
			r.y = r.planetMap.yMax
		} else {
			r.y = r.y - 1
		}
	} else if r.heading == "E" {
		if r.x <= 0 {
			r.x = r.planetMap.xMax
		} else {
			r.x = r.x - 1
		}
	} else if r.heading == "S" {
		if r.y >= r.planetMap.yMax {
			r.y = 0
		} else {
			r.y = r.y + 1
		}
	} else if r.heading == "W" {
		if r.x >= r.planetMap.xMax {
			r.x = 0
		} else {
			r.x = r.x + 1
		}
	}
}

func (r *Rover) moveForwards() {
	if r.heading == "N" {
		if r.y >= r.planetMap.yMax {
			r.y = 0
		} else {
			r.y = r.y + 1
		}
	} else if r.heading == "E" {
		if r.x >= r.planetMap.xMax {
			r.x = 0
		} else {
			r.x = r.x + 1
		}
	} else if r.heading == "S" {
		if r.y <= 0 {
			r.y = r.planetMap.yMax
		} else {
			r.y = r.y - 1
		}
	} else if r.heading == "W" {
		if r.x <= 0 {
			r.x = r.planetMap.xMax
		} else {
			r.x = r.x - 1
		}
	}
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
			r.moveForwards()
		} else if cmd == "B" {
			r.moveBackwards()
		} else if cmd == "R" {
			r.moveRight()
		} else if cmd == "L" {
			r.moveLeft()
		}
	}
	return nil
}

func main() {

}
