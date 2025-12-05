package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) (sum int) {
	if row, exist := cb[file]; exist {
		for _, occupied := range row {
			if occupied == true {
				sum++
			}
		}
	}
	return
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) (sum int) {
	if 1 <= rank && rank <= 8 {
		for _, file := range cb {
			if file[rank-1] == true {
				sum++
			}
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) (sum int) {
	for _, row := range cb {
		for _, _ = range row {
			sum += 1
		}
	}
	return
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) (sum int) {
	for f, _ := range cb {
		sum += CountInFile(cb, f)
	}
	return
}
