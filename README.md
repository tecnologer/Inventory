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
4. Instalar [Postgres][1]
5. Configurar la base de datos `test` y al usuario `postgres` con su contraseña `postgres`. Utilizar el script `/src/Inventory/queries_postgres.sql` para crear la base de datos y sus tablas.
6. `go run main.go`

Como instalar y configurar Go?... click [aqui][2].

### Usar API

1. Entrar al folder donde esta el codigo `cd $GOPATH/src/github.com/tecnologer/inventory/src/api`
2. Ejecutar el archivo principal `go run main.go`
3. Una vez instalado [Postman][3], abrirlo y presionar importar
4. Seleccionar el archivo `Inventory.postman_collection.json`, en este archivo estan definidas las pruebas por default
5. Seleccionar la prueba a ejecutar y hacer click en `Send`

[1]: https://wiki.postgresql.org/wiki/Main_Page
[2]: https://golang.org/doc/install
[3]: https://www.getpostman.com/