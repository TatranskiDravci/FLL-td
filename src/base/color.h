#ifndef COLOR_H
#define COLOR_H

#include "../drivers/sensor.h"

/*
    color sensor interface
        s - sensor object
        k - color calibration parameter
        l - color calibration parameter
*/
typedef struct Color
{
    sensor s;
    double k, l;
}
color;

/*
    color sensor constructor
        sport - color sensor port
*/
color colorNew(char sport);

/*
    loads color profile
        *cs - color sensor pointer (use `&color_sensor_object`)
*/
void colorProfileLoad(color *cs);

/*
    calibrates and loads color profile
        *cs - color sensor pointer (use `&color_sensor_object`)
*/
void colorProfileCalib(color *cs);

/*
    reads color sensor and normalizes it using k and l calibration parameter
        cs      - color sensor object
        @return - normalized color value
*/
double colorRead(color cs);

#endif