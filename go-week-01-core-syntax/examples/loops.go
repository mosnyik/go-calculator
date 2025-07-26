package coreConcepts

import "fmt"

func Loop(){
	// normal for loop
	for i:=0; i<5; i++{
		fmt.Println("Traditionanl for loop",i)
	}
// loop using range
	for i:=range 5{
	fmt.Printf("Print %d \n", i)
	}

	// looping over arrays
	var names = [4] string {"Moses", "Nyikwagh", "Iorwuese", "Msuega"}

	for index, name := range names{
		fmt.Println(index, name)
	}

	// nested loop
	var number = [4] int {1,2,3,4}

	for _, i:= range number {
		for _, j:= range number{
			fmt.Printf("(%d, %d) \n", i,j)
		}
	}

	// simulating a while loop in go

	var a = 25;

	for a > 20 {
		fmt.Println(a);
		a--
	}

}