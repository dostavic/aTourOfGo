package main

import "fmt"

type Vertex struct {
	lat, Log float64
}

var m = map[string]Vertex{
	"Bell Labs": {40, -74},
	"Google":    {37, -122},
}

func main() {
	fmt.Println(m)
}
