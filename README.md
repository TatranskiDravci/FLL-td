# FLL-td

## Project Structure

```
.
├── build
│   └── runX				--> RG run binary
├── go.mod
├── go.sum
├── Makefile
└── src
    ├── runX.go			--> RG run code
    ├── env.go				--> interaction with environment variables
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
Make instructions for runX.go:
```make
runX:
    $(PLATFORM) go build -o build/runX src/runX.go $(LIB)
```
