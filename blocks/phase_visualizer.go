package blocks

import (
	"fmt"
	"math"
	
	"stb_calc/core"
)

var PhaseVisualizerBlock = core.Block{
	ID: "PhaseVisualizerBlock",
	Match: func(s core.Signal) bool {
		return s.Structure == "sum_result"
	},
	Execute: func(inputs []core.Signal) []core.Signal {
		for _, s := range inputs {
			n := int(math.Round(s.Phase))
			fmt.Print("⚡ Активация: ")
			for i := 0; i < n; i++ {
				fmt.Print("⚡")
			}
			fmt.Printf(" (%d)\n", n)
		}
		return []core.Signal{} // не возбуждает дальше
	},
}
