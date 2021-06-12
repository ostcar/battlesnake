package snake

import (
	"encoding/json"
	"fmt"

	"github.com/BattlesnakeOfficial/rules"
)

type payload struct {
	Game struct {
		ID      string          `json:"id"`
		Ruleset json.RawMessage `json:"ruleset"`
		Timeout int             `json:"timeout"`
	} `json:"game"`
	Turn  int    `json:"turn"`
	Board board  `json:"board"`
	You   *snake `json:"you"`
}

func (p *payload) linkYou() {
	for i := range p.Board.Snakes {
		if p.You.ID == p.Board.Snakes[i].ID {
			p.You = &p.Board.Snakes[i]
			return
		}
	}
	panic("you are not on the board")
}

type board struct {
	Height  int     `json:"height"`
	Width   int     `json:"width"`
	Food    []point `json:"food"`
	Hazards []point `json:"hazards"`
	Snakes  []snake `json:"snakes"`

	tempPoints []point
}

type point rules.Point

func (p point) neighbors() []point {
	return []point{
		{p.X + 1, p.Y},
		{p.X - 1, p.Y},
		{p.X, p.Y + 1},
		{p.X, p.Y - 1},
	}
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

func (p point) distance(other point) int32 {
	a := abs(p.X - other.X)
	b := abs(p.Y - other.Y)
	return a + b
}

func abs(i int32) int32 {
	if i < 0 {
		return -1 * i
	}
	return i
}

type snake struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Health int     `json:"health"`
	Body   []point `json:"body"`
	//Latency string  `json:"latency"`
	Head   point  `json:"head"`
	Length int    `json:"length"`
	Shout  string `json:"shout"`
	Squad  string `json:"squad"`
}

func (s snake) direction() direction {
	if len(s.Body) < 2 {
		return dRight
	}

	return s.Body[1].direction(s.Body[0])
}

func (s snake) possibleNext() []point {
	var ps []point
	for _, p := range s.Head.neighbors() {
		ps = append(ps, p)
	}
	return ps
}

type direction int

const (
	dUp direction = iota
	dRight
	dDown
	dLeft
)

func toDirection(in string) direction {
	switch in {
	case "up":
		return dUp
	case "right":
		return dRight
	case "down":
		return dDown
	case "left":
		return dLeft
	}
	panic(fmt.Sprintf("invalid direction %s", in))
}

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

func nearestPoint(p point, others []rules.Point) int {
	n := 0
	d := int32(10000)
	for i, o := range others {
		if newd := p.distance(point(o)); newd < d {
			d = newd
			n = i
		}
	}
	return n
}
