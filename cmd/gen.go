/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen <template_path> <yaml_path>",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		generate(args[0], args[1])
	},
}

var funcMap = template.FuncMap{
	"args":      args,
	"replace":   replace,
	"encodeURL": encodeURL,
	"decodeURL": decodeURL,
	"isEven":    isEven,
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func generate(tmplPath, yamlPath string) {
	f, err := os.Open(yamlPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := make(map[string]interface{})

	raw, err := ioutil.ReadFile(tmplPath)
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("template").Funcs(funcMap).Parse(string(raw))
	if err != nil {
		log.Fatal(err)
	}

	if err = yaml.NewDecoder(f).Decode(data); err != nil {
		log.Fatal(err)
	}

	name := strings.TrimSuffix(tmplPath, filepath.Ext(tmplPath))
	o, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer o.Close()

	if err = t.Execute(o, data); err != nil {
		log.Fatal(err)
	}
}

func args(vs ...interface{}) []interface{} {
	return vs
}

func replace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}

func encodeURL(raw string) string {
	return url.PathEscape(raw)
}

func decodeURL(encoded string) (string, error) {
	return url.PathUnescape(encoded)
}

func isEven(n int) bool {
	return n%2 == 0
}
