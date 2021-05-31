package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
猜数字
	生成随机数
	提示用户在控制台输入猜测数字
	比较，当用户输入较大，提示太大了
		当用户输入太小，提示太小了
		当用户输入正确，提示他第几次输入对了
	用户最多猜五次，如果5次内都没有猜正确，提示太笨了，游戏结束

*/
func main() {
	// 设置种子
	rand.Seed(time.Now().Unix())
	//num := rand.Int()
	//num := rand.Int() % 101 生成0-100
	num := rand.Intn(100)
	fmt.Println(num)
	for i := 1; i <= 5; i++ {
		inputNum := 0
		fmt.Println("请输入数字")
		fmt.Scan(&inputNum)
		if inputNum == num {
			fmt.Printf("恭喜你，猜对了%d\n", i)
			goto End
		}
		//switch {
		//case num == inputNum:
		//	fmt.Printf("恭喜你，猜对了%d\n", i)
		//	goto End
		//case num > inputNum:
		//	fmt.Printf("猜小了")
		//case num < inputNum:
		//	fmt.Printf("猜大了")
		//
		//}
	}
	fmt.Println("太笨了")

End:
	fmt.Print()
}
