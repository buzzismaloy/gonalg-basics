package randbyte

import (
	"encoding/binary"
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
	for i := 0; i+8 < len(bytes); i += 8 {
		randInt := g.rnd.Int63() // get random positive int in [0, 2^63]
		binary.LittleEndian.PutUint64(bytes[i:i+8], uint64(randInt))
	}

	return len(bytes), nil
}
