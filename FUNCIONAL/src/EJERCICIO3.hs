{-# OPTIONS_GHC -Wno-deferred-out-of-scope-variables #-}
{-# OPTIONS_GHC -Wno-unrecognised-pragmas #-}
{-# HLINT ignore "Use camelCase" #-}
{-# HLINT ignore "Redundant if" #-}
{-# HLINT ignore "Redundant bracket" #-}
{-# HLINT ignore "Use guards" #-}
{-# HLINT ignore "Use null" #-}
{-# HLINT ignore "Use any" #-}
{-# HLINT ignore "Avoid lambda using `infix`" #-}
{-# HLINT ignore "Avoid lambda" #-}
{-# HLINT ignore "Eta reduce" #-}
module EJERCICIO3
    (
        aplanar 
    ) where

    append a xs = xs ++ [a]
    new_list  listAnidada listNueva =
        if(listAnidada == []) then listNueva
        else
            new_list  (tail listAnidada) (append (head listAnidada) listNueva)


    aplanar listAnidada listNueva =
        if(listAnidada == []) then listNueva
        else
            aplanar (tail listAnidada) (new_list  (head listAnidada) listNueva)



        