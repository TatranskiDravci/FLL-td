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

    b = baseNew('D', 'A', '1');
    ctl = pidNew(0.0, 25.0, 25.0, 0.0);
    s = shifterNew('B', 'C');
    cs_f = colorNew('4');
    cs_s = colorNew('3');

    colorProfileLoad(&cs_f, "../data/profile_k_F", "../data/profile_l_F");
    colorProfileLoad(&cs_s, "../data/profile_k_S", "../data/profile_l_S");

    getc(stdin);
    
    //retract crane arm
    moduleAsyncDrive(s, -900, 3);
    
    //get to line
    moveTimed(b, 300, 0.8, &ctl, FWD, NS_STD);
    rotate(b, -50, 400);
    moveColor(b, 700, cs_s, 0, 7, &ctl, FWD, NS_INI);
    moveColor(b, 200, cs_s, 100, 20, &ctl, FWD, NS_FIN);
    moveTimed(b, 200, 1.6, &ctl, FWD, NS_STD);
    rotate(b, -36, 400);

    //bridge
    moveColor(b, 700, cs_s, 0, 5, &ctl, FWD, NS_INI);
    moveColor(b, 700, cs_s, 100, 20, &ctl, FWD, NS_MID);
    moveColor(b, 700, cs_s, 7, 5, &ctl, FWD, NS_MID);
    moveColor(b, 700, cs_f, 0, 5, &ctl, FWD, NS_FIN);
    moveTimed(b, 500, 0.6, &ctl, FWD, NS_STD);
    moveTimed(b, 500, 0.6, &ctl, BWD, NS_STD);
    moveTimed(b, 500, 0.8, &ctl, FWD, NS_STD);

    //inno project
    moduleDrive(s, -120, 2);

    moveTimed(b, 500, 0.5, &ctl, BWD, NS_STD);
    moveColor(b, 500, cs_f, 100, 20, &ctl, BWD, NS_STD);
    rotate(b, -94, 400);
    moveTimed(b, 500, 1, &ctl, BWD, NS_STD);



    //crane
    moduleDrive(s, 900, 3);
    moduleAsyncDrive(s, -900, 3);

    //truck
    moveTimed(b, 500, 0.6, &ctl, FWD, NS_STD);
    rotate(b, 90, 400);
    moduleAsyncDrive(s, 360, 0);
    moveTimed(b, 500, 2, &ctl, BWD, NS_STD);
    //take robot from mat by hand

    shifterShift(s, 0);

    return 0;
}
