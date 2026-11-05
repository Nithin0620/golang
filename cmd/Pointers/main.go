package main
import "fmt"

func main() {
	var p *int32 = new(int32)

	var z *int32 // this is a pointer intiallisation which intialises to nill address but this will give a run time error as no adress is assigned to this variable yet 

	var i int32
	fmt.Printf("The valyue of p points to is %v\n", *p) // prints 0 as the default value of int32 is 0
	fmt.Printf("The address of p is %v\n", p) // prints the address of p
	fmt.Printf("The address of i is %v\n", &i) // prints the address of i
	fmt.Printf("The value of i is %v\n", i) // prints 0 as the default value of int32 is 0

	*p = 10

	p =&i // this is how we can assign the address of a variable to a pointer variable
	*p = 10 // this is how we can assign a value to the variable that the pointer points to

	fmt.printf("The value of p points to is %v\n", *p) // prints 10 as we have assigned the value 10 to the variable that p points to




	var slice = [int32{1,2,3}]
	var sliceCopy = slice 
	sliceCopy[2] = 4 
	fmt.Println(slice) // prints [1 2 4] as the the slice uses pointer under the hood to create copy and hence when we change the value of the sliceCopy it also changes the value of the original slice as they both point to the same underlying array
	fmt.Println(sliceCopy) // prints [1 2 4] as the the slice uses pointer under the hood to create copy and hence when we change the value of the sliceCopy it also changes the value of the original slice as they both point to the same underlying array

	


	var thing1 = [5]float64{1,2,3,4,5}
	fmt.printf("The memory location opf the thing1 array os: %p\n", &thing1) // prints the memory location of the thing1 array
	var result [5]float64 = square(thgiung1) // this is how we can pass an array to a function and get the result back as an array
	fmt.printf("The memory location of the result array is: %p\n", &result) // prints the memory location of the result array
	fmt.printf("Ther memory location of the thing1 array is: %p\n", &thing1) // prints the memory location of the thing1 array

	var result2 [5]float64 = *squarePointer(&thing1) // this is how we can pass an array to a function and get the result back as an array using pointers
	fmt.printf("The memory location of the result2 array is: %p\n", &result2) // prints the memory location of the result2 array
	fmt.printf("The memory location of the thing1 array is: %p\n", &thing1) // prints the memory location of the thing1 array
	// now the addr of result2 and thing1 are same as we have passed the address of thing1 to the function and got the result back as an array using pointers

}


func square(thing2 [5]float64) [5]float64 { // this is how we can pass an array to a function and get the result back as an array this is pass by valuye example here a complete copy of the parameters is created wasting so much memory and time so we can use pointers to pass the address of the array to the function and get the result back as an array this is pass by reference example here only the address of the parameters is passed saving memory and time
	fmt.Printf("The memory location of the thing2 array is: %p\n", &thing2) // prints the memory location of the thing2 array
	for i, v := range thing2 {
		thing2[i] = v * v
	}
	return thing2
}


func squarePointer(thing2 *[5]float64) *[5]float64 { // this is how we can pass an array to a function and get the result back as an array this is pass by reference example here only the address of the parameters is passed saving memory and time
	fmt.Printf("The memory location of the thing2 array is: %p\n", thing2) // prints the memory location of the thing2 array
	for i, v := range thing2 {
		thing2[i] = v * v
	}
	return thing2
}