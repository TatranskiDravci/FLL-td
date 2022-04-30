#ifndef PID_H
#define PID_H

typedef struct PID
{
    double SP;
    double KP;
    double KI;
    double KD;

    double integral;
    double perror;
    double ptime;
    double stime; 
}
pid;

pid pidNew(double SP, double KP, double KI, double KD);

#endif