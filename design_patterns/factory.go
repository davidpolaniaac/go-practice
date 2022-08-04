package main

import "fmt"

type IProduct interface {
	getName() string
	getStock() int
	setName(name string)
	setStock(stock int)
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) getName() string {
	return c.name
}
func (c *Computer) getStock() int {
	return c.stock
}
func (c *Computer) setName(name string) {
	c.name = name
}
func (c *Computer) setStock(stock int) {
	c.stock = stock
}

type Laptop struct {
	Computer
}

type Desktop struct {
	Computer
}

func newLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Dell laptop",
			stock: 35,
		},
	}
}

func newDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Dell desktop",
			stock: 25,
		},
	}
}

func ComputerFactory(name string) (IProduct, error) {
	if name == "laptop" {
		return newLaptop(), nil
	}

	if name == "desktop" {
		return newDesktop(), nil
	}
	return nil, fmt.Errorf("invalid type")
}

func printProduct(product IProduct) {
	fmt.Printf("Product: %s - Stock: %d\n", product.getName(), product.getStock())
}

func main() {

	laptop, _ := ComputerFactory("laptop")
	desktop, _ := ComputerFactory("desktop")
	printProduct(laptop)
	printProduct(desktop)

}
