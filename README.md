# ϻvg-ϝoηtiϝier

[![Go Report Card](https://goreportcard.com/badge/github.com/hajisml/mvg-fontifier)](https://goreportcard.com/report/github.com/hajisml/mvg-fontifier)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A text transformation utility that stylizes Latin characters using a custom mapping of Greek and mathematical symbols.

## ✨ Features

- **Instant Stylization:** Convert standard Latin text into stylized character sets.
- **Command-Line Interface:** Easy integration into terminal workflows via arguments.
- **Tested Core:** Robust transformation logic with full unit test coverage.

## 🛠 Installation

```bash
git clone https://github.com/hajisml/mvg-fontifier.git
cd mvg-fontifier
make build
```

## 📖 Usage

```bash
./bin/mvg "Hello World"
# Output: Hello Worlɗ
```

## 🛠 Development

```bash
make build    # Build the 'mvg' binary
make test     # Run unit tests
make clean    # Remove build artifacts
```

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.
