package Inventory

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

//GetProducts returns list of products
func GetProducts() []Product {
	return inventory
}

//Exists return if the id of product is already in inventory slice
func Exists(id int) bool {
	return indexOf(id) >= 0
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
	var index = indexOf(id)
	inventory = append(inventory[:index], inventory[index+1:]...)
}

//Update infor of product
func Update(id int, desc string, price float32, productType Type) {
	i := indexOf(id)
	inventory[i].Description = desc
	inventory[i].Price = price
	inventory[i].Type = productType
}

//InventoryMovement increase qty in specific product
func InventoryMovement(id int, qty float32) {
	i := indexOf(id)
	inventory[i].Qty += qty
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
