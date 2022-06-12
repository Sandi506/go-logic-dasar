package main

import (
	"fmt"
	"logic_dasar/func_array"
)

func main() {
	fmt.Println("Soal 01")
	Logic03Soal01(9)

	fmt.Println("Soal 02")
	Logic03Soal02(9)

	fmt.Println("Soal array")
	Logic03Soal03WithArray(9)

	fmt.Println("Soal 03")
	Logic03Soal03(9)

	fmt.Println("Soal 04")
	Logic03Soal04(9)

	fmt.Println("Soal 05")
	Logic03Soal05(9)

}

/*func Logic03Soal01(n int) {
	a := 3
	for i := 0; i <= n; i++ {
		//Looping kolom
		for j := 0; j <= n; j++ {
			//Kondisi untuk kotak kosong
			if i >= j && i+j >= n-1 {
				fmt.Print(a, "\t")

			} else {
				fmt.Print(" ", "\t")
				//a -= 3
			}
			if i <= j && i+j <= n-1 {
				fmt.Print(a, "\t")

			}
		}
		fmt.Println("\n")
		a += 3
	}
}

func Logic03Soal02(n int) {
	a := 3
	for i := 0; i <= n; i++ {
		//Looping kolom
		for j := 0; j <= n; j++ {
			//Kondisi untuk kotak kosong
			if i >= j && i+j >= n-1 {
				fmt.Print(a, "\t")
			} else {
				//if i >= j && i+j >= n-1 {
				fmt.Print(" ", "\t")
				//if i >= j && i+j >= n-1 {
			}
		}
		fmt.Println("\n")

	}
}*/

func Logic03Soal01(n int) {
	a := 3
	nt := n / 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i >= j && i+j <= n-1 {
				fmt.Print(a, "\t")
			} else if i <= j && i+j >= n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Print("\n")
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	}
}

func Logic03Soal02(n int) {
	a := 3
	nt := n / 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i <= j && i+j <= n-1 {
				fmt.Print(a, "\t")
			} else if i >= j && i+j >= n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Print("\n")
		if i < nt {
			a += 3
		} else {
			a -= 3
		}

	}
}

func Logic03Soal03(n int) {
	a := 3

	for i := 0; i < n; i++ {
		nt := n / 2
		for j := 0; j < n; j++ {
			if i < j && j <= nt {
				fmt.Print(a, "\t")
			} else if i > j && j >= nt {
				fmt.Print(a, "\t")
			} else if i+j < n-1 && i >= nt {
				fmt.Print(a, "\t")
			} else if i+j > n-1 && i <= nt {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		println("\n")
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	}
}

func Logic03Soal03WithArray(n int) {
	array := func_array.NewNumberArray(n, n) //line ini untuk membuat array

	//mengisi array
	a := 3 // => variabel ini tipe nya int
	nt := n / 2
	for i := 0; i < len(array); i++ { // => mulai dari line ini, me-looping baris
		for j := 0; j < len(array[i]); j++ { // => mulai dari line ini, me-looping kolom
			if i < j && j <= nt {
				array[i][j] = int32(a) // => variabel a dikonvert ke int32
			} else if i > j && j >= nt {
				array[i][j] = int32(a) // => variabel a dikonvert ke int32
			} else if i+j < n-1 && i >= nt {
				array[i][j] = int32(a) // => variabel a dikonvert ke int32
			} else if i+j > n-1 && i <= nt {
				array[i][j] = int32(a) // => variabel a dikonvert ke int32
			}
		} // => pada line ini looping kolom selesai
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	} // => pada line ini looping baris selesai
	func_array.PrintNumberArrayZero(array)
}

func Logic03Soal04(n int) {
	a := 3
	nt := n / 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%4 == 0 {
				fmt.Print(a, "\t")
			} else if i%4 == 1 && j == n-1 {
				fmt.Print(a, "\t")
			} else if i%4 == 2 {
				fmt.Print(a, "\t")
			} else if i%4 == 3 && j == 0 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Println("\n")
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	}
}

func Logic03Soal05(n int) {
	a := 3
	nt := n / 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%4 == 0 {
				fmt.Print(a, "\t")
			} else if i%4 == 1 && j == 0 {
				fmt.Print(a, "\t")
			} else if i%4 == 2 {
				fmt.Print(a, "\t")
			} else if i%4 == 3 && j == n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Println("\n")
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	}
}
