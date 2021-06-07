package snake

func ai(p payload) direction {
	d := p.You.direction()
	if len(p.Board.Food) > 0 && p.You.Health < 30 {
		food := p.Board.Food[nearestPoint(p.You.Head, p.Board.Food)]
		next := p.You.Head.next(d)
		if next.distance(food) >= p.You.Head.distance(food) || !p.Board.free(next) {
			// Going in the same direction would not help comming closer. Or the
			// next field is not free.
			d = p.You.Head.direction(p.Board.Food[0])
		}
	}

	return bestMove(p, d)
}

// bestMove calculates a value for each direction and returns the direction with
// the highest value.
func bestMove(p payload, d direction) direction {
	best := 0
	max := 0
	for i := 0; i < 4; i++ {
		newPoint := p.You.Head.next(direction(int(d)+i) % 4)
		_, movesToDie := isSave(p.Board, newPoint, 10)

		for _, snake := range p.Board.Snakes {
			if snake.ID == p.You.ID {
				continue
			}
			neighbors := snake.Head.neighbors()
			for _, neighbor := range neighbors {
				if newPoint == neighbor {
					movesToDie -= len(neighbors)
				}
			}
		}

		if movesToDie > max {
			max = movesToDie
			best = i
		}
	}

	return direction(int(d)+best) % 4
}

// isSave tells if going to point p is save. If it is not save, it returns a
// value how many moves it would take to die. If it is save, then deep is
// returned.
func isSave(b board, p point, deep int) (bool, int) {
	if !b.free(p) {
		return false, 0
	}

	if deep < 0 {
		return true, 0
	}

	b.tempPoints = append(b.tempPoints, p)
	maxMoves := 0
	for d := dUp; d < 4; d++ {
		save, moves := isSave(b, p.next(d), deep-1)
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

func nearestPoint(p point, others []point) int {
	n := 0
	d := 10000.0
	for i, o := range others {
		if p.distance(o) < d {
			n = i
		}
	}
	return n
}
