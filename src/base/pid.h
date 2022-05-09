#ifndef PID_H
#define PID_H

/*
    general PID interface
        SP - setpoint
        KP - proportional coefficient
        KI - integral     coefficient
        KD - differential coefficient

        integral - transfered integral value
        perror   - transfered previous error value
        ptime    - transfered previous time
        stime    - transfered time on move initialization
*/
typedef struct PID
{
    // constant block
    double SP, KP, KI, KD;
    // variable block
    double integral, perror, ptime, stime;
}
pid;

/*
    `pid` constructor
        SP      - setpoint
        KP      - proportional coefficient
        KI      - integral     coefficient
        KD      - differential coefficient
        @return - pid object
*/
pid pidNew(double SP, double KP, double KI, double KD);

#endif
