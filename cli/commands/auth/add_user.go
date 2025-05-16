package auth

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"github.com/spf13/cobra"
)

var AddUserCmd = &cobra.Command{
	Use:   "adduser",
	Short: "Add a user with first name, last name, email, and designation",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)

		// Prompt for First Name
		fmt.Print("Enter First Name: ")
		scanner.Scan()
		firstName := strings.TrimSpace(scanner.Text())
		if firstName == "" {
			fmt.Fprintln(os.Stderr, "Error: First Name cannot be empty")
			os.Exit(1)
		}

		// Prompt for Last Name
		fmt.Print("Enter Last Name: ")
		scanner.Scan()
		lastName := strings.TrimSpace(scanner.Text())
		if lastName == "" {
			fmt.Fprintln(os.Stderr, "Error: Last Name cannot be empty")
			os.Exit(1)
		}

		// Prompt for Email
		fmt.Print("Enter Email: ")
		scanner.Scan()
		email := strings.TrimSpace(scanner.Text())
		if !isValidEmail(email) {
			fmt.Fprintln(os.Stderr, "Error: Invalid or empty email address")
			os.Exit(1)
		}

		// Prompt for Designation
		fmt.Print("Enter Designation: ")
		scanner.Scan()
		designation := strings.TrimSpace(scanner.Text())
		if designation == "" {
			fmt.Fprintln(os.Stderr, "Error: Designation cannot be empty")
			os.Exit(1)
		}

		// Add User In Database
	},
}

// isValidEmail checks if the email is non-empty and matches a basic email pattern
func isValidEmail(email string) bool {
	if email == "" {
		return false
	}
	// Basic email regex: allows letters, numbers, dots, and common symbols before @, and a domain after
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}