all:
	go build -o TrafficMonitor-go .
clean:
	rm TrafficMonitor-go
run: all
	./TrafficMonitor-go