package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ids = make([]int, 0)

func GenId() int {
	rand.Seed(time.Now().UnixNano())

	for {
		loop := false
		r := rand.Intn(30) //生成id
		//查重
		for _, v := range ids {
			if r == v {
				loop = true
				fmt.Println("重复：", r)
				break
			}
		}
		if !loop {
			ids = append(ids, r)
			return r
		}
	}
}
