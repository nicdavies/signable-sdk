package model

import (
	"github.com/bxcodec/faker/v3"
)

var Brandings []Branding

type BrandingOption struct {
	Logo   string `json:"branding_logo" faker:""`
	Colour string `json:"branding_colour" faker:""`
}

type BrandingEmail struct {
	Type         string `json:"branding_email_type" faker:""`
	Subject      string `json:"branding_email_subject" faker:""`
	Content      string `json:"branding_email_content" faker:""`
	ContentPlain string `json:"branding_email_content_plain" faker:""`
	Updated      string `json:"branding_email_updated" faker:"timestamp"`
}

type BrandingOptionsListResponse struct {
	ReadResponse
	BrandingOption
}

type BrandingEmailsListResponse struct {
	ListResponse
	BrandingEmails []BrandingEmail `json:"branding_emails"`
}

type BrandingOptionUpdateResponse struct {
	UpdateResponse
	BrandingOption
}

type BrandingEmailUpdateResponse struct {
	UpdateResponse
	BrandingEmail
}

type Branding struct {
	BrandingOption []BrandingOption
	BrandingEmail  []BrandingEmail
}

func InitBranding() ([]Branding, error) {
	// create 20 contacts
	for i := 0; i < 20; i++ {
		b := Branding{}

		err := faker.FakeData(&b)
		if err != nil {
			return Brandings, err
		}

		// add the newly created contact to the slice
		Brandings = append(Brandings, b)
	}

	return Brandings, nil
}
