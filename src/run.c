#include "base/base.h"
#include "base/move.h"
#include "base/pid.h"
#include "base/color.h"
#include "shifter/shifter.h"
#include "module/module.h"
#include <stdio.h>

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

    moveTimed(b, 600, 1.5, &ctl, FWD, NS_STD);

    moveLine(b, 600, cs_f, cs_s, 100.0, 7.0, -41, FWD, LBRW, NS_STD);

    moduleDrive(s, -450, 1);

    moduleDrive(s, 400, 3);

    return 0;
}