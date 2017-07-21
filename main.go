package main

type Rover struct {
	x, y    int
	heading string
}

func (r *Rover) Move(cmd string) {
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

func main() {

}
