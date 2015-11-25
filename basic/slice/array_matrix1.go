package main
//n*n矩阵沿对角线对调元素数据
import (
	"fmt"
)

func echoMatrix(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j ++ {
			fmt.Printf("%02d ", arr[i][j])
		}
		fmt.Println()
	}
}

//初始化一个n*n矩阵
func createMatrix(matrix_len int) ([][]int) {
	arr := make([][]int, matrix_len)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, matrix_len)
		for j := 0; j < len(arr[i]); j ++ {
			arr[i][j] = i * matrix_len + j + 1
		}
	}
	return arr
}
//按水平中线调换
func left_right(arr [][]int) ([][]int) {
	arr_len := len(arr)
	for i := 0; i < arr_len; i++ {
		sub_len := len(arr[i]);
		for j := 0; j < sub_len; j++ {
			if j > sub_len / 2 {
				break
			}
			arr[i][j], arr[i][sub_len - 1 - j] = arr[i][sub_len - 1 - j], arr[i][j]
		}
	}
	return arr
}

//沿左上角对角线调换
func left_top(arr [][]int) ([][]int) {
	arr_len := len(arr)
	for i := 0; i < arr_len; i++ {
		sub_len := len(arr[i]);
		for j := i; j < sub_len; j++ {
			arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
		}
	}
	return arr
}
//沿右上角对角线调换
func right_top(arr [][]int) ([][]int) {
	arr_len := len(arr)
	for i := 0; i < arr_len; i++ {
		sub_len := len(arr[i]);
		for j := 0; j < sub_len; j++ {
			if j + i + 1 == arr_len {
				break
			}
			arr[i][j], arr[arr_len - 1 - j][arr_len - 1 - i] = arr[arr_len - 1 - j][arr_len - 1 - i], arr[i][j]
		}
	}
	return arr
}


func main() {
	matrix_len := 5
	arr1 := createMatrix(matrix_len)
	echoMatrix(arr1)
	fmt.Println()
	echoMatrix(left_top(arr1))

	fmt.Println("-----------------------------")
	arr2 := createMatrix(matrix_len)
	echoMatrix(right_top(arr2))
	fmt.Println("----------------")
	//下面的操作也可以实现right_top的效果，不过效率要差很多了
	arr3 := createMatrix(matrix_len)
	echoMatrix(arr3)
	fmt.Println("----------")
	left_right(arr3)
	echoMatrix(arr3)
	fmt.Println("----------")
	left_top(arr3)
	echoMatrix(arr3)
	fmt.Println("----------")
	left_right(arr3)
	echoMatrix(arr3)
	fmt.Println("----------")
}