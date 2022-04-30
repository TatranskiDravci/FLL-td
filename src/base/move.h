#ifndef MOVE_H
#define MOVE_H

#include <time.h>
#include "base.h"
#include "pid.h"

// define fn_type-s
#define NS_INI 0b01         // initial non-stop function
#define NS_FIN 0b10         // final non-stop function
#define NS_MID 0b00         // midrun non-stop function
#define NS_STD 0b11         // standard movement function

// define directions
#define FWD  1              // move forwards
#define BWD -1              // move backwards

// speed setpoints for rotation
#define FIX_SPEED 55        // rotation correction speed after move functions
#define MIN_SPEED 25        // minimum speed for rotation

#define NANO 0.000000001

int modSpeed(double target, double angle, double speed);
void moveTimed(base b, int speed, double duration, pid *ctl, int direction, int fn_type);
void rotate(base b, int target, int speed);

#endif