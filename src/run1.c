#include "base/base.h"
#include "base/move.h"
#include "base/pid.h"
#include "base/color.h"
#include "shifter/shifter.h"
#include "shifter/module.h"
#include <stdio.h>
#include <unistd.h>

int main(void)
{
    base b;
    pid ctl;
    shifter s;
    color cs_f, cs_s;

    b = baseNew('D', 'A', '2');
    ctl = pidNew(0.0, 25.0, 25.0, 0.0);
    s = shifterNew('B', 'C');
    cs_f = colorNew('4');
    cs_s = colorNew('3');

    colorProfileLoad(&cs_f, "../data/profile_k_F", "../data/profile_l_F");
    colorProfileLoad(&cs_s, "../data/profile_k_S", "../data/profile_l_S");

    getc(stdin);

    // get to position
    moveTimed(b, 600, 1.5, &ctl, FWD, NS_STD);
    moveLine(b, 500, cs_f, cs_s, 100.0, 8.0, -41, FWD, LBRW, NS_STD);
    moveColor(b, 100, cs_s, 0.0, 10.0, &ctl, FWD, NS_STD);

    // cargo plane
    moduleDrive(s, -540, 3);

    // motor setup
    moduleDrive(s,  180, 2);

    // plane
    moduleDrive(s,  550, 1);
    rotate(b, 30, 400);
    rotate(b, -30, 400);
    moduleDrive(s, -550, 1);

    // cargo plane stick up
    moduleDrive(s,  540, 3);

    // nudge motor
    moveTimed(b, 600, 0.3, &ctl, FWD, NS_STD);
    moduleDrive(s, -360, 2);

    // green container
    moduleDrive(s, -180, 0);

    // back
    moveTimed(b, 900, 1.0, &ctl, BWD, NS_STD);
    moveTimed(b, 900, 1.0, &ctl, BWD, NS_STD);

    shifterShift(s, 0);

    return 0;
}
