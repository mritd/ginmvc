package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var versionTpl = `
Name: ginmvc
Version: %s
Arch: %s
BuildTime: %s
CommitID: %s
`

var (
	Version   string
	BuildTime string
	CommitID  string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long: `
Print version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(versionTpl, Version, runtime.GOOS+"/"+runtime.GOARCH, BuildTime, CommitID)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
