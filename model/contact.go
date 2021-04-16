package model

import "time"

type Contact struct {
	Id                   int       `json:"contact_id"`
	Name                 string    `json:"contact_name"`
	Email                string    `json:"contact_email"`
	OutstandingDocuments int       `json:"contact_outstanding_documents"`
	Created              time.Time `json:"contact_created"`
	Apps                 App       `json:"contact_apps"`
}

type App struct {
	CapsuleId  string `json:"capsule_id"`
	KashflowId string `json:"kashflow_id"`
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

// InitContacts sets up the data fixtures for contacts
func InitContacts() {
	// TODO: create contacts
}
