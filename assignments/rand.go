package assignments

import (
	"math/rand"
	"time"
)

/**
Write a program that takes a value “n” and “m” as input; produces “n” random numbers with a uniform distribution
between 1 and m and places them in a linked list. Order does not matter. cost O(n)
*/
func generateUniformDistributedRandomNumber(n int, m int) []int {
	if n < 1 || m < 1 {
		panic("n or m should greater than 1!")
	}
	list := NodeList{}
	rand.Seed(time.Now().UnixNano()) //set up seed
	var randList = make([]int, m)
	for i := 0; i < m; i++ { //O(m)
		randList[i] = i + 1
	}
	var randIndex int
	for i := m; i > m-n; i-- { // O(n)
		if i == 0 { //if n>m, the loop should be stopped in advance
			break
		}
		randIndex = rand.Intn(i) // generates a number between 0 and m-1
		list.Append(randList[randIndex])
		tmp := randList[i-1]
		randList[i-1] = randList[randIndex]
		randList[randIndex] = tmp
	}
	return randList
}

//use 1uniform,2skewed,3my own distributionn generator to generate different number sequence
//distributedType
func randomNumberSequenceGenerator(needNumber int, numberRange int, distributedType int) []int {
	if needNumber < 1 || numberRange < 1 {
		panic("needNumber or numberRange should be greater than 1")
	}
	if distributedType > 3 || distributedType < 1 {
		panic("distributedType should be between 1 and 3")
	}
	rand.Seed(time.Now().UnixNano())
	list := make([]int, needNumber)
	rand.Seed(time.Now().UnixNano())
	var randList = make([]int, numberRange)
	for i := 0; i < numberRange; i++ {
		randList[i] = i + 1
	}
	var randIndex int
	for i := 0; i < needNumber; i++ {
		if numberRange-i == 0 {
			randIndex = 0
		} else if 1 == distributedType {
			randIndex = uniformRandomNumber(numberRange-i) - 1
		} else if 2 == distributedType {
			randIndex = skewedRandomNumber(numberRange-i) - 1
		} else {
			randIndex = amorRandomNumber(numberRange-i) - 1
		}
		list[i] = randList[randIndex]
		randList[randIndex], randList[numberRange-i-1] = randList[numberRange-i-1], randList[randIndex]
	}
	return list
}

/**
generate uniform random number from 1 to max
*/
func uniformRandomNumber(max int) int {
	return rand.Intn(max) + 1
}

/**
generate skewed random number between 1 and max
deprecated because of its low efficient
*/
func skewedRandomNumber(max int) int {
	v := max * (max + 1) / 2
	k := uniformRandomNumber(v)
	m := 0
	for i := 1; i <= max; i++ {
		m += max - i + 1
		if k <= m {
			return i
		}
	}
	return max
}

//create my own random number between 1 and max
func amorRandomNumber(max int) int {
	if max == 1 {
		return 1
	}
	a := uniformRandomNumber(10)
	if a > 2 {
		if max < 5 {
			return 1
		} else {
			return uniformRandomNumber(max / 5)
		}
	} else {
		return max/5 + uniformRandomNumber(max*4/5)
	}
}
