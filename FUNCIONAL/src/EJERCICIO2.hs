{-# OPTIONS_GHC -Wno-unrecognised-pragmas #-}
{-# HLINT ignore "Use camelCase" #-}
{-# HLINT ignore "Redundant if" #-}
{-# HLINT ignore "Redundant bracket" #-}
{-# HLINT ignore "Use guards" #-}
{-# HLINT ignore "Use null" #-}
{-# HLINT ignore "Use any" #-}
{-# HLINT ignore "Avoid lambda using `infix`" #-}
module EJERCICIO2
    (
        sub_conjunto
    ) where
    compare_element head_el list2 =
        if(list2 == [])then False
        else if( head_el == head list2) then True
        else
            compare_element head_el (tail list2)
            
                

    sub_conjunto list1 list2 =
        if(list1 == []) then True
        else
            if(null (filter(\x -> compare_element x list2) list1)) then False
            else True




