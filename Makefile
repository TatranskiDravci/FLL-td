PLATFORM = GOOS=linux GOARCH=arm GOARM=5
LIB = src/fmath.go src/robot.go src/modules.go src/shifter.go

main:
	 $(PLATFORM) go build -o build/main src/main.go $(LIB)
