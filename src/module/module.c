#include "module.h"

void moduleAsyncDrive(shifter s, int target, int id)
{
    shifterShift(s, id);
    shifterAsyncDrive(s, target);
}