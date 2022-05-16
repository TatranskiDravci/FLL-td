#include "move.h"
#include <time.h>
#include <unistd.h>
#include <math.h>
#include <stdio.h>

int modSpeed(double target, double angle, double speed)
{
    double dangle, dspeed;
    dangle = fabs(target - angle);
    dspeed = speed - MIN_SPEED;

    return (int) (min(0.005556 * dspeed * dangle + MIN_SPEED, speed));
}

void rotate(base b, int target, int speed)
{
    baseResetGyro(b);

    while (1)
    {
        int angle, mod_speed;
        angle = sensorRead(b.gyro, '0');
        mod_speed = modSpeed(target, angle, speed);

        if (angle == target)
        {
            baseStop(b);
            sleep(0.1);

            // "2x meraj a 1x rež"
            if (sensorRead(b.gyro, '0') == target) return;
        }
        else if (angle > target) baseRunTank(b,  mod_speed, -mod_speed);
        else                     baseRunTank(b, -mod_speed,  mod_speed);
    }
}

void moveTimed(base b, int speed, double duration, pid *ctl, int direction, int fn_type)
{
    // initialize non-stop movement
    if (fn_type & NS_INI)
    {
        baseResetGyro(b);

        ctl->stime = timeSeconds();
        ctl->ptime = 0.0;
        ctl->perror = 0.0;
        ctl->integral = 0.0;
    }

    double stime, ptime, ctime;
    double perror, integral;
    stime = ctl->stime;
    ptime = ctl->ptime;
    perror = ctl->perror;
    integral = ctl->integral;
    baseRunSteering(b, speed, 0, direction);

    while ((ctime = timeSeconds() - stime) < duration)
    {
        // course correction
        double dtime;
        double error, x;

        error = ctl->SP - sensorReadDecimal(b.gyro, '0');

        dtime = ctime - ptime;                              // time delta
        ptime = ctime;

        // calculate steering parameter
        integral += 0.5 * dtime * (perror + error);
        x = ctl->KP * error + ctl->KI * integral; // * ctl->KD * (error - perror) / dtime;

        baseRunSteering(b, speed, -x * direction, direction);

        perror = error;
    }

    ctl->ptime = ptime;
    ctl->perror = perror;
    ctl->integral = integral;

    // finalize non-stop movement
    if (fn_type & NS_FIN)
    {
        baseStop(b);
        rotate(b, -sensorRead(b.gyro, '0'), FIX_SPEED);
    }
}

void moveColor(base b, int speed, color cs, double value, double delta, pid *ctl, int direction, int fn_type)
{
    // initialize non-stop movement
    if (fn_type & NS_INI)
    {
        baseResetGyro(b);

        ctl->stime = timeSeconds();
        ctl->ptime = 0;
        ctl->perror = 0.0;
        ctl->integral = 0.0;
    }

    double stime, ptime;
    double perror, integral;
    stime = ctl->stime;
    ptime = ctl->ptime;
    perror = ctl->perror;
    integral = ctl->integral;
    baseRunSteering(b, speed, 0, direction);

    while (fabs(colorRead(cs) - value) > delta)
    {
        // course correction
        double ctime, dtime;
        double error, x;

        error = ctl->SP - sensorReadDecimal(b.gyro, '0');
        ctime = timeSeconds() - stime;                      // elapsed time

        dtime = ctime - ptime;                              // time delta
        ptime = ctime;

        // calculate steering parameter
        integral += 0.5 * dtime * (perror + error);
        x = ctl->KP * error + ctl->KI * integral; // * ctl->KD * (error - perror) / dtime;

        baseRunSteering(b, speed, -x * direction, direction);

        perror = error;
    }

    ctl->ptime = ptime;
    ctl->perror = perror;
    ctl->integral = integral;

    // finalize non-stop movement
    if (fn_type & NS_FIN)
    {
        baseStop(b);
        rotate(b, -sensorRead(b.gyro, '0'), FIX_SPEED);
    }
}

void moveLine(base b, int speed, color cs_f, color cs_s, double value, double delta, int course, int lf_mod, int direction, int fn_type)
{
    // initialize non-stop movement
    if (fn_type & NS_INI)
    {
        baseResetGyro(b);
    }

    baseRunSteering(b, speed, 0, direction);

    while (fabs(colorRead(cs_s) - value) > delta)
    {
        if (colorRead(cs_f) <= 35.0) baseRunSteering(b, speed,  lf_mod * STEERING_MOD * speed, direction);
        else                         baseRunSteering(b, speed, -lf_mod * STEERING_MOD * speed, direction);
    }

    // finalize non-stop movement
    if (fn_type & NS_FIN)
    {
        baseStop(b);
        rotate(b, course - sensorRead(b.gyro, '0'), FIX_SPEED);
    }
}
