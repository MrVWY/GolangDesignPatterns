package main

import "fmt"

//玩家行为
type PlayerAction interface {
	Win()
	Lose()
	AddGame() *PlayDirector
	Die() *PlayDirector
	ChangeTeam(color string) *PlayDirector
}

var director = &PlayDirector{
	playlist :make(map[string]map[string]string),
	removelist: make(map[string]map[string]string),
}
//玩家
type Play struct {
	name string
	teamcolor string
	state string
}

func (P *Play) Win() {
	fmt.Printf("Play:%s team:%s win this game",P.name,P.teamcolor)
}

func (P *Play) Lose()  {
	fmt.Printf("Play:%s team:%s lose this game",P.name,P.teamcolor)
}

func (P *Play) AddGame() *PlayDirector {
	Ps := director.AddPlay(P)
	return Ps
}

func (P *Play) Die() *PlayDirector {
	Ps := director.Die(P)
	return Ps
}

func (P *Play) ChangeTeam(color string) *PlayDirector {
	Ps := director.ChangeTeam(P,color)
	return Ps
}

type PlayDirector struct {
	playlist map[string]map[string]string
	removelist map[string]map[string]string
}

func (P *PlayDirector) AddPlay(play *Play) *PlayDirector {
	A := make(map[string]string)
	A["TeamColor"] = play.teamcolor
	A["State"] = "live"
	P.playlist[play.name] = A
	return P
}

func (P *PlayDirector) Remove(play *Play,A map[string]string) *PlayDirector {
	P.removelist[play.name] = A
	delete(P.playlist,play.name)
	return P
}

func (P *PlayDirector) Die(play *Play) *PlayDirector {
	A := make(map[string]string)
	_, ok := P.playlist[play.name]
	if ok {
		A["TeamColor"] = play.teamcolor
		A["State"] = "dead"
		P.Remove(play,A)
		delete(P.playlist,play.name)
	}else {
		fmt.Println("This play have been dead!!")
	}
	return P
}

func (P *PlayDirector) ChangeTeam(play *Play,changeTeamColor string) *PlayDirector {
	A := make(map[string]string)
	_, ok := P.playlist[play.name]
	if ok {
		A["TeamColor"] = changeTeamColor
		A["State"] = "live"
		P.playlist[play.name] = A
		return P
	}else {
		A["TeamColor"] = changeTeamColor
		A["State"] = "dead"
		P.removelist[play.name] = A
		return P
	}
}

func main() {
	Play1 := &Play{name:"A", teamcolor: "red", state:"live"}
	Play2 := &Play{name:"B", teamcolor: "blue", state:"live"}
	Play3 := &Play{name:"C", teamcolor: "blue", state:"live"}
	Play2.AddGame()
	Play1.AddGame()
	b := Play3.AddGame()
	fmt.Println(b)
	Play3.Die()
	a := Play2.ChangeTeam("red")
	fmt.Println(a)
}