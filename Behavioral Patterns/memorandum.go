package main

type Text struct {
	Value string
}

//文档
func (T *Text) Save(content string)  {
	T.Value = content
}

func (T *Text) GetContent() string  {
	return T.Value
}

func (T *Text) CreateMemorandum() *Memorandum  {
	return &Memorandum{Values:T.Value}
}

//备忘录
type Memorandum struct {
	Values string
}

func (M *Memorandum) SetState(state string)  {
	M.Values = state
}

func (M *Memorandum) GetState() string  {
	return M.Values
}

//负责人
type Caretaker struct {
	Memorandums *Memorandum
}

func (C *Caretaker) GetMemorandum() *Memorandum {
	return C.Memorandums
}

func (C *Caretaker) SetMemorandum(M *Memorandum)  {
	C.Memorandums = M
}

func main() {
	text := &Text{Value:"Hello A!"}
	Caretaker := new(Caretaker)
	Caretaker.SetMemorandum(text.CreateMemorandum())
	text.Save("world")

	//恢复备忘
	text.Save(Caretaker.GetMemorandum().GetState())
}