package main

import (
	"app/pigeon"
	"fmt"
)

func main() {
	//p := pigeon.Pigeon{"Tweety", 10} - This will not  compile
	p := pigeon.Pigeon{Name: "Tweety"}
	p.SetFeatherLength(10)

	fmt.Println(p.Name)
	fmt.Println(p.GetFeatherLength())
	//fmt.Println(p.featherLength) - This will not  compile

}
