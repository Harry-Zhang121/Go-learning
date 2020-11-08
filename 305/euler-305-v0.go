package main

import (
	"fmt"
	"math"
)

func printArray(array []int) {
	for i := 0; i < len(array); i++ {
		println(array[i])
	}

}

func arrayGrow(arr []int, num int) []int {

	tmpArr := make([]int, 0)
	for ; num > 0; num /= 10 {
		tmpArr = append(tmpArr, num%10)
	}

	for i := len(tmpArr) - 1; i >= 0; i-- {
		arr = append(arr, tmpArr[i])
	}

	return arr
}

func intToArray(num int) []int {
	tmpArr := make([]int, 0)
	for ; num > 0; num /= 10 {
		tmpArr = append(tmpArr, num%10)
	}

	arr := make([]int, 0)
	for i := len(tmpArr) - 1; i >= 0; i-- {
		arr = append(arr, tmpArr[i])
	}

	return arr

}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func f(num int) int {
	arr := make([]int, 0)
	var counter int = 0
	var position int = 0

	numArr := intToArray(num)
	var numArrLen int = len(numArr)

	for i := 1; counter < num; i++ {
		arr = arrayGrow(arr, i)
		for ; position < len(arr)-numArrLen; position++ {
			if len(arr) < numArrLen {
				continue
			}
			if equalSlice(arr[position:position+numArrLen], numArr) {
				counter++

			}
		}
	}

	// fmt.Printf("%v\n", arr)

	return position
}

func job(left int, right int) int {
	var sum int = 0

	for i := left; i <= right; i++ {
		sum += f(int(math.Pow(3, float64(i))))
		fmt.Printf("Calulating f(%v)\n", int(math.Pow(3, float64(i))))
	}

	return sum
}

// func main() {
// 	// var result int = 0

// 	// result = f()

// 	// fmt.Println(job(1, 13))
// 	fmt.Println(f(7780))
// }

// ~/D/e/g/305 ❯❯❯ time go run euler-305.go
// 111111365

// ________________________________________________________
// Executed in    5.62 secs   fish           external
//    usr time    6.17 secs  651.00 micros    6.17 secs
//    sys time    1.46 secs    0.00 micros    1.46 secs
