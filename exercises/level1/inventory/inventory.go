package inventory

import "fmt"

type Product struct {
	Id   int
	Name string
}

// create a new product
func NewProduct(id int, name string) Product {
	prod := Product{
		Id:   id,
		Name: name,
	}
	return prod
}

type Inventory struct {
	Products map[string]Product
}

func exists(p map[string]Product, id int) bool {
	for _, v := range p {
		if v.Id == id {
			return true
		}
	}
	return false
}

// add a product to inventory
func (i *Inventory) AddItem(sku string, p Product) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Product %v %v \n", sku, r)
		}
	}()
	if len(i.Products) == 0 {
		i.Products = map[string]Product{}
	}

	if p.Id == 0 {
		panic("Id cannot be empty")
	}

	if !exists(i.Products, p.Id) {
		i.Products[sku] = p
	} else {
		panic("already exists!")
	}
}

func (i Inventory) GetProducts() map[string]Product {
	return i.Products
}
