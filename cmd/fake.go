package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"github.com/jorge07/play/post"
)

var fakeCommand = &cobra.Command{
	Use:   "fake",
	Short: "Display fake data given an identifer",
	Long: "Display fake data from https://jsonplaceholder.typicode.com/posts fir the id given",
	Run: func(cmd *cobra.Command, args []string) {

		identifier, err := cmd.Flags().GetString("identifier")
		if err != nil {
			log.Fatal(err)
		}

		if identifier == "" {
			log.Fatal("An identifier has to be provided")
		}

		repository := post.GetRepository()

		p := repository.Call(identifier)

		p.Render()
	},
}

func init() {

	fakeCommand.Flags().StringP("identifier", "i", "", "A numberic identifier")
	RootCmd.AddCommand(fakeCommand)
}
