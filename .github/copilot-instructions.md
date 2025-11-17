<!-- .github/copilot-instructions.md - guidance for AI coding agents working in this repo -->
# Copilot / AI Agent Instructions

Purpose: Help an AI agent be immediately productive in this learning-oriented Go repository.

**Quick Overview**
- **Repo shape:** A set of small, mostly standalone Go examples/exercises under `go-basic/` (folders `01-hello-world` .. `10-cli-app`). Each example is its own `package main` program and may include a `go.mod`.
- **Primary interactive example:** `10-cli-app` contains a CLI student manager (`10-cli-app/main.go`) — this is the most feature-complete sample and the main focus for changes.

**How to run examples**
- Change into the folder for the example you want to run. Example for the CLI:
  - `cd 10-cli-app`
  - `go run .`
  - Or build: `go build -o app && ./app`

**Key project-specific patterns (discoverable in code)**
- `10-cli-app/main.go` uses an in-memory slice `daftarMahasiswa []Mahasiswa` as the data store — there is no persistence; state is lost on exit.
- Validation pattern: `validasiIPK` returns `(float64, error)`; project uses both custom `ValidationError` (struct implementing `error`) and sentinel `var` errors like `ErrNotFound`, `ErrInvalidInput`.
- CLI I/O: input reading is done with `bufio.Scanner` and `scanner.Scan()`; many helper functions accept `*bufio.Scanner` to re-use the scanner.
- Error handling: prefer returning sentinel errors for control-flow checks (e.g., `ErrNotFound`) and use typed `ValidationError` for field-level messages — preserve this distinction when extending.

**Important implementation gotchas & examples for agents**
- Avoid taking the address of the loop variable in range loops. Example bug in this repo: `cariByNIM` returns `&m` from a `for _, m := range daftarMahasiswa` loop. That returns a pointer to the loop variable (a copy) and can lead to subtle bugs. Prefer returning the index or a pointer to the slice element:
  - Correct lookup pattern to return pointer to slice element:
    - `for i := range daftarMahasiswa { if daftarMahasiswa[i].NIM == nim { return &daftarMahasiswa[i], nil } }`
- Mutations: code that updates or deletes an item finds the index and modifies `daftarMahasiswa` in-place (see `updateIPK` and `hapusMahasiswa`). Follow the same pattern when adding features.

**Conventions & style specific to this repo**
- Message/identifiers use Indonesian (e.g., `Mahasiswa`, `IPK`, `daftarMahasiswa`) — preserve language in user-facing text.
- Each exercise is standalone. Make edits inside the specific folder (e.g., `10-cli-app/`) rather than changing cross-example structure.
- Formatting: use `gofmt`/`gofmt -w .` and `go vet` before submitting changes.

**Build/test/debug workflows**
- No test suite is present in the CLI example. To validate behavior run the program manually (`go run .`) and exercise CLI flows.
- Use `go build` to check for compile issues. Use `go vet` for additional static checks.

**When adding features or fixing bugs**
- Keep changes localized to the example folder. If adding shared utilities, create a `utils` package only if multiple examples truly need it.
- Preserve the current error types and messages unless unifying them is part of the task; user-facing text should remain Indonesian unless instructed otherwise.

**Integration & dependencies**
- There are no external third-party dependencies in `10-cli-app`; other folders are simple Go examples. Use the Go toolchain only.

**Where to look for examples and behaviors**
- `10-cli-app/main.go` — CLI, validation, sentinel errors, `Mahasiswa` struct.
- `09-defer-panic` and `08-go-routines` — examples for defer/panic and goroutines if you need concurrency examples.

If anything in this file is unclear or you want more detail (run scripts, tests, or a refactor suggestion), tell me which part to expand and I'll update this file.
