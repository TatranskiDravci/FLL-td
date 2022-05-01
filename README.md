# FLL-td
## Linux
### Requirements
 - git
 - make
 - arm-linux-gnueabi-gcc
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

## Reports
To create a new report, [use this new issue form](https://github.com/TatranskiDravci/FLL-td/issues/new).
### Template
Title:
```md
Report [report number]: [short description/note]
```
Body:
```md
# Quantitative section
## Number of runs
 - total executed: [total number of executed runs/tests]  
 - by the number of unsuccessful missions out of [total number of missions in a run]:
   - [total number of missions]: [count of complete failures]
   - ...: ...
   - 0: [count of perfect runs]
## Points collected
 - max: [number of points collected in the best run]
 - min: [number of points collected in the worst run]
 - avg: [average number of points collected per all executed runs]
# Qualitative section
## Issues
[detailed description of problems encountered]
## Possible improvements
[improvement suggestions]
```
