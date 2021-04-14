from pybricks import ev3brick as brick
from pybricks.ev3devices import Motor, GyroSensor, UltrasonicSensor
from pybricks.parameters import Port, Stop
from pybricks.robotics import DriveBase
from time import sleep, time

# TODO: write code for dropper module
# TODO: write code for stick module

# contains methods for modules connected to the shifter
class Shifter:
    # initialize Shifter object
    def __init__(self, sm, rm, angs=[0, 150, 300, 450]):
        self.sm = Motor(sm)
        self.rm = Motor(rm)
        self.angs = angs

    # general module function :
    def run(self, index, spd, ang):
        self.sm.run_target(150, self.angs[index])
        self.rm.run_target(spd, ang)

    def dropper(self):
        # code for dropper module
        pass

    def stick(self):
        # code for stick module
        pass

# contains methods for movement
class Robot:
    # initialize Robot object
    def __init__(self, leftMotor, rightMotor, gyro, sonic, wheels):
        # hardware connections
        self.leftMotor = Motor(leftMotor)
        self.rightMotor = Motor(rightMotor)
        self.gyro = GyroSensor(gyro)
        self.sonic = UltrasonicSensor(sonic)
        self.gyro.reset_angle(0)
        self.robot = DriveBase(self.leftMotor, self.rightMotor, wheels[0], wheels[1])

    # parabolic velocity - velocity with respect to distance measured
    def velocity(self, currentDist, startDist, endDist, maxVel):
        s = (startDist - endDist)**2
        # quadratic coeficient
        a = (-4 * maxVel) / s
        # linear coeficient
        b = ((4 * maxVel) / s) * (startDist + endDist)
        # constant
        c = -((4 * startDist * endDist * maxVel) / s)
        return (a * (currentDist**2)) + (b * currentDist) + c

    # rotate in place ; spd : deg/s, ang : degs
    def rotate(self, spd=40, ang=0):
        while(True):
            if(self.gyro.angle() > ang):
                self.rightMotor.run(spd)
                self.leftMotor.run(spd * -1)
            elif(self.gyro.angle() < ang):
                    self.leftMotor.run(spd)
                    self.rightMotor.run(spd * -1)
            else:
                self.leftMotor.stop(Stop.HOLD)
                self.rightMotor.stop(Stop.HOLD)
                sleep(0.5)
                if(self.gyro.angle() == ang):
                    self.gyro.reset_angle(0)
                    break

    # move robot to specified distance ; spd : "mm/s", dist : cm, err : mm
    def move(self, maxVel=100, endDist=20, error=10, fixConst=1):
        # convert distance from cm to mm
        data = []
        endDist = endDist * 10
        startDist = self.sonic.distance() - 1
        while(True):
            currentDist = self.sonic.distance()
            if(currentDist < endDist - error or currentDist > endDist + error):
                originalAng = self.gyro.angle()
                vel = self.velocity(currentDist, startDist, endDist, maxVel)
                self.robot.drive(
                    self.velocity(currentDist, startDist, endDist, maxVel),
                    self.gyro.angle() * -(fixConst)
                )
                data.append((time(), originalAng + (self.gyro.angle() * -(fixConst)), vel, currentDist))
            else:
                self.robot.stop()
                self.gyro.reset_angle(0)
                sleep(0.5) # wait to ensure the measurement is correct
                if(self.sonic.distance() < endDist + error or self.sonic.distance() > endDist - error):
                    break
            return data, "t Δα v μs"

def saveData(data, types):
    rawFile = open("data.dat", "r")
    number = (rawFile.readlines()).count("%")
    rawFile.close()
    file = open("data.dat", "a")
    write = ("% " + str(number) + " " + types + "\n")
    for point in data:
        printed = ""
        for val in point:
            printed += (val + " ")
        write += (printed + "\n")
    write += ("$\n")
    file.write(write)
    file.close()
