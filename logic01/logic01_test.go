package logic01

import (
	"fmt"
	"math"
	"testing"
)

func TestLogic01(t *testing.T) {
	n := 10
	angka := 1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(3, "\t")
		}
	}
}

func TestLogic02(t *testing.T) {
	n := 10
	angka := 1
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(3, "\t")
		}
	}
}

func TestLogic03(t *testing.T) {
	n := 10
	angka := 1
	x := 99
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(x, "\t")
		}
	}
}

func TestLogic04(t *testing.T) {
	n := 10
	angka := 1
	x := 777
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(x, "\t")
		}
	}
}

func TestLogic05(t *testing.T) {
	n := 15
	angka := 1
	for i := 0; i < n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("Fizz,Buzz", "\t")
		} else if i%3 == 0 {
			fmt.Print("Fizz", "\t")
			angka++
		} else if i%5 == 0 {
			fmt.Print("Buzz", "\t")
			angka++
		} else {
			fmt.Print(angka, "\t")
			angka++
		}
	}
}

func TestLogic06(t *testing.T) {

	n := 15
	nt := n / 2
	x := 3
	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			z := math.Pow(float64(i), 2)
			fmt.Print(z, "\t")
		} else {
			fmt.Print(x, "\t")
		}
		if i <= nt {
			x += 3
		} else {
			x -= 3
		}
	}
}

func TestLogic07(t *testing.T) {

}

func TestLogic08(t *testing.T) {

}

func TestLogic09(t *testing.T) {
	n := 12
	a := 1
	b := 2
	c := 3

	for i := 1; i < n; i++ {
		if i%3 == 2 {
			fmt.Print(b, "\t")
			b *= 2
		} else if i%3 == 0 {
			fmt.Print(c, "\t")
			c *= 3
		} else {
			fmt.Print(a, "\t")
		}
	}
}

func TestLogic10(t *testing.T) {
	n := 12
	a := 1
	b := 2
	c := 3

	for i := 0; i < n; i++ {
		if i%4 == 1 {
			fmt.Print(c, "\t")
			c += 3
		} else if i%4 == 2 {
			fmt.Print(b, "\t")
			b += 2
		} else if i%4 == 3 {
			fmt.Print(a, "\t")
			a += 1
		} else {
			fmt.Print(999, "\t")
		}
	}
}
