/*打印杨辉三角形*/
package main

import "fmt"

// const line int =10
// var a [line][line] int
// func main() {
// 	for i := 1; i <=line; i++ {
// 		a[i][1]=a[i][i]

// 	}
// }
// const LINES int = 10

// func ShowYanghui() {
// 	var yh [LINES][LINES]int
// 	for i := 0; i < LINES; i++ {
// 		for j := 0; j < i+1; j++ {
// 			if i < 2 { //两行以内三角中的数字都是1
// 				yh[i][j] = 1
// 			} else { //第三行开始，正式计算数值写入数组
// 				if j == 0 || j == i {
// 					yh[i][j] = 1 //所有行的第一列和最后一列都是1
// 				} else {
// 					yh[i][j] = yh[i-1][j-1] + yh[i-1][j] //当前数组元素是上一行的前一个元素加上上一行的当前列元素
// 				}
// 			}
// 			fmt.Printf("%d\t", yh[i][j]) //格式化输出一行
// 		}
// 		fmt.Print("\n") //换行
// 	}
// }

const line int = 10

var a [line][line]int //数组必须被初始化

func yanghuiTriangle() {
	// REENTRY:
	// fmt.Println("please input yanghuitriangle lines:")
	// fmt.Scan(&line)
	// if line > 0 && line <= 10 {

	// } else {
	// 	fmt.Println("the input number is out of the range,please re-entry")
	// goto REENTRY

	for i := 1; i < line; i++ {

		a[i][1] = 1
		a[i][i] = a[i][1]
	}
	for i := 3; i < line; i++ {
		// for k := line-i; k > i; k-- {
		// 	fmt.Println("")
		for j := 2; j < i; j++ {
			a[i][j] = a[i-1][j-1] + a[i-1][j]
			fmt.Print("1")
			fmt.Printf("%d\t", a[i][j]) //格式化输出一行
		}
	}
}

func main() {
	yanghuiTriangle()
}
