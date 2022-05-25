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

    getc(stdin);

    moveTimed(b, 500, 0.7, &ctl, FWD, NS_STD);
    rotate(b, -90, 400);
    moveTimed(b, 850, 3, &ctl, FWD, NS_INI);
    moveColor(b, 850, cs_l, 0.0, 8.0, &ctl, FWD, NS_FIN);

    shifterShift(s, 0);

    return 0;
}
