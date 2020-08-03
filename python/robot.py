from pybricks.ev3devices import Motor, GyroSensor, UltrasonicSensor
from pybricks.parameters import Port, Stop
from pybricks.robotics import DriveBase
from time import sleep

class Robot:

    # initialize Robot
    def __init__(self, left, right, gyro, sonic, wheels):
        self.left = Motor(Port.B)
        self.right = Motor(Port.C)
        self.gyro = GyroSensor(Port.S1)
        self.sonic = UltrasonicSensor(Port.S2)
        self.gyro.reset_angle(0)
        self.robot = DriveBase(self.left, self.right, wheels[0], wheels[1])

    # reset distances and recalibrate gyroscope
    def reset():
        self.gyro.speed()
        self.gyro.angle()
        sleep(1)
        self.left.reset_angle(0)
        self.gyro.reset_angle(0)

    # max (speed) : deg/s, ang : deg, base (speed) : deg/s
    def rotate(max=40, ang=0, base=20):
        backwards = False
        spd = base
        diff = max / base
        while(True):
            if(ang - 3 < self.gyro.angle() > 3 and spd < max):
                spd = spd + diff
            elif(ang - 3 > self.gyro.angle() or self.gyro.angle() < 3 and spd > base):
                spd = spd - diff
            if(self.gyro.angle() != ang):
                if(backwards == False):
                    if(self.gyro.angle() > ang):
                        self.right.run(spd)
                        self.left.run(spd * -1)
                    else:
                        self.left.run(spd)
                        self.right.run(spd * -1)
                else:
                    if(self.gyro.angle() > ang):
                        self.right.run(spd)
                        self.left.run(spd * -1)
                    else:
                        self.left.run(spd)
                        self.right.run(spd * -1)
            else:
                self.left.stop(Stop.HOLD)
                self.right.stop(Stop.HOLD)
                sleep(0.5)
                if(self.gyro.angle() == ang):
                    self.gyro.reset_angle(0)
                    break
                else:
                    backwards = True

    # max (speed) : mm/s, dist (from nearest wall) : cm, base (speed) : mm/s, err (max error allowed) : mm
    def move(max=40, dist=20, base=20, err=10):
        spd = base
        cdist = dist * 10
        diff = max / base
        sdist = self.sonic.distance()
        while(True):
            rdist = self.sonic.distance()
            if(cdist - 40 < rdist > 40 + sdist and spd < max):
                spd = spd + diff
            elif(cdist - 40 > rdist or rdist < 40 + sdist and spd > base):
                spd = spd - diff

            if(rdist > cdist + err or rdist < cdist - err):
                if(rdist < cdist):
                    self.robot.drive(spd, self.gyro.angle() * -1)
                else:
                    self.robot.drive(spd * -1, self.gyro.angle() * -1)
            else:
                self.robot.stop()
                self.rotate(25, 0)
                sleep(0.5)
                if(self.sonic.distance() < cdist + err or self.sonic.distance() > cdist - err):
                    break
