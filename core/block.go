package core

// Block — реактивный блок STB
type Block struct {
	ID      string                         // Имя блока
	Match   func(Signal) bool              // Условие возбуждения
	Execute func([]Signal) []Signal        // Реакция (новые сигналы)
}
