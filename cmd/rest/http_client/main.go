package main

import (
	"fmt"
	gofakeit "github.com/brianvoe/gofakeit/v7"
)

type Beer struct {
	Name    string `fake:"{firstname}"`
	Alcohol string `fake:"{}"`
}

func main() {
	fmt.Println(gofakeit.Name())
	fmt.Println(gofakeit.BeerName())

	//BeerAlcohol() string
	//BeerBlg() string
	//BeerHop() string
	//BeerIbu() string
	//BeerMalt() string
	//BeerName() string
	//BeerStyle() string
	//BeerYeast() string
}
