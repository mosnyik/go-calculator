package coreConcepts

import (
	"fmt"
	"os"
)

func AverageCalculator() {
	var sum float64
	var n int

	fmt.Println("We have started...")
	
	for {
		var val float64

		_, err := fmt.Fscan(os.Stdin, &val)

		if err != nil{
			break
			}

		sum += val
		n++
		}
    if n == 0{
		fmt.Fprintln(os.Stderr, "No values provided")
		os.Exit(-1)
	 }
 
	average := sum/float64(n)
	fmt.Printf("The avarage is: %f\n", average)
}