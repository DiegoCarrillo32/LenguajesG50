package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	var i int
	var flag bool
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			flag = true
			fmt.Println("Producto ya existe, se aumentará el stock")
			(*l)[i].cantidad = (*l)[i].cantidad + 1
		}
	}
	if !flag {
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}

	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
}

func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaProductos) venderProducto(nombre string, cant int) {
	var prod = l.buscarProducto(nombre)
	if prod != -1 && cant > 0 {
		if (*l)[prod].cantidad >= cant {
			(*l)[prod].cantidad = (*l)[prod].cantidad - cant
		} else {
			fmt.Println("No se puede vayor cantidad de productos que los que hay en existencia")
		}

		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista"
	}
	if (*l)[prod].cantidad == 0 {
		*l = append((*l)[:prod], (*l)[prod+1:]...)

		fmt.Println("Producto eliminado por falta de existencias")

	}
}
func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 5, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
	lProductos.agregarProducto("café", 12, 4500)

}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima
	var existencias listaProductos
	for i := 0; i < len((*l)); i++ {
		if (*l)[i].cantidad <= existenciaMinima {
			existencias = append(existencias, (*l)[i])
		}
	}

	return existencias
}

func saveToFile(prod producto) {

	f, err := os.OpenFile("SEMANA3/base.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s2 := strconv.Itoa(prod.precio)
	s3 := strconv.Itoa(prod.cantidad)
	if _, err = f.WriteString(prod.nombre + " " + s2 + " " + s3 + "\n"); err != nil {
		panic(err)
	}

}

func (l *listaProductos) aumentarInventarioDeMinimos(listaMinimos listaProductos) {
	for i := 0; i < len(*l); i++ {
		for x := 0; x < len(listaMinimos); x++ {

			if (*l)[i].nombre == listaMinimos[x].nombre {

				agregar := existenciaMinima - (*l)[i].cantidad
				(*l)[i].cantidad += (agregar + 1)
			}

		}
	}
}

func (l *listaProductos) ordenarLista() {
	//El paquete que se utilizó es sort
	//https://pkg.go.dev/sort

	fmt.Println(*l, "LIST SIN ORDENAR")
	sort.Slice(*l, func(i, j int) bool {
		return (*l)[i].precio > (*l)[j].precio
	})
	fmt.Println(*l, "LISTA ORDENADA")
}

func main() {
	llenarDatos()

	lProductos.venderProducto("arroz", 4)

	lProductos.agregarProducto("azucar", 20, 1500)

	lProductos.venderProducto("frijoles", 4)

	lProductos.venderProducto("leche", 10)

	fmt.Println(lProductos.listarProductosMínimos(), "PRODUCTOS MINIMOS")

	lProductos.aumentarInventarioDeMinimos(lProductos.listarProductosMínimos())
	fmt.Println(lProductos.listarProductosMínimos(), "PRODUCTOS MINIMOS")
	lProductos.ordenarLista()
}

//RESULTADOS DE LA TERMINAL

//Producto ya existe, se aumentará el stock
//No se puede vayor cantidad de productos que los que hay en existencia
//[{frijoles 1 2000} {leche 8 1200}] PRODUCTOS MINIMOS
//[] PRODUCTOS MINIMOS
//[{arroz 11 2500} {frijoles 11 2000} {leche 11 1200} {café 13 4500} {azucar 20 1500}] LIST SIN ORDENAR
//[{café 13 4500} {arroz 11 2500} {frijoles 11 2000} {azucar 20 1500} {leche 11 1200}] LISTA ORDENADA
