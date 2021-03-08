package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//中文出现次数
	str := "我每天都在学习,Visual Studio"
	var count int
	for _, c := range str {
		if unicode.Is(unicode.Han, c) {
			count++
		}

	}
	fmt.Println(count)
	fmt.Println(str)

	//统计单词出现次数
	s2 := "how do you do ? my name is xiaoming from sichuan , is Southerner , like Eat spicy"

	sList := strings.Split(s2, " ") //分割字符串
	m1 := make(map[string]int, 50)  //建立map
	for _, w := range sList {       //循环遍历分割后
		if _, ok := m1[w]; !ok { //判断map中是否存在此key
			m1[w] = 1 //不存在计数1次
		} else {
			m1[w]++ //存在则加1
		}
	}
	for k, v := range m1 { //打印出现次数
		fmt.Printf("%s:%d \t\n", k, v)
	}
	//回文判断
	ss := "上海自来水来自海上"
	r := make([]rune, 0, len(ss))
	for _, v := range ss {
		r = append(r, v)
	}
	fmt.Println("[]rune", r)
	for i := 0; i < len(r)/2; i++ {
		if r[i] == r[len(r)-1-i] {
			fmt.Println("是回文！")
			return
		}
		fmt.Println("不是回文！")

	}
}
