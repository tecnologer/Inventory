package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Tecnologer/Inventory/src/Inventory"

	"github.com/Tecnologer/Inventory/src/UI"
	"github.com/fatih/color"
)

var menu = UI.MainMenu{
	Title: "Inventory",
}

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
	case 1:
		reader := bufio.NewReader(os.Stdin)
		var id int
		var desc string
		var price float32
		var _type Inventory.Type
		var qty float32

		var err error
		for err != nil || id <= 0 || Inventory.Exists(id) {
			fmt.Print("ID: ")
			_, err = fmt.Scan(&id)

			if id <= 0 || Inventory.Exists(id) {
				fmt.Println("Id invalido, intenta de nuevo.")
			}
		}

		for desc == "" {
			fmt.Print("Descripcion: ")
			desc, _ = reader.ReadString('\n')

			desc = strings.Replace(desc, "\n", "", -1)
			desc = strings.Replace(desc, "\r", "", -1)
			if desc == "" {
				fmt.Println("Descripcion invalida, intenta de nuevo.")
			}
		}

		for err != nil || price <= 0 {
			fmt.Print("Precio: ")
			_, err = fmt.Scan(&price)

			if price <= 0 {
				fmt.Println("Precio invalido, intenta de nuevo.")
			}
		}

		for err != nil || _type == Inventory.Invalid || _type > Inventory.Service {
			fmt.Print("Tipo (1= Producto, 2= Servicio): ")
			_, err = fmt.Scan(&_type)

			if _type == Inventory.Invalid || _type > Inventory.Service {
				fmt.Println("Tipo invalido, intenta de nuevo.")
			}
		}

		fmt.Print("Cantidad inicial: ")
		_, err = fmt.Scan(&qty)

		for err != nil || qty < 0 {
			fmt.Println("Cantidad invalida, intenta de nuevo.")
			fmt.Print("Cantidad inicial: ")
			_, err = fmt.Scan(&qty)
		}

		Inventory.New(id, desc, price, _type, qty)
		return true, true
	case 2:
		Inventory.GetProducts()
		return true, true
	case 3:
		return true, true
	case 4:
		return true, true
	case 5: //exit
		fmt.Println("Cerrando sesion...")
		for i := 0; i < 3; i++ {
			fmt.Println(3 - i)
			time.Sleep(time.Second)
		}
		return true, false
	default:
		return false, false
	}
}
