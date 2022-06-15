package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	counter := 0
	for _, v := range cb[rank] {
		if v {
			counter++
		}
	}
	return counter
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	counter := 0
	if file > 0 && file < 9 {
		for k := range cb {
			if cb[k][file-1] {
				counter++
			}
		}
	}
	return counter
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	counter := 0
	for range cb {
		for range cb {
			counter++
		}
	}
	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	counter := 0
	for k := range cb {
		for v := range cb[k] {
			if cb[k][v] {
				counter++
			}
		}
	}
	return counter
}
