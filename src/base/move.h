/*
    devnote:
        Keep fwd. and bwd. directions as well as non-stop
        and standard functions unified in one function, if possible.

        p.v. - possible values
*/
#ifndef MOVE_H
#define MOVE_H

#include <time.h>
#include "base.h"
#include "pid.h"

// define fn_type-s
#define NS_INI 0b01         // initial non-stop function   (no stoppage)
#define NS_FIN 0b10         // final non-stop function     (stoppage)
#define NS_MID 0b00         // midrun non-stop function    (no stoppage)
#define NS_STD 0b11         // standard movement function  (immediate stoppage)

// define directions
#define FWD  1              // move forwards
#define BWD -1              // move backwards

// speed setpoints for rotation
#define FIX_SPEED 55        // rotation correction speed after move functions
#define MIN_SPEED 25        // minimum speed for rotation

// unit conversions
#define NANO 0.000000001    // nano to unit conversion constant

/*
    provides speed modulation (constant accel. and deccel.) for rotate-like functions
        target  - target angle in degrees
        angle   - current angle in degrees
        speed   - maximum speed on the motors in u/s
        @return - appropriate speed for rotations based on difference of current and target angle
*/
int modSpeed(double target, double angle, double speed);

/*
    timed movement function
        b         - base object
        speed     - maximum speed on the motors in u/s
        duration  - the duration of movement in s
        *ctl      - pointer to pid object (use `&name_of_pid_object` when calling this function)
        direction - direction in which the movement is to be performed (p.v. FWD (also 1), and BWD (also -1))
        fn_type   - type of movement function to be ran (p.v. NS_INI, NS_FIN, NS_MID, NS_STD)
*/
void moveTimed(base b, int speed, double duration, pid *ctl, int direction, int fn_type);

/*
    simple rotation (rotation axis = centre of the wheel axle)
        b      - base object
        target - target angle in degrees
        speed  - maximum speed in u/s
*/
void rotate(base b, int target, int speed);

#endif
