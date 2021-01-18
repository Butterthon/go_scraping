package tools

import (
	"log"

	"github.com/tealeg/xlsx"
)

// Excel エクセルファイルの構造体
type Excel struct {
	file  *xlsx.File // エクセルファイル
	row   *xlsx.Row  // 行情報
	sheet *xlsx.Sheet
}

// CreateExcel エクセルファイルを作成する
func CreateExcel() Excel {
	excel := &Excel{}

	excel.file = xlsx.NewFile()

	// シートの作成
	sheet, err := excel.file.AddSheet("sheet1")
	if err != nil {
		log.Fatal("シートの作成に失敗しました", err)
	}
	excel.sheet = sheet

	// 保存
	excel.Save()

	return *excel
}

// NewRow 行追加
func (excel *Excel) NewRow() {
	excel.row = excel.sheet.AddRow()
}

// AddValue 値の書き込み
func (excel *Excel) AddValue(value string) {
	cell := excel.row.AddCell()
	cell.Value = value
}

// Save エクセルファイルの保存
func (excel *Excel) Save() {
	err := excel.file.Save("sample.xlsx")
	if err != nil {
		log.Fatal("エクセルファイルの作成に失敗しました")
	}
}
