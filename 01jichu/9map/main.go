package main

import "fmt"

func main() {
	mymap := make(map[string]int, 9) //此时候map为null需进行初始化
	//上面的[]中为键，后为值，数字为长度。
	//类型可改！！
	mymap["张三"] = 34 //使用=进行赋值
	mymap["李四"] = 35
	mymap["王五"] = 54
	mymap["赵六"] = 27
	mymap["夏明"] = 19 //此方法的map无序？
	fmt.Println(mymap)
	map1 := make(map[int]string, 8)
	map1[1] = "张三"
	map1[2] = "李四"
	map1[3] = "王五"
	map1[4] = "赵六"
	fmt.Println(map1)

	k, ok := map1[7] //如果该key没有，则拿到对应值类型的零值
	if !ok {
		fmt.Println("不包含此key")
	} else {
		fmt.Println(k)
	}
	//map的遍历
	for k, v := range map1 {
		fmt.Printf("%d:%s\n", k, v)
	}
	for k := range map1 { //取key
		fmt.Println(k)
	}
	for _, v := range map1 { //取值
		fmt.Println(v)
	}
	//删除
	delete(mymap, "夏明")
	delete(mymap, "小明") //当key不存在时不执行任何操作，因不存在对应值
	fmt.Println(mymap)
}
