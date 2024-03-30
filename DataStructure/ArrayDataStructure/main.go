package main

import (
	"fmt"
)

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	revArray := ReverseArray(arr)

	fmt.Println("Reverse Array is ", revArray)

	rotatedArray := RotateArray(arr)

	fmt.Println("Rotated Array is ", rotatedArray)
	rotatedArray = RotateArray(RotateArray(rotatedArray))

	fmt.Println("Rotated Array is ", rotatedArray)

	array := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	sum := MaxConitigousSum(array)
	fmt.Println("Max is  ", sum)

	// Polidram number test

	fmt.Println("Is Number piliddraom ", IsPalindramNumber(123121))

	fmt.Println("Max So far is ", FindConsigativeSubarrayWithMaxSum(array))

	seq := []int{1, 0, 1, 1, 0, 1, 0, 0, 1}
	fmt.Println("Sequence the number ", Segregate1andO(seq))

	seq1 := []int{1, 0, 2, 1, 1, 0, 1, 2, 0, 0, 1, 2, 2}
	fmt.Println("Sequence the number ", Segragate012(seq1))

	ballList := []string{"red", "blue", "white", "blue", "red", "white"}
	fmt.Println("Sorted ball is ", SortBall(ballList))

	Practce1(seq1)
}

func ReverseArray(input []int) []int {
	var revArra = make([]int, len(input))
	n := len(input)
	for i := 0; i < n; i++ {
		revArra[i] = input[(n-1)-i]

	}
	return revArra
}

func RotateArray(input []int) []int {
	n := len(input)
	temp := input[n-1:]
	temp = append(temp, input[:n-1]...)

	// for i := n - 1; i > 0; i-- {
	// 	input[i] = input[i-1]
	// }
	// input[0] = lastElement

	// temp := input[1:]
	// temp = append(temp, first)
	return temp
}

func RotateArraySpace(input []int) []int {
	n := len(input)
	lastElement := input[n-1]
	for i := n - 1; i > 0; i-- {
		input[i] = input[i-1]
	}
	input[0] = lastElement

	return input
}

func MaxConitigousSum(input []int) int {
	max_so_far := input[0]
	maxhere := 0
	start := 0
	end := 0
	s := 0
	fmt.Println("Input", input)
	for i := 0; i < len(input); i++ {

		fmt.Println("Step O ", i, maxhere, input[i], max_so_far)
		maxhere = maxhere + input[i]

		if maxhere < 0 {
			maxhere = 0
			s = i + 1
		}

		if max_so_far < maxhere {
			max_so_far = maxhere
			start = s
			end = i
		}
		fmt.Println(" ", max_so_far, maxhere)
	}
	fmt.Println("Start and end us ", start, end, input[start:end+1])
	return max_so_far
}

func IsPalindramNumber1(num int64) bool {
	arr := []int64{}
	for {
		rem := num % 10
		num = num / 10
		arr = append(arr, rem)
		if num == 0 {
			break
		}
	}

	for i := 0; i <= len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}
	// fmt.Println("Next polidram ", arr)
	return true
}
func IsPalindramNumber(num int64) bool {
	var newNum int64
	// arr := []int{}
	for {
		rem := num % 10
		num = num / 10
		newNum = newNum*10 + rem

		if num == 0 {
			break
		}
	}
	return newNum == num
}

func FindConsigativeSubarrayWithMaxSum(input []int) int {

	var max_so_far int = input[0]
	var max_for_here int = 0
	var s, start, end int
	for i := 0; i < len(input); i++ {

		max_for_here = max_for_here + input[i]

		if max_for_here > max_so_far {
			max_so_far = max_for_here
			start = s
			end = i
		}
		if max_for_here < 0 {
			max_for_here = 0
			s = i + 1
		}
	}
	fmt.Println("Start index is ", start, " End is ", end, input[start:end+1])
	return max_so_far
}

func Segregate1andO(input []int) []int {

	start := 0
	end := len(input) - 1

	for start < end {
		if input[start] == 1 {
			if input[end] != 1 {
				input[start], input[end] = input[end], input[start]
			}
			end--
		} else {
			start++
		}
	}

	return input
}
func Segragate012(input []int) []int {
	low := 0
	mid := 0
	hi := len(input) - 1

	for i := 0; i < len(input); i++ {
		if input[mid] == 0 {
			input[low], input[mid] = input[mid], input[low]
			low++
			mid++
		} else if input[mid] == 1 {
			mid++
		} else {
			input[mid], input[hi] = input[hi], input[mid]
			hi--
		}

	}
	return input
}

func SortBall(input []string) []string {

	var low, mid int = 0, 0
	var high = len(input) - 1

	for _, ball := range input {
		if ball == "red" {
			fmt.Println(" Red", low, mid)
			fmt.Println("Before 11 ", input)
			input[low], input[mid] = swap(input[low], input[mid])
			fmt.Println("After 11 ", input)
			low++
			mid++
		} else if ball == "blue" {
			fmt.Println("Mid only ", mid)
			fmt.Println(" Blue ")
			mid++
		} else {
			fmt.Println(" white ", high)
			fmt.Println("Before 22 ", input)
			input[mid], input[high] = swap(input[mid], input[high])
			fmt.Println("After  22 ", input)
			high--
		}
	}
	return input
}

func swap(a string, b string) (string, string) {

	return b, a
}

func Practce1(input []int) {
	var low, mid int
	var high = len(input) - 1

	for _, v := range input {
		if v == 0 {
			input[low], input[mid] = input[mid], input[low]
			mid++
			low++
		} else if v == 1 {
			mid++
		} else if v == 2 {
			input[mid], input[high] = input[high], input[mid]
			high--
		}

	}
	fmt.Println("Result is Practises ", input)
	PrintNumber()
}

func PrintNumber() {
	var x, y int

	fmt.Scanf("%d %d", &x, &y)

	fmt.Println(x * y)
}
