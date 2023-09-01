package cmd

import (
	"strconv"
	"time"

	"github.com/eryajf/gcs/pkg/logger"
	"github.com/xanzy/go-gitlab"
	"github.com/xuri/excelize/v2"
)

var gc *gitlab.Client

func InitGitlabCli(token, url string) {
	var err error
	gc, err = gitlab.NewClient(token, gitlab.WithBaseURL(url+"/api/v4"))
	if err != nil {
		logger.Fatal("init git client err: ", err)
	}
}

func GetAllProject() ([]*gitlab.Project, error) {
	lbo := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 50,
		},
		Simple:   gitlab.Bool(true),
		Archived: gitlab.Bool(false),
	}
	var pro []*gitlab.Project
	for {
		p, _, err := gc.Projects.ListProjects(lbo)
		if err != nil {
			return nil, err
		}
		pro = append(pro, p...)
		if len(p) < 50 {
			break
		}
		lbo.ListOptions.Page++
	}
	return pro, nil
}

func SearchKeyWord(branch, keyword string, pid int) ([]*gitlab.Blob, error) {
	if branch == "" {
		branch = "master"
	}
	sbo := &gitlab.SearchOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 50,
		},
		Ref: gitlab.String(branch),
	}

	var bs []*gitlab.Blob
	for {
		p, _, err := gc.Search.BlobsByProject(pid, keyword, sbo)
		if err != nil {
			return nil, err
		}
		bs = append(bs, p...)
		if len(p) < 50 {
			break
		}
		sbo.ListOptions.Page++
	}
	return bs, nil
}

func GetProjectBydid(pid int) (*gitlab.Project, error) {
	p, _, err := gc.Projects.GetProject(pid, &gitlab.GetProjectOptions{})
	if err != nil {
		return nil, err
	}
	return p, nil
}

type SearchResult struct {
	ProjectId   int
	ProjectName string
	ProjectUrl  string
	FileName    string
	LineUrl     string
	Data        string
}

func outxlsx(msg []SearchResult) error {
	xx := excelize.NewFile()
	_ = xx.SetCellValue("Sheet1", "A1", "项目ID")
	_ = xx.SetCellValue("Sheet1", "B1", "项目名")
	_ = xx.SetCellValue("Sheet1", "C1", "项目地址")
	_ = xx.SetCellValue("Sheet1", "D1", "文件名")
	_ = xx.SetCellValue("Sheet1", "E1", "代码地址")
	_ = xx.SetCellValue("Sheet1", "F1", "具体信息")
	for k, v := range msg {
		_ = xx.SetCellValue("Sheet1", "A"+strconv.Itoa(k+2), v.ProjectId)
		_ = xx.SetCellValue("Sheet1", "B"+strconv.Itoa(k+2), v.ProjectName)
		_ = xx.SetCellValue("Sheet1", "C"+strconv.Itoa(k+2), v.ProjectUrl)
		_ = xx.SetCellValue("Sheet1", "D"+strconv.Itoa(k+2), v.FileName)
		_ = xx.SetCellValue("Sheet1", "E"+strconv.Itoa(k+2), v.LineUrl)
		_ = xx.SetCellValue("Sheet1", "F"+strconv.Itoa(k+2), v.Data)
	}
	err := xx.SaveAs(GetFileName())
	if err != nil {
		return err
	}
	return nil
}

func GetFileName() string {
	return time.Now().Format("2006_01_02_15_04") + ".xlsx"
}
