package pkg_test

import (
	"math/rand"
	"testing"

	"github.com/joubertredrat-tests/premium-minds-dev-test-golang-2k21/pkg"
)

func TestStorageStart(t *testing.T) {
	storage := pkg.NewPokemonStorage()

	if 0 != storage.Count() {
		t.Errorf("storage.Count() expected 0, got %d", storage.Count())
	}

	countExpected := rand.Intn(99-10) + 99

	for i := 1; i <= countExpected; i++ {
		storage.AddOne()
	}

	if countExpected != storage.Count() {
		t.Errorf("storage.Count() expected %d, got %d", countExpected, storage.Count())
	}
}
