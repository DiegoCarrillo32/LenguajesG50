package main

import "fmt"

//Escriba una función que permita rotar una secuencia de elementos N posiciones a la izquierda o a la derecha según sea el caso en función al parámetro que se reciba. Los parámetros deben ser un puntero a un arreglo previamente creado, la cantidad de movimiento de rotación y la dirección (0 = izq y 1 = der).
//
//A partir de esta función, escriba un programa que haga varias rotaciones cualesquiera de una secuencia de elementos previamente creada que sea inmutable. Al final debe imprimirse el resultado de cada rotación además de la secuencia original.
//	Un ejemplo de rotación puede ser el siguiente:
//Secuencia Original = [a, b, c, d, e, f, g, h,]
//Cantidad de rotaciones = 3
//Dirección = izq
//Secuencia final rotada = [d, e, f, g, h, a, b, c]   Nótese que cada iteración, el elemento más a la izquierda pasó a formar parte del final de la secuencia, si la rotación fuera a la derecha, sería lo contrario
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func rotarCaracteres(cantMovimientos int, direccion bool, lista *[]string) *[]string {

	if !direccion {
		for i := 0; i < cantMovimientos; i++ {
			*lista = append(*lista, (*lista)[0])
			*lista = remove(*lista, 0)

		}
	} else {
		for i := 0; i < cantMovimientos; i++ {
			*lista = append([]string{(*lista)[len(*lista)-1]}, *lista...)
			*lista = remove(*lista, len(*lista)-1)

		}
	}
	return lista
}
func main() {
	original := &[]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	original = rotarCaracteres(3, true, original)
	fmt.Println(*original)
}
