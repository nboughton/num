package num

import (
	"fmt"
)

// Matrix is a slice of slices of Int
type Matrix []Set

// NewMatrix creates an empty matrix of 0s with dimensions rowsXcolumns
func NewMatrix(rows Int, cols Int) Matrix {
	m := Matrix{}

	for r := Int(0); r < rows; r++ {
		m = append(m, make(Set, cols))
	}

	return m
}

// Implement the sort interface
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

// NumberSpiral creates a square grid number spiral of width size. If size is even it is incremented
// to become odd.
/*
func NumberSpiral(size Int) Matrix {
	if size%2 != 0 {
		size++
	}

	var (
		// Lets declare our bits here in a nice orderly list
		m = NewMatrix(size, size)
	//			vecOrder = []Direction{UP, LTR, DOWN, RTL}
		//	vecIdx   = 0
		//	vecLen   = Int(1)
		r = size / 2
		c = size / 2
	)

	// LET THE FUCKERY BEGIN
	m[r][c] = 1
	for i := Int(1); i <= size*size; i++ {
		if i%2 != 0 {
			v, _, err := m.Vector(Coord{r, c}, i+1, LTR)
			if err != nil {
				break
			}
		} else {

		}

	}

	return m
}
*/
