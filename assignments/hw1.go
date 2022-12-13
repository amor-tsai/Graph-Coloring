package assignments

import (
	"fmt"
	"math/rand"
	"time"
)

// Write a program that takes a value “n” as input and prints “Hello, World” n times.

func RunHw1ProgramOne(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Hello, World") // cost n unit_time
	}
}

/**
Write a program that takes a value “n” as input; produces “n” random numbers with a uniform distribution between 1 and n and
places them in an array. Order does not matter.
*/
func RunHw1ProgramTwo(n int) {
	arr := make([]int, n)
	rand.Seed(time.Now().UnixNano()) //cost 1 unit time
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n-1) + 1 //cost n unit time
	}
}

/**
Write a program that takes a value “n” as input; produces “n” random numbers with a uniform distribution between 1 and n and
places them in an array in sorted order. Place them in the array in order, do not sort the array after placing them there.
*/
func RunHw1ProgramThree(n int) {
	if n < 1 {
		panic("n couldn't less than 1")
	}
	arr := make([]int, n, n)
	rand.Seed(time.Now().UnixNano())
	var num int
	for i := 0; i < n; i++ { //O(n)
		num = rand.Intn(n) + 1
		j := i - 1
		for ; j >= 0; j-- { // O(n-1)
			if arr[j] > num {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = num
	}
}
