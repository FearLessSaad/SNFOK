build-agent:
	go build -o agent/agent agent/main.go

build-cli:
	go build -o cli/snfokctl cli/main.go

build-server:
	go build -o server main.go

clean:
	rm -f agent/agent cli/snfokctl server

copy-agent:
	scp agent/agent ubuntu@10.251.137.249:~/

copy-k8s:
	scp k8s ubuntu@$(IP):~/

bac-agent:
	make clean
	make build-agent
	make copy-agent

run-server:
	make build-server
	./server

expose-server:
	ngrok http --url=neutral-widely-fawn.ngrok-free.app 8989

git-push:
	make clean
	git add -A
	git commit -m "$(VER)"
	git push