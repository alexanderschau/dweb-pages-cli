package cmd

import (
	"fmt"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/spf13/cobra"
	"go.alxs.xyz/dweb-pages/functions"
	"go.alxs.xyz/dweb-pages/pages"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove tags from a project",
	Long:  `This command removes the given tags from an existing project.`,
	Run: func(cmd *cobra.Command, args []string) {
		settings := functions.GetSettings()

		if len(tags) == 0 {
			sendError(fmt.Errorf("No tag given"))
		}

		sh := shell.NewShell(settings.Api)

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

			delete(links, strings.Join(tagData, "-"))
		}

		dir.Links = pages.ToDirObject(links)

		resCid, err := sh.ObjectPut(dir)
		if err != nil {
			sendError(err)
		}

		if settings.CurrentIPNS {
			settings.Current = fmt.Sprintf("/ipfs/%s", resCid)
			functions.UpdateSettings(settings)
		}

		fmt.Println(resCid)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	rmCmd.Flags().StringArrayVarP(&tags, "tag", "t", []string{}, "tag name (e.g. main/latest)")
}
