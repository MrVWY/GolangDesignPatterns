package main

import "fmt"

//该模式包含6个组件
//proxy: 定义了客户端要调用的Active Object接口。当客户端调用它的方法是，方法调用被转换成method request放入到scheduler的activation queue之中。
//method request: 用来封装方法调用的上下文
//activation queue:待处理的 method request队列
//scheduler:一个独立的线程，管理activation queue，调度方法的执行
//servant:active object的方法执行的具体实现，
//future:当客户端调用方法时，一个future对象会立即返回，允许客户端可以获取返回结果。

type MethodRequest int  //method request

const (
	Incr MethodRequest = iota
	Decr
)

//proxy
type Service struct {
	queue chan MethodRequest  //activation queue
	v int
}

func New(buffer int) *Service  {
	s := &Service{
		queue: make(chan MethodRequest, buffer),
	}
	fmt.Println("This is a New")
	go s.scheduler()
	return s
}

//scheduler
func (s *Service) scheduler() {
	fmt.Println("This is a scheduler")
	for r := range s.queue {
		if r == Incr {
			s.v++
		}else {
			s.v--
		}
	}
}

func (s *Service) Incr() {
	s.queue <- Incr
}

func (s *Service) Decr() {
	s.queue <- Decr
}