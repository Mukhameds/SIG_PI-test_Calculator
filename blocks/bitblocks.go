package blocks

import (
	"fmt"
	"time"
	"stb_calc/core"
)

func MakeBitBlock(n int) core.Block {
	return core.Block{
		ID: fmt.Sprintf("BitBlock_%d", n),
		Match: func(s core.Signal) bool {
			return s.Structure == fmt.Sprintf("bit_%d", n)
		},
		Execute: func(_ []core.Signal) []core.Signal {
			return []core.Signal{{
				ID:        fmt.Sprintf("bit_output_%d", n),
				Structure: "phase_unit",
				Phase:     1.0,
				Time:      time.Now(),
			}}
		},
	}
}

var BitBlocks = []core.Block{
	MakeBitBlock(0), MakeBitBlock(1), MakeBitBlock(2),
	MakeBitBlock(3), MakeBitBlock(4), MakeBitBlock(5),
	MakeBitBlock(6), MakeBitBlock(7), MakeBitBlock(8),
	MakeBitBlock(9),
}

