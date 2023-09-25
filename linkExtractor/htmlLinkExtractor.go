package linkExtractor

import (
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

type HtmlLinkExtractor struct{}

func (le *HtmlLinkExtractor) ExtractLinks(resp *http.Response) ([]string, error) {
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			link := processAnchorNode(n, resp.Request.URL)
			if link != "" {
				links = append(links, link)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links, nil
}

func processAnchorNode(n *html.Node, baseURL *url.URL) string {
	for _, a := range n.Attr {
		if a.Key != "href" {
			continue
		}
		link := normalizeLink(a.Val, baseURL)
		if isValidLink(link) {
			return link
		}
	}
	return ""
}

func normalizeLink(link string, baseURL *url.URL) string {
	parsedLink, err := url.Parse(link)
	if err != nil {
		return ""
	}

	normalizedLink := baseURL.ResolveReference(parsedLink)
	normalizedLink.Fragment = ""
	if normalizedLink.Path == "" {
		normalizedLink.Path = "/"
	}
	return normalizedLink.String()
}

func isValidLink(link string) bool {
	forbiddenExtensions := []string{".mp3", ".wav", ".ogg", ".png", ".pdf", ".gif"}
	for _, ext := range forbiddenExtensions {
		if strings.Contains(link, ext) {
			return false
		}
	}

	if strings.Contains(link, "mailto://") {
		return false
	}

	if strings.Contains(link, "/rss") || strings.Contains(link, "/feed") {
		return false
	}

	return true
}
