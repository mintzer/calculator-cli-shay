# Calculator CLI

A very simple Python CLI calculator targeting Python 3.12+.

## Usage

```bash
# Direct execution
python calculator.py <operation> <a> <b>

# Operations: add, sub, mul, div
python calculator.py add 5 3    # Output: 8.0
python calculator.py sub 10 4   # Output: 6.0
python calculator.py mul 6 7    # Output: 42.0
python calculator.py div 20 4   # Output: 5.0
```

## Installation

```bash
python3.12 -m venv .venv
source .venv/bin/activate
pip install -e .
```

After installation, use the `calc` console script:

```bash
calc add 5 3    # Output: 8.0
```
