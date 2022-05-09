#include "base/color.h"
#include <stdio.h>

int main(void)
{
    color cs;
    cs = colorNew('3');

    colorProfileLoad(&cs, "../data/profile_k_S", "../data/profile_l_S");
    while (1) printf("%f\n", colorRead(cs));

    return 0;
}
