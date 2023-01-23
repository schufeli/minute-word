package random

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/schufeli/minute-word/internal/word"
)

func GetRandomWord(url string) (word.Word, error) {
	response, err := http.Get(url)

	if err != nil {
		return word.Word{}, err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return word.Word{}, err
	}

	headerword := doc.Find(".dwdswb-ft-lemmaansatz").Text()
	if len(headerword) == 0 {
		return word.Word{}, errors.New("Headerword in dwds.de response was empty")
	}

	grammar := doc.Find("div[data-content-piece='Formteil']").Find(".dwdswb-ft-blocktext").First().Text()

	explanation := doc.Find(".dwdswb-definition").Text()

	example := doc.Find(".dwdswb-kompetenzbeispiel").Find(".dwdswb-belegtext").Text()

	usages := []string{}

	doc.Find("div[data-content-piece='Verwendungsbeispiele']").Find(".dwdswb-belegtext").Each(func(i int, element *goquery.Selection) {
		usages = append(usages, element.Text())
	})

	return word.Word{
			Word:        headerword,
			Grammar:     grammar,
			Explanation: explanation,
			Example:     example,
			Usages:      usages},
		nil
}
