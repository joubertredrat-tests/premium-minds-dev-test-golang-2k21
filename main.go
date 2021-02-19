package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joubertredrat-tests/premium-minds-dev-test-golang-2k21/pkg"
)

func showIndex() {
	filename := filepath.Base(os.Args[0])

	fmt.Printf("\n\n")
	fmt.Printf("Pokemon: apanhá-los todos \n\n")
	fmt.Printf("Como jogar: \n")
	fmt.Printf("  %s pokemon:capturar --casas [casas] \n\n", filename)
	fmt.Printf("Nas casas informar \"N\" para norte, \"S\" para sul, \"E\" para este e/ou \"O\" para oeste	\n\n")
	fmt.Printf("Exemplos: \n")
	fmt.Printf("  %s pokemon:capturar --casas N \n", filename)
	fmt.Printf("  %s pokemon:capturar --casas NESO \n", filename)
	fmt.Printf("  %s pokemon:capturar --casas NSNSNSNSNS \n", filename)
	fmt.Printf("\n\n")
	os.Exit(0)
}

func showErrorSteps() {
	fmt.Printf("\n\n")
	fmt.Printf(" [WARNING] Você deve ter informado alguma casa errada, tente novamente! \n")
	fmt.Printf("\n\n")
	os.Exit(255)
}

func showInternalError() {
	fmt.Printf("\n\n")
	fmt.Printf(" [ERROR] Alguma coisa deu errado! Tente novamente \n")
	fmt.Printf("\n\n")
	os.Exit(255)
}

func showResult(count int) {
	fmt.Printf("\n\n")
	fmt.Printf(" [OK] Legal! Você apanhou %d pokemons \n", count)
	fmt.Printf("\n\n")
	os.Exit(0)
}

func isset(arr []string, index int) bool {
	return (len(arr) > index)
}

func main() {
	if !isset(os.Args, 1) || !isset(os.Args, 2) || !isset(os.Args, 3) {
		showIndex()
	}

	if os.Args[1] != "pokemon:capturar" || os.Args[2] != "--casas" {
		showIndex()
	}

	position := pkg.NewPosition()
	storage := pkg.NewPokemonStorage()
	steps := os.Args[3]

	game, errNewGame := pkg.NewPokemonGame(storage, position, steps)

	if errNewGame != nil {
		showErrorSteps()
	}

	errStart := game.Start()

	if errStart != nil {
		showInternalError()
	}

	count, errCount := game.GetPokemonCount()

	if errCount != nil {
		showInternalError()
	}

	showResult(count)
}
