package crawlers

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	excludeExtensions = []string{".gif"}
)

// GetSrc 対象URL中の全IMGタグのsrc属性の値を取得して返す
func GetSrc(url string) []string {
	document, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal("ページの取得に失敗しました")
	}

	results := make([]string, 0, 100)
	document.Find("img").Each(func(_ int, selection *goquery.Selection) {
		src, _ := selection.Attr("src")

		// GIFファイルは除外する
		for _, extension := range excludeExtensions {
			if strings.Index(src, extension) > 0 {
				return
			}
		}

		// 同じsrc属性は除外
		for _, result := range results {
			if strings.EqualFold(src, result) {
				return
			}
		}
		results = append(results, src)
	})
	return results
}
