/*
Copyright © 2021 eryajf

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

// var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gcs",
	Short: "gitlab code search",
	Long:  `通过关键字搜索gitlab所有匹配的项目。`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// 获取配置文件
	rootCmd.AddCommand(GetConfigCmd)
	cset := GetConfigCmd.Flags()
	cset.StringP("branch", "b", "master", "每个项目检索的分支，默认为master")
	cset.StringP("word", "w", "", "检索的关键字")
	cset.StringP("url", "u", "", "gitlab的地址")
	cset.StringP("token", "t", "", "gitlab的认证token")
	_ = GetConfigCmd.MarkFlagRequired("word")
	_ = GetConfigCmd.MarkFlagRequired("url")
	_ = GetConfigCmd.MarkFlagRequired("token")
}
