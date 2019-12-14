package main

import "fmt"

type ActionState interface {
	View()
	Comment()
	Post()
}
//------------------------------------------------------------------------------
type Account struct {
	State ActionState
	HealthValue int
}

func NewAccount(health int) *Account  {
	account := &Account{HealthValue:health}
	account.ChangeState()
	return account
}

func (a *Account) View() {
	a.State.View()
}

func (a *Account) Comment()  {
	a.State.Comment()
}

func (a *Account) Post()  {
	a.State.Post()
}

//----------------------------------------------------------------------------------
type CloseState struct {}

func (C *CloseState) View() {
	fmt.Println("账号被封，无法看帖")
}

func (C *CloseState) Comment()  {
	fmt.Println("抱歉，你的健康值小于-10，不能评论")
}

func (C *CloseState) Post()  {
	fmt.Println("抱歉，你的健康值小于0，不能发帖")
}

//------------------------------------------------------------------------------------
type RestrictedState  struct {}

func (R *RestrictedState) View() {
	fmt.Println("正常看帖")
}

func (R *RestrictedState) Comment()  {
	fmt.Println("正常评论")
}

func (R *RestrictedState) Post()  {
	fmt.Println("抱歉，你的健康值小于0，不能发帖")
}

//-------------------------------------------------------------------------------------
type NormalState   struct {}

func (N *NormalState ) View() {
	fmt.Println("正常看帖")
}

func (N *NormalState ) Comment()  {
	fmt.Println("正常评论")
}

func (N *NormalState ) Post()  {
	fmt.Println("正常发帖")
}

func (a *Account) ChangeState()  {
	if a.HealthValue <= -10 {
		a.State = &CloseState{}
	} else if a.HealthValue > -10 && a.HealthValue <=0 {
		a.State = &RestrictedState{}
	} else if a.HealthValue > 0 {
		a.State = &NormalState{}
	}
}

func (a *Account) SetHealth(value int) {
  	a.HealthValue = value
 	a.ChangeState()
}

func main() {
	account := NewAccount(10)
	account.Post()

	account2 := NewAccount(-1)
	account2.Post()

	account3 := NewAccount(-11)
	account3.Post()
}