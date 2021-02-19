package pkg

type PokemonStorage struct {
	count int
}

func NewPokemonStorage() PokemonStorage {
	return PokemonStorage{
		count: 0,
	}
}

func (p *PokemonStorage) AddOne() {
	p.count++
}

func (p *PokemonStorage) Count() int {
	return p.count
}
