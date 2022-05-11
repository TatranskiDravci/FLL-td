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

#define moduleDriveModularSpeed(s, target, speed, id) \
    ({ moduleAsyncDriveModularSpeed(s, target, speed, id); \
       moduleAwaitDrive(s); })

/*
    drive modules asynchronously
        s      - `shifter` object
        target - module angle target
        id     - module id
*/
void moduleAsyncDrive(shifter s, int target, int id);

/*
    drive modules asynchronously with modular speed
        s      - `shifter` object
        target - module angle target
        speed  - module drive rate
        id     - module id
*/
void moduleAsyncDriveModularSpeed(shifter s, int target, int speed, int id);


#endif