package main

import (
	"github.com/tabata/go_scraping/crawlers"
	"github.com/tabata/go_scraping/tools"
)

var urls = []string{
	"https://www.irasutoya.com/",
}

func main() {
	// エクセルファイル作成
	excel := tools.CreateExcel()

	// クローリング開始
	for i := range urls {
		results := crawlers.GetSrc(urls[i])
		for _, src := range results {
			// 行追加
			excel.NewRow()

			// src属性書き込み
			excel.AddValue(src)
		}
	}

	// 保存
	excel.Save()
}
