package main

import "fmt"

/*在经典汉诺塔问题中，有 3 根柱子及 N 个不同大小的穿孔圆盘，盘子可以滑入任意一根柱子。一开始，所有盘子自上而下按升序依次套在第一根柱子上
(即每一个盘子只能放在更大的盘子上面)。移动圆盘时受到以下限制:
(1) 每次只能移动一个盘子;
(2) 盘子只能从柱子顶端滑出移到下一根柱子;
(3) 盘子只能叠在比它大的盘子上。
n start end
n-1 start tmp end
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/hanota-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
// n个盘子 start（都在） end（移动到） tmp（借助的） layer （多少个层）
// n start ->temp(借助) -> end
// n-1 start -> end(借助)  -> tmp
// 最大的   start - > end
// n-1  temp->start(借助)->end
func tower(start, end, temp string, layer int) {
	if layer == 1 {
		fmt.Println(start, "->", end)
		return // 无返回值，表示函数结束
	}
	tower(start, temp, end, layer-1)
	fmt.Println(start, "->", end)
	tower(temp, end, start, layer-1)
}
func main() {
	tower("towerA", "towerC","towerB",3)
}
