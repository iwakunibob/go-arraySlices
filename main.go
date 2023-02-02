package main

import (
	"fmt"
	"math/rand"
	"time"
)

type card struct {
	rank  string
	suit  rune
	point int
}

type deck []card

func newDeck() deck {
	var c card
	cards := deck{}
	cardRanks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	cardSuits := []rune{'♦', '♥', '♠', '♣'}
	for i, cr := range cardRanks {
		for j, cs := range cardSuits {
			c = card{
				rank:  cr,
				suit:  cs,
				point: 1 + i*4 + j,
			}
			cards = append(cards, c)
		}
	}
	return cards
}

func main() {
	allCards := newDeck()
	allCards.print()
	allCards.shuffle()
	allCards.print()
}

func (d deck) print() {
	for i, card := range d {
		fmt.Printf("%2s%c =%3d points  ", card.rank, card.suit, card.point)
		if (i % 4) == 3 {
			fmt.Println()
		}
	}
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	fmt.Println("\nCards are shuffled:")
}
