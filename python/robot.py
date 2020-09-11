from pybricks.ev3devices import Motor, GyroSensor, UltrasonicSensor
from pybricks.parameters import Port, Stop
from pybricks.robotics import DriveBase
from time import sleep

class Shifter:

    def __init__(self, sm, rm, angs=[0, 150, 300, 450]):
        self.sm = Motor(sm)
        self.rm = Motor(rm)
        self.angs = angs

    def run(self, index, spd, ang):
        self.sm.run_target(150, self.angs[index])
        self.rm.run_target(spd, ang)

class Robot:

    def linear(self, x, A, B):
        y = (((B[1] - A[1]) / (B[0] - A[0])) * x) + (A[1] - (((B[1] - A[1]) / (B[0] - A[0])) * A[0]))
        return y

    # initialize Robot
    def __init__(self, left, right, gyro, sonic, wheels):
        self.on_wp = 0
        self.waypoints = []
        self.midpoints = []
        self.left = Motor(left)
        self.right = Motor(right)
        self.gyro = GyroSensor(gyro)
        self.sonic = UltrasonicSensor(sonic)
        self.gyro.reset_angle(0)
        self.robot = DriveBase(self.left, self.right, wheels[0], wheels[1])

    # recalibrate gyroscope
    def reset(self):
        self.gyro.speed()
        self.gyro.angle()
        sleep(1)
        self.left.reset_angle(0)
        self.gyro.reset_angle(0)

    # rotate in place ; spd : deg/s, ang : degs
    def rotate(self, spd=40, ang=0):
        while(True):
            if(self.gyro.angle() > ang):
                self.right.run(spd)
                self.left.run(spd * -1)
            elif(self.gyro.angle() < ang):
                    self.left.run(spd)
                    self.right.run(spd * -1)
            else:
                self.left.stop(Stop.HOLD)
                self.right.stop(Stop.HOLD)
                sleep(0.5)
                if(self.gyro.angle() == ang):
                    self.gyro.reset_angle(0)
                    break

    # spd : "mm/s", dist : cm, err : mm
    def move(self, max=100, dist=20, err=10, base=50):
        # convert distance from cm to mm
        dist = dist * 10
        sdist = self.sonic.distance()
        while(True):
            
            rdist = self.sonic.distance()

            # speedup and slowdown code
            if(rdist < sdist + 10):
                spd = self.linear(rdist - sdist, (0, base), (10, max))
            elif(rdist < dist - 10):
                spd = max
            else:
                spd = self.linear(rdist - (dist - 10), (0, max), (10, base))

            brick.display.text(rdist)
            if(rdist < dist - err):
                self.robot.drive(spd, self.gyro.angle() * -5 )
            elif(rdist > dist + err):
                self.robot.drive(spd * -1, self.gyro.angle() * -5)
            else:
                self.robot.stop()
                self.gyro.reset_angle(0)
                sleep(1)
                if(self.sonic.distance() < dist + err or self.sonic.distance() > dist - err):
                    break

    # WAYPOINT CODE
    # create waypoint map ; waypoints : list<2-tuple>
    def wp_create(self, waypoints):
        self.waypoints = waypoints
        self.midpoints.append(self.sonic.distance() / 10)
        self.on_wp = 0

    # run waypoint map ; rspd : deg/s, mspd : "mm/s", err : mm
    def wp_exec(self, rspd=40, mspd=40, err=10):
        for i in range(self.on_wp, len(self.waypoints)):
            self.rotate(rspd, self.waypoints[i][0])
            sleep(0.5)
            self.midpoints.append(self.sonic.distance() / 10)
            self.move(mspd, self.waypoints[i][1], err)
        self.on_wp = len(self.waypoints)

    # run waypoint map from x ; 
    def wp_exec_from(self, wpn, rspd=40, mspd=40, err=10):
        for i in range(wpn, len(self.waypoints)):
            self.midpoints.pop()
        for i in range(wpn, len(self.waypoints)):
            self.rotate(rspd, self.waypoints[i][0])
            self.midpoints.append(self.sonic.distance() / 10)
            self.move(mspd, self.waypoints[i][1], err)
        self.on_wp = len(self.waypoints)

    # add new waypoints ; waypoints : list<2-tuple>
    def wp_append(self, waypoints):
        for point in waypoints:
            self.waypoints.append(point)
    
    # revert back by x steps ; steps : int, rspd : deg/s, mspd : "mm/s", err : mm
    def wp_revert(self, steps=1, rspd=40, mspd=40, err=10):
        if(steps > len(self.waypoints)):
            steps = len(self.waypoints)
        for i in range(steps):
            self.move(mspd, self.midpoints[len(self.midpoints) - i - 1], err)
            self.rotate(rspd, self.waypoints[len(self.waypoints) - i - 1][0] * -1)
            self.waypoints.pop()
            self.midpoints.pop()
        self.on_wp = len(self.waypoints)
    # WAYPOINT CODE

