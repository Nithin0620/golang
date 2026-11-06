package main

type gasEngine struct {
	gallons uint8
	mpg uint8
}

type electricEngine struct {
	kwh uint8
	mpkwh uint8
}


type car[T gasEngine | electricEngine] struct {
	carMake string
	carModel string
	engine T
}


func main() {
	var gasCar = car[gasEngine]{
		carMake : "Toyota",
		carModel : "camry",
		engine : gasEngine{
			gallons : 10,
			mpg : 30,
		},
	}
	
	var electricCar = car[electricEngine]{
		carMake : "Tesla",
		carModel : "Model 3",
		engine : electricEngine{
			kwh : 75,
			mpkwh: 4,
		},
		
	}

}