package deck

import (
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

// represent possible card suites
const (
	SPADE Suit = iota
	DIAMOND
	CLUB
	HEART
	JOKER
)

var suits = [...]Suit{SPADE, DIAMOND, CLUB, HEART}

type Rank uint8

// represent possible card ranks
const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

// Create new deck of cards
func New(options ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	// apply all options to new deck
	for _, opt := range options {
		cards = opt(cards)
	}

	return cards
}

// Customize sort with less function
func Sort(less func(cards []Card) func(i, j int) bool) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// Default sort
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func MyShuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	copy(ret, cards)
	rand.Shuffle(len(ret), func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})
	return ret
}

func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	for i, j := range perm {
		ret[i] = cards[j]
	}

	return ret
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return cardDeckRank(cards[i]) < cardDeckRank(cards[j])
	}
}

func cardDeckRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}
