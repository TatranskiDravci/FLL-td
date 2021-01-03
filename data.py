# [WIP] robot data parser
import matplotlib.pyplot as plt
from sys import argv

# read data
file = open(argv[1], "r")
rawData = file.readlines()
entries = rawData.count("%")

# parse data
allData = []
n = 0
markings = []
name = ""
normTime = 0
for line in rawData:
    if("%" in line):
        info = (line.strip()).split(" ")
        info.pop(0)
        name = info.pop(0)
        markings = info
        data = []
    elif("$" in line):
        allData.append((data, name, markings))
    elif(line == "\n"):
        continue
    else:
        point = (line.strip()).split(" ")
        if("t" in markings):
            tID = markings.index("t")
            if(len(data) > 0):
                point[tID] = float(point[tID]) - normTime
            else:
                normTime = float(point[tID])
                point[tID] = 0.0

        for n in range(len(point)):
            if("tID" in locals()):
                if(n == tID):
                    continue
            point[n] = float(point[n])

        data.append(point)

# plot data
while True:
    prompt1 = ""
    _ = 0
    for entry in allData:
        prompt1 += ("[" + str(_) + "⟩ " + str(entry[1]) + "\n")
        _ += 1
    graphs = input(prompt1 + "record ⟩⟩ ")
    if("," in graphs):
        records = graphs.split(",")
        for i in range(len(records)):
            records[i] = int(records[i])
        n = records[0]
    data = allData[n][0]
    prompt2 = str(allData[n][2])
    xAx = int(input(prompt2 + "\n x-axis ⟩⟩ "))
    yAx = int(input(" y-axis ⟩⟩ "))
    print("")
    for record in records:
        X = []; Y = []
        for point in allData[record][0]:
            X.append(point[xAx])
            Y.append(point[yAx])
        plt.plot(X, Y)
    plt.xlabel(allData[n][2][xAx])
    plt.ylabel(allData[n][2][yAx])
    plt.show()
