---
test signal calculator
---

 C:\Documents\test_calc - root folder


---

"C:\Documents\test_calc\main.go"

## main.go:

'''
package main

import (
	"fmt"
	"stb_calc/core"
	"stb_calc/blocks"
)

func main() {
	fmt.Println("🧠 STB: Сложение через возбуждённые биты")
	fmt.Println("▶ Ввод: X+Y (где X и Y — цифры от 0 до 9)")

	// 🔹 Получаем сигналы digit_X и digit_Y
	inputSignals := core.ParseUserInputDual()
	if inputSignals == nil {
		return
	}

	// 🔹 Сложение через возбуждённые бит-блоки
	result := blocks.AdderByBitsBlock.Execute(inputSignals)

	// 🔹 Вывод результата + ⚡ визуализация
	for _, r := range result {
		fmt.Printf("📤 Результат: %.0f\n", r.Phase)
		blocks.PhaseVisualizerBlock.Execute([]core.Signal{r})
	}
}

'''
---

"C:\Documents\test_calc\blocks\adder_by_bits.go"

## adder_by_bits.go:

'''
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

'''

---

"C:\Documents\test_calc\blocks\bitblocks.go"

## bitblocks.go:

'''
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

'''

 ---

 "C:\Documents\test_calc\blocks\digit_to_bits.go"

 ## digit_to_bits.go:

'''
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

'''

---

"C:\Documents\test_calc\blocks\phase_visualizer.go"

## phase_visualizer.go:

'''
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

'''

---

"C:\Documents\test_calc\blocks\summator.go"

'''
## summator.go:

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

'''
---

"C:\Documents\test_calc\core\block.go"

## block.go:

'''
package core

// Block — реактивный блок STB
type Block struct {
	ID      string                         // Имя блока
	Match   func(Signal) bool              // Условие возбуждения
	Execute func([]Signal) []Signal        // Реакция (новые сигналы)
}

'''

---

"C:\Documents\test_calc\core\dictionary.go"


## dictionary.go:

'''
package core

import "time"

// DigitSignals — все цифры 0–9 как сигналы
var DigitSignals = map[string]Signal{
	"0": {ID: "S_digit_0", Structure: "digit_0", Phase: 1.0, Time: time.Now()},
	"1": {ID: "S_digit_1", Structure: "digit_1", Phase: 1.0, Time: time.Now()},
	"2": {ID: "S_digit_2", Structure: "digit_2", Phase: 1.0, Time: time.Now()},
	"3": {ID: "S_digit_3", Structure: "digit_3", Phase: 1.0, Time: time.Now()},
	"4": {ID: "S_digit_4", Structure: "digit_4", Phase: 1.0, Time: time.Now()},
	"5": {ID: "S_digit_5", Structure: "digit_5", Phase: 1.0, Time: time.Now()},
	"6": {ID: "S_digit_6", Structure: "digit_6", Phase: 1.0, Time: time.Now()},
	"7": {ID: "S_digit_7", Structure: "digit_7", Phase: 1.0, Time: time.Now()},
	"8": {ID: "S_digit_8", Structure: "digit_8", Phase: 1.0, Time: time.Now()},
	"9": {ID: "S_digit_9", Structure: "digit_9", Phase: 1.0, Time: time.Now()},
}

// SymbolSignals — операторы
var SymbolSignals = map[string]Signal{
	"+": {ID: "S_plus", Structure: "symbol_plus", Phase: 1.0, Time: time.Now()},
	"=": {ID: "S_equals", Structure: "symbol_equals", Phase: 1.0, Time: time.Now()},
}

'''

---

"C:\Documents\test_calc\core\input.go"

## input.go:

'''
package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// ParseUserInput — принимает одну цифру (0–9) и возвращает сигнал digit_X
func ParseUserInput() []Signal {
	fmt.Print("Ввод: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	if len(text) != 1 || text[0] < '0' || text[0] > '9' {
		fmt.Println("⛔ Введите только одну цифру от 0 до 9")
		return nil
	}

	sym := string(text[0])
	signal, ok := DigitSignals[sym]
	if !ok {
		fmt.Println("⛔ Неизвестная цифра")
		return nil
	}

	return []Signal{{
		ID:        signal.ID,
		Structure: signal.Structure,
		Phase:     1.0,
		Time:      time.Now(),
	}}
}

func ParseUserInputDual() []Signal {
	fmt.Print("Ввод (X+Y): ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	if len(text) != 3 || text[1] != '+' {
		fmt.Println("⛔ Формат: X+Y (0–9)")
		return nil
	}
	a, b := text[0], text[2]
	if a < '0' || a > '9' || b < '0' || b > '9' {
		fmt.Println("⛔ Только цифры от 0 до 9")
		return nil
	}

	sa := DigitSignals[string(a)]
	sb := DigitSignals[string(b)]

	return []Signal{
		{ID: sa.ID, Structure: sa.Structure, Time: time.Now()},
		{ID: sb.ID, Structure: sb.Structure, Time: time.Now()},
	}
}

'''

---

"C:\Documents\test_calc\core\signal.go"

## signal.go:

'''
package core

import "time"

// Signal — сигнальный объект в STB
type Signal struct {
	ID        string    // Уникальный идентификатор
	Structure string    // Тип сигнала, например: "digit_2"
	Phase     float64   // Сила/значение сигнала (если декодировано)
	Time      time.Time // Время возбуждения
}

'''

---
