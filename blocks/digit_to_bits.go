package blocks

import (
	"stb_calc/core"
	"strings"
)

var DigitToBitRouter = core.Block{
	ID: "DigitToBitRouter",
	Match: func(s core.Signal) bool {
		return strings.HasPrefix(s.Structure, "digit_")
	},
	Execute: func(inputs []core.Signal) []core.Signal {
		var result []core.Signal
		for _, s := range inputs {
			digit := strings.TrimPrefix(s.Structure, "digit_")
			switch digit {
			case "0":
				// ничего не возбуждаем
			case "1", "2", "3", "4", "5", "6", "7", "8", "9":
				n := int(digit[0] - '0')
				for i := 0; i < n; i++ {
					result = append(result, core.Signal{
						ID:        "bit_signal",
						Structure: "bit_" + string('0'+i),
						Phase:     0.0,
						Time:      s.Time,
					})
				}
			}
		}
		return result
	},
}
