module Main (main) where

--import Lib
import EJERCICIO1
import EJERCICIO2

main :: IO()
main = do
    let prefijo = "per"
    let lista = ["el perro", "anpdirkperr", "per hola", "ab"]
    let subconjunto = [1,2,3]
    let subconjunto2 = [6,7]
    let conjunto = [3,4,5, 1,2]
    
    print(sub_conjunto subconjunto conjunto)
    --print(sub_cadenas prefijo lista)
