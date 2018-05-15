package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/tecnologer/inventory/src/Inventory"
	"github.com/tecnologer/inventory/src/UI"
)

var err error
var menu = UI.MainMenu{
	Title: "Inventory",
	Options: map[int]string{
		1: "Nuevo producto",
		2: "Mostrar productos",
		3: "Modificar Producto",
		4: "Movimiento de inventario",
		5: "Eliminar producto",
		6: "Mostrar opciones",
		7: "Salir",
	},
}

type MovementType int

const (
	increase MovementType = iota
	decrease MovementType = iota
)

func main() {
	menu.ShowMenu()

	chooseOption()
}

func chooseOption() {
	var option int
	color.Cyan("Seleccione una opcion: ")
	var _, err = fmt.Scan(&option)
	var valid, success = selectOption(option)

	for err != nil || !valid {
		color.Red("Opcion incorrecta. Intenta de nuevo")
		_, err = fmt.Scan(&option)
		valid, success = selectOption(option)
	}

	if success {
		chooseOption()
	}
}

func selectOption(option int) (bool, bool) {
	switch option {
	case 1: //Create
		addProduct()
		return true, true
	case 2: //Read
		showProducts()
		return true, true
	case 3: //Update
		updateProduct()
		return true, true
	case 4: //Inventory Movement
		inventoryMovement()
		return true, true
	case 5: //Delete
		deleteProduct()
		return true, true
	case 6:
		menu.PrintOptions()
		return true, true
	case 7: //exit
		fmt.Println("Cerrando sesion...")
		Inventory.CloseDB()

		for i := 0; i < 3; i++ {
			fmt.Println(3 - i)
			time.Sleep(500 * time.Millisecond)
		}
		return true, false
	default:
		return false, false
	}
}

func addProduct() {
	id := readID()
	desc := readDesc()
	price := readPrice()
	productType := readType()
	qty := readQty()
	Inventory.New(id, desc, price, productType, qty)
}

func updateProduct() {
	showProducts()
	fmt.Println("****** Actualizando producto ******")
	id := readExistingID()
	desc := readDesc()
	price := readPrice()
	productType := readType()
	Inventory.Update(id, desc, price, productType)

	fmt.Println("****** Producto Actualizado ******")
	showProducts()
}

func inventoryMovement() {
	showProducts()
	movementType := readMovementType()
	id := readExistingID()
	qty := readQty()

	if movementType == decrease {
		qty = qty * -1
	}

	Inventory.InventoryMovement(id, qty)
	showProducts()
}

func deleteProduct() {
	showProducts()
	Inventory.Delete(readExistingID())
	showProducts()
}

//#region Input
func readID() int {
	var id int
	for err != nil || id <= 0 {
		fmt.Print("ID: ")
		_, err = fmt.Scan(&id)

		if id <= 0 {
			fmt.Println("Id invalido, intenta de nuevo.")
		} else if Inventory.Exists(id) {
			fmt.Println("Id ya esta registrado.")
			id = 0
		}
	}
	return id
}

func readExistingID() int {
	var id int
	for err != nil || id <= 0 || !Inventory.Exists(id) {
		fmt.Print("ID: ")
		_, err = fmt.Scan(&id)

		if id <= 0 || !Inventory.Exists(id) {
			fmt.Println("Id invalido, intenta de nuevo.")
		}
	}
	return id
}

func readDesc() string {
	reader := bufio.NewReader(os.Stdin)
	var desc string
	for desc == "" {
		fmt.Print("Descripcion: ")
		desc, _ = reader.ReadString('\n')

		desc = strings.Replace(desc, "\n", "", -1)
		desc = strings.Replace(desc, "\r", "", -1)
		if desc == "" {
			fmt.Println("Descripcion invalida, intenta de nuevo.")
		}
	}

	return desc
}

func readPrice() float32 {
	var price float32
	for err != nil || price <= 0 {
		fmt.Print("Precio: ")
		_, err = fmt.Scan(&price)

		if price <= 0 {
			fmt.Println("Precio invalido, intenta de nuevo.")
		}
	}
	return price
}

func readType() Inventory.Type {
	var productType Inventory.Type
	for err != nil || productType == Inventory.Invalid || productType > Inventory.Service {
		fmt.Print("Tipo (1= Producto, 2= Servicio): ")
		_, err = fmt.Scan(&productType)

		if productType == Inventory.Invalid || productType > Inventory.Service {
			fmt.Println("Tipo invalido, intenta de nuevo.")
		}
	}
	return productType
}

func readQty() float32 {
	var qty float32
	fmt.Print("Cantidad: ")
	_, err = fmt.Scan(&qty)

	for err != nil || qty < 0 {
		fmt.Println("Cantidad invalida, intenta de nuevo.")
		fmt.Print("Cantidad: ")
		_, err = fmt.Scan(&qty)
	}
	return qty
}

func readMovementType() MovementType {
	var movementType MovementType
	var loop = 0
	for err != nil || loop == 0 {
		fmt.Print("Movimiento (0= Entrada, 1= Salida): ")
		_, err = fmt.Scan(&movementType)
		loop++
		if movementType < 0 || movementType > decrease {
			fmt.Println("Movimiento invalido, intenta de nuevo.")
			loop = 0
		}
	}
	return movementType
}

//#endregion
//ShowProducts prints the products saved
func showProducts() {
	var header = fmt.Sprintf("|%-10s | %-25s | %-16s | %-16s | %-10s|", "Id", "Descripcion", "Precio", "Cantidad", "Tipo")
	color.Green(header)
	var reg, _ = regexp.Compile(`.`)
	var headerSeparator = reg.ReplaceAllString(header, "-")
	color.Green(headerSeparator)
	for _, p := range Inventory.GetProductsList() {
		color.White(fmt.Sprintf("|%10d | %-25s | %16f | %16f | %-10s|\n", p.ID, p.Description, p.Price, p.Qty, p.TypeString()))
	}
	color.Green(headerSeparator)
}
