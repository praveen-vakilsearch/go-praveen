package main
import "fmt"

func main() {

	fmt.Printf("Farenheit: ")
	var f float64
	fmt.Scanf("%f", &f)
	c := (f - 32) * 5/9
	fmt.Println("Celsius : ",  c)
	
}