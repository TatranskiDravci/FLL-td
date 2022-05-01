#ifndef PID_H
#define PID_H

/*
    general PID interface
        SP - setpoint
        KP - proportional coefficient
        KI - integral     coefficient
        KD - differential coefficient

        integral - transfered integral value
        perror   - transfered preceeding error value
        ptime    - transfered preceeding time
        stime    - transfered time on move initialization
*/
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

/*
    pid constructor
        SP      - setpoint
        KP      - proportional coefficient
        KI      - integral     coefficient
        KD      - differential coefficient
        @return - pid object
*/
pid pidNew(double SP, double KP, double KI, double KD);

#endif