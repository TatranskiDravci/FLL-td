from pybricks import ev3brick as brick
from pybricks.ev3devices import Motor, GyroSensor, UltrasonicSensor
from pybricks.parameters import Port, Stop
from pybricks.robotics import DriveBase
from time import sleep

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
    def __init__(self, left, right, gyro, sonic, wheels):
        self.on_wp = 0
        #   -> index of the last executed waypoint
        self.waypoints = []
        #   -> stores waypoint map
        self.midpoints = []
        #   -> stores distance between waipoints (useful for revert function)
        # hardware connections
        self.left = Motor(left)
        self.right = Motor(right)
        self.gyro = GyroSensor(gyro)
        self.sonic = UltrasonicSensor(sonic)
        self.gyro.reset_angle(0)
        self.robot = DriveBase(self.left, self.right, wheels[0], wheels[1])

    # [experimental] recalibrate gyroscope
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

    # move robot to specified distance ; spd : "mm/s", dist : cm, err : mm
    def move(self, spd=100, dist=20, err=10):
        # convert distance from cm to mm
        dist = dist * 10
        while(True):
            rdist = self.sonic.distance()
            brick.display.text(rdist)
            if(rdist < dist - err):
                self.robot.drive(spd, self.gyro.angle() * -5 )
            elif(rdist > dist + err):
                self.robot.drive(spd * -1, self.gyro.angle() * -5)
            else:
                self.robot.stop()
                self.gyro.reset_angle(0)
                sleep(1) # wait to ensure the measurement is correct
                if(self.sonic.distance() < dist + err or self.sonic.distance() > dist - err):
                    break

    # WAYPOINT CODE
    # create waypoint map ; waypoints : list<2-tuple>
    def wp_create(self, waypoints):
        self.waypoints = waypoints
        self.midpoints.append(self.sonic.distance() / 10) # create a midpoint
        self.on_wp = 0

    # execute waypoint map ; rspd : deg/s, mspd : "mm/s", err : mm
    def wp_exec(self, rspd=40, mspd=40, err=10):
        for i in range(self.on_wp, len(self.waypoints)):
            self.rotate(rspd, self.waypoints[i][0])
            sleep(0.5) # wait to ensure correct measurement
            self.midpoints.append(self.sonic.distance() / 10) # create a midpoint
            self.move(mspd, self.waypoints[i][1], err)
        self.on_wp = len(self.waypoints)

    # add new waypoints ; waypoints : list<2-tuple>
    def wp_append(self, waypoints):
        for point in waypoints:
            self.waypoints.append(point)

    # FIXME: didn't execute last run, unknown issue
    # revert back by x steps ; steps : int, rspd : deg/s, mspd : "mm/s", err : mm
    def wp_revert(self, steps=1, rspd=40, mspd=40, err=10):
        # check if number of steps doesn't exceed the length of waypoint map
        if(steps > len(self.waypoints)):
            steps = len(self.waypoints)

        for i in range(steps):
            self.move(mspd, self.midpoints[len(self.midpoints) - i - 1], err) # move back to the selected midpoint
            self.rotate(rspd, self.waypoints[len(self.waypoints) - i - 1][0] * -1)
            self.waypoints.pop()
            self.midpoints.pop()
        self.on_wp = len(self.waypoints)
    # WAYPOINT CODE
