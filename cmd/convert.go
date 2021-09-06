/*
Copyright © 2021 Kiefer Chang <kiefer.chang@suse.com>

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
	"fmt"
	"io/ioutil"
	"os"

	"github.com/harvester/harvester-installer/pkg/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert Harvester config to cOS config",
	Long:  "Convert Harvester config to cOS config",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runConvert(args[0], args[1]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
	Args: cobra.ExactArgs(2),
}

func init() {
	rootCmd.AddCommand(convertCmd)
}

func runConvert(fromFile, toFile string) error {
	b, err := ioutil.ReadFile(fromFile)
	if err != nil {
		return err
	}
	harvConfig, err := config.LoadHarvesterConfig(b)
	if err != nil {
		return err
	}

	cosConfig, err := config.ConvertToCOS(harvConfig)
	if err != nil {
		return err
	}

	b, err = yaml.Marshal(cosConfig)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(toFile, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
