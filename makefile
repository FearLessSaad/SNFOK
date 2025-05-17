build_agent:
	go build -o agent/agent agent/main.go

build_cli:
	go build -o cli/snfokctl cli/main.go

build_server:
	go build -o server main.go

clean:
	rm -f agent/agent cli/snfokctl server

copy_agent:
	scp agent/agent ubuntu@10.251.137.249:~/

bac_agent:
	make clean
	make build_agent
	make copy_agent

run_server:
	make build_server
	./server