package main
import "fmt"
import "unicode/utf8"


func main() {
	fmt.Println("Hello World!")

	fmt.Println("The value of intNum is:")
	var intNum int =32767

	intNum = intNum + 1 

	fmt.Println(intNum)


	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)



	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2 
	var result float32 = floatNum32 + intNum32  //this is wrong
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var myString string = "Hello World" //with double quotes u cant write ulti line string and with ` these quotes u can write and priont a multi line string`
	var myString1 string= "Hello" + " " + "World"


	fmt.Println(len(myString))  //>> 11 this is not length this is byte of the string 

	fmt.Println(utf8.RuneCountInString(myString))

	var myRune rune = 'a' //this is a single character and it is a int32 type
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var inNum3 int
	fmt.Println(inNum3)  //>> 0 this is default value of int type for all int and related types its 0 
	// for string its "" and for bool its false and for float its 0.0

	var myString2 = "text" //this is type inference and it will automatically assign the type of variable based on the value assigned to it

	myvar := 10 //this is short hand declaration and it can only be used inside a function and not outside of it

	var1,var2 := 10,20 //this is also short hand declaration and it can only be used inside a function and not outside of it
	fmt.Println(var1,var2)


	const myConst string = "This is a constant vlaue and its value cant be changed one written "
	myConst = "This is a new value" //this will give an error because we cant change the value of a constant once it is assigned
	fmt.Println(myConst)





} 