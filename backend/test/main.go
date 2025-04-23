package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	// SSH configuration
	config := &ssh.ClientConfig{
		User: "your_username",
		Auth: []ssh.AuthMethod{
			ssh.Password("your_password"), // or use ssh.PublicKeys() for key-based auth
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // use a proper callback in production
	}

	// Connect to SSH server
	client, err := ssh.Dial("tcp", "your.server.ip:22", config)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	defer client.Close()

	// Create a new session
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	// Pipe output to stdout
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	// Run command
	cmd := "uptime"
	fmt.Println("Running command:", cmd)
	err = session.Run(cmd)
	if err != nil {
		log.Fatalf("Failed to run: %s", err)
	}
}
