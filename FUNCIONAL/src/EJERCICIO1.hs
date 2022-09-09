{-# OPTIONS_GHC -Wno-unrecognised-pragmas #-}
{-# HLINT ignore "" #-}
module EJERCICIO1
    (
        
        ejercicio1,
        detectarCadena
    ) where

    detectarCadena prefix palabra cont cont2 flag =
        
        if (prefix !! cont) == (palabra !! cont2) then
            --cuando se encuentra pasamos a la siguiente letra del prefijo
            if (cont + 1  > (length prefix)-1) then True
            else
                detectarCadena prefix palabra (cont+1) (cont2+1) True
        else 
            --se llama la funcion hasta encontrar la letra igual a la priemra letra de prefix
            if(flag == True && cont2+1 > (length palabra)-1) then False
            else
                if(cont2 + 1  > (length palabra)-1) then False
                else
                    detectarCadena prefix palabra (cont) (cont2+1) False

        

    ejercicio1 prefix word_list = filter(\x -> detectarCadena prefix x 0 0 False) word_list

    

 