package cmd

import (
	"strconv"

	"github.com/eryajf/gcs/pkg/logger"
	"github.com/spf13/cobra"
)

var GetConfigCmd = &cobra.Command{
	Use:   "search",
	Short: "通过关键字搜索gitlab所有匹配的项目。",
	Run: func(cmd *cobra.Command, args []string) {
		word, _ := cmd.Flags().GetString("word")
		url, _ := cmd.Flags().GetString("url")
		token, _ := cmd.Flags().GetString("token")
		branch, _ := cmd.Flags().GetString("branch")

		InitGitlabCli(token, url)
		projects, err := GetAllProject()
		if err != nil {
			logger.Fatal("获取所有项目失败: ", err)
		}
		var dtmp []SearchResult
		for _, project := range projects {
			bs, err := SearchKeyWord(branch, word, project.ID)
			if err != nil {
				logger.Error("搜索失败: ", err)
			}
			if len(bs) != 0 {
				for _, v := range bs {
					dtmp = append(dtmp, SearchResult{
						ProjectId:   project.ID,
						ProjectName: project.Name,
						ProjectUrl:  project.WebURL,
						FileName:    v.Filename,
						LineUrl:     project.WebURL + "/-/blob/master/" + v.Filename + "#L" + strconv.Itoa(v.Startline),
						Data:        v.Data,
					})
				}
			}
		}
		err = outxlsx(dtmp)
		if err != nil {
			logger.Error("将内容写入到Excel失败: ", err)
		} else {
			logger.Info("🥳 搜索结果已生成，请查看最新的表格。")
		}
	},
}
