package main

/*
    devnote:
        USEFUL FUNCTIONS:
            @ shifter.go
                - (Shifter).BeginShifting(id)
                - (Shifter).AwaitShifting()
                - (Shifter).DriveAbsolute(target, rate)
                - (Shifter).DriveRelative(target, rate)
*/

/*
    general Module interface
        shifter - Shifter object
*/
type Module struct {
    shifter Shifter
}

/*
    Module constructor
        shifter - Shifter object
*/
func NewModule(shifter Shifter) Module {
    return Module {
        shifter: shifter,
    }
}

/*
    provides Drive for synchronous module control
        target  - target angle on driver motor
        id      - module's id on the shifter
*/
func (this Module) Drive(target, id int) {
    this.shifter.BeginShifting(id)
    this.shifter.AwaitShifting()
    this.shifter.DriveRelative(target, 500)
}

/*
    provides Drive for asynchronous module control
        target  - target angle on driver motor
        id      - module's id on the shifter
*/
func (this Module) BeginDrive(target, id int) {
    this.shifter.BeginShifting(id)
    this.shifter.AwaitShifting()
    this.shifter.BeginDriveRelative(target, 500)
}

/*
    provides await "operator" for Drive for asynchronous module control
*/
func (this Module) AwaitDrive() {
    this.shifter.AwaitDriveRelative()
}
