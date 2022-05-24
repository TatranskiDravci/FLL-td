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

    b = baseNew('D', 'A', '1');
    ctl = pidNew(0.0, 25.0, 25.0, 0.0);
    s = shifterNew('B', 'C');
    cs_m = colorNew('4');
    cs_l = colorNew('3');

    colorProfileLoad(&cs_m, "../data/profile_k_F", "../data/profile_l_F");
    colorProfileLoad(&cs_l, "../data/profile_k_S", "../data/profile_l_S");

    getc(stdin);
    

    // get to line
    moveTimed(b, 300, 0.5, &ctl, FWD, NS_STD);
    rotate(b, -45, 400);
    moveColor(b, 700, cs_l, 0.0, 7.0, &ctl, FWD, NS_INI);
    moveColor(b, 200, cs_l, 100.0, 20.0, &ctl, FWD, NS_FIN);
    moveTimed(b, 200, 1.8, &ctl, FWD, NS_STD);
    rotate(b, -39, 400);
    moduleAwaitDrive(s);
    moduleAsyncDrive(s, -900, 3); // retract crane arm

    // bridge
    moveTimed(b, 700, 1.3, &ctl, FWD, NS_INI);
    moveColor(b, 700, cs_m, 35.0, 5.0, &ctl, FWD, NS_MID);
    moveColor(b, 700, cs_m, 0.0, 5.0, &ctl, FWD, NS_FIN);
    moduleAwaitDrive(s);
    
    // container
    moveTimed(b, 500, 1.4, &ctl, FWD, NS_STD);
    moduleDrive(s, -120, 1);
    moveTimed(b, 500, 0.55, &ctl, BWD, NS_STD);

    // inno project
    moduleDrive(s, -120, 2);
    sleep(0.3);
    moduleDrive(s, 110, 2);
    moduleDrive(s, 30, 2);

    //moveTimed(b, 500, 0.3, &ctl, BWD, NS_STD);
    moveColor(b, 500, cs_m, 0.0, 7.0, &ctl, BWD, NS_STD);
    rotate(b, -94, 400);
    moveTimed(b, 500, 1.0, &ctl, BWD, NS_STD);

    // crane
    moduleAsyncDrive(s, 900, 3);
    rotate(b, 0, FIX_SPEED);
    moduleAwaitDrive(s);

    // truck
    moveTimed(b, 500, 1.0, &ctl, FWD, NS_STD);
    moveTimed(b, 500, 0.35, &ctl, BWD, NS_STD);
    rotate(b, 90, 400);
    moveColor(b, 500, cs_l, 100.0, 20.0, &ctl, BWD, NS_INI);
    moveColor(b, 500, cs_l, 0.0, 5.0, &ctl, BWD, NS_FIN);
    moveTimed(b, 300, 0.5, &ctl, BWD, NS_STD);
    moduleDrive(s, 360, 0);
    moveTimed(b, 700, 0.6, &ctl, BWD, NS_STD);
    moveTimed(b, 500, 0.5, &ctl, FWD, NS_STD);

    // take robot from mat by hand

    shifterShift(s, 0);

    return 0;
}
