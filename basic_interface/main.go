package main

import (
	"fmt"
	"go_basic/basic_interface/model"
)

type Bear struct {
	Name  string
	Color string
}

func (b *Bear) GetName() string {
	return b.Name
}

func (b *Bear) GetColor() string {
	b.Color = "black"
	return b.Color
}

type Dog struct {
	Name  string
	Color string
}

func (d *Dog) GetName() string {
	return d.Name
}

func (d *Dog) GetColor() string {
	return d.Color
}

type Bird struct {
	Name  string
	Color string
}

func (b *Bird) GetName() string {
	return b.Name
}

func (b *Bird) GetColor() string {
	return b.Color
}

type Fish struct {
	Name  string
	Color string
}

func (f *Fish) GetName() string {
	return f.Name
}

func (f *Fish) GetColor() string {
	return f.Color
}

type Employee struct {
	Name      string
	Color     string
	FirstName string
	LastName  string
}

func (e *Employee) GetName() string {
	return e.Name
}

func (e *Employee) GetColor() string {
	return e.Color
}

func (e *Employee) GetFirstName() string {
	return e.FirstName
}

func (e *Employee) GetLastName() string {
	return e.LastName
}

func main() {

	wb := Bear{
		Name:  "Polar Bear",
		Color: "white",
	}

	dbm := Dog{
		Name:  "Doberman",
		Color: "black",
	}

	bd := Bird{
		Name:  "Twitter",
		Color: "blue",
	}

	fh := Fish{
		Name:  "Nemo",
		Color: "orange",
	}

	em := Employee{
		Name:      "Pipe",
		Color:     "white",
		FirstName: "Chayakorn",
		LastName:  "Pamon",
	}

	fmt.Println("All Animals")
	var animals []model.Animal

	animals = append(animals, &wb, &dbm, &bd, &fh, &em)

	for _, v := range animals {
		ta := v.(model.Animal)
		fmt.Printf("%s/%s \n", ta.GetName(), ta.GetColor())

	}

	fmt.Println("white bear after : ", wb.Color)

}
