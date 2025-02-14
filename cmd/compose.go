/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/cobra"
)

// composeCmd represents the compose command
var composeCmd = &cobra.Command{
	Use:   "compose",
	Short: "Sets up composition of your applications",
	Long:  ``,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	ps := prompts.Prompts {
	// 		"serviceName": &prompts.UserInput{
	// 			Label: "name of service",
	// 		},
	// 		"image": &prompts.UserInput{
	// 			Label: "image name[:tag]",
	// 		},
	// 	}
	// 	answers := ps.Run()
	// 	fmt.Print(answers)

	// 	dc, _ := configs.NewDockerCompose()

	// 	dc.Services[answers["serviceName"]] = &configs.Service{
	// 		Image: answers["image"],
	// 	}

	// 	configs.Render("docker-compose.yml", dc)
	// },
}

func init() {
	rootCmd.AddCommand(composeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// composeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// composeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
