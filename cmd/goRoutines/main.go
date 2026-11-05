package main

/*
concurrency is the ability of a program to manage multiple tasks at the same time. 
In Go, concurrency is achieved through goroutines and channels.
A goroutine is a lightweight thread managed by the Go runtime, 
and channels are used to communicate between goroutines.
*/

// parallel Execution != concurrency

// as parallel execution means both task as execcutong parallely in diff ahrdware 
// and in concurrency both tasks are executed in the same hardware but they are executed in a way that they are interleaved and they are not executed in parallel
 
import ("fmt" "math/rand" "time" "sync")

var dbData = []string{"data1", "data2", "data3", "data4", "data5"}

func main() {
	// fmt.Println("Hello, World!")

	t0 := time.Now()
	for i:= 0; i < len(dbData); i++ { // without go keyword it will have to wait fdor the call return before going to the next loop and next iteration and hence it will take more time to complete the loop as it will have to wait for each dbCall to complete before going to the next iteration of the loop
		go dbCall(i)   // this is how we can create a goroutine in go and this will run the dbCall function in a separate goroutine and the main function will not wait for the dbCall function to complete and it will move to the next iteration of the loop and create another goroutine for the next iteration and so on
	}// but this will not work as expected as the main function will exit before the goroutines complete and hence we need to use wait groups to wait for the goroutines to complete before exiting the main function
	fmt.Printf("Time taken: %v\n", time.Since(t0))
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000 //simulating DB call delay
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("Data: %v, Delay: %v\n", dbData[i], delay)
}

var wg = sync.WaitGroup{}
//Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
var mu = sync.Mutex{} // this is how we can create a mutex in go and this will be used to lock the memory location when a goroutine is accessing it and unlock it when it is done accessing it
var results = []string{}
func main2 () {
	t0 := time.Now()
	for i:= 0; i < len(dbData); i++ { 
		wg.Add(1) // this is how we can add a goroutine to the wait group and this will increment the counter of the wait group by 1
		go dbCall2(i)  

	}
	wg.Wait() // this is how we can wait for all the goroutines in the wait group to complete and this will block the main function until all the goroutines in the wait group have completed
	fmt.Printf("Time taken: %v\n", time.Since(t0))
}

func dbCall2(i int) {
	var delay float32 = rand.Float32() * 2000 //simulating DB call delay
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("Data: %v, Delay: %v\n", dbData[i], delay)
	m.Lock() // this is how we can lock the memory location when a goroutine is accessing it and unlock it when it is done accessing it
	results = append(results, dbData[i])
	m.Unlock() // this is how we can unlock the memory location when a goroutine is done accessing it and this will allow other goroutines to access the memory location
	wg.Done() // this is how we can mark a goroutine as done and this will decrement the counter of the wait group by 1
}// ther is a huge issue in this as they are running concurrently when they access ther same memory location at the same time there could be any number of issues like race conditions and data corruption and hence we need to use mutex to lock the memory location when a goroutine is accessing it and unlock it when it is done accessing it

// we also have a RWMutex which is a read write mutex and this will Not allow multiple goroutines to read the memory location at the same time and also only one goroutine can write to the memory location at a time and this will prevent race conditions and data corruption
// this works exactly same as mutex but this will not have 2 more functions RLock() and RUnlock() which will be used to lock the memory location for reading and unlock it for reading 


var mrw = sync.RWMutex{}

func save(result string){
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func Log(result string){
	mrw.RLock() // so this will be allowing multiple goroutines to read the memory location at the same time but it will not allow read when any write is happing means any normal lock is attained
	fmt.Printf("The current result is: %v\n", results)
	mrw.RUnlock()
}