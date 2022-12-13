package assignments

import (
	"fmt"
	"testing"
)

func Test_RandomNumberGeneratorWithskewedRandomNumber(t *testing.T) {
	t.Skip()
	fmt.Println("Test_RandomNumberGeneratorWithskewedRandomNumber")
	count := [102]int{0}
	for i := 0; i < 1000000; i++ {
		count[skewedRandomNumber(100)]++
	}
	for k, v := range count {
		if v > 0 {
			fmt.Println(fmt.Sprint(k) + ": " + fmt.Sprint(v))
		}
	}
}

func Test_RandomNumberGeneratorWithUniformRandomNumber(t *testing.T) {
	t.Skip()
	fmt.Println("Test_RandomNumberGeneratorWithUniformRandomNumber")
	count := [102]int{0}
	for i := 0; i < 100000; i++ {
		count[uniformRandomNumber(100)]++
	}
	for k, v := range count {
		if v > 0 {
			fmt.Println(fmt.Sprint(k) + ": " + fmt.Sprint(v))
		}
	}
}

func Test_RandomNumberGeneratorWithAmorRandomNumber(t *testing.T) {
	t.Skip()
	fmt.Println("Test_RandomNumberGeneratorWithAmorRandomNumber")
	count := [101]int{0}
	for i := 0; i < 10000; i++ {
		count[amorRandomNumber(100)]++
	}
	for k, v := range count {
		if v > 0 {
			fmt.Println(fmt.Sprint(k) + ": " + fmt.Sprint(v))
		}
	}
}

func TestRandomNumberSequenceGenerator(t *testing.T) {
	t.Skip()
	fmt.Println("TestRandomNumberSequenceGenerator")
	s1 := randomNumberSequenceGenerator(10, 100, 1)
	fmt.Println("uniform distribution:")
	fmt.Println(s1)
	s2 := randomNumberSequenceGenerator(10, 100, 2)
	fmt.Println("skewed distribution")
	fmt.Println(s2)
	s3 := randomNumberSequenceGenerator(10, 100, 3)
	fmt.Println("my own distribution")
	fmt.Println(s3)

}
