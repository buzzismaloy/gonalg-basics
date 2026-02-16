package randbyte

import (
	"io"
	"math/rand"
)

type generator struct {
	rnd rand.Source
}

// New — note that we return a generator assigned to the io.Reader interface; the generator structure itself is unexported.
// We've hidden all the details inside the package.
func New(seed int64) io.Reader {
	return &generator{
		rnd: rand.NewSource(seed),
	}
}

func (g *generator) Read(bytes []byte) (n int, err error) {
	for i := range bytes {
		randInt := g.rnd.Int63() // get random positive int in [0, 2^63]
		randByte := byte(randInt)
		bytes[i] = randByte
	}

	return len(bytes), nil
}
