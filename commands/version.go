package commands

import (
	"encoding/json"
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
	"gopkg.in/resty.v1"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

const latestReleaseUrl = "https://api.github.com/repos/gochrono/chrono/releases/latest"

var banner = `   _____ _
  / ____| |
 | |    | |__  _ __ ___  _ __   ___
 | |    | '_ \| '__/ _ \| '_ \ / _ \ 
 | |____| | | | | | (_) | | | | (_) |
  \_____|_| |_|_|  \___/|_| |_|\___/`

func newVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Prints the version then exits",
		Long:  "Prints the version then exits",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(banner)
			fmt.Printf("\nversion: %s\ncommit: %s\nbuilt: %s\n", version, commit, date)
		},
	}

	versionCmd.AddCommand(newVersionCheckCmd())
	return versionCmd
}

func newVersionCheckCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "check",
		Short: "Checks if there is a new version",
		Long:  "Checks if there is a new version release on GitHub",
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := resty.R().Get(latestReleaseUrl)
			if err != nil {
				panic(err)
			}
			if resp.StatusCode() == 200 {
				var latestRelease chronolib.GithubLatestRelease
				json.Unmarshal(resp.Body(), &latestRelease)
				if version != latestRelease.TagName {
					fmt.Printf("Found a new release [%s]!\nVisit %s and download the latest version\n", latestRelease.TagName, latestRelease.HTMLURL)
				} else {
					fmt.Println("You are on the latest version!")
				}
			} else {
				fmt.Println("No releases found.")
			}

		},
	}
}
