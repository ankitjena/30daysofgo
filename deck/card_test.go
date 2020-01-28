package deck

import "fmt"

import "testing"

import "math/rand"

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})

	// Output:
	// Ace of Hearts
}

func TestNew(t *testing.T) {
	cards := New()

	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in the new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of spades as first card. Received: ", cards[0])
	}
}

func TestShuffle(t *testing.T) {
	//make shuffleRand deterministic so we know the output
	shuffleRand = rand.New(rand.NewSource(0))

	orig := New()
	first := orig[40]
	second := orig[35]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("Expected the first card to be %s, received %s", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected the first card to be %s, received %s", second, cards[1])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Expected 3 Jokers received: ", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}

	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Errorf("Expected all twos and threes to be filtered out")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 13 * 4 * 3 {
		t.Errorf("Expected %d cards, received %d cards", 13 * 4 * 3, len(cards))
	}
}