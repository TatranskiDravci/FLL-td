# run2:
# 	GOOS=linux GOARCH=arm GOARM=5 go build -o run2 run2.go modules.go fmath.go shifter.go robot.go

# run1:
# 	GOOS=linux GOARCH=arm GOARM=5 go build -o run1 run1.go modules.go fmath.go shifter.go robot.go

main:
	GOOS=linux GOARCH=arm GOARM=5 go build -o build/main src/main.go src/fmath.go src/robot.go
