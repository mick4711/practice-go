package chess

import (
	"errors"
	"strings"
)

func CanKnightAttack(white, black string) (bool, error) {
	// validate entry
	// square must have 2 characters
	if len(white) != 2 || len(black) != 2 {
		return true, errors.New("invalid square")
	}

	// cannot be on the same
	if white == black {
		return true, errors.New("invalid square")
	}
	// parse rank & file
	files := "abcdefgh"
	ranks := "12345678"
	wf := string(white[0])
	wr := string(white[1])
	bf := string(black[0])
	br := string(black[1])
	if !strings.Contains(files, wf) ||
		!strings.Contains(files, bf) ||
		!strings.Contains(ranks, wr) ||
		!strings.Contains(ranks, br) {
		return true, errors.New("invalid square")
	}
	// check for colour match

	return false, nil
}
