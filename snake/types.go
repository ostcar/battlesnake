package snake

import (
	"encoding/json"
	"math"
)

type payload struct {
	Game struct {
		ID      string          `json:"id"`
		Ruleset json.RawMessage `json:"ruleset"`
		Timeout int             `json:"timeout"`
	} `json:"game"`
	Turn  int   `json:"turn"`
	Board board `json:"board"`
	You   snake `json:"you"`
}

type board struct {
	Height  int     `json:"height"`
	Width   int     `json:"width"`
	Food    []point `json:"food"`
	Hazards []point `json:"hazards"`
	Snakes  []snake `json:"snakes"`

	tempPoints []point
}

func (b board) free(p point) bool {
	if p.X >= b.Height || p.Y >= b.Width || p.X < 0 || p.Y < 0 {
		// Out of the board
		return false
	}

	for _, snake := range b.Snakes {
		for _, p2 := range snake.Body {
			if p == p2 {
				return false
			}
		}
	}

	for _, p2 := range b.tempPoints {
		if p == p2 {
			return false
		}
	}
	return true
}

type point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// next returns the point that is next to this point at the given direction.
func (p point) next(d direction) point {
	switch d {
	case dUp:
		p.Y++
	case dRight:
		p.X++
	case dDown:
		p.Y--
	case dLeft:
		p.X--
	}
	return p
}

// direction returns the direction to another point.
//
// It first goes on the X axis and then on the Y axis.
//
// It returns dUp, if p == other.
func (p point) direction(other point) direction {
	if p.X < other.X {
		return dRight
	}

	if p.X > other.X {
		return dLeft
	}

	if p.Y > other.Y {
		return dDown
	}

	return dUp
}

func (p point) distance(other point) float64 {
	a := p.X - other.X
	b := p.Y - other.Y
	return math.Sqrt(float64(a*a + b*b))
}

type snake struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Health  int     `json:"health"`
	Body    []point `json:"body"`
	Latency string  `json:"latency"`
	Head    point   `json:"head"`
	Length  int     `json:"length"`
	Shout   string  `json:"shout"`
	Squad   string  `json:"squad"`
}

func (s snake) direction() direction {
	if len(s.Body) < 2 {
		return dRight
	}

	return s.Body[1].direction(s.Body[0])
}

type direction int

const (
	dUp direction = iota
	dRight
	dDown
	dLeft
)

func (d direction) String() string {
	switch d % 4 {
	case 0:
		return "up"
	case 1:
		return "right"
	case 2:
		return "down"
	case 3:
		return "left"
	default:
		panic("Invalid direction")
	}
}
