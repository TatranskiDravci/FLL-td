#ifndef SHIFTER_H
#define SHIFTER_H

#include "../drivers/motor.h"

// shifter constants
#define SHIFTER_RATE 400
#define SHIFTER_OFFSET 90
#define DRIVE_RATE 700            // rate of `dmotor` on `shifter` 

/*
    shifter interface
        smotor - shifter motor
        dmotor - drive   motor
*/
typedef struct Shifter
{
    motor smotor, dmotor;
}
shifter;

/*
    `shifter` constructor
        sport - shifter motor port
        dport - drive   motor port
*/
shifter shifterNew(char sport, char dport);

/*
    shift
        s  - `shifter` object
        id - module id
*/
void shifterShift(shifter s, int id);

/*
    drive modules asynchronously
        s      - `shifter` object
        target - module angle target
*/
void shifterAsyncDrive(shifter s, int target);

/*
    await asynchronous drive (call after shifterAsyncDrive(...) to wait for it to finish) 
        s - `shifter` object
*/
void shifterAwaitDrive(shifter s);

#endif