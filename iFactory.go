package main

import "fmt"

type iFurnitureFactory interface {
	makeChair() iChair
	makeSofa() iSofa
}

func getSportsFactory(brand string) (iFurnitureFactory, error) {
	if brand == "adidas" {
		return &kafka{}, nil

	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
