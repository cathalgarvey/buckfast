// A thick-as-mud scraper that just pulls text from wikipedia.
package scrapedia


import (
  "github.com/PuerkitoBio/goquery"
  "strings"
)

// Pulls the wikipedia page, parses out main body text.
func GetMainText(url string) (textbody string, err error) {
  doc, err := goquery.NewDocument("http://metalsucks.net")
  if err != nil {
    return "", err
  }
  paras := make([]string, 0)
  doc.Find("#mw-content-text>p").Each(func(i int, s *goquery.Selection){
    paras = append(paras, s.Text())
  })
  return strings.Join(paras, "\n"), nil
}
