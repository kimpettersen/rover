package main

type Rover struct {
	x, y    int
	heading string
}

func (r *Rover) Move(cmd string) {
	r.y = r.y + 1
}

func main() {

}
