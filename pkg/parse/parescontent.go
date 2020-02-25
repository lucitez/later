package parse

import (
	"errors"
	"net/http"
	"regexp"

	"golang.org/x/net/html"

	"later.co/pkg/later/content"
	"later.co/pkg/repository/domainrepo"
	"later.co/pkg/util/wrappers"
)

// ContentFromURL scrapes the data found at the url's address to find elements to populate Content with
func ContentFromURL(url string) (*content.Content, error) {

	domainRegex := regexp.MustCompile(`.*[\./]([^\.]+)\.(com|co|org)`)
	matches := domainRegex.FindStringSubmatch(url)
	urlDomain := string(matches[1])

	var contentType *string

	domain, err := domainrepo.ByDomain(urlDomain)

	if err != nil {
		return nil, err
	}

	var title, description, imageURL *string

	switch {
	case domain == nil:
		title, description, imageURL, err = parseContentDefault(url, urlDomain)
	}

	if err != nil {
		return nil, err
	}

	if domain != nil {
		contentType = &domain.ContentType
	}

	newContent, err := content.New(
		*title,
		*wrappers.NewNullString(description),
		*wrappers.NewNullString(imageURL),
		*wrappers.NewNullString(contentType),
		url,
		urlDomain)

	if err != nil {
		return nil, err
	}

	return newContent, nil
}

func parseContentDefault(url string, urlDomain string) (*string, *string, *string, error) {

	resp, err := http.Get(url)

	if resp.StatusCode != 200 {
		return nil, nil, nil, err
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, nil, nil, err
	}

	var title, description, imageURL *string
	var parse func(*html.Node)

	parse = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "head" {
			title, description, imageURL, err = parseHead(node.FirstChild)
		}
		for currNode := node.FirstChild; currNode != nil; currNode = currNode.NextSibling {
			parse(currNode)
		}
	}

	parse(doc)

	if err != nil {
		return nil, nil, nil, err
	}

	if title == nil {
		return nil, nil, nil, errors.New("Title could not be found")
	}

	return title, description, imageURL, nil
}

func parseHead(node *html.Node) (*string, *string, *string, error) {

	var title *string
	var description *string
	var imageURL *string

	for c := node; c != nil; c = c.NextSibling {

		if c.Type == html.ElementNode && c.Data == "title" {
			title = &c.FirstChild.Data
		}

		if c.Type == html.ElementNode && c.Data == "meta" {

			var name string
			var property string
			var content *string

			for _, attr := range c.Attr {
				switch attr.Key {
				case "name":
					name = attr.Val
				case "property":
					property = attr.Val
				case "content":
					content = &attr.Val
				}
			}

			switch {
			case name == "description":
				description = content
			case property == "og:image":
				imageURL = content
			}
		}
	}

	return title, description, imageURL, nil
}
