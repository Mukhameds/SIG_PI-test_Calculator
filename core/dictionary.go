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
