package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"go.alxs.xyz/dweb-pages/types"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if _, err := os.Stat(".dweb-pages/settings.json"); !os.IsNotExist(err) {
			rewrite := false
			prompt := &survey.Confirm{
				Message: "Project already exists, do you really want to overwrite it?",
			}
			survey.AskOne(prompt, &rewrite)
			if !rewrite {
				return
			}
		}

		settings := types.Settings{
			Api:       "localhost:5001",
			Current:   "/ipfs/bafybeiczsscdsbs7ffqz55asqdf3smv6klcw3gofszvwlyarci47bgf354",
			OutputDir: ".",
		}

		survey.AskOne(&survey.Input{
			Message: "API endpoint",
			Default: settings.Api,
		}, &settings.Api)
		survey.AskOne(&survey.Input{
			Message: "Initial project",
			Default: settings.Current,
		}, &settings.Current)
		survey.AskOne(&survey.Input{
			Message: "Output directory",
			Default: settings.OutputDir,
		}, &settings.OutputDir)

		jsn, _ := json.Marshal(settings)
		os.Mkdir(".dweb-pages", 0755)
		ioutil.WriteFile(".dweb-pages/settings.json", jsn, 0644)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
