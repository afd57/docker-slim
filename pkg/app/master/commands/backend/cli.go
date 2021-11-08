package backend

import (
	"fmt"

	"github.com/docker-slim/docker-slim/pkg/app/master/commands"

	"github.com/urfave/cli"
)

const (
	Name  = "backend"
	Usage = "backend command"
	Alias = "b"
)

var CLI = cli.Command{
	Name:    Name,
	Aliases: []string{Alias},
	Usage:   Usage,
	Action: func(ctx *cli.Context) error {
		backendEnable := ctx.GlobalBool(commands.FlagBackendEnable)
		fmt.Println("Backend Command")
		fmt.Println(backendEnable)
		OnCommand(backendEnable)

		return nil
	},
}
