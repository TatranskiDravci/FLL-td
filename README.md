# FLL-td
## Linux and WSL
### Requirements
 - git
 - make
 - arm-linux-gnueabi-gcc

To install required packages on debian/ubuntu (apt) based systems, use:
```sh
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
and build all targets,
```sh
make
```
### Makefile
To create build instructions for `runX.c`, use this template:
```make
runX: runX.c $(DRIVERS) $(INCLUDE)
	$(CC) $^ -o build/runX
```
To clean `build/` directory, use `make clean`.

## Reports
To create a new report, [use this new issue form](https://github.com/TatranskiDravci/FLL-td/issues/new?assignees=LukasDrsman&labels=report&template=report-template.md&title=Report+%5Breport+number%5D%3A+%5Bshort+description%2Fnote%5D).

## Documentation
The documentation for the code, provided in subdirectories of `src/` can be found [here](https://github.com/TatranskiDravci/FLL-td/wiki). One may also refer to the header files (`.h` files) in the subdirectories of `src/`. These contain relatively detailed descriptions of the usage and parameters of the provided functions as well as some useful constants. (Note, the documentation in `src/drivers/*.h` is kept relatively simple. This is because a detailed description of every function would be beyond useless in this specific case.)

## Styleguide
A relatively comprehensive styleguide with examples can be found [here](https://github.com/TatranskiDravci/FLL-td/blob/main/STYLEGUIDE.md) or on the [wiki](https://github.com/TatranskiDravci/FLL-td/wiki/Styleguide).
