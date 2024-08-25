/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"log/slog"

	"github.com/iancoleman/strcase"
	"github.com/juju/errors"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "b2 uploader",
	Long:  `b2 uploader`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		if *bucket == "" {
			slog.Error("bucket name is required")
			fmt.Println(cmd.UsageString())
			return
		}
		if *folder == "" {
			slog.Error("folder name is required")
			fmt.Println(cmd.UsageString())
			return
		}

		lsOut, err := exec.CommandContext(ctx, "b2", "ls", "--recursive", "b2://"+*bucket).CombinedOutput()
		if err != nil {
			slog.Error("b2 ls failed", "error", err)
			return
		}
		existingFiles := strings.Split(string(lsOut), "\n")

		fp, un, err := recursiveFolderLS(*folder)
		if err != nil {
			slog.Error("recursiveFolderLS err", "error", err)
			return
		}
		group, ctx := errgroup.WithContext(ctx)
		for i := range fp {
			group.Go(func() error {
				f, u := fp[i], un[i]
				if slices.Contains(existingFiles, u) {
					slog.Info("file already exists", "file", f)
					return nil
				}
				command := exec.CommandContext(ctx, "b2", "file", "upload", *bucket, f, u)
				if *dryRun {
					slog.Info("dry run", "command", command.String())
					return nil
				}
				out, err := command.CombinedOutput()
				if err != nil {
					return errors.Annotatef(err, "file upload failed: %s", f)
				}
				fmt.Printf("File uploaded:\n%s\n", string(out))
				return nil
			})
		}
		if err := group.Wait(); err != nil {
			slog.Error("file upload failed", "error", err)
			return
		}
		if *dryRun {
			slog.Info("dry run")
		} else {
			slog.Info("all files uploaded")
		}
	},
}

func recursiveFolderLS(folder string) (filepaths, uploadnames []string, err error) {
	entries, err := os.ReadDir(folder)
	if err != nil {
		return
	}
	filepaths = slices.Grow(filepaths, len(entries))
	uploadnames = slices.Grow(uploadnames, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			subfolder := filepath.Join(folder, entry.Name())
			fp, un, err := recursiveFolderLS(subfolder)
			if err != nil {
				return nil, nil, err
			}
			filepaths = append(filepaths, fp...)
			uploadnames = append(uploadnames, un...)
		} else {
			fp := filepath.Join(folder, entry.Name())
			var un string
			if *snake {
				un = strcase.ToSnakeWithIgnore(fp, ".")
			}
			if *replaceStr != "" {
				split := strings.Split(*replaceStr, "/")
				if len(split) != 2 {
					err = fmt.Errorf("replace string must have 2 parts separated by /: %s", *replaceStr)
					return
				}
				before := split[0]
				after := split[1]
				un = strings.ReplaceAll(un, before, after)
			}
			if len(*skips) > 0 {
				for _, skipStr := range *skips {
					if strings.Contains(fp, skipStr) {
						continue
					}
				}
			}
			filepaths = append(filepaths, fp)
			uploadnames = append(uploadnames, un)

		}
	}
	return
}

var (
	bucket     *string
	folder     *string
	replaceStr *string
	snake      *bool
	dryRun     *bool
	skips      *[]string
)

func init() {
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	bucket = uploadCmd.Flags().StringP("bucket", "b", "", "bucket name")
	folder = uploadCmd.Flags().StringP("folder", "f", "", "folder name")
	snake = uploadCmd.Flags().Bool("snake", false, "convert upload filepath to snake_case")
	replaceStr = uploadCmd.Flags().StringP("replace", "r", "", "replace string before/after")
	skips = uploadCmd.Flags().StringArrayP("skip", "s", []string{}, "skip files that contain these strings")
	dryRun = uploadCmd.Flags().Bool("dry", false, "dry run")
}
