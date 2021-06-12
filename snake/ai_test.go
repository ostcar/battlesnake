package snake

import (
	"strings"
	"testing"

	"github.com/BattlesnakeOfficial/rules"
)

func TestBeatenSnakesPoints(t *testing.T) {
	for _, tt := range []struct {
		startalive     int
		alive          int
		expectedPoints int
	}{
		{
			3,
			3,
			0,
		},
		{
			3,
			2,
			100,
		},
		{
			4,
			2,
			100,
		},
		{
			4,
			3,
			50,
		},
	} {
		got := beatenSnakesPoints(tt.startalive, tt.alive)

		if got != tt.expectedPoints {
			t.Errorf("beatenSnakesPoints(%d, %d) == %d, expected %d", tt.startalive, tt.alive, got, tt.expectedPoints)
		}
	}
}

func TestPossibleMoves(t *testing.T) {
	for _, tt := range []struct {
		name        string
		board       rules.BoardState
		point       point
		deep        int
		expectMoves int
	}{
		{
			"test on point",
			rules.BoardState{
				Width:  11,
				Height: 11,
				Snakes: []rules.Snake{
					{
						Body: points(`
						....xxx.
						....xvx.
						`),
					},
				},
			},
			point{5, 0},
			2,
			0,
		},

		{
			"test in dead end",
			rules.BoardState{
				Width:  11,
				Height: 11,
				Snakes: []rules.Snake{
					{
						Body: points(`
						....xxxx
						....xv.x
						`),
					},
				},
			},
			point{5, 0},
			2,
			1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, moves := possibleMoves(&tt.board, tt.point, []point{}, tt.deep)
			if moves != tt.expectMoves {
				t.Errorf("possibleMoves() returned %d, expected %d", moves, tt.expectMoves)
			}
		})
	}
}

func points(s string) []rules.Point {
	lines := strings.Split(s, "\n")
	revert(lines)
	var ps []rules.Point
	y := int32(-1)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		y++
		x := int32(-1)
		for _, c := range line {
			x++
			if c == '.' {
				continue
			}
			ps = append(ps, rules.Point{X: x, Y: y})
		}
	}
	return ps
}

func revert(in []string) {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
}
