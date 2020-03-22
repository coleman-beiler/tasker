package commands

import (
	"fmt"
	"os"

	"../common"
	"github.com/urfave/cli"
)

func getRegisterFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "username",
			Aliases:  []string{"u"},
			Usage:    "Name of the user.",
			Required: true,
		},
	}
}

func RegisterIt(c *cli.Context) error {
	username := c.String("username")
	var conf common.TaskerConfig
	err := conf.SetUsername(username)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Print("\n")
	fmt.Println("--------")
	fmt.Println("User Changed To: " + username)
	fmt.Println("--------")
	fmt.Print("\n")
	return err
}

func init() {
	flags := getRegisterFlags()
	common.RegisterCommand(&cli.Command{
		Name:   "register",
		Usage:  "Enter options to configure tasker",
		Action: RegisterIt,
		Flags:  flags,
	})
}
