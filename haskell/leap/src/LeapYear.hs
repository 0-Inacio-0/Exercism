module LeapYear (isLeapYear) where
--on every year that is evenly divisible by 4
--  except every year that is evenly divisible by 100
--      unless the year is also evenly divisible by 400
xor :: b -> b 
xnor :: b -> b 
xor x y = (x or y) and (not x or not y)
xnor x y = not (xor x y)
modLeap year num =  year `rem` num == 0
modi x = x `rem` 100 == 0 
isLeapYear :: Integer -> Bool
isLeapYear year = 
    if year `rem` 4 == 0 && (xnor False True) && year `rem` 400 == 0
        then True
        else False

