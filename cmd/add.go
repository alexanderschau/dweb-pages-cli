package cmd

import (
	"fmt"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/spf13/cobra"
	"go.alxs.xyz/dweb-pages/functions"
	"go.alxs.xyz/dweb-pages/pages"
)

var tags []string

func sendError(err error) {
	fmt.Printf("Error: %s\n", err.Error())
	os.Exit(1)
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add content to an existing project",
	Long:  `This command adds a folder to an existing project.`,
	Run: func(cmd *cobra.Command, args []string) {
		settings := functions.GetSettings()

		if len(tags) == 0 {
			sendError(fmt.Errorf("No tag given"))
		}

		sh := shell.NewShell("localhost:5001")

		hash, err := sh.AddDir(settings.OutputDir)
		if err != nil {
			panic(err)
		}

		dir, err := sh.ObjectGet(settings.Current)
		if err != nil {
			sendError(err)
		}
		links := pages.ReadDirObject(dir)

		for _, tag := range tags {
			tagData, err := pages.ResolveTag(tag)
			if err != nil {
				sendError(err)
			}

			links[strings.Join(tagData, "-")] = hash
		}

		dir.Links = pages.ToDirObject(links)

		resCid, err := sh.ObjectPut(dir)
		if err != nil {
			sendError(err)
		}

		settings.Current = fmt.Sprintf("/ipfs/%s", resCid)
		functions.UpdateSettings(settings, resCid)

		fmt.Println(resCid)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringArrayVarP(&tags, "tag", "t", []string{}, "tag name (e.g. main/latest, default)")
}
