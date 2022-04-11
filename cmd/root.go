/*

 */
package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"net/url"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var (
	urlString string

	rootCmd = &cobra.Command{
		Use:   "dc-audit",
		Short: "This application will scan a docker compose file and list the installed software.",
		Long:  `This application will scan a docker compose file and recursively list the installed software including the pulled docker image.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	dcUrl, err := url.Parse(urlString)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(dcUrl.String())
	/***************************************************/

	resp, err := http.Get(dcUrl.String())
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	composeText := buf.String()

	defer resp.Body.Close()

	//fmt.Println(composeText)

	data := make(map[interface{}]interface{})

	err2 := yaml.Unmarshal([]byte(composeText), &data)
	if err2 != nil {
		log.Fatal(err2)
	}

	for k, v := range data {
		fmt.Printf("%s -> %d\n", k, v)
	}

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dc-audit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVar(&urlString, "url", "", "Compose file URL")
}
