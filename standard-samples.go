package main

import "fmt"
import "math"

//this is my comment

func main() {

	rect := Rectangle{20,50}
	circ := Circle{4}

	fmt.Println("Rectangle:", getArea(rect))
	fmt.Println("Circle:", getArea(circ))

}

type Shape interface {
	area() float64
}

type Rectangle struct{
	height float64
	width float64
}

type Circle struct{
	radius float64
}

func (r Rectangle) area() float64{
	return r.height * r.width
}

func (c Circle) area() float64{
	return math.Pi * math.Pow(c.radius,2)
}

func getArea(shape Shape) float64{
	return shape.area()
}

func ImpRectangle(){

	//rec1 := Rectangle{leftX:0, topY:50, height:10, width:10}
	//rec2 := Rectangle{20,501,101,110}

	//fmt.Println(rec1.width)
	//fmt.Println(rec2.width)

	//rec3 := Rectangle{20,501,101,110}
	//fmt.Println("Area of rectangle is", rec3.area())

}

//type Rectangle struct {
	
//	leftX float64
//	topY float64
//	height float64
//	width float64
//}

//func (rect *Rectangle) area() float64{
//
//	return rect.width * rect.height
//
//}

func CallXValNow(){
	x := 0

	changeXValNow(&x)
	fmt.Println("x = ", x)
}

func changeXValNow(x *int){

	*x = 2

}

func demPanic(){

	defer func(){
		fmt.Println(recover())
	}()

	panic("PANIC")

}

func safeDiv(num1, num2 int) int{

	//fmt.Println(safeDiv(3,0))
	//fmt.Println(safeDiv(3,2))

	defer func(){
		fmt.Println(recover())  //allows us to continue even with fatal error
		//recover()   this just allows you to continue without showing the error message
	}()

	solution := num1 / num2
	return solution

}

func DeferPlease(){

	defer printTwo()
	printOne()

}

func printOne() {fmt.Println(1)}
func printTwo() {fmt.Println(2)}

func factorial(num int) int {

//fmt.Println(factorial(3))

	if num == 0 {
		return 1
	}

	return num * factorial(num - 1)

}

func FunctionwithinFunction(){

	num3 := 3

	doubleNum := func() int {
	
		num3 *= 2
		return num3
	
	}

	fmt.Println(doubleNum())

}

func CallSubtractThem(){

fmt.Println(subtractThem(1,2,3,4,5))

}

func subtractThem(args ...int) int {

	finalValue := 0

	for _, value := range args {
		finalValue -= value
	}

	return finalValue

}

func CallAddThemUp(){

	listNums := []float64{1,2,3,4,5}

	fmt.Println("Sum: ", addThemUp(listNums))

}

func addThemUp(numbers []float64) float64 {

	sum := 0.0

	for _, val := range numbers {
	
		sum += val

	}

	return sum

}

func myMaps() {

	presAge := make(map[string] int)

	presAge["Mark Pownell"] = 43
	presAge["John Smith"] = 44

	fmt.Println(presAge["Mark Pownell"])

	delete(presAge, "John Smith")

	fmt.Println(len(presAge))

}

func mySlice(){

	numSlice := []int  {5,4,3,2,1}

	numSlice2 := numSlice[3:5]

	fmt.Println("numSlice2[2] = ", numSlice2[1])

	numSlice3 := make([]int, 5, 10)

	copy(numSlice3, numSlice)

	fmt.Println(numSlice3[0])

	numSlice3 = append(numSlice3, 0, -1)

	fmt.Println(numSlice3[6])

}

func test() {

	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true && false = ", !true && false)
	
	favNums3 := [5] float64 {1,2,3,4,5}

	for _, value := range favNums3 {
		fmt.Println(value)
	}

	//for i, value := range favNums3 {
	//	fmt.Println(value, i)
	//}

	var favNums2[5] float64

	favNums2[0] = 1232132
	favNums2[1] = 1232132
	favNums2[2] = 1232132
	favNums2[3] = 434543543534
	favNums2[4] = 1232132

	fmt.Println(favNums2[3])

	yourAge := 15

	if yourAge > 16 {
		fmt.Println("You can drive")
	} else if yourAge == 16 {
		fmt.Println("You can drive restricted")
	} else {
		fmt.Println("You can't drive")
	}

		switch yourAge {
		case 16 : fmt.Println("Drive")
		case 18 : fmt.Println("Yes")
		default : fmt.Println("No Drive")
	}
	
	for j := 0; j < 5; j++ {

		fmt.Println(j)

	}

		i := 1
	
	for i <= 10 {

		fmt.Println(i)

		i++
	}

    fmt.Println("Hello World")

	var age int = 40

	var favNum float64 = 1.6843434

	randNum := 1

	fmt.Println(randNum, age, favNum)

	var one1 = 1.000
	var two2 = .999999

	fmt.Println(one1 - two2)

	fmt.Println(6/5, 6%5)

	const pi float64 = 3.1445434

	var (
		varA = 1445434
		varB = 32432
	)

	var myName = "Mark Pownell"

	fmt.Println(len(myName), varA, varB)

	fmt.Println(myName + " is a robot \n and is very nice")

	var isOver40 bool = true

	fmt.Printf("%.5f \n", pi)

	fmt.Printf("%T \n", isOver40)

}