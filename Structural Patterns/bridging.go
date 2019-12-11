package main

import "fmt"

type WorkerAction interface {
	Leave()
	Description() string
}

type Worker struct {
	WorkerAction
	name string
}

type GoodWorker struct {
	*Worker
}

func NewGoodWorker() *BadWorker  {
	return &BadWorker{&Worker{name:"Good  Worker",}}
}

func (G *GoodWorker) Leave() {
	fmt.Printf("%s 走了", G.name)
}

func (G *GoodWorker) Description() string {
	return G.name
}

type BadWorker struct {
	*Worker
}

func NewBadWorker() *BadWorker  {
	return &BadWorker{&Worker{name:"Bad  Worker",}}
}

func (B *BadWorker) Leave() {
	fmt.Printf("%s 走了\n", B.name)
}

func (B *BadWorker) Description() string {
	return B.name
}


type Company interface {
	Running()
}

type BigCompant struct {
	WorkerAction
}

func (B *BigCompant) Running() {
	fmt.Printf("%s 正在工作\n",B.WorkerAction.Description())
	B.WorkerAction.Leave()
}

type SmallCompant struct {
	WorkerAction
}

func (S *SmallCompant) Running() {
	fmt.Printf("%s 正在工作\n",S.WorkerAction.Description())
	S.WorkerAction.Leave()
}

func main() {
	GoodWorker := NewGoodWorker()
	BigCompant := BigCompant{WorkerAction:GoodWorker}
	BigCompant.Running()

	BadWorker := NewBadWorker()
	SmallCompant := SmallCompant{WorkerAction:BadWorker}
	SmallCompant.Running()
}