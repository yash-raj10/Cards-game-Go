package main

import "fmt"

type deck []string

func (d deck) printIt(){
	for i, card := range d{
		fmt.Println(i,card)
	} 
}

func newDeck() deck{
	cards := deck{}
	suitCards := []string{"Spades", "Hearts", "Dimaonds","Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suitCard := range suitCards{
		for _, cardValue := range cardValues{
			cards = append(cards,cardValue+"of"+suitCard )
		}  
	}
	return cards
}
