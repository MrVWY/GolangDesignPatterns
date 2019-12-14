package main

import "fmt"

type CashSuper interface {
	AccpetMoney(float64) float64
}

type CashNormal struct {}

func NewCashNormal() *CashNormal {
	instance := new(CashNormal)
	return instance
}

func (C *CashNormal) AccpetMoney(money float64) float64 {
	return money
}

type CashRebate struct {
	Rebate float64
}

func NewCashRebate(rebate float64) *CashRebate {
	instance := new(CashRebate)
	instance.Rebate = rebate
	return instance
}

func (C *CashRebate) AccpetMoney(money float64) float64 {
	return money * C.Rebate
}

type CashReturn struct {
	MoneyCondition float64
	MoneyReturn float64
}

func NewCashReturn(moneyCondition float64, moneyReturn float64) *CashReturn {
	instance := new(CashReturn)
	instance.MoneyCondition = moneyCondition
	instance.MoneyReturn = moneyReturn
	return instance
}

func (C *CashReturn) AccpetMoney(money float64) float64 {
	if money > C.MoneyCondition{
		n := int(money / C.MoneyCondition)
		return money - float64(n) * C.MoneyReturn
	}
	return money
}

type CashContext struct {
	CashSuper CashSuper
}

func NewCashConText(CashSuper CashSuper) *CashContext  {
	instance := new(CashContext)
	instance.CashSuper = CashSuper
	return instance
}

func (C *CashContext) GetMoney(money float64) float64 {
	return C.CashSuper.AccpetMoney(money)
}

func main() {
	C := NewCashConText(&CashReturn{
		MoneyCondition: 100,
		MoneyReturn:    20,
	})
	a := C.CashSuper.AccpetMoney(100)
	fmt.Println(a)
	b := C.GetMoney(100)
	fmt.Println(b)
}