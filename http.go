package main

import "fmt"
import (
	"github.com/zhshch2002/goreq"
)

func main() {
	res := goreq.Do(goreq.Get("https://www.bilibili.com/"))
	if res.Err != nil {
		fmt.Println(res.Err)
	} else {
		fmt.Println(res.Text)
	}
}
