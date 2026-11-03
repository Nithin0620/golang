package main

import "fmt"

func main(){
	var intArr [3]int32 = [3]int32{1,2,3}
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Printf("%v , %v , %v  \n",intArr[0],intArr[1],intArr[2])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
}


func mainNew() {

	arr := [3]int32{1,2,3}
	arr2:= [...]int32{1,2,3,4,5} // size 5 array 
	fmt.Println(arr)

	fmt.Println(arr2)

	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)
	fmt.Printf("the Length is %v with capacity %v",len(intSlice),cap(intSlice)) //len is 3 and cap a;lso 3
	intSlice = append(intSlice,7)
	fmt.Printf("The Length ios %v with capacity %v",len(intSlice),cap(intSlice)) //llength is 4 with cap at 6


	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice,intSlice2...) //spread operator

	var intSlice3 []int32 = make(int32[],3,8)  //type,len,cap

	for i,v:= range arr{
		fmt.Printf("Index: %v, value:%v \n",i,v)
	}

}