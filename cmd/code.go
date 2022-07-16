/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"uzo/utils"

	"github.com/spf13/cobra"
)

var File string

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code <file_to_unzip_and_open>",
	Short: "It will open the directory in visual studio code",
	// Args:                  cobra.ExactArgs(1),
	Args: func(cmd *cobra.Command, args []string) error {
		if File == "" && len(args) < 1 {
			return errors.New("require atleast one argument or -f flag")
		}
		return nil
	},
	DisableFlagsInUseLine: true,
	Example:               "uzo code /home/<user>/Downloads/demo.zip",
	RunE: func(cmd *cobra.Command, args []string) error {
		var fileName string
		var err error
		var argument string
		// argument = args[0]

		if File != "" {
			argument = File
		} else {
			argument = args[0]
		}

		fileExists, err := utils.FileExists(argument)
		if err != nil {
			fmt.Println("File doesn't exist")
			return err
		}
		if fileExists {
			fileName, err = filepath.Abs(argument)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
		} else {
			fmt.Println("File does not exist")
			return errors.New("file does not exist")
		}
		fmt.Println("Ok till here 48")

		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
			return errors.New(err.Error())
		}
		fmt.Println("Ok till here 55", wd)

		utils.UnzipSource(fileName, wd)
		fmt.Println("Ok till here 58 - unzipped folder")

		// fileNameWithoutExtension := strings.Trim(fileName, path.Ext(fileName))
		// err = os.Mkdir(fileNameWithoutExtension, 0777)
		// if err != nil {
		// 	fmt.Println("Error creating the directory :", err.Error())
		// 	return err
		// }

		// err = os.Chdir(fileNameWithoutExtension)
		// if err != nil {
		// 	fmt.Println("Error changing the directory :", err.Error())
		// 	return err
		// }
		// fmt.Println("Changed directory")

		wd, err = os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
			return errors.New(err.Error())
		}
		fmt.Println("Ok till here 78", wd)

		commandCode := exec.Command("code", wd)
		err = commandCode.Run()
		if err != nil {
			fmt.Println("NOT OKLine 83", err.Error())
			return errors.New(err.Error())
		}
		fmt.Println("Probably opened the file by now")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	codeCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "A file name to unzip and open in VSCode")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
