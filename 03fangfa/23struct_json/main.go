package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"` //大写为公有值
	Age  int    `json:"age"`
}

//json
func main() {
	p := person{
		Name: "赵玉",
		Age:  29,
	}
	//json序列化
	v, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("err:%s\n", err)
		return
	}
	fmt.Printf("%v\n", string(v))

	//
	var p2 person
	str := `"name":"赵玉","age":15`
	fmt.Println(str)
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2) //????
}
