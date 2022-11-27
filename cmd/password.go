package cmd

import (
	"fmt"
	"go_web/api/handler"
	"go_web/config"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

var runHashPassword = &cobra.Command{
	Use:   "bcrypt",
	Short: "hash a password",
	Long:  `Tbd`,
	Run:   hash,
}
var rawPassword string

func init() {
	flag.StringVarP(&rawPassword, "raw-password", "r", "", "help message")
}
func hash(cmd *cobra.Command, args []string) {
	if rawPassword == "" {
		fmt.Println("nothing to do")
		return
	}
	cfg := config.LoadConfig()
	authService := handler.InitAuthService(nil, cfg)
	hashPassword, err := authService.HashPassword(rawPassword)
	if err != nil {
		fmt.Printf("error: %v \n", err)
		return
	}

	fmt.Println(hashPassword)
}
