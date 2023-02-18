package main

// import (
// 	"fmt"
// )

// import "fmt"


func main() {
	filepath := "./cards/results/my_cards"
	cards := newDeckFromFile(filepath)
	cards.suffle()
	cards.print()

	// err := cards.saveToFile(filepath)
	// if err != nil {
	// 	fmt.Println(err)
	// }

}
