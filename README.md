# Inventario [![Build Status](https://travis-ci.org/Tecnologer/Inventory.svg?branch=master)](https://travis-ci.org/Tecnologer/Inventory)

Sencillo programa para Crear, Leer, Modificar y Eliminar productos. 
Utilizando una base de datos en Postgres.

## Dependencias

```None
//banner
go get github.com/tecnologer/asciiart

//color para la consola
go get github.com/fatih/color

//postgres
go get github.com/lib/pq
```

### Uso

1. Primero hay que clonar el repositorio en `$GOPATH/src/github.com/tecnologer` <br/>
    `git clone https://github.com/Tecnologer/Inventory.git inventory`
2. Entrar a la carpeta del repositorio: <br/> `cd $GOPATH/src/github.com/tecnologer/inventory`
3. Instalar dependencias
4. `go run main.go`
