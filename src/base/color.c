#include "color.h"
#include <stdio.h>
#include <stdlib.h>

color colorNew(char sport)
{
    color cs;
    cs.s = sensorNew(sport);
    cs.k = 1.0;
    cs.l = 0.0;
    return cs;
}

void colorProfileLoad(color *cs)
{
    FILE *fp;
    char reading[50];

    fp = fopen("../data/profile_k", "r");
    fgets(reading, 50, fp);
    cs->k = atof(reading);
    fclose(fp);

    fp = fopen("../data/profile_l", "r");
    fgets(reading, 50, fp);
    cs->l = atof(reading);
    fclose(fp);
}

void colorProfileCalib(color *cs)
{
    double w, b;

    scanf("%s\n", NULL);
    for (int i = 0; i < 50; i++) w += sensorReadDecimal(cs->s);
    scanf("%s\n", NULL);
    for (int i = 0; i < 50; i++) b += sensorReadDecimal(cs->s);

    w *= 0.02;
    b *= 0.02;

    cs->k = 1.0 / (w - b);
    cs->l = -b;

    FILE *fp;

    fp = fopen("../data/profile_k", "w");
    fprintf(fp, "%f\n", cs->k);
    fclose(fp);

    fp = fopen("../data/profile_l", "w");
    fprintf(fp, "%f\n", cs->l);
    fclose(fp);
}

double colorRead(color cs)
{
    return cs.k * (sensorReadDecimal(cs.s, '0') + cs.l);
}