package parse

import (
	"errors"
	"net/http"
	"regexp"

	"golang.org/x/net/html"

	"later/pkg/model"
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
	valid   bool
	message string
}

// NewContent for wire generation
func NewContent() Content {
	return Content{
		valid: true,
	}
}

// DomainFromURL extracts the domain from the url
func DomainFromURL(url string) string {
	domainRegex := regexp.MustCompile(`.*[\./]([^\.]+)\.(com|co|org)`)
	matches := domainRegex.FindStringSubmatch(url)
	urlDomain := string(matches[1])

	return urlDomain
}

// Err was there an error that occured while parsing
func (parser *Content) Err() error {
	if parser.valid {
		return nil
	}

	return errors.New(parser.message)
}

// ContentFromURL scrapes the data found at the url's address to find elements to populate Content with
func (parser *Content) ContentFromURL(url string, domain *model.Domain) model.Content {
	parser.message = ""
	parser.valid = true
	var contentType *string
	var contentMetadata contentMetadata

	// TODO just pass a pointer in as an argument
	switch {
	default:
		parser.contentMetadataDefault(&contentMetadata, url)
	}

	if domain != nil {
		contentType = &domain.ContentType
	}

	newContent := model.NewContent(
		wrappers.NewNullString(contentMetadata.title),
		wrappers.NewNullString(contentMetadata.description),
		wrappers.NewNullString(contentMetadata.imageURL),
		wrappers.NewNullString(contentType),
		url,
		DomainFromURL(url),
	)

	return newContent
}

func (parser *Content) contentMetadataDefault(metadata *contentMetadata, url string) {
	resp, err := http.Get(url)

	if err != nil {
		parser.valid = false
		parser.message = err.Error()
		return
	}

	if resp.StatusCode != 200 {
		parser.valid = false
		parser.message = "Failed to retrieve URL content"
		return
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		parser.valid = false
		parser.message = "Failed to parse URL content"
		return
	}

	head := findHead(doc)

	headContent := parseHead(head)

	metadata.title = headContent.title
	metadata.description = headContent.description
	metadata.imageURL = headContent.imageURL
}

func findHead(node *html.Node) *html.Node {
	if node.Type == html.ElementNode && node.Data == "head" {
		return node
	}

	if node.FirstChild == nil {
		if node.NextSibling != nil {
			return findHead(node.NextSibling)
		}
		return nil
	}

	return findHead(node.FirstChild)
}

func parseHead(head *html.Node) headerContent {

	var headerContent headerContent

	if head == nil {
		return headerContent
	}

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

	return headerContent
}
