CC = arm-linux-gnueabi-gcc

# build all targets
all: run calib colors

run: run.c sensor.c motor.c shifter.c base.c module.c move.c pid.c color.c
	$(CC) $^ -o build/run

calib: calib.c sensor.c color.c
	$(CC) $^ -o build/calib

colors: colors.c sensor.c color.c
	$(CC) $^ -o build/colors

vpath %.c src
vpath %.c src/drivers
vpath %.c src/base
vpath %.c src/shifter
vpath %.c src/module

clean:
	rm build/*
	rm data/*

init:
	mkdir build
	mkdir data
	touch data/placeholder		# required for folder creation on the robot
