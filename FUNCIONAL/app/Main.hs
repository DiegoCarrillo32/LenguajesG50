module Main (main) where

--import Lib
import EJERCICIO1
import EJERCICIO2
import EJERCICIO3

main :: IO()
main = do
    let prefijo = "per"
    let lista = ["el perro", "anpdirkperr", "per hola", "ab"]
    let subconjunto = [1,2,3]
    let subconjunto2 = [6,7]
    let conjunto = [3,4,5, 1,2]
    let listaanidada = [[3,2,1], [3,4,5], [7,8,9]]

    print(aplanar listaanidada [])
    print(sub_conjunto subconjunto conjunto)
    print(sub_cadenas prefijo lista)
