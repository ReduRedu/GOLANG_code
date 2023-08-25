package insert

import
(
"bufio"
"os"
"strconv"
"strings"
"log"
)

func InsertFloat64() float64 {
var floatconversed float64
reader := bufio.NewReader(os.Stdin)
input, err := reader.ReadString('\n')
	if err != nil {
	 log.Fatal(err)
	 }
input = strings.TrimSpace(input)
floatconversed, err = strconv.ParseFloat(input, 64)
	if err != nil {
	 log.Fatal(err)
	 }
	 return floatconversed
	 }
	 
func InsertInt () int {
reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString ('\n')
		if err !=nil {
			log.Fatal (err) 
			}
	input = strings.TrimSpace (input)
	intconversed, err := strconv.Atoi (input)
		if err !=nil {
			log.Fatal (err) 
		}
		return intconversed
		}
func InsertString () string {
reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString ('\n')
		if err !=nil {
			log.Fatal (err) 
			}
	input = strings.TrimSpace (input)
	return input
}