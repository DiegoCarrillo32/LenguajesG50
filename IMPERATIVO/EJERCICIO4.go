package main

import "fmt"

//scriba un programa que utilice una estructura y un arreglo/slice para almacenar en memoria un inventario de una tienda que vende zapatos.
//Cada zapato debe contar con la información de su modelo (marca), su precio y la talla del mismo que debe ir únicamente del 34 al 44.
//La estructura debe llamarse "calzado". El programa debe poder almacenar la información “quemada” para 10+ zapatos del inventario y un stock de N cantidad de unidades de dicho zapato
//(quiere decir por ejemplo que puede existir en inventario el zapato marca Nike, talla 42 y que cuesta 50000 colones, pero además que puede existir no solo un par de esos,
//si no N pares del mismo zapato).
//
//Haga una función que venda zapatos (eliminando del stock cada vez que haya venta e indicando que no se puede vender porque ya o hay stock)
//y pruébela dentro del main haciendo las ventas que usted considere para demostrar su funcionamiento.
type calzado struct {
	marca    string
	talla    int
	precio   int
	cantidad int
}

func venderZapato(marca string, almacena []calzado) []calzado {

	for i := 0; i < len(almacena); i++ {
		if almacena[i].marca == marca && almacena[i].cantidad > 0 {
			almacena[i].cantidad--
			fmt.Println("Se vendio un zapato de marca: ", marca, " Quedan disponibles: ", almacena[i].cantidad)
		}
		if almacena[i].marca == marca && almacena[i].cantidad == 0 {
			fmt.Println("No se puede vender porque no hay stock")
		}

	}
	return almacena
}
func main() {
	calzados := []calzado{}
	calzados = append(calzados, calzado{"Nike", 42, 50000, 10})
	calzados = append(calzados, calzado{"Puma", 44, 50000, 10})
	calzados = append(calzados, calzado{"Under", 42, 50000, 0})
	calzados = venderZapato("Nike", calzados)
	calzados = venderZapato("Nike", calzados)
	calzados = venderZapato("Nike", calzados)
	calzados = venderZapato("Nike", calzados)
	calzados = venderZapato("Nike", calzados)
	calzados = venderZapato("Nike", calzados)
	calzados = venderZapato("Nike", calzados)
	calzados = venderZapato("Nike", calzados)
	calzados = venderZapato("Nike", calzados)
	calzados = venderZapato("Nike", calzados)
	calzados = venderZapato("Under", calzados)

}
