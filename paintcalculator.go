package main

import 
( 
"fmt"
"insert"

)


func paintNeeded (width float64, height float64) (float64, error){
if width < 0 {
return 0, fmt.Errorf("a width of %0.2f is invalid", width)
}
if height < 0 {
return 0, fmt.Errorf("a height of %0.2f is invalid", height)
}
area := width * height
return area / 10.0, nil
}

func main() {
var a float64  = 0
var b float64 = 0
fmt.Println ("Enter the width:")
a = insert.InsertFloat64 ()
 fmt.Println ("Enter height:")
b = insert.InsertFloat64 ()
amount, err:= paintNeeded (a, b)
	if err == nil {
	fmt.Printf("%0.2f liters needed\n", amount)

	} else {
	fmt.Println(err)
	}
}
