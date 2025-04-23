# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# Static Variables
# ==============================================================================
REMOTE_SERVER ?= ghostx64@192.168.2.10
PASSWORD ?= 12345

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
	sshpass -p $(PASSWORD) ssh $(REMOTE_SERVER) "rm -rf ~/SNFOK/*"

copy-to-dev-server:
	echo "[+] Copying The Files"
	sshpass -p $(PASSWORD) scp -r ./client $(REMOTE_SERVER):~/SNFOK/
	sshpass -p $(PASSWORD) scp -r ./vendor $(REMOTE_SERVER):~/SNFOK/
	sshpass -p $(PASSWORD) scp ./go.sum $(REMOTE_SERVER):~/SNFOK/
	sshpass -p $(PASSWORD) scp ./go.mod $(REMOTE_SERVER):~/SNFOK/