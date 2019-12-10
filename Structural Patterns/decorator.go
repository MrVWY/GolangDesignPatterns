package main
//参考博客 ： https://www.cnblogs.com/cinlap/p/11654927.html
import "fmt"

//饮料
type Beverage interface {
	Cost() int
	Description() string
}
//-----------------------------------------------------------------
//茶
type Tea struct {
	name string
	price int
}

func (T *Tea) Cost() int  {
	return T.price
}

func (T *Tea) Description() string  {
	return T.name
}
//Moli
type Moli struct {
	*Tea
}

func NewMoli() *Moli  {
	return &Moli{&Tea{name:  "茉莉", price: 20,}}
}

func (M *Moli) Cost() int {
	return M.price
}

func (M *Moli) Description() string  {
	return M.name
}

//Puer
type Puer struct {
	*Tea
}

func NewPuer() *Puer  {
	return &Puer{&Tea{name:"普洱",price:30}}
}

func (P *Puer) Cost() int {
	return P.price
}

func (P *Puer) Description() string  {
	return P.name
}
//-----------------------------------------------------------------
//配料
type Condiment struct {
	*Tea
	Beverage
	name string
	price int
}

func (C *Condiment) Cost() int  {
	return C.price
}

func (C *Condiment) Description() string  {
	return C.name
}

type Sugar struct {
	*Condiment
}

func NewSugar(B Beverage) *Sugar {
	return &Sugar{&Condiment{Beverage:B,name:"糖",price:10}}
}

func (S *Sugar) Cost() int  {
	return S.Beverage.Cost() + S.price
}

func (S *Sugar) Description() string  {
	return S.Beverage.Description() + S.name
}

type Ice struct {
	*Condiment
}

func NewIce(B Beverage) *Ice {
	return &Ice{&Condiment{Beverage:B,name:"冰",price:10}}
}

func (I *Ice) Cost() int  {
	return I.Beverage.Cost() + I.price
}

func (I *Ice) Description() string  {
	return I.Beverage.Description() + I.name
}

func main() {
	moli := NewMoli()
	puer := NewPuer()

	fmt.Printf("第 %v 杯是 %s 售价 %v 元\n",1,moli.Description(),moli.Cost())
	fmt.Printf("第 %v 杯是 %s 售价 %v 元\n",2,puer.Description(),puer.Cost())

	Moli_sugar := NewSugar(moli)
	Moli_Ice := NewIce(moli)

	fmt.Printf("第 %v 杯是 %s 售价 %v 元\n",1,Moli_sugar.Description(),Moli_sugar.Cost())
	fmt.Printf("第 %v 杯是 %s 售价 %v 元\n",2,Moli_Ice.Description(),Moli_Ice.Cost())

	Puer_sugar := NewSugar(puer)
	Puer_Ice := NewIce(puer)

	fmt.Printf("第 %v 杯是 %s 售价 %v 元\n",1,Puer_sugar.Description(),Puer_sugar.Cost())
	fmt.Printf("第 %v 杯是 %s 售价 %v 元\n",2,Puer_Ice.Description(),Puer_Ice.Cost())
}