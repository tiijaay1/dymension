package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/dymensionxyz/dymension/app"
	"github.com/ignite/cli/ignite/pkg/cosmoscmd"
)

func main() {

	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.ShortName,
		app.ShortName,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)
	// see git issue: https://github.com/dymensionxyz/dymension/issues/99
	rootCmd.Short =
		`
	 ______   __   __  __   __  _______  __    _  _______  ___   _______  __    _    __   __  __   __  _______ 
	|      | |  | |  ||  |_|  ||       ||  |  | ||       ||   | |       ||  |  | |  |  | |  ||  | |  ||  _    |
	|  _    ||  |_|  ||       ||    ___||   |_| ||  _____||   | |   _   ||   |_| |  |  |_|  ||  | |  || |_|   |
	| | |   ||       ||       ||   |___ |       || |_____ |   | |  | |  ||       |  |       ||  |_|  ||       |
	| |_|   ||_     _||       ||    ___||  _    ||_____  ||   | |  |_|  ||  _    |  |       ||       ||  _   | 
	|       |  |   |  | ||_|| ||   |___ | | |   | _____| ||   | |       || | |   |  |   _   ||       || |_|   |
	|______|   |___|  |_|   |_||_______||_|  |__||_______||___| |_______||_|  |__|  |__| |__||_______||_______|
	`
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
