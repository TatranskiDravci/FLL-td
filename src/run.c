#include "base/base.h"
#include "base/move.h"
#include <stdio.h>
#include <unistd.h>


int main(void)
{
    base b;
    b = baseNew('D', 'A', '2');

    getc(stdin);

    int error;
    error = 0;

    for (int i = 0; i < 4; i++)
    {
        rotate(b, -90, 400);
        sleep(1);
        error += sensorRead(b.gyro, '0');
        printf("%d\n", error);
    }

    return 0;
}
