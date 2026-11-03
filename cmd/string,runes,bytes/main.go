package main

import "fmt"

//strings are stored as utf-8 values in go unlike some languages 
//which store characters as ascii value 
//utf-8 is a 8bit number with variable byte size it extends or reduces its 
//size as per the letter of characters 
// ascii is not prefered as there coul be only 233 somehting kind of vlayue passible 
// as ascii vlaues uses 7bits only hence we ahve other option of using 
// utf32 utf 16 and all but they are wastage of so much space 
// hence the go lang uses utf8 with extendable size as per the character 

func main() {
	var myString = "Nithin"
	var indexed = myString[1]
	fmt.Printf("%v,%T\n",indexed,indexed) //%T is for type of the vlaue prints:=105,uint8

	for i,v := range myString{
		fmt.Println(i,v)
	} //print index , value of the character in utf8 format

	fmt.Printf("%v,%T\n",string(indexed),string(indexed)) //converting the uint8 value to string and printing it
	//prints:= i,string

	fmt.Printf("\n The length of the string is %v",len(myString)) // this prints the number of bytes and not the lenth of the actual string as the string is stored in utf8 format and hence the length of the string is not equal to the number of bytes in the string
}