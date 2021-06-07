package snake

import (
	"strings"
	"testing"
)

func TestBestMove(t *testing.T) {
	for _, tt := range []struct {
		name            string
		board           payload
		direction       direction
		expectDirection direction
	}{
		{
			"best suicide",
			payload{
				Board: board{
					Width:  11,
					Height: 11,
					Snakes: []snake{
						{
							Body: points(`
						xxxx.
						x..x.
						xvxx.
						xx...
						`),
						},
					},
				},
				You: snake{
					Head: point{1, 1},
				},
			},
			direction(dLeft),
			dUp,
		},

		{
			"go around",
			payload{
				Board: board{
					Width:  7,
					Height: 7,
					Snakes: []snake{
						{
							Body: points(`
							...xx.x
							xxxxxvx
							x.....x
							x.....x
							x.....x
							x.....x
							xxxxxxx
							`),
						},
					},
				},
				You: snake{
					Head: point{5, 5},
				},
			},
			direction(dUp),
			dDown,
		},

		{
			"long sackway",
			payload{
				Board: board{
					Width:  11,
					Height: 11,
					Snakes: []snake{
						{
							Body: points(`
							.
							xxx
							x.x
							x.x
							x.x
							x.x
							x.x
							x.x
							x.x
							x.v
							xxxxxxx
							`),
						},
					},
				},
				You: snake{
					Head: point{2, 1},
				},
			},
			direction(dLeft),
			dRight,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := bestMove(tt.board, tt.direction)

			if got != tt.expectDirection {
				t.Errorf("saveMove() returned %s, expected %s", got, tt.expectDirection)
			}
		})
	}
}

func TestIsSave(t *testing.T) {
	for _, tt := range []struct {
		name        string
		board       board
		point       point
		deep        int
		expectSave  bool
		expectMoves int
	}{
		{
			"test on point",
			board{
				Width:  11,
				Height: 11,
				Snakes: []snake{
					{
						Body: points(`
						........
						....vxx.
						`),
					},
				},
			},
			point{5, 0},
			2,
			false,
			0,
		},

		{
			"test in dead end",
			board{
				Width:  11,
				Height: 11,
				Snakes: []snake{
					{
						Body: points(`
						....xxx.
						....v.x.
						`),
					},
				},
			},
			point{5, 0},
			2,
			false,
			1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			save, moves := isSave(tt.board, tt.point, tt.deep)
			if save != tt.expectSave || moves != tt.expectMoves {
				t.Errorf("isSave() returned (%t, %d), expected (%t, %d)", save, moves, tt.expectSave, tt.expectMoves)
			}
		})
	}
}

func points(s string) []point {
	lines := strings.Split(s, "\n")
	revert(lines)
	var ps []point
	y := -1
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		y++
		x := -1
		for _, c := range line {
			x++
			if c == '.' {
				continue
			}
			ps = append(ps, point{x, y})
		}
	}
	return ps
}

func revert(in []string) {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
}
