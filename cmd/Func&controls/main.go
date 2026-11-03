package main

import ("fmt" "errors")

func main() {
	// printMe()

	// PrintParameter("This is a Print Function taking a parameter from main function")

	var numerator int = 20
	var denominator int = 5

	var result int
	var reminder int
	result, reminder,err = intDivision(numerator,denominator)
	// printInt(result)
	// printInt(reminder)

	if err!=nil {
		fmt.Printf(err.Error())
	}else if reminder == 0 {
		fmt.Printf("The result of the integer division is %v",result)
	}else {
		fmt.Printf("The result of the integer division is %v with reminder %v  \n",result,reminder)
	}


	switch {
		case err!=nil:
			fmt.Printf(err.Error())
		case reminder==0:
			fmt.Printf("The result of the integer division is%v",result)
		default:
			fmt.Printf("The result of the integer division is %v and reminder is %v",result,reminder)	
	}

	switch reminder{
	case 0:
		fmt.Printf("The division was exact")
	case 1,2:  
		fmt.Printf("The division was close")
	default :
		fmt.Printf("The division was not close at all")
	}


}

func printInt(value int) {
	fmt.Println(value)
}

func printMe() {
	fmt.Println("This is a Hello World message from PrintMe function")
}

func PrintParameter(myString string) {
	fmt.Println(myString)
}

func intDivision(numerator int,denominator int) (int,int,error) {
	var err error // default value of this error is nil
	if denominator==0{
		err = errors.New("Cannot Divide by Zero")

		return 0,0,err
	}

	var result int = numerator/denominator
	var reminder int = numerator%denominator
	return result, reminder,err
}