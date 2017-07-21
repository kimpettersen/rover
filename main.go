package main

type Rover struct {
	x, y    int
	heading string
}

func (r *Rover) moveBackwards() {
	if r.heading == "N" {
		r.y = r.y - 1
	} else if r.heading == "E" {
		r.x = r.x - 1
	} else if r.heading == "S" {
		r.y = r.y + 1
	} else if r.heading == "W" {
		r.x = r.x + 1
	}
}

func (r *Rover) moveForwards() {
	if r.heading == "N" {
		r.y = r.y + 1
	} else if r.heading == "E" {
		r.x = r.x + 1
	} else if r.heading == "S" {
		r.y = r.y - 1
	} else if r.heading == "W" {
		r.x = r.x - 1
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

func (r *Rover) Move(cmd string) {
	if cmd == "F" {
		r.moveForwards()
	} else if cmd == "B" {
		r.moveBackwards()
	} else if cmd == "R" {
		r.moveRight()
	}
}

func main() {

}
