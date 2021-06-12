package snake

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/BattlesnakeOfficial/rules"
)

const lookAheadMoves = 5

func ai(s state) direction {
	debugLog("call ai on turn %d", s.turn)
	if len(s.me().Body) < 2 || s.me().Body[0] == s.me().Body[1] {
		// First move.
		return dUp
	}

	d, _, err := lookAhead(s, s.snakesAlive(), lookAheadMoves)
	if err != nil {
		log.Printf("Error: %v", err)
		return dUp
	}
	debugLog("turn %d: send %s\nxxxxxxxxx", s.turn, d.String())
	return d
}

// lookAhead looks some moves in the future and tells, which direction gives the
// highest points for the player.
//
// Returns the direction and the points.
//
// Does not work on the first move when the snakes do not have a direction.
func lookAhead(s state, startSnakeCount, moveCount int) (direction, int, error) {
	debugLog("lookAhead %d", moveCount)
	if moveCount <= 0 {
		points := movePoints(s, startSnakeCount)
		debugLog("points: %d", points)
		return dUp, points, nil
	}

	// Initialize snakeMove with id of the snakes. Move all snakes upwards.
	snakeMove := make([]rules.SnakeMove, 1, len(s.boardState.Snakes))
	snakeMove[0].ID = s.me().ID
	for _, snake := range s.boardState.Snakes {
		if snake.ID == s.me().ID {
			continue
		}
		snakeMove = append(snakeMove, rules.SnakeMove{ID: snake.ID, Move: dUp.String()})
	}

	maxPoints := 0
	maxDirection := []direction{dUp}
	for d := dUp; d < 4; d++ {
		debugLog("l%d: test direction %s", moveCount, d)
		// TODO: Continue if against direction.

		snakeMove[0].Move = d.String()

		minPoints := 1000
		// TODO: Fix solo game
		// if len(s.boardState.Snakes) == 1 {
		// 	// Solo Game
		// 	nextState, err := s.ruleset.CreateNextBoardState(s.boardState, snakeMove)
		// 	if err != nil {
		// 		return 0, 0, fmt.Errorf("creating next state: %w", err)
		// 	}

		// 	newState := state{
		// 		boardState: nextState,
		// 		indexes:    s.indexes,
		// 		ruleset:    s.ruleset,
		// 	}

		// 	_, minPoints, err = lookAhead(newState, startSnakeCount, moveCount-1)
		// 	if err != nil {
		// 		return 0, 0, fmt.Errorf("look ahead(%d): %w", moveCount-1, err)
		// 	}
		// }

		pos := 1
		for pos < len(snakeMove) {
			nextState, err := s.ruleset.CreateNextBoardState(s.boardState, snakeMove)
			if err != nil {
				return 0, 0, fmt.Errorf("creating next state: %w", err)
			}

			newState := state{
				boardState: nextState,
				indexes:    s.indexes,
				ruleset:    s.ruleset,
			}

			points := 0
			if newState.me().EliminatedCause == rules.NotEliminated {
				_, points, err = lookAhead(newState, startSnakeCount, moveCount-1)
				if err != nil {
					return 0, 0, fmt.Errorf("look ahead(%d): %w", moveCount-1, err)
				}
			}
			debugLog("Test moves: %v: points: %d", snakeMove, points)

			if points < minPoints {
				minPoints = points
				if minPoints <= 0 {
					break
				}
			}

			pos = nextMoves(snakeMove, pos)
		}

		debugLog("l%d: would give min points %d", moveCount, minPoints)

		if minPoints == maxPoints {
			maxDirection = append(maxDirection, d)
		} else if minPoints > maxPoints {
			maxPoints = minPoints
			maxDirection = []direction{d}
			if maxPoints >= 100 {
				break
			}
		}
	}
	debugLog("lookahead %d returns %d points on directions %v", moveCount, maxPoints, maxDirection)
	return maxDirection[rand.Intn(len(maxDirection))], maxPoints, nil
}

// nextMoves alters the snakeMove slice.
func nextMoves(snakeMove []rules.SnakeMove, pos int) int {
	// TODO: skip moves where snake runs backwards
	// TODO: skip dead snakes
	snakeMove[pos].Move = ((toDirection(snakeMove[pos].Move) + 1) % 4).String()
	if snakeMove[pos].Move == rules.MoveUp {
		return pos + 1
	}
	return 1
}

func movePoints(s state, startSnakeCount int) int {
	if over, _ := s.ruleset.IsGameOver(s.boardState); over {
		debugLog("Game is over")
		if s.me().EliminatedCause != rules.NotEliminated {
			log.Print(s.me().EliminatedCause)
			return 0
		}
		return 100
	}

	// All Points are a value between 0 and 100
	beatenPoints := beatenSnakesPoints(s.snakesAlive(), startSnakeCount)
	movePoints := possibleMovePoints(s.boardState, point(s.me().Body[0]))
	foodPoints := findFoodPoint(s)
	// TODO: be small but the biggest

	// All weights together have to be 100
	beatenPointsWeight := 40
	movePointsWeight := 20
	foodPointsWeight := 40

	return beatenPointsWeight*beatenPoints/100 + movePointsWeight*movePoints/100 + foodPointsWeight*foodPoints/100
}

// beatenSnakesPoints returns a value between 0 and 100. 0 means, no snake were
// beeten, 100 means, all-1 other snakes where beaten.
//
// The -1 is necessary, because if all other snakes are beaten, then the game is won.
func beatenSnakesPoints(startSnakeCount int, snakesAlive int) int {
	if startSnakeCount <= 2 {
		return 0
	}

	lessSnakes := startSnakeCount - snakesAlive

	return 100 * lessSnakes / (startSnakeCount - 2)
}

func possibleMovePoints(b *rules.BoardState, p point) int {
	_, possibleMovePoints := possibleMoves(b, p, []point{}, 10)
	return possibleMovePoints * 10
}

// possibleMoves tells if going to point p is save. If it is not save, it
// returns a value how many moves it would take to die. If it is save, then deep
// is returned.
func possibleMoves(b *rules.BoardState, p point, tempPoints []point, deep int) (bool, int) {
	if deep < 0 {
		return true, 0
	}

	tempPoints = append(tempPoints, p)
	maxMoves := 0
	for d := dUp; d < 4; d++ {
		next := p.next(d)
		if !boardFree(b, next, tempPoints) {
			continue
		}

		tempPoints[len(tempPoints)-1] = next
		save, moves := possibleMoves(b, next, tempPoints, deep-1)
		moves++
		if save {
			return true, deep
		}

		if moves > maxMoves {
			maxMoves = moves
		}
	}

	return false, maxMoves
}

func findFoodPoint(s state) int {
	if s.me().Health > 30 {
		return 100
	}

	ni := nearestPoint(point(s.me().Body[0]), s.boardState.Food)
	distance := point(s.me().Body[0]).distance(point(s.boardState.Food[ni]))

	if distance == 0 {
		return 100
	}

	points := (20 - int(distance)) * 5

	if points < 0 {
		return 0
	}

	return points
}
