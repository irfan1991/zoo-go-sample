package main

import (
	"fmt"

	"gitlab.com/ari-salt/zoo/animal"
	"gitlab.com/ari-salt/zoo/animal/cat"
	"gitlab.com/ari-salt/zoo/animal/cat/tiger"
)

func main() {
	fmt.Println("Zoo!")

	c1 := cat.New("Tom", "cat")
	c2 := tiger.New("Long", "tiger")

	var ca1 animal.Animal = c1
	var ci1 cat.CatInterface = c1
	var ca2 animal.Animal = c2
	var ci2 cat.CatInterface = c2

	fmt.Printf("%s is a %s.\n%s\n", ca1.Name(), ca1.Kind(), ci1.Rawr())
	fmt.Printf("%s is a %s.\n%s\n", ca2.Name(), ca2.Kind(), ci2.Rawr())
}
