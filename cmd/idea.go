package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"github.com/marzelwidmer/uzo/util"

	"github.com/spf13/cobra"
)

var ideaCmd = &cobra.Command{
	Use:   "idea",
	Short: "It will open the directory in IntelliJ IDEA",
	Long: `This command will help to open the unzipped folder
	to IntelliJ IDEA.
	In order for this command to work, IntelliJ IDEA should be installed in your system`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(File)
		var filename string
		var err error
		var argument string

		if File != "" {
			argument = File
		} else {
			argument = args[0]
		}

		fileExists, err := util.FileExists(argument)
		if err != nil {
			fmt.Println(err)
		}
		if fileExists {
			filename, err = filepath.Abs(argument)
			if err != nil {
				fmt.Println(err.Error())

			}
		} else {
			fmt.Printf("File %v doest not Exists", argument)
			return
		}

		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}

		util.Unzip(filename, wd)

		os.Chdir(util.FilenameWithoutExtension(filename))

		wd, err = os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		commandCode := exec.Command("idea", wd)
		err = commandCode.Run()

		if err != nil {
			log.Fatal("Intellij Idea executable file not found in %PATH%")
		}
	},
}

func init() {
	rootCmd.AddCommand(ideaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	ideaCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "A File name to unzip and open in IDE")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ideaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
