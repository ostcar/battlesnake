package main

func ai(b payload) direction {
	d := b.You.direction()
	if len(b.Board.Food) > 0 {
		food := b.Board.Food[0]
		next := b.You.Head.next(d)
		if next.distance(food) >= b.You.Head.distance(food) || !b.Board.free(next) {
			// Going in the same direction would not help comming closer. Or the
			// next field is not free.
			d = b.You.Head.direction(b.Board.Food[0])
		}
	}

	return saveMove(b, d, 10)
}

func saveMove(b payload, d direction, steps int) direction {
	for i := 0; i < 4; i++ {
		if b.Board.free(b.You.Head.next(d)) {
			break
		}
		d = (d + 1) % 4
	}
	return d
}
