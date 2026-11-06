package main

import "fmt"

func main() {
	var intSlice = []int{1,2,3,4,5}
	fmt.Println(sumSlice[int](intSlice)) // this is how we can call a generic function in go and pass the type of the slice as a parameter to the function and this will return the sum of the slice

	var floatSlice = []float64{1.1,2.2,3.3,4.4,5.5}
	fmt.Println(sumSlice[float64](floatSlice)) // this is how we can call a generic function in go and pass the type of the slice as a parameter to the function and this will return the sum of the slice

}


func sumSlice[T int| float64 | float32](slice []T) T { // this is how we can create a generic function in go and this will take in a slice of any type and return the sum of the slice and we can use the type parameter T to specify the type of the slice and we can use the type constraint to specify the types that are allowed for the type parameter T
	var sum T
	for _,v := range slice{
		sum +=v
	}

	return sum
}

// we also have an any type in this generics but it might not be compatible with all options alike + is not compatible with any type and hence we need to use the type constraint to specify the types that are allowed for the type parameter T
//