package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
	"gitlab.eng.vmware.com/marketplace-partner-eng/chart-mover/v2/lib"
)

func init() {
	rootCmd.AddCommand(ListImagesCmd)
	ListImagesCmd.SetOut(os.Stdout)
}

var ListImagesCmd = &cobra.Command{
	Use:   "list-images <chart>",
	Short: "Renders and lists the images, found using the image list file",
	//Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		var images []string

		for _, imageTemplate := range ImageTemplates {
			image, err := imageTemplate.Render(Chart, []*lib.RewriteAction{})
			if err != nil {
				cmd.PrintErrln(err.Error())
				return
			}

			images = append(images, image.Remote())
		}

		encoded, err := json.Marshal(images)
		if err != nil {
			cmd.PrintErrln(err.Error())
			return
		}

		cmd.Println(string(encoded))
	},
}
