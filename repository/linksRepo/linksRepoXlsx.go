package linksRepo

import (
	"fmt"
	"links-crawler/models"

	"github.com/xuri/excelize/v2"
)

type XlsxLinksRepo struct{}

func (linksRepo *XlsxLinksRepo) StoreLink(link models.Link) error {
	return nil
}

func (linksRepo *XlsxLinksRepo) StoreLinks(links []models.Link) error {
	f, err := excelize.OpenFile("public/links.xlsx")
	if err != nil {
		f = excelize.NewFile()
	}
	index := f.NewSheet("Sheet1")

	for _, link := range links {
		f.SetCellValue("Sheet1", "A2", link.Url)
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("public/links.xlsx"); err != nil {
		fmt.Println(err)
	}

	return nil
}

func (linksRepo *XlsxLinksRepo) IsLinkExist(href string) bool {
	return false
}

func (linksRepo *XlsxLinksRepo) All() []models.Link {
	return nil
}
