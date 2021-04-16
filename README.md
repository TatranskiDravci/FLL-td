# FLL-td

## Project Structure
```
.
├── LICENSE
├── whatever.x      additional utilities and files
├── python
│   ├── README.md
│   ├── robot.py
│   └── main.py
├── 'language name'
│   ├── robot.x     function definitions
│   ├── paths.x     collection of paths
│   ├── main.x      executable program
│   └── README.md   documentation
└── README.md
```
## Python
Version: `pybricks-micropython` <br>
Packages: `pybricks` <br>
Structure:
```
./python/
├── README.md  documentation
├── robot.py   function definitions
└── main.py    executable
```

## GO
Version: `go, linux on ARM5` <br>
Packages: `github.com/ev3go/ev3dev` <br>
Structure:
```
./golang/
├── ...        additional files
├── Makefile
├── main       executable
├── robot.go   function definitions
└── main.go    main function
```
### Compiling
```sh
make
# or
make build
```
