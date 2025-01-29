# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# Static Variables
# ==============================================================================



# ==============================================================================
# Go Tooling
# ==============================================================================

dev-go-tooling:
	go install github.com/divan/expvarmon@latest

tidy:
	go mod tidy
	go mod vendor


# ==============================================================================
# SNFOK Server
# ==============================================================================

server-run:
	go run ./api/main.go

# ==============================================================================
# Copy Client To Dev Ubuntu Server
# ==============================================================================
clean-server:
	echo "[+] Clearning the previous build"
	sshpass -p 12345 ssh vm@192.168.2.200 "rm -rf ~/SNFOK/*"

copy-to-dev-server:
	echo "[+] Copying The Files"
	sshpass -p 12345 scp -r ./client vm@192.168.2.200:~/SNFOK/
	sshpass -p 12345 scp -r ./vendor vm@192.168.2.200:~/SNFOK/
	sshpass -p 12345 scp ./go.sum vm@192.168.2.200:~/SNFOK/
	sshpass -p 12345 scp ./go.mod vm@192.168.2.200:~/SNFOK/
