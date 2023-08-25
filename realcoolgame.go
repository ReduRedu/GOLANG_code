package main
import (
"fmt"
"math/rand"
"insert"

)

func main() {
var tries int
target := rand.Intn(1000) + 1 

fmt.Println("Slish loh ugaday chislo ot 1 do 1000")

for tries = 0 ; tries<10; tries++ {
	fmt.Println ("Make a guess.")
		guess:=insert.InsertInt ()
		if guess < 0 {
		fmt.Println ("eblan pishi chislo bolshe 0")
		continue
		}
		if guess > 1000 {
		fmt.Println ("CYKA NAUCHIS CHITAT CHISLO OT 1 DO 1000")
		continue
		}
		if guess < target {
		fmt.Println ("Eblan bolshe")
		}
		if guess > target {
		fmt.Println("eblan menshe")
		}
		if guess == target {
		fmt.Println ("Krasava")
			break
		}


}
fmt.Println ("Tebe ponadobilos", tries, "pop-it-ok, ti realno debs")
fmt.Println("Zagadannoe peslo:", target)
}