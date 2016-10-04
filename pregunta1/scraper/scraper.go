package scraper

import (
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

const (
	url = "http://www.sii.cl/pagina/valores/dolar/dolar2015.htm"
)

var (
	commaToDot   = strings.NewReplacer(",", ".")
	summaryRegex = regexp.MustCompile("(?i)promedio")
)

// GetCSV returns a CSV string with the USD value in CLP currency
// for each day on an specific year.
// Each column represents the month of the year and each line
// represents the day of month.
func GetCSV() (string, error) {
	res, err := makeRequest()
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		return "", err
	}

	csv := ""
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "tr" && len(n.Attr) > 0 {
			title := n.FirstChild.NextSibling.FirstChild.Data
			if summaryRegex.MatchString(title) {
				return
			}
			var vals []string
			for s := n.FirstChild; s != nil; s = s.NextSibling {
				if s.Type == html.ElementNode && s.Data == "td" {
					value := commaToDot.Replace(strings.TrimSpace(s.FirstChild.Data))
					vals = append(vals, value)
				}
			}

			csv += strings.Join(vals, ",")
			csv += "\n"
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return strings.TrimSpace(csv), nil
}

func makeRequest() (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}
