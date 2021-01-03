#!/usr/bin/env pybricks-micropython

import robot as rb
from pybricks.parameters import Port
from time import time

#
r = rb.Robot(Port.D, Port.A, Port.S2, Port.S3, (30, 140))
s = rb.Shifter(Port.B, Port.C)

# move robot and collect movement data
data, types = r.move()
# save movement data
rb.saveData(data, types)
