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
