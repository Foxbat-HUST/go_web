package cmd

import (
	"fmt"
	"go_web/config"
	"go_web/sql"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/spf13/cobra"
)

var runMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "run migrate",
	Long:  `Tbd`,
	Run:   doMigrate,
}

func doMigrate(cmd *cobra.Command, args []string) {
	d, err := iofs.New(sql.Fs, ".")
	if err != nil {
		cmd.Println(err)
		return
	}
	cfg := config.LoadConfig()
	dbUrl := fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s", cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Db)
	m, err := migrate.NewWithSourceInstance("iofs", d, dbUrl)
	if err != nil {
		cmd.Println(err)
		return
	}
	if !ask("do migration?(Y/N)") {
		cmd.Println("migration was canceled!")
		return
	}
	err = m.Up()
	if err != nil {
		cmd.Println(err)
		return
	}
	cmd.Println("migration successful!")
}
