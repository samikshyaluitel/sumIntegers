package main

import (
	"fmt"
	"math"
)

func Add(num1 []int, num2 []int) string {
	// implement this
	var l1 = len(num1)
	var l2 = len(num2)

	var greaterLength = 0
	if l1 > l2 {
		greaterLength = l1 - 1
	} else {
		greaterLength = l2 - 1
	}
	// for reversal of second number
	indexNum2 := 0
	indexNum1 := l1 - 1
	carry := 0
	total := 0
	for greaterLength >= 0 {
		var i1 = 0
		var i2 = 0
		if indexNum1 < l1 && indexNum1 >= 0 {
			i1 = num1[indexNum1]
		}
		if indexNum2 < l2 && indexNum2 >= 0 {
			i2 = num2[indexNum2]
		}
		indexSum := i1 + i2 + carry
		remainder := indexSum % 10
		carry = indexSum / 10
		total = total + remainder*int(math.Pow10(indexNum2))
		indexNum2++
		greaterLength--
		indexNum1--
	}

	return fmt.Sprintf("%d", total)
}

func main() {

	var num1 []int
	var num2 []int

	var input1 int
	var input2 int

	i := 0

	fmt.Println("Number of digits is first input:")
	_, err := fmt.Scanf("%d", &i)

	for j := 1; j <= i; j++ {
		fmt.Printf("Enter the  digit number %v : ", j)
		_, err = fmt.Scanf("%d", &input1)
		num1 = append(num1, input1)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("The slice is:", num1)

	fmt.Println("Number of digits is second input:")
	_, err = fmt.Scanf("%d", &i)

	for k := 1; k <= i; k++ {
		fmt.Printf("Enter the digit number %v: ", k)
		_, err = fmt.Scanf("%d", &input2)
		num2 = append(num2, input2)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("The slice is:", num2)

	result := Add(num1, num2)

	fmt.Println("Result:", result)
}
