#!/usr/bin/env pybricks-micropython

from robot import *
from pybricks.parameters import Port

r = Robot(Port.D, Port.A, Port.S2, Port.S3, (30, 140))
s = Shifter(Port.B, Port.C)

r.wp_create(
    [
        (0, 20),
        (90, 30)
    ]
)
r.wp_exec(100, 150, 2)
s.run(0, 300, -1000)
r.revert(1, 100, 150, 2)
