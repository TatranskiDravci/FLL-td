#include "module.h"

void moduleAsyncDrive(shifter s, int target, int id)
{
    shifterShift(s, id);
    shifterAsyncDrive(s, target);
}

void moduleAsyncDriveModularSpeed(shifter s, int target, int speed, int id)
{
    shifterShift(s, id);
    shifterAsyncDriveModularSpeed(s, target, speed);
}