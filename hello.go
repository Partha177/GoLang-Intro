package main //user created package

import (
	"fmt"
	"math"
	"time"
) //in build package

var c = 8 //package level variable

func main()  {
	var num1 int =2  //function level variable
	var num3 int
	var num4, num5 int
	num6 := 6 //var num6 = 6
	var num2 = 3
	const num7 = 7
	// num7=0 //not allowed

	num3 = 4
	num4, num5 = 5, 6
	// fmt.Println("Hello, World")
	
	
	var result = num1 + num2 + num3 + num4 + num5 + num6
	fmt.Println(result)
	fmt.Println("main c=", c)

	//There is only 1 loop in Go i.e. for loop in 3 formats

	//infinite for
	// for {
	// 	fmt.Println("GoLang")
	// }

	//2nd format
	// for i:=0; i<=5; i++ {
	// 	fmt.Println("GoLang", i)
	// }

	//3rd format
	j := 2
	for j<=3 {
		fmt.Println("j is ", j)
		j++
	}

	// function
	result3, result4 := calc(num3, num4)
	
	fmt.Println(result3, result4) //Printf(), Print() also available😊😂🧇🚙⛳🚩
	result2 := add(num1, num2)
	fmt.Println(result2)
	var result5, result6 = remainder(num5, num6)
	fmt.Println(result5, result6)


	//use of mth package
	var num float64 = 9
	var n2 float64 = 12
	
	var res = math.Sqrt(num)  //Sqrt() only takes float input
	var res2 = math.Sqrt(n2)
	var res3 = math.Round(res2)
	var res4 = math.Ceil(res2)
	var res5 = math.Floor(res2)
	fmt.Println("3 is square root is =", res)
	//round to 2 decimal points
	fmt.Printf("Result is %.2f thanks\n", res2)
	//rounding off
	fmt.Printf("Result is %.2g thanks\n", res2)  //using placeholder
	fmt.Printf("Result is %.2f thanks\n", res3)  
	fmt.Printf("Result is %.2f thanks\n", res4)  
	fmt.Printf("Result is %.2f thanks\n", res5)  


	n := 5
	if n<5 {
		fmt.Println("Hi", n)
	} else if n==5 {
		fmt.Println("Oh", n)
	} else {
		fmt.Println("Bye ", n)
	}

	p := 4
	switch p {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("None")
	}


	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	fmt.Println(time.Saturday)
	fmt.Println(today + 2)


}

func add(x int, y int) int {
	c := 4
	fmt.Println("add c=", c)
	var sum = x + y
	return sum
}

//if both same type  //if 2 returns
func calc(a, b int) (int, int){
	prod := a*b
	var div = a/b
	return prod, div
}

func remainder(x, y int) (out1, out2 int){
	out1 = x%y
	out2 = y%x
	return
}

//exported names/functions
//If a function name starts with capital letter then we can access that outside the package
//if a function name starts with small letter then we can't access it outside that package/file
//func Demo()  it can be accessed from separate package/file




                   

