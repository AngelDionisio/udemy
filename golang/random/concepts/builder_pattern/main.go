package main

import (
	"fmt"
)

func main() {
	b := &Builder{}
	employee123 := b.SetName("Steve Wozniak").
		SetRole("manager").
		Build()
	fmt.Println(employee123)
}

type Employee struct {
	Name      string
	Role      string
	MinSalary int
	MaxSalary int
}

type Builder struct {
	e Employee
}

func (b *Builder) Build() *Employee {
	return &b.e
}

func (b *Builder) SetName(name string) *Builder {
	b.e.Name = name
	return b
}

func (b *Builder) SetRole(role string) *Builder {
	if b.e.Role == "manager" {
		b.e.MinSalary = 200000
		b.e.MaxSalary = 400000
	}
	b.e.Role = role
	return b
}
