/*
    devnote:
        base.h should provide only bare-bones movement functions,
        all more complex functions shoud be be reserved for move.h.

        p.v. - possible values
*/
#ifndef BASE_H
#define BASE_H

#include <time.h>
#include "../drivers/motor.h"
#include "../drivers/sensor.h"

// define arithmetic function max()
#define max(a, b) \
    ({ __typeof__ (a) _a = (a); \
        __typeof__ (b) _b = (b); \
        _a > _b ? _a : _b; })

// define arithmetic function min()
#define min(a, b) \
    ({ __typeof__ (a) _a = (a); \
        __typeof__ (b) _b = (b); \
        _a < _b ? _a : _b; })

// unit conversions
#define NANO 0.000000001    // nano to unit conversion constant

/*
    base interface for robot movement control
        lmotor - left  motor object
        rmotor - right motor object
        gyro   - gyro sensor object
*/
typedef struct Base
{
    motor  lmotor, rmotor;
    sensor gyro;
}
base;

/*
    `base` constructor - creates base object
        lport    - left  motor port (p.v. 'A', 'B', 'C', 'D')
        rport    - right motor port (p.v. 'A', 'B', 'C', 'D')
        gyroport - gyro sensor port (p.v. '1', '2', '3', '4')
        @return  - base object
*/
base baseNew(char lport, char rport, char gyroport);

/*
    gyro reset function
        b - base object
*/
void baseResetGyro(base b);

/*
    tank-like movement - allows control of speeds of individual motors
        b      - base object
        lspeed - left  motor speed
        rspeed - right motor speed
*/
void baseRunTank(base b, int lspeed, int rspeed);

/*
    steering movement - compound control of motor speeds using steering parameter
        b         - base object
        speed     - maximum speed
        x         - steering parameter (usually in a form of pid error; when 0, the robot should not steer)
        direction - movement direction (p.v. 1 (also FWD from move.h), and -1 (also BWD from move.h))

*/
void baseRunSteering(base b, double speed, double x, int direction);

/*
    stop motors
        b - base object
*/
void baseStop(base b);

/*
    returns epoch timestamp in seconds with nanosecond accuracy
        @return - time in seconds
*/
double timeSeconds(void);

#endif
