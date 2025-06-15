package core

import "time"

// Signal — сигнальный объект в STB
type Signal struct {
	ID        string    // Уникальный идентификатор
	Structure string    // Тип сигнала, например: "digit_2"
	Phase     float64   // Сила/значение сигнала (если декодировано)
	Time      time.Time // Время возбуждения
}
