package pkg_test

import (
	"testing"

	"github.com/joubertredrat-tests/premium-minds-dev-test-golang-2k21/pkg"
)

func TestGaming(t *testing.T) {
	tests := []struct {
		name          string
		steps         string
		countExpected int
	}{
		{
			name:          "Test input1",
			steps:         "N",
			countExpected: 2,
		},
		{
			name:          "Test input2",
			steps:         "NESO",
			countExpected: 4,
		},
		{
			name:          "Test input3",
			steps:         "NSNSNSNSNS",
			countExpected: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			game, _ := pkg.NewPokemonGame(
				pkg.NewPokemonStorage(),
				pkg.NewPosition(),
				test.steps,
			)

			game.Start()
			countGot, _ := game.GetPokemonCount()

			if test.countExpected != countGot {
				t.Errorf("game.GetPokemonCount() expected %d, got %d", test.countExpected, countGot)
			}
		})
	}
}

func TestInvalidSteps(t *testing.T) {
	_, err := pkg.NewPokemonGame(
		pkg.NewPokemonStorage(),
		pkg.NewPosition(),
		"NESAO",
	)

	if err == nil {
		t.Errorf("pkg.NewPokemonGame() expected error, got nil")
	}
}

func TestGameAlreadyStarted(t *testing.T) {
	game, _ := pkg.NewPokemonGame(
		pkg.NewPokemonStorage(),
		pkg.NewPosition(),
		"NESO",
	)

	game.Start()
	err := game.Start()

	if err == nil {
		t.Errorf("game.Start() expected error, got nil")
	}
}

func TestGameNotFinished(t *testing.T) {
	game, _ := pkg.NewPokemonGame(
		pkg.NewPokemonStorage(),
		pkg.NewPosition(),
		"NESO",
	)

	_, err := game.GetPokemonCount()

	if err == nil {
		t.Errorf("game.GetPokemonCount() expected error, got nil")
	}
}
