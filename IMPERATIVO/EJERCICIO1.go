package main

import "fmt"

//Haga un programa que cuente e indique el número de caracteres, el número de palabras y el número de líneas de un texto
//cualquiera (obtenido de cualquier forma que considere). Considere hacer el cálculo de las tres variables en el mismo proceso.

//CONTAR EL NUMERO DE CARACTERES // no contar los espacios en blanco
//CONTAR EL NUMERO DE PALABRAS
//CONTAR EL NUMERO DE LINEAS

func contarCaracteres(text string) (int, int, int) {
	contadorPalabras := 0
	contadorCaracteres := 0
	contadorLineas := 0
	for caracter := 0; caracter < len(text); caracter++ {
		//get the character at the current index
		fmt.Println(string(text[caracter]))
		//si es un espacio no se cuenta y se aumenta el contador de palabras, si es un salto de linea se aumenta el contador de lineas
		if text[caracter] != 32 || string(text[caracter]) != "\n" {
			contadorCaracteres++
		}
		if text[caracter] == 32 && caracter+1 < len(text) && string(text[caracter+1]) != "\n" {

			contadorPalabras++
		}
		if string(text[caracter]) == "\n" {
			contadorLineas++
		}

	}
	return contadorCaracteres, contadorPalabras, contadorLineas

}
func main() {
	a, b, c := contarCaracteres("Hola mundo dios vamos a ver si esto funciona bien \n otra linea se añadio catorce quince \n y seiscientos")

	fmt.Println("El numero de caracteres es: ", a)
	fmt.Println("El numero de palabras es: ", b)
	fmt.Println("El numero de lineas es: ", c)
}
