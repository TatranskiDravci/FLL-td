#include "base/color.h"

int main(void)
{
    color cs_f, cs_s;
    cs_f = colorNew('4');
    cs_s = colorNew('3');

    colorProfileCalib(&cs_f, "../data/profile_k_F", "../data/profile_l_F");
    colorProfileCalib(&cs_s, "../data/profile_k_S", "../data/profile_l_S");

    return 0;
}
