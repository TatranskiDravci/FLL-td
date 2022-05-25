CC = arm-linux-gnueabi-gcc
DRIVERS = sensor.c motor.c
INCLUDE = shifter.c base.c module.c move.c pid.c color.c

# build all targets
all: run1 run2 run3 run4 calib colors

run1: run1.c $(DRIVERS) $(INCLUDE)
	$(CC) $^ -o build/run1

run2: run2.c $(DRIVERS) $(INCLUDE)
	$(CC) $^ -o build/run2

run3: run3.c $(DRIVERS) $(INCLUDE)
	$(CC) $^ -o build/run3

run4: run4.c $(DRIVERS) $(INCLUDE)
	$(CC) $^ -o build/run4

run: run.c $(DRIVERS) $(INCLUDE)
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

clear:
	rm build/*
	rm data/*

init:
	mkdir build
	mkdir data
	touch data/placeholder		# required for folder creation on the robot
