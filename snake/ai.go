package snake

func ai(p payload) direction {
	d := p.You.direction()
	if len(p.Board.Food) > 0 {
		food := p.Board.Food[nearestPoint(p.You.Head, p.Board.Food)]
		next := p.You.Head.next(d)
		if next.distance(food) >= p.You.Head.distance(food) || !p.Board.free(next) {
			// Going in the same direction would not help comming closer. Or the
			// next field is not free.
			d = p.You.Head.direction(p.Board.Food[0])
		}
	}

	return saveMove(p.Board, p.You.Head, d)
}

func saveMove(b board, p point, d direction) direction {
	best := 0
	max := 0
	for i := 0; i < 4; i++ {
		save, moves := isSave(b, p.next(direction(int(d)+i)%4), 10)
		if save {
			return direction(int(d)+i) % 4
		}

		if moves > max {
			max = moves
			best = i
		}
	}
	return direction(int(d)+best) % 4
}

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
			return true, 0
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
