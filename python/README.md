# Python
## Project structure
Python interpreter: `pybricks-micropython` <br>
Packages: `pybricks` <br>
Structure:
```
./python/
├── README.md  documentation
├── robot.py  function definitions
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
* **left_motor** (Port) - port to witch the left motor is connected
* **right_motor** (Port) - port to witch the right motor is connected
* **gyro** (Port) - port to witch the gyro sensor is connected
* **ultrasonic** (Port) - port to witch the ultrasonic sensor is connected
* **wheels** (Tuple) - diameter of wheel and distance between wheels (axle length)

##### Methods:
####  move(*max, distance, base, error*)
* **max** (Integer : *mm/s*) - max speed
* **distance** (Float : *cm*) - distance from wall
* **base** (Integer : *mm/s*) - base speed
* **error** (Float : *mm*) - allowed error in distance

####  rotate(*max, angle, base*)
* **max** (Integer : *mm/s*) - max speed
* **angle** (Integer : *deg*) - target angle
* **base** (Integer : *mm/s*) - base speed
