package model

import (
	"github.com/bxcodec/faker/v3"
	"time"
)

type Contact struct {
	Id                   int    `json:"contact_id"`
	Name                 string `json:"contact_name" faker:"name"`
	Email                string `json:"contact_email" faker:"email"`
	OutstandingDocuments uint   `json:"contact_outstanding_documents" faker:"oneof: 0, 1, 2, 3, 4"`
	Created              string `json:"contact_created" faker:"timestamp"`
	//Apps                 App    `json:"contact_apps" faker:"-"`
}

type App struct {
	CapsuleId  string `json:"capsule_id" faker:"oneof: 0, 1, 2, 3, 4"`
	KashflowId string `json:"kashflow_id" faker:"oneof: 0, 1, 2, 3, 4"`
}

// ContactListResponse response payload for the list endpoint
type ContactListResponse struct {
	PaginateResponse
	TotalContacts int       `json:"total_contacts"`
	Contacts      []Contact `json:"contacts"`
}

// ContactCreateResponse response payload for the create endpoint
type ContactCreateResponse struct {
	CreateResponse
	Id      int       `json:"contact_id"`
	Name    string    `json:"contact_name"`
	Email   string    `json:"contact_email"`
	Created time.Time `json:"contact_created"`
}

// ContactReadResponse response payload for the read endpoint
type ContactReadResponse struct {
	ReadResponse
	Contact
}

// ContactUpdateResponse response payload for the update endpoint
type ContactUpdateResponse struct {
	UpdateResponse
	Id      int       `json:"contact_id"`
	Name    string    `json:"contact_name"`
	Email   string    `json:"contact_email"`
	Created time.Time `json:"contact_created"`
}

// ContactDeleteResponse response payload for the delete endpoint
type ContactDeleteResponse struct {
	DeleteResponse
	Id    int    `json:"contact_id"`
	Name  string `json:"contact_name"`
	Email string `json:"contact_email"`
}

type ContactEnvelopeListResponse struct {
	// TODO...
}

// InitContact sets up the data fixtures for contacts
func InitContact() error {
	contact := Contact{}
	err := faker.FakeData(&contact)

	return err
}
