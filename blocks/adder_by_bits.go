package blocks

import (
	"strings"
	"time"
	"stb_calc/core"
)

var AdderByBitsBlock = core.Block{
	ID: "AdderByBitsBlock",
	Match: func(s core.Signal) bool {
		// Ждём два сигнала: digit_X + digit_Y
		return strings.HasPrefix(s.Structure, "digit_")
	},
	Execute: func(inputs []core.Signal) []core.Signal {
		if len(inputs) != 2 {
			return nil // нужен ровно 2 сигнала
		}

		var bitSignals []core.Signal
		for _, s := range inputs {
			digit := strings.TrimPrefix(s.Structure, "digit_")
			if len(digit) != 1 || digit[0] < '0' || digit[0] > '9' {
				continue
			}
			n := int(digit[0] - '0')
			for i := 0; i < n; i++ {
				bitSignals = append(bitSignals, core.Signal{
					ID:        "bit_signal",
					Structure: "bit_" + string('0'+i),
					Phase:     0.0,
					Time:      time.Now(),
				})
			}
		}

		// активируем каждый BitBlock
		var phaseUnits []core.Signal
		for _, bit := range BitBlocks {
			for _, s := range bitSignals {
				if bit.Match(s) {
					phaseUnits = append(phaseUnits, bit.Execute([]core.Signal{s})...)
				}
			}
		}

		// передаём в SummatorBlock
		return SummatorBlock.Execute(phaseUnits)
	},
}
