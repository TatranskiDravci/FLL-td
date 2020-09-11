#!/usr/bin/env pybricks-micropython

from robot import *
from pybricks.parameters import Port

r = Robot(Port.A, Port.D, Port.S2, Port.S3, (30, 140))
s = Shifter(Port.B, Port.C)

r.wp_create(
    [
        (0, 20),
        (90, 60)
    ]
)
r.wp_exec(100, 120, 2)
s.run(1, 350, 1500)
r.wp_revert(2, 100, 120, 2)
