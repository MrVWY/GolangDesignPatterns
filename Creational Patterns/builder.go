package main

import "fmt"

type Builder interface {
	AnimalName(name string) Builder
	Animalcolor(color string) Builder
	AnimalType(Type string) Builder
	Build() Animal
}

type Animal struct {
	Name string
	Color string
	Type string
}

func (this *Animal) Drive() error {
	fmt.Printf(" %s %s %s Animal is running on the road!\n", this.Name,this.Color, this.Type)
	return nil
}


type ConcrateBuild struct {
	CrateAnimal Animal
}

type Directos struct {
	Builder Builder
}

func (ConcrateBuild *ConcrateBuild) AnimalName(name string) Builder  {
	ConcrateBuild.CrateAnimal.Name = name
	return ConcrateBuild
}


func (ConcrateBuild *ConcrateBuild) Animalcolor(color string) Builder   {
	ConcrateBuild.CrateAnimal.Color = color
	return ConcrateBuild
}

func (ConcrateBuild *ConcrateBuild) AnimalType(Type string) Builder  {
	ConcrateBuild.CrateAnimal.Type = Type
	return ConcrateBuild
}

func (ConcrateBuild *ConcrateBuild) Build() Animal  {
	return ConcrateBuild.CrateAnimal
}

func main() {
	d := Directos{&ConcrateBuild{}}
	s := d.Builder.Animalcolor("red").AnimalName("A").AnimalType("cat").Build()
	s.Drive()

}