# Python
## Structure
Python interpreter: `pybricks-micropython` <br>
Packages: `pybricks` <br>
Structure:
```
./python/
├── README.md  documentation
├── robot.py   function definitions
└── main.py    executable
```
## Usage
```python
#!/usr/bin/env pybricks-micropython
from robot import *
from pybricks.parameters import Port

robot = Robot(Port.B, Port.C, Port.S1, Port.S2, (44, 180))

robot.move(60, 10)
robot.rotate(60, 90)
```

* import everything from `robot.py`
* import necessary objects from `pybricks`

### *class* Robot(*left_motor, right_motor, gyro, ultrasonic, wheels*)
##### Parameters:
* **left_motor** (Port) - port to which the left motor is connected
* **right_motor** (Port) - port to which the right motor is connected
* **gyro** (Port) - port to witch the gyro sensor is connected
* **ultrasonic** (Port) - port to which the ultrasonic sensor is connected
* **wheels** (Tuple) - diameter of wheel and distance between wheels (axle length)

##### Methods:
####  move(*max, distance, error, base*)
* **max** (Integer : *mm/s*) - max speed
* **distance** (Float : *cm*) - distance from wall
* **base** (Integer : *mm/s*) - base speed
* **error** (Float : *mm*) - tolerable error in distance

####  rotate(*speed, angle*)
* **speed** (Integer : *degs/s*) - rotation speed
* **angle** (Integer : *degs*) - target angle

####  wp_create(*waypoints*)
* **waypoints** (List > 2-Tuple > Integer : *deg*, Integer : *cm*) - set of waypoints

####  wp_exec(*r_spd, m_max, error*)
* **r_spd** (Integer : *degs/s*) - rotation speed
* **m_max** (Integer : *mm/s*) - max movement speed
* **error** (Integer : *mm*) - tolerable error in distance

#### wp_revert(*steps, r_spd, m_max, error**)
* **steps** (Integer : *none*) - revert by **steps** number of waypoints
* **r_spd** (Integer : *degs/s*) - rotation speed
* **m_max** (Integer : *mm/s*) - max movement speed
* **error** (Integer : *mm*) - tolerable error in distance

#### wp_append(*waypoints*)
* **waypoints** (List > 2-Tuple > Integer : *deg*, Integer : *cm*) - set of waypoints

### *class* Shifter(*shifter_motor, driver_motor, module_angles*)
##### Parameters:
* **shifter_motor** (Port) - port to which the shifter motor is connected
* **driver_motor** (Port) - port to which the driver motor is connected
* **module_angles** (List > Integers : *degs*) - set of angles, on which different modules are engaged

##### Methods:
####  run(*index, speed, angle*)
* **index** (Integer : *list index*) - number of module
* **speed** (Integer : *degs/s*) - rotation speed of driver motor
* **angle** (Integer : *degs*) - target angle of the driver motor

