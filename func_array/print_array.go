package func_array

import "fmt"

func PrintNumberArray(array [][]int32) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			fmt.Print(array[i][j], "\t")
		}
		fmt.Println("\n")
	}
}

func PrintStringArray(array [][]string) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			fmt.Print(array[i][j], "\t")
		}
		fmt.Println("\n")
	}
}

func PrintNumberArrayZero(array [][]int32) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if array[i][j] > 0 {
				fmt.Print(array[i][j], "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Println("\n")
	}
}
