package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) name() string {
	return "circle"
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	length float64
}

func (s square) name() string {
	return "square"
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	name() string
	area() float64
}

type outputter struct {
}

func (o outputter) Text(s shape) string {
	return fmt.Sprintf("area of the %s: %v", s.name(), s.area())
}

type calcuclator struct {
}

func (c calcuclator) areaSum(shapes ...shape) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.area()
	}
	return sum
}

func (o outputter) JSON(s shape) string {
	res := struct {
		Name string  `json:"shape"`
		Area float64 `json:"area"`
	}{
		Name: s.name(),
		Area: s.area(),
	}

	bs, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}

func main() {
	c := circle{
		radius: 5,
	}
	s := square{
		length: 7,
	}

	out := outputter{}
	calc := calcuclator{}

	fmt.Println(out.Text(c))
	fmt.Println(out.JSON(c))
	fmt.Println(out.Text(s))
	fmt.Println(out.JSON(s))
	fmt.Println("area sum:", calc.areaSum(c, s))
}
