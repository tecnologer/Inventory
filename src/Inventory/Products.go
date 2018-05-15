package Inventory

import "fmt"

type Product struct {
	ID          int     `json:"id,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price,omitempty"`
	Type        Type    `json:"type,omitempty"`
	Qty         float32 `json:"qty,omitempty"`
}

type Type int

const (
	Invalid Type = iota
	Item    Type = iota
	Service Type = iota
)

var inventory []Product

//New adds a new product to inventory
func New(id int, desc string, price float32, productType Type, qty float32) {

	if Exists(id) {
		return
	}

	newProduct := Product{
		ID:          id,
		Description: desc,
		Price:       price,
		Type:        productType,
		Qty:         qty,
	}

	_, err := InsertNewProduct(newProduct)

	if err != nil {
		fmt.Printf("Error when try insert new product. %s\n", err)
	}
	//inventory = append(inventory, newProduct)
}

//GetProducts returns list of products
func GetProductsList() []Product {
	list, err := GetProducts()
	if err != nil {
		fmt.Printf("Error when try get products %s\n", err)
		return nil
	}

	return list
}

//Exists return if the id of product is already in inventory slice
func Exists(id int) bool {
	exists, err := ExistsProduct(id)
	if err != nil {
		fmt.Printf("Error when try chek if product exists. %s\n", err)
		return true
	}

	return exists
}

func indexOf(id int) int {
	for i, p := range inventory {
		if p.ID == id {
			return i
		}
	}

	return -1
}

//Delete product of the list
func Delete(id int) {
	// var index = indexOf(id)
	// inventory = append(inventory[:index], inventory[index+1:]...)
	err := DeleteProduct(id)
	if err != nil {
		fmt.Printf("Error when try delete product. %s\n", err)
	}
}

//Update infor of product
func Update(id int, desc string, price float32, productType Type) {
	// i := indexOf(id)
	// inventory[i].Description = desc
	// inventory[i].Price = price
	// inventory[i].Type = productType
	err := UpdateProduct(id, desc, price, productType)
	if err != nil {
		fmt.Printf("Error when try update %d. %s\n", id, err)
	}
}

//InventoryMovement increase qty in specific product
func InventoryMovement(id int, qty float32) {
	// i := indexOf(id)
	// inventory[i].Qty += qty
	err := AddMovement(id, qty)
	if err != nil {
		fmt.Printf("Error when try update quantity of %d. %s\n", id, err)
	}
}

func (p Product) TypeString() string {
	switch p.Type {
	case Item:
		return "Producto"
	case Service:
		return "Servicio"
	default:
		return "Invalido"
	}
}
