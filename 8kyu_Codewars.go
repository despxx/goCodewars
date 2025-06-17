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
