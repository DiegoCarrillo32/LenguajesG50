module Main (main) where

--import Lib
import EJERCICIO1

main :: IO()
main = do
    let prefijo = "per"
    let lista = ["el perro", "anpdirkperr", "per hola", "ab"]
    
    print(ejercicio1 prefijo lista)