package main

import (
	"fmt"
	"go-spider/domain/dongqiudi"
	"go-spider/domain/gushiwen"
	"go-spider/infra/initial"
	"go-spider/src/model"
)

func init() {
	initial.DBInit()
	CreateTable()
	Start()

}

func CreateTable() {
	initial.DataBase.AutoMigrate(
		// 古诗文
		&model.RankListCBOInfo{},
		&model.Dynasty{},
		&model.PoetryType{},
		&model.Poet{},
		&model.PoetryInfo{},
		// github 仓库
		&model.Repositories{},
		// pexles
		&model.ImageSize{},
		&model.ImageAddress{},
	)
}
func Start() {
	{
		//cbooo.SaveCBOIntoDB()
	}

	{
		// 古诗文抓取
		var engine gushiwen.SimpleEngine
		var results []gushiwen.Request
		var rootURL = "https://www.gushiwen.org/shiwen/default.aspx?page=%d&type=0&id=0"
		for index := 1; index <= 50; index++ {
			url := fmt.Sprintf(rootURL, index)
			results = append(results, gushiwen.Request{
				URL: url,
			})
		}
		engine.Run(
			results,
		)

	}

}

func main() {
	//defer initial.DataBase.Close()
	//meizitu.Start()
	dongqiudi.StartDongQiuDi()
	//githubtrending.TrendingStart()
	//pexels.Start()
	//cmd.Execute()
	//Start()

}
