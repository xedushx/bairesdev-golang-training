package main

import "fmt"

type MusicalPlayer interface {
	Greetings()
}

type Trumpeter struct {
	Name string
}

type Violinist struct {
	Name string
}

func (t Trumpeter) Greetings() {
	fmt.Printf("Hi, I'm the Trumpeter %v \n", t.Name)
}

func (v Violinist) Greetings() {
	fmt.Printf("Hi, I'm the Violinist %v \n", v.Name)
}

func main() {

	musicalPlayers := []MusicalPlayer{
		Trumpeter{Name: "Eduardo"},
		Violinist{Name: "Michel"},
		Trumpeter{Name: "Carolina"},
		Violinist{Name: "Maya"},
		Trumpeter{Name: "Rodrigo"},
	}

	for _, musmusicalPlayer := range musicalPlayers {
		musmusicalPlayer.Greetings()
	}

}
