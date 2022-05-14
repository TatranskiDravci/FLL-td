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


    //get to line
    moveTimed(b, 300, 0.4, &ctl, FWD, NS_STD);
    rotate(b, -50, 400);
    moveColor(b, 700, cs_s, 0, 7, &ctl, FWD, NS_INI);
    moveColor(b, 200, cs_s, 100, 20, &ctl, FWD, NS_FIN);
    moveTimed(b, 200, 2, &ctl, FWD, NS_STD);
    rotate(b, -38, 400);

    //move on line
    moveColor(b, 700, cs_f, 0, 8, &ctl, FWD, NS_INI);
    moveColor(b, 700, cs_f, 100, 20, &ctl, FWD, NS_MID);
    moveColor(b, 700, cs_f, 0, 8, &ctl, FWD, NS_FIN);
    moveTimed(b, 500, 2, &ctl, FWD, NS_STD);
    moveColor(b, 500, cs_s, 0, 8, &ctl, FWD, NS_STD);

    //put down boxes
    moduleDrive(s, -3960, 0);

    //go back
    moveTimed(b, 500, 0.3, &ctl, BWD, NS_STD);

    //turn left and do the heli
    moduleAsyncDrive(s, 3960, 0);
    rotate(b, 88, 400);
    moveTimed(b, 500, 0.8, &ctl, FWD, NS_STD);
    rotate(b, -60, 400);
    moveTimed(b, 500, 0.7, &ctl, FWD, NS_STD);
    moveTimed(b, 500, 0.6, &ctl, BWD, NS_STD);
    rotate(b, 60, 400);
    moduleAwaitDrive(s);

    //put down forks
    moduleDrive(s, -900, 1);
    
    //lower pusher stick
    moduleDrive(s, -400, 2);

    //reverse
    moveTimed(b, 500, 1, &ctl, BWD, NS_STD);

    //raise forks
    moduleDrive(s, 900, 1);

    //adjust push stick
    moduleDrive(s, 45, 2);

    //reverse a bit more
    moveTimed(b, 500, 0.3, &ctl, BWD, NS_STD);



    shifterShift(s, 0);

    return 0;
}
