package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "app",
		Short: "TBD",
		Long:  `TBD`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(
		runServer,
		runMigrate,
		runHashPassword,
	)
}

func ask(q string) bool {
	fmt.Println(q)
	reader := bufio.NewReader(os.Stdin)
	for {
		answer, _ := reader.ReadString('\n')
		answer = strings.ToLower(strings.TrimSpace(answer))
		if answer == "y" || answer == "yes" {
			return true
		} else if answer == "n" || answer == "no" {
			break
		} else {
			continue
		}
	}
	return false
}
