/*猜数字游戏 启动程序时生成一个随机数(target)让用户猜测guess(让用户输入数据) 猜测太大了 提示 太大了 猜测太小了
提示 太小了 想等 提示 猜对了 => 程序退出
最多猜测5次
进阶：猜对了，询问是否继续，继续则重新生成随机数，规则一致
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var number int
	rand.Seed(time.Now().UnixNano()) //设置随机数种子，种子只设置1次
	//随机数范围为1-100
START:
	var randNumber int = rand.Int() % 101
	var answer string
	fmt.Println(randNumber)
	fmt.Println("猜猜数字是什么,请输入：")
	fmt.Scan(&number)
	for i := 4; i > 0; i-- {
		if number != randNumber {
			fmt.Println("抱歉，答错了，还有", i, "次机会,请重新输入：")
			//fmt.Print("test")
			fmt.Scan(&number)
		} else {
			fmt.Println("恭喜你答对了，是否需要继续(y/n)?")
			fmt.Scan(&answer)
			if answer == "y" {
				goto START
			} else {
				fmt.Println("不需要，程序退出。")
				return
			}
		}
	}
	fmt.Println("抱歉，你没有机会了。答案是", randNumber, "程序退出。")
}
