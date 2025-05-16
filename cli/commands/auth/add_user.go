package auth

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/FearLessSaad/SNFOK/cli/tooling"
	"github.com/FearLessSaad/SNFOK/constants/auth_constants"
	"github.com/FearLessSaad/SNFOK/db"
	"github.com/FearLessSaad/SNFOK/db/models/auth"
	"github.com/FearLessSaad/SNFOK/tooling/logger"
	"github.com/spf13/cobra"
	"github.com/uptrace/bun"
)

var AddUserCmd = &cobra.Command{
	Use:   "adduser",
	Short: "Add a user with first name, last name, email, and designation",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)

		// Prompt for First Name
		fmt.Print("[=] Enter First Name >> ")
		scanner.Scan()
		firstName := strings.TrimSpace(scanner.Text())
		if firstName == "" {
			fmt.Fprintln(os.Stderr, "Error: First Name cannot be empty")
			os.Exit(1)
		}

		// Prompt for Last Name
		fmt.Print("[=] Enter Last Name >> ")
		scanner.Scan()
		lastName := strings.TrimSpace(scanner.Text())
		if lastName == "" {
			fmt.Fprintln(os.Stderr, "Error: Last Name cannot be empty")
			os.Exit(1)
		}

		// Prompt for Email
		fmt.Print("[=] Enter Email >> ")
		scanner.Scan()
		email := strings.TrimSpace(scanner.Text())
		if !isValidEmail(email) {
			fmt.Fprintln(os.Stderr, "Error: Invalid or empty email address")
			os.Exit(1)
		}

		// Prompt for Designation
		fmt.Print("[=] Enter Designation >> ")
		scanner.Scan()
		designation := strings.TrimSpace(scanner.Text())
		if designation == "" {
			fmt.Fprintln(os.Stderr, "Error: Designation cannot be empty")
			os.Exit(1)
		}

		// Add User In Database
		ctx := context.Background()
		conn := db.GetDB()

		chk_user := new(auth.Users)
		exists, _ := conn.NewSelect().Model(chk_user).Where("email = ?", email).Exists(ctx)

		if !exists {
			fmt.Println("[!] User creation failed becuase user with entered email is already exists.")
			return
		}

		user := new(auth.Users)
		user.FirstName = firstName
		user.LastName = lastName
		user.Email = email
		user.Designation = designation

		token, _ := tooling.GenerateRandomToken()
		user.Token = token
		user.Status = auth.UserStatusActive

		user.AuditFields.CreatedAt = time.Now()
		user.AuditFields.CreatedBy = auth_constants.SNFOK_CLI

		_, err := conn.NewInsert().Model(user).Exec(ctx)

		if err != nil {
			logger.Log(logger.ERROR, "Unable To Create New User Using CLI", logger.Field{Key: "error", Value: err.Error()})
			return
		}

		fmt.Println("[+] New user is addedd successfully.")
		fmt.Printf("[+] User this token email '%s' and token '%s' to login.\n", email, token)

	},
}

var RotateTokenCmd = &cobra.Command{
	Use:   "rotatetoken",
	Short: "Rotate token for access web ui of SNFOK.",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)

		// Prompt for Email
		fmt.Print("[=] Enter Email >> ")
		scanner.Scan()
		email := strings.TrimSpace(scanner.Text())
		if !isValidEmail(email) {
			fmt.Fprintln(os.Stderr, "Error: Invalid or empty email address")
			os.Exit(1)
		}

		// Add User In Database
		ctx := context.Background()
		conn := db.GetDB()

		chk_user := new(auth.Users)
		err := conn.NewSelect().Model(chk_user).Where("email = ?", email).Limit(1).Scan(ctx)

		if err != nil {
			fmt.Println("[!] No User Is Found!")
			return
		}

		token, _ := tooling.GenerateRandomToken()
		chk_user.Token = token

		chk_user.AuditFields.UpdatedAt = bun.NullTime{Time: time.Now()}
		chk_user.AuditFields.UpdatedBy = auth_constants.SNFOK_CLI

		_, err = conn.NewUpdate().Model(chk_user).Where("email = ?", email).Exec(ctx)

		if err != nil {
			logger.Log(logger.ERROR, "Unable To Create New User Using CLI", logger.Field{Key: "error", Value: err.Error()})
			return
		}

		fmt.Println("[+] User is rotated successfully.")
		fmt.Printf("[+] User this token email '%s' and token '%s' to login.\n", email, token)

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
