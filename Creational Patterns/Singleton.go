package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	once sync.Once
	instance *Fish
)

type Fish struct {
	str string
}

type Instance interface {
	Action()
}

func (fish *Fish) Action()  {
	fmt.Printf("I am fish %s !",fish.str)
}

func GetInstance(star string) *Fish {
	once.Do(func() {
		instance = &Fish{str:star}
	})
	return instance
}

func main() {
	for i:=0; i<10 ;i++  {
		go func() {
			s := GetInstance(strconv.Itoa(i))  //获取不到就一直阻塞，而s是第一次获取得值
			s.Action()
		}()
	}
	time.Sleep(1e5)
}