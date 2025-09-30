package asset

import (
	"NotMe/config"
	"NotMe/utils"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func loadXlsx(cfg []config.Xlsx) {

	utils := utils.Asset
	for _, c := range cfg {
		f, err := excelize.OpenFile(c.Path, excelize.Options{
			Password: c.PASS,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println(err)
				return
			}
		}()

		cols, err := f.GetCols(c.Sheet)
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(cols) != 0 {
			for i, cell := range cols[utils.ColLetterToIndex(c.Col)] {
				if utils.IsValidIPv4(cell) {
					if isTrue := addIP(cell, cols[utils.ColLetterToIndex(c.DescCol)][i]); isTrue == false {
						fmt.Println("添加IP失败:", cell)
					}
					continue
				}
				if utils.IsValidCIDR(cell) {
					if isTrue := addCIDR(cell, cols[utils.ColLetterToIndex(c.DescCol)][i]); isTrue == false {
						fmt.Println("添加CIDR失败:", cell)
					}
					continue
				}
				// TODO: 未识别的处理逻辑
			}
		}
	}
}
