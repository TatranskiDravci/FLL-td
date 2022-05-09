/*
    devnote:
        Keep fwd. and bwd. directions as well as non-stop
        and standard functions unified in one function, if possible.

        p.v. - possible values

    todo:
        - find fitting STEERING_MOD for bang-bang
        - test lf_mod-s
            - if the robot steers in the opposite direction than desired,
              make LBRW 1 and LWRB -1
*/
#ifndef MOVE_H
#define MOVE_H

#include <time.h>
#include "base.h"
#include "pid.h"
#include "color.h"

// define fn_type-s
#define NS_INI 0b01         // initial non-stop function   (no stoppage)
#define NS_FIN 0b10         // final non-stop function     (stoppage)
#define NS_MID 0b00         // midrun non-stop function    (no stoppage)
#define NS_STD 0b11         // standard movement function  (immediate stoppage)

// define directions
#define FWD  1              // move forwards
#define BWD -1              // move backwards

// speed setpoints for rotation
#define FIX_SPEED 75        // rotation correction speed after move functions
#define MIN_SPEED 50        // minimum speed for rotation

// line following modifiers (lf_mod-s + steering modifier)
#define LBRW  1             // black line on the left, white line on the right
#define LWRB -1             // white line on the left, black line on the right

#define STEERING_MOD 20     // modifies the agressivness of steering when following lines (bang-bang agression constant)

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
        duration  - the duration of movement in seconds
        *ctl      - pointer to pid object (use `&name_of_pid_object` when calling this function)
        direction - direction in which the movement is to be performed (p.v. FWD (also 1), and BWD (also -1))
        fn_type   - type of movement function to be ran (p.v. NS_INI, NS_FIN, NS_MID, NS_STD)
*/
void moveTimed(base b, int speed, double duration, pid *ctl, int direction, int fn_type);

/*
    color stop-action movement function
        b         - base object
        speed     - maximum speed on the motors in u/s
        cs        - color sensor object
        value     - color value
        delta     - allowed color error
        *ctl      - pointer to pid object (use `&name_of_pid_object` when calling this function)
        direction - direction in which the movement is to be performed (p.v. FWD (also 1), and BWD (also -1))
        fn_type   - type of movement function to be ran (p.v. NS_INI, NS_FIN, NS_MID, NS_STD)
*/
void moveColor(base b, int speed, color cs, double value, double delta, pid *ctl, int direction, int fn_type);

/*
    color stop-action line guided movement function (cannot be used as non-stop function in combination with non-moveLine functions)
        b         - base object
        speed     - maximum speed on the motors in u/s
        cs_f      - color sensor object for line following
        cs_s      - color sensor object for stopping
        value     - color value
        delta     - allowed color error
        course    - course when stopping (for angle correction)
        lf_mod    - line following modifier (modifies which line is to the left and which to the right of the sensor)
        direction - direction in which the movement is to be performed (p.v. FWD (also 1), and BWD (also -1))
        fn_type   - type of movement function to be ran (p.v. NS_INI, NS_FIN, NS_MID, NS_STD)
*/
void moveLine(base b, int speed, color cs_f, color cs_s, double value, double delta, int course, int lf_mod, int direction, int fn_type);

/*
    simple rotation (rotation axis = centre of the wheel axle)
        b      - base object
        target - target angle in degrees
        speed  - maximum speed in u/s
*/
void rotate(base b, int target, int speed);

#endif
