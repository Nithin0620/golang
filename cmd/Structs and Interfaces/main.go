package main
import ("fmt")

type gasEngine struct{
	mpg uint8
	gallons uint8
}


type petrolEngine struct{
	mpg uint8
	gallons uint8
	ownerInfo Owner
}

type Owner struct {
	name string
}

func (e gasEngine) getMiles() uint8{
	return e.gallons * e.mpg
} // this is a method of this struct gas engoine we have assiend the struct and given nthe varibales access to this function using th epart (e gasEgine) which is called the receiver of this function and we can use this to access the variables of the struct

func (e electricEngine) getMiles() uint8{
	return e.batteryLife * e.efficiency
} // this is a method of this struct electric engoine we have assiend the struct and given nthe varibales access to this function using th epart (e electricEngine) which is called the receiver of this function and we can use this to access the variables of the struct

type electricEngine struct{
	batteryLife uint8
	efficiency uint8
}

func canmakeIT (e engine , miles uint8) { // notice this can only take in the single struct of type gasEngine only and it cant takr the petrolEngione struct this isn where interface comes into play as we can create an interface and then use that interface to take in any struct that implements the methods of that interface
	if e.getMiles() >= miles {
		fmt.Println("Yes, you can make it!")
	} else {
		fmt.Println("No, you cannot make it.")
	}
}

type engine interface{
	getMiles() uint8
}

func main() {
	var myEngine gasEngine

	fmt.Println(myEngine.mpg,myEngine.gallons) //prints 0 0 as the default value of uint8 is 0

	myEngine.mpg = 30
	myEngine.gallons = 10

	fmt.Printf("%v,%T\n",myEngine,myEngine) //prints {30 10},main.gasEngine as the type of the struct is main.gasEngine

	var newEngine gasEngine = gasEngine{mpg:25,gallons:25}
	var newEngine2 gasEngine = gasEngine{20,20} //this is also valid as the order of the values is same as the order of the fields in the struct

	var newpetrolEngine petrolEngine = petrolEngine{20,20,Owner{"Nithin"}} //this is also valid as the order of the values is same as the order of the fields in the struct

	fmt.Printf("Total miles left in the tank: %v",myEngine.getMiles())//this is how we can call the method of the struct and get the value of the total miles left in the tank


	var myEngine3 = struct{
		mpg uint8
		gallons uint8
	}{25,30} //this is an anonymous struct and can be used when we don't want to create a new struct type
	


	var myEngineElectric electricEngine = electricEngine{batteryLife: 100, efficiency: 2}

	canmakeIT(myEngine, 200) //this will print Yes, you can make it! as the total miles left in the tank is 300 which is greater than 200
	canmakeIT(myEngineElectric, 200) //this will print No, you cannot make it. as the total miles left in the tank is 200 which is equal to 200

}  