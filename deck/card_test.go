package deck

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := New()
	if len(deck) != 13*4 {
		t.Error("Invalid number of cards in the deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	deck := New(DefaultSort)
	// should be Ace of Spades
	aceCard := Card{Suit: SPADE, Rank: Ace}
	if deck[0] != aceCard {
		t.Error("Ace of Spades expected, got: ", deck[0])
	}
}

func TestSort(t *testing.T) {
	deck := New(Sort(Less))
	// should be Ace of Spades
	aceCard := Card{Suit: SPADE, Rank: Ace}
	if deck[0] != aceCard {
		t.Error("Ace of Spades expected, got: ", deck[0])
	}
}
