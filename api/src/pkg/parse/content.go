package parse

import (
	"later/pkg/response"
	"log"
	"net/http"
	"regexp"

	"github.com/google/uuid"

	"golang.org/x/net/html"

	"later/pkg/model"
	"later/pkg/util/wrappers"
)

type ContentMetadata struct {
	url         string
	title       *string
	description *string
	imageURL    *string
	contentType *string
}

func (c *ContentMetadata) ToContent(userID uuid.UUID) model.Content {
	return model.NewContent(
		wrappers.NewNullString(c.title),
		wrappers.NewNullString(c.description),
		wrappers.NewNullString(c.imageURL),
		wrappers.NewNullString(c.contentType),
		c.url,
		DomainFromURL(c.url),
		userID,
	)
}

func (c *ContentMetadata) ToContentPreview() response.ContentPreview {
	return response.ContentPreview{
		URL:         c.url,
		Title:       wrappers.NewNullString(c.title),
		Description: wrappers.NewNullString(c.description),
		ImageURL:    wrappers.NewNullString(c.imageURL),
		ContentType: wrappers.NewNullString(c.contentType),
	}
}

type headerContent struct {
	title       *string
	description *string
	imageURL    *string
}

// DomainFromURL extracts the domain from the url
func DomainFromURL(url string) string {
	domainRegex := regexp.MustCompile(`.*[\./]([^\.]+)\.(com|co|org)`)
	matches := domainRegex.FindStringSubmatch(url)
	urlDomain := string(matches[1])

	return urlDomain
}

// ContentFromURL scrapes the data found at the url's address to find elements to populate Content with
// return a content preview, not a content model obj
func ContentFromURL(url string) ContentMetadata {
	var contentMetadata = ContentMetadata{
		url: url,
	}

	switch {
	default:
		contentMetadataDefault(&contentMetadata, url)
	}

	return contentMetadata
}

func contentMetadataDefault(metadata *ContentMetadata, url string) {
	resp, err := http.Get(url)

	if err != nil {
		log.Printf("Error retrieving URL content. Error: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Failed to retrieve URL content, response status %d", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		log.Printf("Failed to parse URL content. Error: %v", err)
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
