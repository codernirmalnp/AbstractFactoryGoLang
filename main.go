package main

import "fmt"

func main() {
	kafFactory, _ := getSportsFactory("adidas")

	productA := kafFactory.makeChair()
	makeSofa := kafFactory.makeSofa()

	printChairDetails(productA)
	printSoftDetails(makeSofa)
}

func printChairDetails(s iChair) {
	fmt.Printf("Logo: %s", s.getChair())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printSoftDetails(s iSofa) {
	fmt.Printf("Logo: %s", s.getSofa())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
