package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// lookCmd represents the look command
var lookCmd = &cobra.Command{
	Use:   "look",
	Short: "Perform advanced search and replace operations",
	Long: `Perform search and replace related operations
Pass path using -p or --path`,
	Run: func(cmd *cobra.Command, args []string) {
		//Getting data of the flag passed
		filePath, _ := cmd.Flags().GetString("path")
		isRecursive, _ := cmd.Flags().GetBool("recursive")
		if filePath == "" {
			log.Fatal("Empty file path passed")
		}
		//	if recursive is passed, read all files under that directory and subdirectory
		//  else list the files present only in that particular directory
		//  else a single path is passed
		if isDir(filePath) {
			fmt.Println(listAllFiles(filePath, isRecursive))
		}

	},
}

func init() {
	// Initializing root command
	rootCmd.AddCommand(lookCmd)

	//Adding a flag to the command
	lookCmd.Flags().StringP("path", "p", "", "File path in which search needs to be done")
	lookCmd.Flags().BoolP("recursive", "r", false, "Read files recursively")
}

//returns if passed path is a file or a directory
func isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	return fileInfo.IsDir()
}

//list files recursively
func listAllFiles(path string, isRecursive bool) []string {
	var fileList []string
	l, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	if isRecursive {
		for _, file := range l {
			if file.IsDir() {
				recList := listAllFiles(filepath.Join(path, file.Name()), true)
				fileList = append(fileList, recList...)
			} else {
				fileList = append(fileList, filepath.Join(path, file.Name()))
			}
		}
	} else {
		for _, file := range l {
			if !file.IsDir() {
				fileList = append(fileList, filepath.Join(path, file.Name()))
			}
		}
	}
	return fileList
}

//Read file of a given path
func readFile(path string) {

}
