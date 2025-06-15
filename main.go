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
