package logic04

import "logic_dasar/func_array"

func Logic04Soal04(n int) {
	array := func_array.NewNumberArray(n, n) //=> ini untuk membuat array

	a := 1

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if i%4 == 0 { // => line ini untuk membuat deret dari kiri ke kanan
				array[i][n-1-j] = int32(a)
				a += 3
			} else if i%4 == 1 && j == 0 { // => untuk membuat berisi pada kolom 0
				array[i][j] = int32(a)
				a += 3
			} else if i%4 == 2 { // => untuk membuat baris 3 dan baris 6 bergerak mundur
				array[i][j] = int32(a)
				a += 3
			} else if i%4 == 3 && j == n-1 { // => untuk membuat kotak kosong pada baris 1 dan berisi pada kolom 8
				array[i][j] = int32(a)
				a += 3
			}
		}
	}
	func_array.PrintNumberArrayZero(array) // => ini untuk print array
}
