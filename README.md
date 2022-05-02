# FLL-td
## Linux and WSL
### Requirements
 - git
 - make
 - arm-linux-gnueabi-gcc

To install required packages on debian/ubuntu (apt) based systems, use:
```zsh
sudo apt-get install git make libc6-armel-cross libc6-dev-armel-cross binutils-arm-linux-gnueabi libncurses5-dev build-essential bison flex libssl-dev bc gcc-arm-linux-gnueabi g++-arm-linux-gnueabi
```
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
To create a new report, [use this new issue form](https://github.com/TatranskiDravci/FLL-td/issues/new?assignees=LukasDrsman&labels=report&template=report-template.md&title=Report+%5Breport+number%5D%3A+%5Bshort+description%2Fnote%5D).
