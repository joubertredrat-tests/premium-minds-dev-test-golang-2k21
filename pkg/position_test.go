package pkg_test

import (
	"testing"

	"github.com/joubertredrat-tests/premium-minds-dev-test-golang-2k21/pkg"
)

func TestPositionStart(t *testing.T) {
	position := pkg.NewPosition()

	if position.IsAlreadyVisited() {
		t.Error("position.IsAlreadyVisited() expected false, got true")
	}

	position.RegisterPosition()

	if !position.IsAlreadyVisited() {
		t.Error("position.IsAlreadyVisited() expected true, got false")
	}
}

func TestPositionsMove(t *testing.T) {
	tests := []struct {
		name        string
		getPosition func() pkg.Position
	}{
		{
			name: "Testing moving to North",
			getPosition: func() pkg.Position {
				position := pkg.NewPosition()
				position.MoveToNorth()
				return position
			},
		},
		{
			name: "Testing moving to South",
			getPosition: func() pkg.Position {
				position := pkg.NewPosition()
				position.MoveToSouth()
				return position
			},
		},
		{
			name: "Testing moving to East",
			getPosition: func() pkg.Position {
				position := pkg.NewPosition()
				position.MoveToEast()
				return position
			},
		},
		{
			name: "Testing moving to West",
			getPosition: func() pkg.Position {
				position := pkg.NewPosition()
				position.MoveToWest()
				return position
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			position := test.getPosition()

			if position.IsAlreadyVisited() {
				t.Error("position.IsAlreadyVisited() expected false, got true")
			}

			position.RegisterPosition()

			if !position.IsAlreadyVisited() {
				t.Error("position.IsAlreadyVisited() expected true, got false")
			}
		})
	}
}
