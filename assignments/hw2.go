package assignments

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/**
Write a program that takes a value “n” and “m” as input; produces “n” random numbers with a uniform distribution
between 1 and m and places them in a linked list. Order does not matter. cost O(n)
*/
func RunHw2ProgramOne(n int, m int) {
	list := NodeList{}
	rand.Seed(time.Now().UnixNano()) //cost 1 unit time
	for n > 0 {
		if m == 1 {
			list.Append(1) //rand.Intn(x) x must >=1	O(1)
		} else {
			list.Append(rand.Intn(m-1) + 1) //O(1)
		}
		n--
	}
}

/**
Write a program that takes a value “n” and “m” as input; produces “n” random numbers with a uniform distribution
between 1 and m and places them in a linked list. Order does not matter. In this case, however ensure the list has
no duplicates
*/
func RunHw2ProgramTwo(n int, m int) []int{
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

func Run2testHw1ProgramOne(ch chan int) {
	i := 10
	for i <= 10000 {
		start1 := time.Now()
		RunHw1ProgramOne(i)
		end1 := time.Since(start1).Microseconds()
		writeFile("RunHw1ProgramOne-", fmt.Sprint(i)+"\t"+strconv.FormatInt(end1, 10)+"\n")
		i++
	}
	ch <- 1
}

func Run2testHw1ProgramTwo(ch chan int) {
	i := 10
	for i <= 10000 {
		start1 := time.Now()
		RunHw1ProgramTwo(i)
		end1 := time.Since(start1).Microseconds()
		writeFile("RunHw1ProgramTwo-", fmt.Sprint(i)+"\t"+strconv.FormatInt(end1, 10)+"\n")
		i++
	}
	ch <- 1
}

func Run2testHw1ProgramThree(ch chan int) {
	i := 10
	for i <= 10000 {
		start1 := time.Now()
		RunHw1ProgramThree(i)
		end1 := time.Since(start1).Microseconds()
		writeFile("RunHw1ProgramThree-", fmt.Sprint(i)+"\t"+strconv.FormatInt(end1, 10)+"\n")
		i++
	}
	ch <- 1
}

func RunHw2Main() {
	chs := make([]chan int, 2)
	chs[0] = make(chan int)
	chs[1] = make(chan int)
	// chs[2] = make(chan int)

	go Run2testHw1ProgramOne(chs[0])
	// go Run2testHw1ProgramTwo(chs[1])
	go Run2testHw1ProgramThree(chs[1])

	for _, ch := range chs {
		<-ch
	}
	fmt.Println("done")
}
