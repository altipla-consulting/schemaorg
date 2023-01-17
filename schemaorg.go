package schemaorg

import (
	"encoding/json"
	"fmt"
)

const (
	Context = DescriptorContext("http://schema.org")

	TypeNewsArticle        = DescriptorType("NewsArticle")
	TypeBreadcrumbList     = DescriptorType("BreadcrumbList")
	TypeOpinionNewsArticle = DescriptorType("OpinionNewsArticle")
	TypeLiveBlogPosting    = DescriptorType("LiveBlogPosting")

	AuthorPerson = AuthorType("Person")

	ImageObject = ImageType("ImageObject")

	PublisherOrganization = PublisherType("Organization")

	BreadcrumbItem = BreadcrumbElementType("ListItem")

	PaywallType = CreativeWorkType("WebPageElement")

	ContentNotFree = AccesibleForFree("False")
	ContentFree    = AccesibleForFree("True")
)

type DescriptorContext string
type DescriptorType string

type Descriptor struct {
	Context DescriptorContext `json:"@context"`
	Type    DescriptorType    `json:"@type"`
}

type AuthorType string
type ImageType string
type PublisherType string
type BreadcrumbElementType string
type CreativeWorkType string
type AccesibleForFree string

type NewsArticle struct {
	Descriptor
	Headline            string           `json:"headline"`
	AlternativeHeadline string           `json:"alternativeHeadline"`
	Image               *Image           `json:"image,omitempty"`
	DatePublished       string           `json:"datePublished"`
	DateModified        string           `json:"dateModified,omitempty"`
	Description         string           `json:"description"`
	Author              *Author          `json:"author,omitempty"`
	Publisher           *Publisher       `json:"publisher"`
	MainEntity          string           `json:"mainEntityOfPage"`
	AccesibleForFree    AccesibleForFree `json:"isAccesibleForFree,omitempty"`
	Paywall             *Paywall         `json:"hasPart,omitempty"`
}

type Image struct {
	Type   ImageType `json:"@type"`
	URL    string    `json:"url"`
	Width  int32     `json:"width"`
	Height int32     `json:"height"`
}

type Author struct {
	Type AuthorType `json:"@type"`
	Name string     `json:"name"`
}

type Publisher struct {
	Type PublisherType `json:"@type"`
	Name string        `json:"name"`
	URL  string        `json:"url"`
	Logo *Image        `json:"logo"`
}

type BreadcrumbList struct {
	Descriptor
	Elements []*BreadcrumbElement `json:"itemListElement"`
}

type BreadcrumbElement struct {
	Type     BreadcrumbElementType `json:"@type"`
	Position int32                 `json:"position"`
	Name     string                `json:"name"`
	Item     string                `json:"item"`
}

type Paywall struct {
	Type             CreativeWorkType `json:"@type"`
	AccesibleForFree AccesibleForFree `json:"isAccessibleForFree"`
	Selector         string           `json:"cssSelector"`
}

func RenderString(thing any) (string, error) {
	data, err := json.Marshal(thing)
	if err != nil {
		return "", fmt.Errorf("cannot render schema.org: %w", err)
	}
	return string(data), nil
}
