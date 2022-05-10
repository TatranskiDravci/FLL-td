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

void colorProfileLoad(color *cs, char *profile_k, char *profile_l)
{
    FILE *fp;
    char reading[50];

    fp = fopen(profile_k, "r");
    fgets(reading, 50, fp);
    cs->k = atof(reading);
    fclose(fp);

    fp = fopen(profile_l, "r");
    fgets(reading, 50, fp);
    cs->l = atof(reading);
    fclose(fp);
}

void colorProfileCalib(color *cs, char *profile_k, char *profile_l)
{
    double w, b;

    printf("scan white\n");
    getc(stdin);
    for (int i = 0; i < 50; i++) w += sensorReadDecimal(cs->s, '0');
    printf("scanned\n");

    printf("scan black\n");
    getc(stdin);
    printf("scanned\n");

    // read and sum some number of values
    for (int i = 0; i < 50; i++) b += sensorReadDecimal(cs->s, '0');

    // divide by number of values read
    w *= 0.02;
    b *= 0.02;

    cs->k = 100.0 / (w - b);
    cs->l = -b;

    FILE *fp;

    fp = fopen(profile_k, "w");
    fprintf(fp, "%f\n", cs->k);
    fclose(fp);

    fp = fopen(profile_l, "w");
    fprintf(fp, "%f\n", cs->l);
    fclose(fp);
}

double colorRead(color cs)
{
    return cs.k * (sensorReadDecimal(cs.s, '0') + cs.l);
}
