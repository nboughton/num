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
func (m Matrix) Len() int           { return len(m) }
func (m Matrix) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Matrix) Less(i, j int) bool { return len(m[i]) < len(m[j]) }

// Find returns all sets in m that contain n
func (m Matrix) Find(n Int) chan Set {
	c := make(chan Set)

	go func() {
		defer close(c)

		for _, set := range m {
			if set.Contains(n) {
				c <- set
			}
		}
	}()

	return c
}

// MaxPathSum returns the maximum value available in a path through
// a numerical grid, i.e a Set of sets
func (m Matrix) MaxPathSum() Int {
	for row := len(m) - 2; row >= 0; row-- {
		for col := 0; col < len(m[row])-1; col++ {
			if m[row+1][col] > m[row+1][col+1] {
				m[row][col] += m[row+1][col]
			} else {
				m[row][col] += m[row+1][col+1]
			}
		}
	}

	return m[0][0]
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
	RIGHT  Direction = iota // Left To Right
	LEFT                    // Right To Left
	UP                      // Up
	DOWN                    // Down
	RIGHTU                  // Left To Right Up (diagonal)
	RIGHTD                  // Left To Right Down (diagonal)
	LEFTU                   // Right To Left Up (diagonal)
	LEFTD                   // Right To Left Down (diagonal)
)

// Vector returns a ln length set of values starting at row/col extending in Direction d.
// Vector also returns the coordinates of those values.
// If supplied Vector will set the values to replace (in order)
func (m Matrix) Vector(pos Coord, ln Int, d Direction, replaceWith ...Int) (Set, []Coord, error) {
	var (
		res  Set
		crds []Coord
	)

	for i := Int(0); i < ln; i++ {
		crd := Coord{}

		switch d {
		case RIGHT:
			crd = Coord{pos.Row, pos.Col + i}
		case LEFT:
			crd = Coord{pos.Row, pos.Col - i}
		case DOWN:
			crd = Coord{pos.Row + i, pos.Col}
		case UP:
			crd = Coord{pos.Row - i, pos.Col}
		case RIGHTD:
			crd = Coord{pos.Row + i, pos.Col + i}
		case LEFTD:
			crd = Coord{pos.Row + i, pos.Col - i}
		case RIGHTU:
			crd = Coord{pos.Row - i, pos.Col + i}
		case LEFTU:
			crd = Coord{pos.Row - i, pos.Col - i}
		}

		if crd.Row >= Int(len(m)) || crd.Row < 0 || crd.Col >= Int(len(m[crd.Row])) || crd.Col < 0 {
			return res, crds, fmt.Errorf("Vector out of bounds [ROW|COL]:[%d|%d]", crd.Row, crd.Col)
		}

		if i < Int(len(replaceWith)) {
			m[crd.Row][crd.Col] = replaceWith[i]
		}

		res = append(res, m[crd.Row][crd.Col])
		crds = append(crds, crd)
	}

	return res, crds, nil
}

// Spiral creates a square grid number spiral of width size. If size is even it is incremented
// to become odd.
func Spiral(size Int) Matrix {
	if size%2 == 0 {
		size++
	}

	var (
		// Lets declare our bits here in a nice orderly list
		m = NewMatrix(size, size)
		r = size / 2
		c = size / 2
	)

	// LET THE FUCKERY BEGIN
	m[r][c] = 1
	for inc := Int(1); true; inc++ {
		done := false

		if inc%2 != 0 {
			for _, d := range []Direction{RIGHT, DOWN} {
				_, vec, err := m.Vector(Coord{r, c}, inc+1, d, Range(m[r][c], m[r][c]+inc)...)
				if err != nil {
					done = true
					break
				}
				last := len(vec) - 1
				r, c = vec[last].Row, vec[last].Col
			}

		} else {
			for _, d := range []Direction{LEFT, UP} {
				_, vec, err := m.Vector(Coord{r, c}, inc+1, d, Range(m[r][c], m[r][c]+inc)...)
				if err != nil {
					done = true
					break
				}
				last := len(vec) - 1
				r, c = vec[last].Row, vec[last].Col
			}

		}

		if done {
			break
		}
	}

	return m
}
