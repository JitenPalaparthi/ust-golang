create a slice as 

slice1:= make([]any,10)

assign all defined types values to the slice1 (int,uint8-uint64,uint, int8-int64,float32-64,bool,string,rune,byte)

write a function to give sumOfSlice(slice []any)float64

- if bool type is true add 1 else (false) add 0
- if string has a number(int and float) convert and add it, else add 0

str1 = "123" , then 123 to be added       --> convert to int
str2 ="abcd123" , then 0 to be added
str3 = "123.123" then 123.123 to be added --> convert to float64