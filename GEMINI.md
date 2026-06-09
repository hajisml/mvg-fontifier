# mvg-fontifier

A lightweight Go-based utility for transforming text using a custom Greek-inspired stylized alphabet.

## Project Overview

*   **Purpose:** Stylizes text by substituting specific Latin characters with visual lookalikes from various alphabets (primarily Greek and mathematical symbols).
*   **Primary Language:** Go
*   **Structure:** A single-file application (`mvg.go`) that contains the substitution logic and a simple demonstration in the `main` function.

## Building and Running

Since this project currently lacks a `go.mod` file and external dependencies, you can run it directly using the Go toolchain.

*   **Run:** `go run mvg.go`
*   **Build:** `go build -o mvg-fontifier mvg.go`
*   **Test:** No tests currently exist. TODO: Add unit tests for `substituteWithCustomAlphabet`.

## Development Conventions

*   **Simplicity:** The project currently follows a minimalist single-file structure.
*   **Character Mapping:** Substitutions are defined in a global `customAlphabet` map.
*   **Standard Library:** Relies exclusively on Go's standard library (`fmt`, `strings`).
