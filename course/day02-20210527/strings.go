package main

import (
	"fmt"
	"strings"
)

func main() {
	// 比较 1,0,-1
	res := strings.Compare("s", "d") //1
	fmt.Println(res)
	// 包含 true false
	res1 := strings.Contains("dfdfd", "kk")
	fmt.Println(res1)
	// 某个字符串出现多少次
	res2 := strings.Count("sdsdsd", "d")
	fmt.Println(res2)

	// 不区分大小写
	res4 := strings.EqualFold("a", "A")
	fmt.Println(res4)

	// 空白符
	res5 := strings.Fields("dsdsd sdsdsd\ndfdf")
	fmt.Println(res5)

	// 分割·
	res6 := strings.Split("sdsdsd", "s")
	fmt.Println(res6)

	// 开头结尾
	res7 := strings.HasPrefix("sdsdsd", "sd")
	fmt.Println(res7)
	res8 := strings.HasSuffix("sdsdsd", "sd")
	fmt.Println(res8)
	// 字符串出现索引位置， 没有-1
	fmt.Println(strings.Index("sdsd", "ds"))
	fmt.Println(strings.LastIndex("sdsd", "d"))
	fmt.Println(strings.Join([]string{"a", "b"}, "|"))
	fmt.Println(strings.SplitN("sdsdsds", "d", 2))
	// 重复
	fmt.Println(strings.Repeat("*", 5))
	// 替换
	// 全部替换
	fmt.Println(strings.ReplaceAll("sdsddssd", "s", "d"))
	fmt.Println(strings.Replace("sdsddssd", "s", "d", -1))
	// 仅替换第一个
	fmt.Println(strings.Replace("sdsddssd", "s", "d", 1))

	// 首字母大写
	fmt.Println(strings.Title("sddsd"))
	// 全部小写
	fmt.Println(strings.ToLower("FEDFD"))
	// 全部大写
	fmt.Println(strings.ToUpper("sddsd"))
	// 替除
	fmt.Println(strings.Trim("sdsdsd", "s"))
	fmt.Println(strings.TrimSpace("      d     sds         "))
	fmt.Println(strings.TrimLeft("dsdsds", "ds")) //把每个字符匹配
	fmt.Println(strings.TrimRight("dsdsds", "ds"))
	fmt.Println(strings.TrimPrefix("dsdsds", "ds")) // 整体
	fmt.Println(strings.TrimSuffix("dsdsds", "ds"))

}
