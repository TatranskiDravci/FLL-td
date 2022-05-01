# FLL-td
## Linux
### Quickstart
Clone,
```sh
git clone https://github.com/TatranskiDravci/FLL-td.git
cd FLL-td
```
initialize build directory,
```sh
make init
```
and build run (e.g. `run.c`),
```sh
make run
```
### Makefile
To create build instructions for `runX.c`, use this template:
```make
runX: runX.c sensor.c motor.c shifter.c base.c module.c move.c pid.c
	$(CC) $^ -o build/runX
```
To clean `build/` directory, use `make clean`.
