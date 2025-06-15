package blocks

import (
	"time"
	"stb_calc/core"
)

var SummatorBlock = core.Block{
	ID: "SummatorBlock",
	Match: func(s core.Signal) bool {
		return s.Structure == "phase_unit"
	},
	Execute: func(inputs []core.Signal) []core.Signal {
		var sum float64
		for _, s := range inputs {
			sum += s.Phase
		}
		return []core.Signal{{
			ID:        "sum_output",
			Structure: "sum_result",
			Phase:     sum,
			Time:      time.Now(),
		}}
	},
}
