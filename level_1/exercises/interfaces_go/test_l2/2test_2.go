package main

import "fmt"

type Product interface {
	PrintInformation()
	ApplyDiscount() float64
}

type Book struct {
	Name  string
	Price float64
}

func (b Book) PrintInformation() {
	fmt.Printf("BOOK: %-8v ===>>> $%v\n", b.Name, b.Price)
}

func (b Book) ApplyDiscount() float64 {
	return b.Price * 10 / 100
}

type Game struct {
	Name  string
	Price float64
}

func (g Game) PrintInformation() {
	fmt.Printf("GAME: %-8v ===>>> $%v\n", g.Name, g.Price)
}

func (g Game) ApplyDiscount() float64 {
	return g.Price * 20 / 100
}

func main() {
	products := []Product{
		Book{Name: "B1", Price: 125.25},
		Book{Name: "B2", Price: 30.12},
		Game{Name: "G1", Price: 45.10},
		Game{Name: "G2", Price: 255.03},
	}

	for _, product := range products {
		product.PrintInformation()
		fmt.Printf("\t Discount: %0.2f \n", product.ApplyDiscount())
		fmt.Println("**********************")
	}
}
