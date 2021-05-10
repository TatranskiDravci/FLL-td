#!/usr/bin/env python3

import sys
import numpy as np
import matplotlib.pyplot as plt

def main(args=[]):
    extracted = []
    with open(args[0], "r") as data:
        callnum = -1
        count = 0
        for line in data:
            if "call to" in line:
                callnum += 1
                count = 0
                continue
            if callnum == int(args[1]):
                count += 1
                inline = line.split(" ")
                for i in range(len(inline)):
                    try:
                        inline[i] = int(inline[i].strip())
                    except:
                        inline[i] = float(inline[i].strip())
                inline.insert(0, count)
                extracted.append(
                    inline
                )

    npext = np.array(extracted)
    plt.plot(npext[:,int(args[2])], npext[:,int(args[3])])
    plt.show()


if __name__ == "__main__":
    args = sys.argv
    args.pop(0)
    main(args)
