#ifndef BASE_H
#define BASE_H

#include "../drivers/motor.h"
#include "../drivers/sensor.h"

#define max(a, b) \
    ({ __typeof__ (a) _a = (a); \
        __typeof__ (b) _b = (b); \
        _a > _b ? _a : _b; })

#define min(a, b) \
    ({ __typeof__ (a) _a = (a); \
        __typeof__ (b) _b = (b); \
        _a < _b ? _a : _b; })

typedef struct Base
{
    motor  lmotor;
    motor  rmotor;
    sensor gyro;
}
base;

base baseNew(char lport, char rport, char gyroport);
void baseResetGyro(base b);
void baseRunTank(base b, int lspeed, int rspeed);
void baseRunSteering(base b, double maxspeed, double x, int direction);
void baseStop(base b);

#endif