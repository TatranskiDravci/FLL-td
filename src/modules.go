package main

/*
    TODO:
        [ ] implement modules (OOP)

    STRUCTURE:

    type ModuleName struct {
        shifter Shifter

        field ...
    }

    func (this ModuleName) FunctionName(args... type) {
        this.shifter.BeginShifting(ID)
        this.shifter.AwaitShifting()

        implementation ...
    }

    USEFUL FUNCTIONS:
        @ shifter.go
            - (Shifter).BeginShifting(id)
            - (Shifter).AwaitShifting()
            - (Shifter).DriveAbsolute(target, rate)
            - (Shifter).DriveRelative(target, rate)
*/
// type M1S0 struct {
//     shifter Shifter
// }
//
// func NewM1S0(shifter Shifter) M1S0 {
//     return M1S0 {
//         shifter: shifter,
//     }
// }
//
// func (this M1S0) Function() {
//     this.shifter.BeginShifting(0)
//     this.shifter.AwaitShifting()
//     this.shifter.DriveRelative(530, 400)
// }

type M1S0 struct {
    shifter Shifter
}

func NewM1S0(shifter Shifter) M1S0 {
    return M1S0 {
        shifter: shifter,
    }
}

func (this M1S0) Extend() {
    this.shifter.BeginShifting(0)
    this.shifter.AwaitShifting()
    this.shifter.DriveRelative(-1700, 500)
}

func (this M1S0) ParitalExtend() {
    this.shifter.BeginShifting(0)
    this.shifter.AwaitShifting()
    this.shifter.DriveRelative(-400, 500)
}

func (this M1S0) Retract() {
    this.shifter.BeginShifting(0)
    this.shifter.AwaitShifting()
    this.shifter.DriveRelative(1700, 500)
}

func (this M1S0) BeginRetract() {
    this.shifter.BeginDriveRelative(500, 500)
}

func (this M1S0) AwaitRetract() {
    this.shifter.AwaitDriveRelative()
}

type M1S1 struct {
    shifter Shifter
}

func NewM1S1(shifter Shifter) M1S1 {
    return M1S1 {
        shifter: shifter,
    }
}

func (this M1S1) Flip() {
    this.shifter.BeginShifting(1)
    this.shifter.AwaitShifting()
    this.shifter.DriveRelative(900, 500)
}

type M1S3 struct {
    shifter Shifter
}

func NewM1S3(shifter Shifter) M1S3 {
    return M1S3 {
        shifter: shifter,
    }
}

func (this M1S3) Open() {
    this.shifter.BeginShifting(3)
    this.shifter.AwaitShifting()
    this.shifter.DriveRelative(600, 500)
}

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
    this.shifter.DriveRelative(target, 400)
}

func (this Module) BeginDrive(target, id int) {
    this.shifter.BeginShifting(id)
    this.shifter.AwaitShifting()
    this.shifter.BeginDriveRelative(target, 400)
}

func (this Module) AwaitDrive() {
    this.shifter.AwaitDriveRelative()
}
