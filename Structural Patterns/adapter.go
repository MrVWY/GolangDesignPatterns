package main

import "fmt"

type Player interface {
	PlayMusic()
}

func Play(P Player)  {
	P.PlayMusic()
}

type GameMusciPlayer struct {
	Src string
}

func (G GameMusciPlayer) PlaySound() {
	fmt.Printf("Play '%s' music",G.Src)
}

type GameSoundAdapter struct {
	GameMusciPlayer GameMusciPlayer
}

func (G GameSoundAdapter) PlayMusic() {
	G.GameMusciPlayer.PlaySound()
}

func main() {
	GameSound := GameMusciPlayer{Src:"The truth that you leave"}
	GameSoundAdapters := GameSoundAdapter{GameMusciPlayer:GameSound}
	Play(GameSoundAdapters)
}