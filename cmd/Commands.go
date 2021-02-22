package cmd

import (
	"bulk-upload-to-consul/consul"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var consulURL string
var domain string
var file string

// PutnewKeyValue : CMD to upload
func PutnewKeyValue() *cobra.Command {

	var local2consulcmd = &cobra.Command{
		Use:   "put",
		Short: "A friendly way to upload large amounts of key/value in Consul ",
	}

	local2consulcmd.Flags().StringVarP(&domain, "domain", "d", "", "Domain ( required  )")

	local2consulcmd.Flags().StringVarP(&file, "file", "f", "", "This file must contain the key values to put on consul with sintax key=value")

	local2consulcmd.RunE = func(command *cobra.Command, args []string) error {

		switch fileExtension := filepath.Ext(file); fileExtension {
		case ".json":
			keyvalues := consul.Unmarshalconfig(file)

			err := consul.PutKeyValueJson(keyvalues, domain)
			if err != nil {
				return err
			}
		case ".txt":
			consulURL, _ := rootCmd.Flags().GetString("consulURL")
			err := consul.PutKeyValue(consulURL, domain, file)
			if err != nil {
				panic(err)
			}
		default:
			panic("Error to get failed, extension not supported ")

		}
		return nil
	}
	return local2consulcmd
}

var rootCmd = &cobra.Command{
	Use:   "local2consul",
	Short: "local2consul is a friendly tool to put new key/values to consul",

	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		cmd.Help()
	},
}

func init() {
	fmt.Print(local2consulMessage)
	rootCmd.AddCommand(PutnewKeyValue())
	rootCmd.PersistentFlags().StringVar(&consulURL, "consulURL", "http://localhost:8500", "Define Consul URL ")

}

// Execute : execute cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)

	}

}

const local2consulMessage = `
 _                    _ ____   ____                      _ 
| |    ___   ___ __ _| |___ \ / ___|___  _ __  ___ _   _| |
| |   / _ \ / __/ _  | | __) | |   / _ \| '_ \/ __| | | | |
| |__| (_) | (_| (_| | |/ __/| |__| (_) | | | \__ \ |_| | |
|_____\___/ \___\__,_|_|_____|\____\___/|_| |_|___/\__,_|_|

`
