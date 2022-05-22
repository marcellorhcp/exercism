package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sumOfPlayerCards := ParseCard(card1) + ParseCard(card2)
	parsedDealerCard := ParseCard(dealerCard)
	switch {
	case sumOfPlayerCards == 22:
		return "P"
	case sumOfPlayerCards == 21 && !(parsedDealerCard == 11 || parsedDealerCard == 10):
		return "W"
	case (sumOfPlayerCards >= 12 && sumOfPlayerCards <= 16 && parsedDealerCard >= 7) || sumOfPlayerCards <= 11:
		return "H"
	case (sumOfPlayerCards == 21 && (parsedDealerCard == 11 || parsedDealerCard == 10)) || (sumOfPlayerCards >= 17 && sumOfPlayerCards <= 20) || (sumOfPlayerCards >= 12 && sumOfPlayerCards <= 16):
		return "S"
	default:
		return ""
	}
}
