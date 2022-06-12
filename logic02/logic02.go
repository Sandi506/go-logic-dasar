package main

import (
	"fmt"
)

func main() {

	fmt.Println("Logic02soal01")
	Logic02soal01(9)

	fmt.Println("Logic02soal02")
	Logic02soal02(9)

	fmt.Println("Logic02soal03")
	Logic02soal03(9)

	fmt.Println("Logic02soal04")
	Logic02soal04(9)

	fmt.Println("Logic02soal05")
	Logic02soal05(9)

	fmt.Println("Logic02soal06")
	Logic02soal06(9)

	fmt.Println("Logic02soal07")
	Logic02soal07(9)

	fmt.Println("Logic02soal08")
	Logic02soal08(9)

	fmt.Println("Logic02soal09")
	Logic02soal09(9)

	fmt.Println("Logic02soal10")
	Logic02soal10(9)

}

func Logic02soal01(n int) {
	// 3 6 9 12 15 18 21 24 27
	//n := 9 adalan banyak nya jumlah nilai atau baris
	a := 3 // nilai variable
	// i sebagai loop baris
	for i := 0; i < n; i++ {
		// j sebagai loop kolom
		for j := 0; j < n; j++ {
			//if i == j {
			fmt.Print(a, "\t")
		}
		// kolom selesai dan melanjutkan ke baris
		fmt.Println("\n")
		a += 3 // tambahan nilai utk nilai a

	}

}

func Logic02soal02(n int) {
	// 3 6 9 12 15 18 21 24 27

	for i := 0; i < n; i++ { // loop baris ke pinggir
		a := 3
		for j := 0; j < n; j++ { // loop kolom ke bawah
			fmt.Print(a, "\t")
			a += 3
		}
		fmt.Println("\n")

	}
}

func Logic02soal03(n int) {
	for i := 0; i < n; i++ {
		a := 3 * n
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
			a -= 3

		}
		fmt.Println("\n")

	}
}

func Logic02soal04(n int) {
	a := 3 * n // nilai variable
	// i sebagai loop baris
	for i := 0; i < n; i++ {
		// j sebagai loop kolom
		for j := 0; j < n; j++ {
			//if i == j {
			fmt.Print(a, "\t")
		}
		// kolom selesai dan melanjutkan ke baris
		fmt.Println("\n")
		a -= 3 // tambahan nilai utk nilai a

	}

}

func Logic02soal05(n int) {
	nilaitengah := n / 2
	y := 3
	for i := 0; i < n; i++ { // looping baris "bawah / line"
		for j := 0; j < n; j++ { // looping kolom "pinggir"
			fmt.Print(y, "\t")
		}
		if i < nilaitengah {
			fmt.Print(y, "\t")
			y += 3
		} else {
			fmt.Print(y, "\t")
			y -= 3
		}
		fmt.Println("\n")
	}
}

func Logic02soal06(n int) {
	tengah := n / 2

	for i := 0; i < n; i++ {
		a := 3
		for j := 0; j < n; j++ {
			if j < tengah {
				fmt.Print(a, "\t")
				a += 3
			} else {
				fmt.Print(a, "\t")
				a -= 3
			}

		}
		fmt.Println("\n")
	}
}

func Logic02soal07(n int) {
	a := 3 // nilai variable
	// i sebagai loop baris
	for i := 0; i < n; i++ {
		// j sebagai loop kolom
		for j := 0; j < n; j++ {
			if i < j {
				fmt.Print(" ", "\t")
			} else {
				fmt.Print(a, "\t")

			}

		}
		// kolom selesai dan melanjutkan ke baris
		fmt.Println("\n")
		a += 3 // tambahan nilai utk nilai a

	}

}

func Logic02soal08(n int) {
	a := 3 // nilai variable
	// i sebagai loop baris
	for i := 0; i < n; i++ {
		// j sebagai loop kolom
		for j := 0; j < n; j++ {
			if i > j {
				fmt.Print(" ", "\t")
			} else {
				fmt.Print(a, "\t")
			}
		}
		// kolom selesai dan melanjutkan ke baris
		fmt.Println("\n")
		a += 3 // tambahan nilai utk nilai a
	}
}

func Logic02soal09(n int) {
	// nilai variable
	// i sebagai loop baris

	for i := 0; i < n; i++ {
		a := 0 // nilai 27
		// j sebagai loop kolom
		for j := 0; j < n; j++ {
			a += 3
			if i+j == n-1 {
				fmt.Print(a, "\t")

			} else if i+j >= n-1 {
				fmt.Print(a, "\t")

			} else {
				fmt.Print(" ", "\t")
			}
		}
		// kolom selesai dan melanjutkan ke baris
		fmt.Println("\n")
		//a += 3 // tambahan nilai utk nilai a
	}
}

func Logic02soal10(n int) {
	for i := 0; i < n; i++ {
		a := 0 // nilai 27
		// j sebagai loop kolom
		for j := 0; j < n; j++ {
			a += 3
			if i+j == n-1 {
				fmt.Print(a, "\t")

			} else if i+j <= n-1 {
				fmt.Print(a, "\t")

			} else {
				fmt.Print(" ", "\t")
			}
		}
		// kolom selesai dan melanjutkan ke baris
		fmt.Println("\n")
		//a += 3 // tambahan nilai utk nilai a
	}
}

/**
  for i := 0; i < n; i++ { // loop baris ke pinggir
  		a := 3
  		for j := 0; j < n; j++ { // loop kolom ke bawah
  			fmt.Print(a, "\t")
  			a += 3
  		}
  		fmt.Println("\n")
*/
