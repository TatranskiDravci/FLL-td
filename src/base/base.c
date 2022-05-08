#include "base.h"
#include <stdio.h>

base baseNew(char lport, char rport, char gyroport)
{
    base b;
    b.lmotor = motorNew(lport);
    b.rmotor = motorNew(rport);
    b.gyro   = sensorNew(gyroport);

    motorSetStopAction(b.lmotor, "hold");
    motorSetStopAction(b.rmotor, "hold");
    sensorSetMode(b.gyro, "GYRO-RATE");
    sensorSetMode(b.gyro, "GYRO-ANG");

    while (sensorRead(b.gyro, '0'));        // wait until sensor gets recalibrated

    return b;
}

void baseResetGyro(base b)
{
    sensorSetMode(b.gyro, "GYRO-RATE");
    sensorSetMode(b.gyro, "GYRO-ANG");

    while (sensorRead(b.gyro, '0'));        // wait until sensor gets recalibrated
}

void baseRunTank(base b, int lspeed, int rspeed)
{
    motorSetSpeed(b.lmotor, lspeed);
    motorSetSpeed(b.rmotor, rspeed);
    motorCommand(b.lmotor, "run-forever");
    motorCommand(b.rmotor, "run-forever");
}

void baseRunSteering(base b, double speed, double x, int direction)
{
    motorSetSpeed(b.lmotor, direction * (int) max(min( x + speed, speed), -speed));
    motorSetSpeed(b.rmotor, direction * (int) max(min(-x + speed, speed), -speed));
    motorCommand(b.lmotor, "run-forever");
    motorCommand(b.rmotor, "run-forever");
}

void baseStop(base b)
{
    motorCommand(b.lmotor, "stop");
    motorCommand(b.rmotor, "stop");
}

double timeSeconds()
{
    struct timespec measured_time;
    clock_gettime(CLOCK_REALTIME, &measured_time);
    return (double) measured_time.tv_sec + NANO * measured_time.tv_nsec;
}