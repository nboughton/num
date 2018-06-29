package num

import (
	"fmt"
)

// Matrix is a slice of slices of Int
type Matrix []Set

func (s Matrix) Len() int           { return len(s) }
func (s Matrix) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Matrix) Less(i, j int) bool { return len(s[i]) < len(s[j]) }

// Find returns all sets in s that contain n
func (s Matrix) Find(n Int) chan Set {
	c := make(chan Set)

	go func() {
		defer close(c)

		for _, set := range s {
			if set.Contains(n) {
				c <- set
			}
		}
	}()

	return c
}

// MaxPathSum returns the maximum value available in a path through
// a numerical grid, i.e a Set of sets
func (s Matrix) MaxPathSum() Int {
	for row := len(s) - 2; row >= 0; row-- {
		for col := 0; col < len(s[row])-1; col++ {
			if s[row+1][col] > s[row+1][col+1] {
				s[row][col] += s[row+1][col]
			} else {
				s[row][col] += s[row+1][col+1]
			}
		}
	}

	return s[0][0]
}

// Coord represents the values of coordinates within a grid
type Coord struct {
	Row Int
	Col Int
}

// Direction represents an identifier for vector direction
type Direction int

// Vector Directions constants
const (
	LTR  Direction = iota // Left To Right
	RTL                   // Right To Left
	UP                    // Up
	DOWN                  // Down
	LTRU                  // Left To Right Up (diagonal)
	LTRD                  // Left To Right Down (diagonal)
	RTLU                  // Right To Left Up (diagonal)
	RTLD                  // Right To Left Down (diagonal)
)

// Vector returns a ln length set of values starting at row/col extending in Direction d.
// Vector also returns the coordinates of those values.
// If supplied Vector will set the values to replace (in order)
func (s Matrix) Vector(pos Coord, ln Int, d Direction, replaceWith ...Int) (Set, []Coord, error) {
	var (
		res  Set
		crds = make([]Coord, ln)
	)

	for i := Int(0); i < ln; i++ {
		crd := Coord{}

		switch d {
		case LTR:
			crd = Coord{pos.Row, pos.Col + i}
		case RTL:
			crd = Coord{pos.Row, pos.Col - i}
		case DOWN:
			crd = Coord{pos.Row + i, pos.Col}
		case UP:
			crd = Coord{pos.Row - i, pos.Col}
		case LTRD:
			crd = Coord{pos.Row + i, pos.Col + i}
		case RTLD:
			crd = Coord{pos.Row + i, pos.Col - i}
		case LTRU:
			crd = Coord{pos.Row - i, pos.Col + i}
		case RTLU:
			crd = Coord{pos.Row - i, pos.Col - i}
		}

		if crd.Row >= Int(len(s)) || crd.Row < 0 || crd.Col >= Int(len(s[crd.Row])) || crd.Col < 0 {
			return nil, nil, fmt.Errorf("Vector out of bounds [ROW|COL]:[%d|%d]", crd.Row, crd.Col)
		}

		if i < Int(len(replaceWith)) {
			s[crd.Row][crd.Col] = replaceWith[i]
		}

		res = append(res, s[crd.Row][crd.Col])
		crds = append(crds, crd)
	}

	return res, crds, nil
}

/* Using Matrix as a grid
Consider the following:
Set{
	Set{25,10,11,12,13},
	Set{24,09,02,03,14},
	Set{23,08,01,04,15},
	Set{22,07,06,05,16},
	Set{21,20,19,18,17}
}
*/

// NewNumberSpiral creates a square grid number spiral of width size. If size is even it is incremented
// to become odd.
/*
CONSIDER:
for i := 1; i < max; i += inc {
	inc increases every 2nd and 4th turn
	use vector, supply n = i..{i+inc}
}
*/
/*
func NewNumberSpiral(size Int) Matrix {
	if size%2 == 0 {
		size++
	}

	grid := make(Matrix, size)
	for row := range grid {
		grid[row] = make(Set, size)
	}

	// Starting from the center head up 1...
	crd, max := Coord{Row: size / 2, Col: size / 2}, size*size
	i, inc := Int(1), Int(1)
	for i <= max {
		_, crds, err := grid.Vector(crd.Row, crd.Col, inc, UP, Range(i, i+inc)...)
		if err != nil {
			break
		}
		crd = crds[len(crds)-1]
		i += inc

		_, crds, err = grid.Vector(crd.Row, crd.Col, inc, LTR, Range(i, i+inc)...)
		if err != nil {
			break
		}
		crd = crds[len(crds)-1]
		i += inc
		inc++

		_, crds, err = grid.Vector(crd.Row, crd.Col, inc, DOWN, Range(i, i+inc)...)
		if err != nil {
			break
		}
		crd = crds[len(crds)-1]
		i += inc

		_, crds, err = grid.Vector(crd.Row, crd.Col, inc, RTL, Range(i, i+inc)...)
		if err != nil {
			break
		}
		crd = crds[len(crds)-1]
		i += inc
		inc++
	}

	return grid
}
*/

/*
	switch {
			case d == UP:
				grid[row-1][col] = i
			case d == LTR:
				grid[row][col+1] = i
			case d == DOWN:
				grid[row+1][col] = i
			case d == RTL:
				grid[row][col-1] = i
			}
*/
