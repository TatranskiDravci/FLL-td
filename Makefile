PLATFORM = GOOS=linux GOARCH=arm GOARM=5
LIB = src/exten.go src/robot.go src/modules.go src/shifter.go src/env.go

calibrate:
	$(PLATFORM) go build -o build/calibrate src/calibrate.go $(LIB)

run_1:
	$(PLATFORM) go build -o build/run_1 src/run_1.go $(LIB)

run_2:
	$(PLATFORM) go build -o build/run_2 src/run_2.go $(LIB)

run_3:
	$(PLATFORM) go build -o build/run_3 src/run_3.go $(LIB)

clean:
	rm build/*