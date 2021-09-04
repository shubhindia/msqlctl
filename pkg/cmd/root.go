package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var kubeconfigFlag string
var versionFlag bool
var version string

var rootCmd = &cobra.Command{
	Use:   "msqlctl",
	Short: "A simple CLI tool to control mysql instances deployed using mysql operator",
	RunE: func(_ *cobra.Command, args []string) error {
		if versionFlag {
			fmt.Printf("msqlctl version %s\n", version)
		}
		return nil
	},
}

func init() {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	rootCmd.PersistentFlags().StringVar(&kubeconfigFlag, "kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	rootCmd.Flags().BoolVar(&versionFlag, "version", true, "--version")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
