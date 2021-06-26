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
	Short: "Initialize a new project",
	Run: func(cmd *cobra.Command, args []string) {

		settings := types.Settings{
			Api:         "localhost:5001",
			Current:     "/ipns/pages-init.dweb.party",
			CurrentIPNS: true,
			OutputDir:   ".",
			Then:        ".dweb-pages/then.sh",
		}

		if _, err := os.Stat(".dweb-pages/settings.json"); !os.IsNotExist(err) {
			rewrite := false
			prompt := &survey.Confirm{
				Message: "Project already exists, do you really want to overwrite it?",
			}
			survey.AskOne(prompt, &rewrite)
			if !rewrite {
				return
			}

			out, err := ioutil.ReadFile(".dweb-pages/settings.json")
			if err != nil {
				panic(err)
			}
			json.Unmarshal(out, &settings)
		}

		survey.AskOne(&survey.Input{
			Message: "API endpoint",
			Default: settings.Api,
		}, &settings.Api)
		survey.AskOne(&survey.Input{
			Message: "Initial project",
			Default: settings.Current,
		}, &settings.Current)
		survey.AskOne(&survey.Confirm{
			Message: "Should we update this value on each change?",
			Default: settings.CurrentIPNS,
		}, &settings.CurrentIPNS)
		survey.AskOne(&survey.Input{
			Message: "Output directory",
			Default: settings.OutputDir,
		}, &settings.OutputDir)

		jsn, err := json.Marshal(settings)
		if err != nil {
			panic(err)
		}
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
