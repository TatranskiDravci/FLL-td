PLATFORM = GOOS=linux GOARCH=arm GOARM=5
LIB = src/exten.go src/robot.go src/modules.go src/shifter.go

main:
	 $(PLATFORM) go build -o build/main src/main.go $(LIB)

calib:
	$(PLATFORM) go build -o build/calib src/main.go

run_1:
	$(PLATFORM) go build -o build/run_1 src/run_1.go $(LIB)