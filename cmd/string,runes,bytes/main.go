package main

import ("fmt" "strings")

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

	var myString = []rune("Nithin") //Even I didnt understand this part of the code but it is used to convert the string into runes and store it in a slice of runes



	var strSlice= []string {"Nithin","is","a","good","boy"} //this is a slice of strings which is a collection of strings and it is used to store multiple strings in a single variable
	var catString = ""

	for i:= range strSlice{
		catString+=strSlice[i]
	}
	fmt.Printf("\n The concatenated string is %v",catString)
	// the string built using this kind of concatination the sting build is not editable they are immutable hence we use go built in strings package to build the string using the strings.Builder which is a more efficient way of building strings as it uses less memory and less time to build the string


	var strSlice= []string {"Nithin","is","a","bad","boy"} //this is a slice of strings which is a collection of strings and it is used to store multiple strings in a single variable
	var strBuilder strings.Builder //this is a struct which is used to build the string using the strings.Builder which is a more efficient way of building strings as it uses less memory and less time to build the string

	for i:= range strSlice{
		strBuilder.writeString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n The concatenated string using the string builder struct is %v",catStr)
	// so in this method internally an array is created and the strings are added to the array and then the array is converted to string and returned as a single string which is more efficient than the previous method of concatenation as it uses less memory and less time to build the string

}