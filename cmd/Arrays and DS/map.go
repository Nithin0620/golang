package main

import "fmt"

func main() {

	var myMap map[string]uint8 = make(map[string]uint8)

	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam":23,:"Sarah":45}
	fmt.Println(myMap2[Adam])
	fmt.Println(myMap2[Janson]) // return default value that is 0
	var age,ok = myMap2[Adam] //ok is true
	if ok{
		fmt.printf("The Age is %v",age)
	}else {
		fmt.Printf("Invalid Name")
	}
	var age,ok = myMap2[Janson] // ok is false	

	delete(myMap2,"Adam")
	
	// Create a map to store student names and their corresponding grades
	grades := make(map[string]int)\







	for name:= range myMap2{
		fmt.Printf("Name: %v\n",name)
	}

	for name ,age:= range myMap2{
		fmt.Printf("Name: %v, Age:%v \n",name,age)
	}


	var i int = 0

	for i<10{
		fmt.Println(i)
		i++
	}

	for i:=0;i<10;i++ {
		fmt.Println(i)
	}

}
	