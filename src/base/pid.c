#include "pid.h"

pid pidNew(double SP, double KP, double KI, double KD)
{
    pid ctl;
    ctl.SP = SP;
    ctl.KP = KP;
    ctl.KI = KI;
    ctl.KD = KD;
    return ctl;
}