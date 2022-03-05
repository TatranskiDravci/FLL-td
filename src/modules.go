package main

/*
    USEFUL FUNCTIONS:
        @ shifter.go
            - (Shifter).BeginShifting(id)
            - (Shifter).AwaitShifting()
            - (Shifter).DriveAbsolute(target, rate)
            - (Shifter).DriveRelative(target, rate)
*/

type Module struct {
    shifter Shifter
}

func NewModule(shifter Shifter) Module {
    return Module {
        shifter: shifter,
    }
}

func (this Module) Drive(target, id int) {
    this.shifter.BeginShifting(id)
    this.shifter.AwaitShifting()
    this.shifter.DriveRelative(target, 500)
}

func (this Module) BeginDrive(target, id int) {
    this.shifter.BeginShifting(id)
    this.shifter.AwaitShifting()
    this.shifter.BeginDriveRelative(target, 500)
}

func (this Module) AwaitDrive() {
    this.shifter.AwaitDriveRelative()
}
