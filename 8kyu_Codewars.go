// package main

// import "fmt"

// func main() {
// 	var number int
// 	fmt.Println("Введите число: ")
// 	fmt.Scan(&number)

// 	if number%2 == 0 {
// 		fmt.Println("Even")
// 	} else {
// 		fmt.Println("Odd")
// 	}
// }

// ---------------------------------------------------
package kata

// MARK: - 8kyu

/*
Even or Odd
*/

func EvenOrOdd(number int) string {

	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

//---------------------------------------------------

/*
Multiply
*/

func Multiply(a, b int) int {
	return a * b
}

//---------------------------------------------------

/*
Return Negative
*/

func MakeNegative(x int) int {
	if x > 0 {
		return -x
	}
	return x
}

//---------------------------------------------------

/*
Sum of positive
*/

func PositiveSum(numbers []int) (sum int) {
	for _, num := range numbers {
		if num > 0 {
			sum = sum + num
		}
	}
	return
}

//---------------------------------------------------
