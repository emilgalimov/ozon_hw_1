package chess

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type coordinates struct {
	x rune
	y int
}

func newCoordinates(x rune, y int) (*coordinates, error) {

	if x != 'a' && x != 'b' && x != 'c' && x != 'd' && x != 'e' && x != 'f' && x != 'g' && x != 'h' {
		return nil, errors.New("invalid coordinate")
	}

	if y < 1 || y > 8 {
		return nil, errors.New("invalid coordinate")
	}

	return &coordinates{
		x: x,
		y: y,
	}, nil
}

func (c *coordinates) equal(compare coordinates) bool {
	if c.x == compare.x && c.y == compare.y {
		return true
	}

	return false
}

func (c *coordinates) move(x int, y int) (*coordinates, error) {
	return newCoordinates(rune(int(c.x)+x), c.y+y)
}

func CanKnightAttack(white, black string) (bool, error) {
	whiteCoo, err := stringToCoord(white)
	if err != nil {
		return false, err
	}

	blackCoo, err := stringToCoord(black)
	if err != nil {
		return false, err
	}

	if whiteCoo.equal(*blackCoo) {
		return false, errors.New("invalid coordinates")
	}

	moves := getAllPossibleMoves(whiteCoo)
	for _, move := range moves {
		if move.equal(*blackCoo) {
			return true, nil
		}
	}
	return false, nil
}

func stringToCoord(position string) (coords *coordinates, err error) {
	if len(position) != 2 {
		return nil, errors.New("invalid coordinate")
	}
	x := []rune(strings.ToLower(position))[0]

	yRune := []rune(position)[1]
	if false == unicode.IsDigit(yRune) {
		return nil, errors.New("invalid coordinate")
	}

	yString := string(yRune)
	y, _ := strconv.Atoi(yString)

	return newCoordinates(x, y)
}

func getAllPossibleMoves(coord *coordinates) []*coordinates {
	var moves []*coordinates

	for x := -2; x <= 2; x++ {
		for y := -2; y <= 2; y++ {
			if abs(x)+abs(y) == 3 {
				move, err := coord.move(x, y)
				if err == nil {
					moves = append(moves, move)
				}
			}
		}
	}
	return moves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
