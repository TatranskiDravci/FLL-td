#include "shifter.h"

shifter shifterNew(char sport, char dport)
{
    shifter s;
    s.smotor = motorNew(sport);
    s.dmotor = motorNew(dport);

    motorSetStopAction(s.smotor, "hold");
    motorSetStopAction(s.dmotor, "hold");

    motorReset(s.smotor);

    return s;
}

void shifterShift(shifter s, int id)
{
    motorSetSpeed(s.smotor, SHIFTER_RATE);
    motorSetTarget(s.smotor, id * SHIFTER_OFFSET);
    motorCommand(s.smotor, "run-to-abs-pos");

    while (motorState(s.smotor) & RUNNING);
}
