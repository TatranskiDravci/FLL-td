#ifndef SHIFTER_H
#define SHIFTER_H

#include "../drivers/motor.h"

// shifter constants
#define SHIFTER_RATE 400
#define SHIFTER_OFFSET 90
#define SHIFTER_CORRECT -3

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

#endif
