#include "move.h"
#include <time.h>
#include <unistd.h>
#include <math.h>

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
        else if (angle > target) baseRunTank(b, mod_speed, -mod_speed);
	    else baseRunTank(b, -mod_speed, mod_speed);
	}
}

void moveTimed(base b, int speed, double duration, pid *ctl, int direction, int fn_type)
{
    // initialize non-stop movement
    if (fn_type & NS_INI)
    {
        baseResetGyro(b);

        struct timespec measured_time;
        clock_gettime(CLOCK_REALTIME, &measured_time);

        ctl->stime = NANO * measured_time.tv_nsec;
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

    while (1)
    {
        // course correction
        double ctime, dtime;
        double error, x;

        struct timespec measured_time;
        clock_gettime(CLOCK_REALTIME, &measured_time);

        error = ctl->SP - sensorReadDecimal(b.gyro, '0');
        ctime = NANO * measured_time.tv_nsec - stime;       // elapsed time

        dtime = ctime - ptime;                              // time delta
        ptime = ctime;

        // calculate steering parameter
        integral += 0.5 * dtime * (perror + error);
        x = ctl->KP * error + ctl->KI * integral; // * ctl->KD * (error - perror) / dtime;
        baseRunSteering(b, speed, -x * direction, direction);

        perror = error;

        // breakpoint
        if (ctime >= duration) break;
    }

    ctl->ptime = ptime;
    ctl->perror = perror;
    ctl->integral = integral;

    if (fn_type & NS_FIN)
    {
        baseStop(b);
        rotate(b, -sensorRead(b.gyro, '0'), FIX_SPEED);
    }
}