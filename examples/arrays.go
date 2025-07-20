package coreConcepts

import "fmt" 

func Array(){
	// declare and populate array
	var students [3] string
	students = [3]string {"Iyanu", "Eniola","Taliat"}

	bank:= [3]string {"GTB", "Zenith", "Access"}
	for position, name:= range students{
		fmt.Printf("Student with name %s is number %d \n", name, position+1)
	}

	for i:=0; i< len(bank); i++{
		fmt.Println(bank[i])
	}
	// nested array
	nestedArray := [3][5] int {
		{1,2,3,4,5},
		{2,2,4,5,7},
		{4,5,1,3,5},
	}

	// create a slice
	newSlice:= bank[1:];

	fmt.Println("Slice:",newSlice)

	makeSlice := make([]string, 5, 10);

	fmt.Println("Make slice", makeSlice)
	fmt.Println("Make slice length", cap(makeSlice))

	// flatten array
	position:= 0;
	for range len(nestedArray){
		for i:= 0; i< len(nestedArray[position]); i++{
		fmt.Println(nestedArray[position][i])
		}
		position++
	}
}

