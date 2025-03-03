package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "file info [path]",
		Short: "Get information about files and directories",
		Long:  `fileinfo is a CLI tool that provides detailed information about files and directories`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := "."
			if len(args) > 0 {
				path = args[0]
			}

			recursive, _ := cmd.Flags().GetBool("recursive")
			fileType, _ := cmd.Flags().GetString("type")

			info, err := os.Stat(path)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			green := color.New(color.FgGreen).SprintFunc()
			yellow := color.New(color.FgYellow).SprintFunc()

			if info.IsDir() {
				if recursive {
					count, size := countFilesInDir(path, fileType)
					// fmt.Printf("Directory: %s\n", path)
					// fmt.Printf("Total files: %d\n", count)
					// fmt.Printf("Total size: %s\n", formatSize(size))
					fmt.Printf("%s: %s\n", green("Directory"), path)
					fmt.Printf("%s: %d\n", yellow("Total files"), count)
					fmt.Printf("%s: %s\n", yellow("Total size"), formatSize(size))
				} else {
					// fmt.Printf("Directory: %s\n", path)
					// fmt.Printf("Size: %s\n", formatSize(info.Size()))
					// fmt.Printf("Modified: %s\n", info.ModTime().Format("2006-01-02 15:04:05"))
					// fmt.Printf("Permissions: %s\n", info.Mode())
					fmt.Printf("%s: %s\n", green("Directory"), path)
					fmt.Printf("%s: %s\n", yellow("Size"), formatSize(info.Size()))
					fmt.Printf("%s: %s\n", yellow("Modified"), info.ModTime().Format("2006-01-02 15:04:05"))
					fmt.Printf("%s: %s\n", yellow("Permissions"), info.Mode())
				}
			} else {
				// fmt.Printf("File: %s\n", path)
				// fmt.Printf("Size: %s\n", formatSize(info.Size()))
				// fmt.Printf("Modified: %s\n", info.ModTime().Format("2006-01-02 15:04:05"))
				// fmt.Printf("Permissions: %s\n", info.Mode())
				fmt.Printf("%s: %s\n", green("File"), path)
				fmt.Printf("%s: %s\n", yellow("Size"), formatSize(info.Size()))
				fmt.Printf("%s: %s\n", yellow("Modified"), info.ModTime().Format("2006-01-02 15:04:05"))
				fmt.Printf("%s: %s\n", yellow("Permissions"), info.Mode())
			}
		},
	}

	rootCmd.Flags().BoolP("recursive", "r", false, "Recursively count files in directories")
	rootCmd.Flags().StringP("type", "t", "", "Filter by file extension(e.g., .go, .txt)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func countFilesInDir(path, fileType string) (int, int64) {
	var count int
	var totalSize int64

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if fileType != "" || strings.HasSuffix(filePath, fileType) {
				count++
				totalSize += info.Size()
			}
		}
		return nil

	})
	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
	}
	return count, totalSize
}

func formatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}

	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])

}
