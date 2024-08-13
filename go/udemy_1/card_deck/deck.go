package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string


func newDeck() deck {
	cards := deck{}

	suits := []string { "Spades", "Hearts", "Diamonds", "Clubs"}
	values := []string { "Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack","Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card);
	}
}

func deal(d deck, handSize int) (deck, deck) {
	hand := d[:handSize]
	rem := d[handSize:]

	return hand, rem
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) writeToFile(filename string) error {
	str := d.toString()
	return os.WriteFile(filename, []byte(str), 0777)
}

func newDeckFromFile(filename string) deck {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	strs := strings.Split(string(bytes), ",")
	cards := deck(strs)

	return cards
}

func (d deck) shuffle() {
	for i := range d {
		seed := rand.NewSource(time.Now().UnixNano())
		r := rand.New(seed)
		swapPos := r.Intn(len(d) - 1)
		d[i], d[swapPos] = d[swapPos], d[i]
	}
}