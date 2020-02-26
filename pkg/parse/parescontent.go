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

type contentMetadata struct {
	title       *string
	description *string
	imageURL    *string
}

type headerContent struct {
	title       *string
	description *string
	imageURL    *string
}

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

	var contentMetadata *contentMetadata

	switch {
	case domain == nil:
		contentMetadata, err = contentMetadataDefault(url, urlDomain)
	}

	if err != nil {
		return nil, err
	}

	if domain != nil {
		contentType = &domain.ContentType
	}

	newContent, err := content.New(
		*contentMetadata.title,
		*wrappers.NewNullString(contentMetadata.description),
		*wrappers.NewNullString(contentMetadata.imageURL),
		*wrappers.NewNullString(contentType),
		url,
		urlDomain)

	if err != nil {
		return nil, err
	}

	return newContent, nil
}

// TODO dont parse entire doc, just go to header and go thru all meta tags
func contentMetadataDefault(url string, urlDomain string) (*contentMetadata, error) {

	resp, err := http.Get(url)

	if resp.StatusCode != 200 {
		return nil, err
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var headContent headerContent
	var parse func(*html.Node)

	parse = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "head" {
			headContent, err = parseHead(node.FirstChild)
		}
		for currNode := node.FirstChild; currNode != nil; currNode = currNode.NextSibling {
			parse(currNode)
		}
	}

	parse(doc)

	if err != nil {
		return nil, err
	}

	if headContent.title == nil {
		return nil, errors.New("Title could not be found")
	}

	contentMetadata := contentMetadata{
		title:       headContent.title,
		description: headContent.description,
		imageURL:    headContent.imageURL}

	return &contentMetadata, nil
}

func parseHead(node *html.Node) (headerContent, error) {

	var headerContent headerContent

	for currNode := node; currNode != nil; currNode = currNode.NextSibling {

		if currNode.Type == html.ElementNode && currNode.Data == "title" {
			headerContent.title = &currNode.FirstChild.Data
		}

		if currNode.Type == html.ElementNode && currNode.Data == "meta" {

			var name string
			var property string
			var content *string

			for _, attr := range currNode.Attr {
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
				headerContent.description = content
			case property == "og:image":
				headerContent.imageURL = content
			}
		}
	}

	return headerContent, nil
}
