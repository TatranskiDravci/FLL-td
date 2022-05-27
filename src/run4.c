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
    color cs_m, cs_l;

    b = baseNew('D', 'A', '2');
    ctl = pidNew(0.0, 25.0, 25.0, 0.0);
    s = shifterNew('B', 'C');
    cs_m = colorNew('4');
    cs_l = colorNew('3');

    colorProfileLoad(&cs_m, "../data/profile_k_F", "../data/profile_l_F");
    colorProfileLoad(&cs_l, "../data/profile_k_S", "../data/profile_l_S");

    // nudge fwd
    moveTimed(b, 500, 0.85, &ctl, FWD, NS_STD);
    rotate(b, -90, 400);

    // get to missions
    moveTimed(b, 850, 3, &ctl, FWD, NS_INI);
    moveColor(b, 850, cs_l, 0.0, 8.0, &ctl, FWD, NS_FIN);
    moveTimed(b, 300, 0.3, &ctl, BWD, NS_STD);
    rotate(b, 90, 400);
    moveTimed(b, 500, 0.8, &ctl, BWD, NS_STD);

    // pckg
    moduleDrive(s, -180, 2);

    // get to wing
    moveTimed(b, 500, 0.52, &ctl, FWD, NS_STD);
    rotate(b, -90, 400);
    moveTimed(b, 500, 0.2, &ctl, BWD, NS_STD);
    rotate(b, 210, 400);
    moveTimed(b, 500, 0.4, &ctl, FWD, NS_STD);
    rotate(b, -30, 400);
    moveTimed(b, 500, 0.2, &ctl, BWD, NS_STD);
    moduleDrive(s, -1200, 0);
    moveTimed(b, 500, 0.3, &ctl, FWD, NS_STD);
    rotate(b, -30, 400);
    moveTimed(b, 500, 0.6, &ctl, FWD, NS_STD);
    moduleDrive(s, -180, 3);


    shifterShift(s, 0);

    return 0;
}
