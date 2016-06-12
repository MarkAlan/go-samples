package main

import ("fmt"
"time"
"strconv"
)

/*
Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

By default, sends and receives block until the other side is ready. 

This allows goroutines to synchronize without explicit locks or condition variables.

*/

var pizzNum = 0
var pizzaName = ""

func makeDough(stringChan chan string){
	pizzNum++

	pizzaName = "Pizza #" + strconv.Itoa(pizzNum)

	fmt.Println("Make Dough and Send for Sauce")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)

}

func addSauce(stringChan chan string){
	pizza := <- stringChan 

	fmt.Println("Add Sauce and Send", pizza, "for toppings")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)

}

func addToppings(stringChan chan string){
	pizza := <- stringChan
	fmt.Println("Add Toppings to", pizza, "and ship")
	time.Sleep(time.Millisecond * 10)
}

func main(){

	stringChan := make(chan string)

	for i := 0; i < 3; i++{
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)
	}
	time.Sleep(time.Millisecond * 5000)
}

