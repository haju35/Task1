package main

import "fmt"


func sumOfIntegers(numbers []int) int{
	if len(numbers) == 0{
		return 0
	}

	sum := 0
	for _, n := range numbers{
		sum += n
	}
	return sum

}

func main(){
	num1 := []int{-10, -5, 0, 5, 10, 15}
	num2 := []int{}

	fmt.Println(sumOfIntegers(num1))
	fmt.Println(sumOfIntegers(num2))
}
