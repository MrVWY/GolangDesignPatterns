package main

import "fmt"

type Property interface {
	GetHPLimit() int
	GetMPLimit() int
	GetAttack_power() int
}

type PeopleBase struct {
	MAX_HP int
	MAX_MP int
	Attack_power int
}

func NewPeopleBase() *PeopleBase {
	return &PeopleBase{MAX_HP: 1500, MAX_MP: 500,Attack_power:50}
}

func (P *PeopleBase) GetHPLimit() int  {
	return P.MAX_HP
}

func (P *PeopleBase) GetMPLimit() int  {
	return P.MAX_MP
}

func (P *PeopleBase) GetAttack_power() int  {
	return P.Attack_power
}

//头盔
type Helmet struct {
	character Property
	HP_ADD int
	MP_ADD int
	Attack_power int
}

func (H *Helmet) GetHPLimit() int  {
	return H.HP_ADD + H.character.GetHPLimit()
}

func (H *Helmet) GetMPLimit() int  {
	return H.MP_ADD + H.character.GetMPLimit()
}

func (H *Helmet) GetAttack_power() int  {
	return H.Attack_power + H.character.GetAttack_power()
}

func NewPeople_Helmet(P Property,HP_ADD, MP_ADD,Attack_power int) *PeopleBase {
	People_get_Helmet:= NewHelmet(P,HP_ADD,MP_ADD,Attack_power)
	PeopleBase := NewHelmet_after_people(People_get_Helmet.GetHPLimit(),People_get_Helmet.GetMPLimit(),People_get_Helmet.GetAttack_power())
	return PeopleBase
}

func NewHelmet(P Property,HP_ADD, MP_ADD, Attack_power int)*Helmet{
	return &Helmet{P,HP_ADD,MP_ADD,Attack_power}
}

func NewHelmet_after_people(HP_ADD int, MP_ADD int, Attack_power int ) *PeopleBase  {
	return &PeopleBase{MAX_HP:HP_ADD, MAX_MP:MP_ADD,Attack_power:Attack_power}
}
//剑
type Sword struct {
	character Property
	HP_ADD int
	MP_ADD int
	Attack_power int
}

func (S *Sword) GetHPLimit() int  {
	return S.HP_ADD + S.character.GetHPLimit()
}

func (S *Sword) GetMPLimit() int  {
	return S.MP_ADD + S.character.GetMPLimit()
}

func (S *Sword) GetAttack_power() int  {
	return S.Attack_power + S.character.GetAttack_power()
}

func NewPeople_Sword(P Property,HP_ADD, MP_ADD, Attack_power int) *PeopleBase {
	People_get_Sword:= NewSword(P,HP_ADD,MP_ADD,Attack_power)
	PeopleBase := NewSword_after_people(People_get_Sword.GetHPLimit(),People_get_Sword.GetMPLimit(),People_get_Sword.GetAttack_power())
	return PeopleBase
}

func NewSword(P Property,HP_ADD, MP_ADD,Attack_power int)*Sword{
	return &Sword{P,HP_ADD,MP_ADD,Attack_power}
}

func NewSword_after_people(HP_ADD int, MP_ADD int, Attack_power int ) *PeopleBase  {
	return &PeopleBase{MAX_HP:HP_ADD, MAX_MP:MP_ADD,Attack_power:Attack_power}
}

//享元工厂
type Element struct {
	Value interface{}
}

func NewElement(Value interface{}) *Element {
	return &Element{Value:Value}
}

type FlyweightFactory struct {
	pool map[string]*Element
}

func (this *FlyweightFactory) GetElement(key string) interface{} {
	return this.pool[key].Value
}

func (this *FlyweightFactory)SetElement(key string,value interface{}){
	ne := NewElement(value)
	this.pool[key]=ne
}

func NewFlyweight()*FlyweightFactory{
	flyweight := FlyweightFactory{}
	flyweight.pool=make(map[string]*Element)
	return &flyweight
}

func main() {
	F := NewFlyweight()
	PeopleBase := NewPeopleBase()
	F.SetElement("People",PeopleBase)  //享元PeopleBase

	//生产俩个人物
	people_1 := F.GetElement("People").(Property)
	people_2 := F.GetElement("People").(Property)
	fmt.Printf("people_1：HP=%d , MP=%d , Attack_power=%d . ||||||| people_2：HP=%d , MP=%d , Attack_power=%d.\n",people_1.GetHPLimit(),people_1.GetMPLimit(),people_1.GetAttack_power(),people_2.GetHPLimit(),people_2.GetMPLimit(),people_2.GetAttack_power())

	//获得头盔
	people_1 = NewPeople_Helmet(people_1, 200, 200,10)
	people_2 = NewPeople_Helmet(people_2,100,100,20)
	fmt.Printf("people_1：HP=%d , MP=%d , Attack_power=%d . ||||||| people_2：HP-%d , MP=%d , Attack_power=%d.\n",people_1.GetHPLimit(),people_1.GetMPLimit(),people_1.GetAttack_power(),people_2.GetHPLimit(),people_2.GetMPLimit(),people_2.GetAttack_power())

	//获得头盔之后，在得到剑
	people_1 = NewPeople_Sword(people_1,100,200,20)
	people_2 = NewPeople_Sword(people_1,200,300,30)
	fmt.Printf("people_1：HP=%d , MP=%d , Attack_power=%d . ||||||| people_2：HP-%d , MP=%d , Attack_power=%d.\n",people_1.GetHPLimit(),people_1.GetMPLimit(),people_1.GetAttack_power(),people_2.GetHPLimit(),people_2.GetMPLimit(),people_2.GetAttack_power())
}