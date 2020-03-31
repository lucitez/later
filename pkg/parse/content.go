package parse

import (
	"errors"
	"net/http"
	"regexp"

	"golang.org/x/net/html"

	"later/pkg/model"
	"later/pkg/service"
	"later/pkg/util/wrappers"
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

// Content handles parsing html at a url to extract content data
type Content struct {
	DomainManager service.DomainManager
}

// NewContent for wire generation
func NewContent(domainManager service.DomainManager) Content {
	return Content{DomainManager: domainManager}
}

// ContentFromURL scrapes the data found at the url's address to find elements to populate Content with
func (parser *Content) ContentFromURL(url string) (*model.Content, error) {

	domainRegex := regexp.MustCompile(`.*[\./]([^\.]+)\.(com|co|org)`)
	matches := domainRegex.FindStringSubmatch(url)
	urlDomain := string(matches[1])

	var contentType *string

	domain, err := parser.DomainManager.ByDomain(urlDomain)

	if err != nil {
		return nil, err
	}

	var contentMetadata *contentMetadata

	// TODO just pass a pointer in as an argument
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

	newContent, err := model.NewContent(
		*contentMetadata.title,
		wrappers.NewNullString(contentMetadata.description),
		wrappers.NewNullString(contentMetadata.imageURL),
		wrappers.NewNullString(contentType),
		url,
		urlDomain)

	if err != nil {
		return nil, err
	}

	return newContent, nil
}

func contentMetadataDefault(url string, urlDomain string) (*contentMetadata, error) {
	resp, err := http.Get(url)

	if resp.StatusCode != 200 {
		return nil, err
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		return nil, err
	}

	head, err := findHead(doc)

	if err != nil {
		return nil, err
	}

	headContent, err := parseHead(head)

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

func findHead(node *html.Node) (*html.Node, error) {

	if node.Type == html.ElementNode && node.Data == "head" {
		return node, nil
	}

	if node.FirstChild == nil {
		if node.NextSibling != nil {
			return findHead(node.NextSibling)
		}
		return nil, errors.New("Head could not be found")
	}

	return findHead(node.FirstChild)
}

func parseHead(head *html.Node) (headerContent, error) {

	var headerContent headerContent

	for currNode := head.FirstChild; currNode != nil; currNode = currNode.NextSibling {

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
