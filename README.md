# FLL-td
A (hopefully) successful successor to FLL-evicka. Written in golang

## Project Structure

```
.
├── build
│   └── run_X				--> RG run binary
├── go.mod
├── go.sum
├── Makefile
└── src
    ├── run_X.go			--> RG run code
    ├── exten.go			--> math & miscelaneous functions
    ├── modules.go			--> code for modules on robot
    ├── robot.go			--> movement code
    └── shifter.go			--> code for shifter
```

## Makefile Structure
```make
PLATFORM = GOOS=linux GOARCH=arm GOARM=5
LIB = src/exten.go src/robot.go src/modules.go src/shifter.go

foo:
    bar
```
 - ***PLATFORM*** - platform specifications
 - ***LIB*** - ofter reused, non-RG-run go files from src/

### Example:
Make instructions for run_X.go:
```make
run_X:
    $(PLATFORM) go build -o build/run_X src/run_X.go $(LIB)
```
