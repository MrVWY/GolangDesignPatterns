package main

import (
	"fmt"
	"strings"
)

type Handler interface {
	Handle(string)
	Next(Handler,string)
}

type Sensitive_Words struct {
	handler Handler
}

func (S *Sensitive_Words) Handle(content string)  {
	fmt.Println("敏感词过滤!")
	NewContent := strings.Replace(content,"敏感词","***",1)
	S.Next(S.handler,NewContent)
}

func (S *Sensitive_Words) Next(handler Handler, content string)  {
	if S.handler != nil{
		 S.handler.Handle(content)
	}
}

type Advertisement struct {
	handler Handler
}

func (A *Advertisement) Handle(content string)  {
	fmt.Println("广告词过滤!")
	NewContent := strings.Replace(content,"广告词","$$$",1)
	A.Next(A.handler,NewContent)
}

func (A *Advertisement) Next(handler Handler, content string)  {
	if A.handler != nil{
		 A.handler.Handle(content)
	}
}

type Incorrect struct {
	handler Handler
}

func (I *Incorrect) Handle(content string)  {
	fmt.Println("不正确词过滤!")
	NewContent := strings.Replace(content,"不正确词","&&&",1)
	fmt.Println(NewContent)
	I.Next(I.handler,NewContent)
}

func (I *Incorrect) Next(handler Handler, content string)  {
	if I.handler != nil{
		I.handler.Handle(content)
	}
}

func main() {
	sensitive_Words := &Sensitive_Words{}
	advertisement := &Advertisement{}
	incorrect := &Incorrect{}

	sensitive_Words.handler = advertisement
	advertisement.handler = incorrect

	sensitive_Words.Handle("我是正常词，我是敏感词，我是广告词，我是不正确词")
}