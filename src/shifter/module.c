#include "module.h"

void moduleAsyncDrive(shifter s, int target, int id)
{
    shifterShift(s, id);
    motorSetSpeed(s.dmotor, DRIVE_RATE);
    motorSetTarget(s.dmotor, target);
    motorCommand(s.dmotor, "run-to-rel-pos");
}

void moduleAwaitDrive(shifter s)
{
    while (motorState(s.dmotor) & RUNNING);
}
