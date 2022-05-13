#ifndef MODULE_H
#define MODULE_H

#include "shifter.h"

// module constants
#define DRIVE_RATE 700            // rate of `dmotor` on `shifter`

/*
    synchronous module drive
        s      - `shifter` object
        target - module target
        id     - module id
*/
#define moduleDrive(s, target, id) \
    ({ moduleAsyncDrive(s, target, id); \
       moduleAwaitDrive(s); })

/*
    drive modules asynchronously
        s      - `shifter` object
        target - module angle target
        id     - module id
*/
void moduleAsyncDrive(shifter s, int target, int id);

/*
    await asynchronous drive (call after moduleAsyncDrive(...) to wait for it to finish)
        s - `shifter` object
*/
void moduleAwaitDrive(shifter s);

#endif
