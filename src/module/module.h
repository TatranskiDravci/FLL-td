#ifndef MODULE_H
#define MODULE_H

#include "../shifter/shifter.h"

/*
    await asynchronous drive (call after moduleAsyncDrive(...) to wait for it to finish) 
        s - `shifter` object
*/
#define moduleAwaitDrive(s) shifterAwaitDrive(s)

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
*/
void moduleAsyncDrive(shifter s, int target, int id);


#endif