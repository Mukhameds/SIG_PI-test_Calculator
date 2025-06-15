
# SIG_PI-test_Calculator

**Phase-based π Calculator Based on the Signal Theory of Being (STB)**

---

## 🧠 Project Overview

This project presents an alternative computational architecture based on the **Signal Theory of Being (STB)**. Instead of conventional arithmetic, π is extracted via reactive signal chains, where each block responds to incoming signals without using variables, loops, or memory.

The entire logic is built on the paradigm:

```

Signal → Block → Reaction → New Signal

```

---

## 🔬 What is SIGPI?

`SIGPI` (Signal Phase π Extractor) is a phase-signal calculator grounded in the **geometric relationship between circumference and diameter**. The system simulates signal propagation along two reactive circuits:

- `L` — a chain of 360 blocks (circumference phase)
- `D` — one of several diameter chains (e.g., 118+ blocks with varied phase growth rates)

The phase ratio `φ = L / D` is passed to a comparison block, where **DigitBlock[n]** are activated when `φ` enters their defined range. This enables digit-by-digit extraction of π.

---

## 📐 Architecture

- **L-Circuit**  
  360 blocks. One full signal cycle = one circumference phase.

- **D-Circuits**  
  Blocks 238–242. Represent 5 variations of diameter phase growth rates.

- **Phase Division**  
  Instead of arithmetic division, the system uses L/D excitation to produce a phase coefficient φ.

- **DigitBlock[n]**  
  Each block activates when φ falls into a defined interval:  
  `[0.0–0.1), [0.1–0.2), ..., [0.9–1.0)` — each interpreted as a single digit.

- **Signal Dictionary**  
  Maps phase to digit via block activation. The system "senses" the number by counting activated units (analogous to lit elements or Braille-like codes).

---

## 📂 Project Structure

```

SIG\_PI-test\_Calculator/
├── main.go                      # Entry point
├── go.mod                      # Go module
├── test\_calc\_stb.md            # STB logic documentation
├── blocks/
│   ├── adder\_by\_bits.go
│   ├── bitblocks.go
│   ├── digit\_to\_bits.go
│   ├── phase\_visualizer.go
│   └── summator.go
├── core/
│   ├── block.go
│   ├── dictionary.go
│   ├── input.go
│   └── signal.go
└── diagrame/
├── STB\_Pi\_diagrame\_1.png
└── STB\_Pi\_diagrame\_2.drawio.png

```

---

## 🎯 Project Goals

1. **Prove** that π can be extracted digit-by-digit through signal structures.
2. **Validate** the STB paradigm as a new computational approach.
3. **Attract engineers, thinkers, and investors** for development of the next system — `ARA-NODE`  
   → [ARA-NODE on GitHub Pages](https://mukhameds.github.io/ARA-NODE_AI-AGENT/)  
4. **Support STB as a scientific direction**  
   → [STB-Physics](https://mukhameds.github.io/STB-Physics/)

---

## 📎 Core Concepts

- Signal Theory of Being (STB)
- Division via phase ratios
- Geometric computation of π
- Digit recognition via associative signal patterns
- Non-arithmetic architecture of reasoning
- A new logic foundation for AGI/RAI

---

## 📜 License

**Proprietary**. Any usage of the code, diagrams, or architecture requires **written permission from the author**. Early contributors will be granted core team access and future development privileges.

---

## 🤝 Want to Contribute?

→ Fork the repository, clone it, and help build the calculator of the future.  
Principle: each block ≠ line of code, each bit = reaction, each digit is a signal.

---
