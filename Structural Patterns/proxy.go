package main

import "fmt"

type ProxyFunc interface {
	SaleMusic()
}

type MusicOwner struct {
	MusicName string
	MusicNumber int
	MusicPrice string
}

type Proxy struct {
	MusicOwner *MusicOwner
}

func (M *MusicOwner) SaleMusic()  {
	fmt.Printf("%s %d %s",M.MusicName,M.MusicNumber,M.MusicPrice)
}

func main() {
	M := &MusicOwner{
		MusicName:   "Hopeless Case",
		MusicNumber: 2,
		MusicPrice:  "15",
	}
	P := &Proxy{
		MusicOwner:M,
	}
	P.MusicOwner.SaleMusic()
}