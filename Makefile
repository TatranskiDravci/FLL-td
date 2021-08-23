PLATFORM = GOOS=linux GOARCH=arm GOARM=5
LIB = src/exten.go src/robot.go src/modules.go src/shifter.go src/env.go

run:
	$(PLATFORM) go build -o build/run src/run.go $(LIB)

clean:
	rm build/*