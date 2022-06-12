package func_array

/*
1 2 3 => baris 0, kolom 0, 1, 2
4 5 6 => baris 1, kolom 0, 1, 2
7 8 9 => baris 2, kolom 0, 1, 2
*/

//[][]string artinya membuat array 2 dimensi dengan tipe data string

func NewStringArray(baris int, kolom int) [][]string {
	result := make([][]string, baris) //line berisi tentang membuat slice sepanjang "baris"
	for i := range result {
		result[i] = make([]string, kolom)
	}
	return result
}

func NewNumberArray(baris int, kolom int) [][]int32 {
	result := make([][]int32, baris)
	for i := range result {
		result[i] = make([]int32, kolom)
	}
	return result
}
