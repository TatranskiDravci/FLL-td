CC = arm-linux-gnueabi-gcc

run: run.c sensor.c motor.c shifter.c base.c module.c move.c pid.c color.c
	$(CC) $^ -o build/run

vpath %.c src
vpath %.c src/drivers
vpath %.c src/base
vpath %.c src/shifter
vpath %.c src/module

clean:
	rm build/*

init:
	mkdir build
