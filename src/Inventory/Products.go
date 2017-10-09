package Inventory

import (
	"fmt"
	"regexp"

	"github.com/fatih/color"
)

type Product struct {
	ID          int
	Description string
	Price       float32
	Type        Type
	Qty         float32
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

	inventory = append(inventory, newProduct)
}

//GetProducts prints the products saved
func GetProducts() {
	var header = fmt.Sprintf("|%-10s | %-25s | %-16s | %-16s | %-7s|", "Id", "Descripcion", "Precio", "Cantidad", "Tipo")
	color.Green(header)
	var reg, _ = regexp.Compile(`.`)
	var headerSeparator = reg.ReplaceAllString(header, "-")
	color.Green(headerSeparator)
	for _, p := range inventory {
		color.White(fmt.Sprintf("|%10d | %-25s | %16f | %16f | %-7s|\n", p.ID, p.Description, p.Price, p.Qty, TypeString(p.Type)))
	}
	color.Green(headerSeparator)
}

//Exists return if the id of product is already in inventory slice
func Exists(id int) bool {
	for _, p := range inventory {
		if p.ID == id {
			return true
		}
	}

	return false
}

//TypeString returns the string value of product type
func TypeString(productType Type) string {
	switch productType {
	case Item:
		return "Producto"
	case Service:
		return "Servicio"
	default:
		return "Invalido"
	}
}
