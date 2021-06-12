package snake

import (
	"github.com/BattlesnakeOfficial/rules"
)

type state struct {
	boardState *rules.BoardState
	ruleset    rules.Ruleset
	indexes    map[string]int
	turn       int
}

func stateFromPayload(p payload) state {
	var s state
	s.boardState = &rules.BoardState{
		Height: int32(p.Board.Height),
		Width:  int32(p.Board.Width),
		Food:   make([]rules.Point, len(p.Board.Food)),
		Snakes: make([]rules.Snake, len(p.Board.Snakes)),
	}
	for i, f := range p.Board.Food {
		s.boardState.Food[i] = rules.Point(f)
	}
	for i, snake := range p.Board.Snakes {
		s.boardState.Snakes[i].ID = p.Board.Snakes[i].ID
		s.boardState.Snakes[i].Health = int32(p.Board.Snakes[i].Health)
		s.boardState.Snakes[i].Body = make([]rules.Point, len(p.Board.Snakes[i].Body))

		for j := range snake.Body {
			s.boardState.Snakes[i].Body[j] = rules.Point(p.Board.Snakes[i].Body[j])
		}
	}

	// TODO: Read from payload. Specialy solo!!!!!2
	s.ruleset = &rules.StandardRuleset{}

	if len(s.boardState.Snakes) == 1 {
		s.ruleset = &rules.SoloRuleset{}
	}

	s.indexes = make(map[string]int)
	for i, snake := range p.Board.Snakes {
		s.indexes[snake.ID] = i
		if snake.ID == p.You.ID {
			s.indexes["me"] = i
		}
	}

	s.turn = p.Turn
	return s
}

func (s state) me() rules.Snake {
	return s.boardState.Snakes[s.indexes["me"]]
}

func (s state) snakesAlive() int {
	count := 0

	for _, snake := range s.boardState.Snakes {
		if snake.EliminatedCause != rules.NotEliminated {
			continue
		}
		count++
	}
	return count
}

func boardFree(b *rules.BoardState, p point, tempPoints []point) bool {
	if p.X >= b.Height || p.Y >= b.Width || p.X < 0 || p.Y < 0 {
		// Out of the board
		return false
	}

	for _, snake := range b.Snakes {
		// The last field of the snake will be free on the next turn.
		for _, p2 := range snake.Body[:len(snake.Body)-1] {
			if p == point(p2) {
				return false
			}
		}
	}

	for _, p2 := range tempPoints {
		if p == p2 {
			return false
		}
	}
	return true
}
