#include "base/base.h"
#include "base/move.h"
#include "base/pid.h"

int main(void)
{
    // initialize control interfaces
    base b; pid ctl;
    b = baseNew('A', 'D', '1');
    ctl = pidNew(0.0, 25.0, 15.0, 0.5);

    moveTimed(b, 300, 3.5, &ctl, FWD, NS_STD);  // move forwards with speed 300u/s for 3.5s with proper stoppage (NS_STD)
    rotate(b, 45, 200);                         // rotate 45 degrees counter-clockwise with max speed of 200u/s


    return 0;
}