package scraping

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ExtractedData struct {
	Title string
	Image []byte
}

// AmazonURL AmazonのURLからExtractedDataをスクレイピング
func AmazonURL(url string) (*ExtractedData, error) {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	// Find the data
	// get Title
	var title string
	doc.Find(".a-size-extra-large").Each(func(_ int, s *goquery.Selection) {
		title = s.Text()
	})

	// get img
	var imgUrl string
	doc.Find(".a-dynamic-image.image-stretch-horizontal.frontImage").Each(func(_ int, s *goquery.Selection) {
		imgUrl, _ = s.Attr("src")
	})
	res, err = http.Get(imgUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("img url status code error: %d %s", res.StatusCode, res.Status)
	}
	img, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	data := ExtractedData{
		Title: strings.TrimSpace(title),
		Image: img,
	}
	return &data, nil
}
