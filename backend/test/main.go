package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	config := &ssh.ClientConfig{
		User: "ubuntu",
		Auth: []ssh.AuthMethod{
			ssh.Password("saad123456"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "10.251.137.198:22", config)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	output, err := session.CombinedOutput("uname -a")
	if err != nil {
		log.Fatalf("Failed to run command: %s", err)
	}

	fmt.Println(string(output))
}
