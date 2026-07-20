# CodeBuddy 🚀

> Generate production-ready Python code from your terminal. No internet. No tokens. No AI costs. Just templates + your input = code.

CodeBuddy is an offline CLI tool that generates clean, ready-to-use Python code. Pick a pattern, answer a few questions, and get code instantly — copied to your clipboard and ready to paste.

## Features

- 🐍 **Python code generation** — loops, functions, classes, DSA, OOP, file I/O, APIs
- 📋 **Auto clipboard copy** — generated code is instantly on your clipboard
- 💾 **File output** — save generated code directly to a file with `--output`
- 🎨 **Colored terminal output** — easy-to-read prompts and results
- ⚡ **Offline** — no internet, API keys, or accounts needed
- 🔒 **Input validation** — no crashes from invalid selections

## Installation

### Prerequisites

- [Go 1.21+](https://go.dev/dl/) installed

### From source

```bash
git clone https://github.com/chiru001/codebuddy.git
cd codebuddy
go mod tidy
go build -o codebuddy.exe ./cmd/codebuddy/
```

The binary will be at `./codebuddy.exe` (Windows) or `./codebuddy` (Linux/macOS).

### Add to PATH (optional)

Move the binary to a directory in your PATH for global access:

**Windows (PowerShell):**
```powershell
Move-Item codebuddy.exe C:\Users\$env:USERNAME\go\bin\
```

**Linux/macOS:**
```bash
go build -o codebuddy ./cmd/codebuddy/
sudo mv codebuddy /usr/local/bin/
```

## Usage

```bash
codebuddy python <command> [flags]
```

### Quick Examples

```bash
# Generate a for loop
codebuddy python for-loop

# Generate a function
codebuddy python function

# Generate a class
codebuddy python class

# Generate DSA code
codebuddy python dsa sort
codebuddy python dsa search
codebuddy python dsa pattern

# OOP patterns
codebuddy python oop inheritance
codebuddy python oop polymorphism
codebuddy python oop design-pattern

# Save output to a file instead of just printing
codebuddy python for-loop --output loop.py
codebuddy python function -o my_function.py

# Other generators
codebuddy python generator
codebuddy python error-handling
codebuddy python file-io
codebuddy python api
```

### Save to File

Use the `--output` / `-o` flag on any python subcommand to write the generated code to a file:

```bash
codebuddy python class -o models.py
codebuddy python dsa pattern --output two_sum.py
```

## Available Commands

| Command | Description |
|---------|-------------|
| `python for-loop` | Generate for loops (list, range, dict, file, nested, comprehension) |
| `python function` | Generate functions with typed params, docstrings, error handling |
| `python class` | Generate classes (regular, dataclass, with custom methods) |
| `python dsa sort` | Sorting algorithms (quick sort, merge sort, bubble sort, built-in) |
| `python dsa search` | Searching algorithms (binary search, linear, filter) |
| `python dsa pattern` | DSA patterns (two sum, sliding window, two pointer, BFS, DFS, DP) |
| `python oop inheritance` | Inheritance with parent/child classes |
| `python oop polymorphism` | Abstract base classes with polymorphism |
| `python oop encapsulation` | Private attributes with @property |
| `python oop abstract` | Abstract base class (ABC) patterns |
| `python oop design-pattern` | Singleton, Factory, Observer, Strategy, Decorator |
| `python generator` | Generator functions (yield, infinite, chunked, pipeline) |
| `python error-handling` | Try/except, custom exceptions, retry, context managers |
| `python file-io` | File read/write, JSON, CSV operations |
| `python api` | HTTP requests, FastAPI/Flask endpoints |
| `version` | Show CLI version |

## How It Works

1. You pick a code category (e.g., `python for-loop`)
2. CodeBuddy asks targeted questions about your specific needs
3. It generates customized, production-ready code
4. Code is printed with syntax highlighting, copied to clipboard, and optionally saved to a file

No generic templates — every output is tailored to your exact input.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License — see [LICENSE](LICENSE) for details.
