package coreConcepts

import "fmt" 

func Calculate(){
	var number1 float64
	var number2 float64
	var operator string
	fmt.Println("Enter the first number: ")
	fmt.Scan(&number1)
	fmt.Println("Enter the Second number: ")
	fmt.Scan(&number2)
	fmt.Println("Enter the operator (+, -, * , / ): ")
	fmt.Scan(&operator)
	switch operator {
	case "+": fmt.Printf("%f %s %f = %f", number1, operator, number2, number1 + number2)
	case "-": fmt.Printf("%f %s %f = %f", number1, operator, number2, number1 - number2)
	case "*": fmt.Printf("%f %s %f = %f", number1, operator, number2, number1 * number2)
	// case "%": fmt.Printf("%f %s %f = %f", number1, operator, number2, number1 % number2)
	case "/": if number2 == 0{
		fmt.Println("We don't support 0 division")
	}else{
	fmt.Printf("%f %s %f = %f", number1, operator, number2, number1 / number2)
	}

	default : fmt.Println("Invalid operation")
	}


}