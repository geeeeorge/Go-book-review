// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Defines values for BookStatus.
const (
	Doing BookStatus = "doing"
	Done  BookStatus = "done"
	Want  BookStatus = "want"
	Will  BookStatus = "will"
)

// Book defines model for book.
type Book struct {
	// BookId uuid
	BookId    string              `json:"book_id"`
	Image     *openapi_types.File `json:"image,omitempty"`
	Status    *BookStatus         `json:"status,omitempty"`
	Summaries *Summaries          `json:"summaries,omitempty"`
	Tags      *Tags               `json:"tags,omitempty"`
	Title     string              `json:"title"`
}

// BookStatus defines model for book-status.
type BookStatus string

// Books defines model for books.
type Books = []Book

// Login defines model for login.
type Login struct {
	// Token jwtのトークン
	Token *string `json:"token,omitempty"`
}

// Summaries defines model for summaries.
type Summaries = []Summary

// Summary defines model for summary.
type Summary struct {
	Content string `json:"content"`

	// Id uuid
	Id string `json:"id"`
}

// Tag defines model for tag.
type Tag struct {
	// Id uuid
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Tags defines model for tags.
type Tags = []Tag

// BookId defines model for book-id.
type BookId = string

// QueryBookId defines model for query-book-id.
type QueryBookId = string

// Status defines model for status.
type Status = BookStatus

// SummaryId defines model for summary-id.
type SummaryId = string

// TagId defines model for tag-id.
type TagId = string

// LoginResponse defines model for login-response.
type LoginResponse = Login

// SummariesResponse defines model for summaries-response.
type SummariesResponse = Summaries

// SummaryResponse defines model for summary-response.
type SummaryResponse = Summary

// TagResponse defines model for tag-response.
type TagResponse = Tag

// TagsResponse defines model for tags-response.
type TagsResponse = Tags

// AuthenticationRequest defines model for authentication-request.
type AuthenticationRequest struct {
	// Password password
	Password *string `json:"password,omitempty"`

	// Username username
	Username *string `json:"username,omitempty"`
}

// BookRequest defines model for book-request.
type BookRequest struct {
	// AmazonUrl amazon_url
	AmazonUrl *string `json:"amazon_url,omitempty"`
}

// PutSummaryRequest defines model for put-summary-request.
type PutSummaryRequest struct {
	// Content summary
	Content *string `json:"content,omitempty"`
}

// PutTagRequest defines model for put-tag-request.
type PutTagRequest struct {
	// Name tag
	Name *string `json:"name,omitempty"`
}

// SummaryRequest defines model for summary-request.
type SummaryRequest struct {
	// BookId book_id
	BookId *string `json:"book_id,omitempty"`

	// Content summary
	Content *string `json:"content,omitempty"`
}

// TagRequest defines model for tag-request.
type TagRequest struct {
	// Name tag
	Name *string `json:"name,omitempty"`
}

// PostBooksJSONBody defines parameters for PostBooks.
type PostBooksJSONBody struct {
	// AmazonUrl amazon_url
	AmazonUrl *string `json:"amazon_url,omitempty"`
}

// PutBookParams defines parameters for PutBook.
type PutBookParams struct {
	Status Status `form:"status" json:"status"`
}

// GetSummariesParams defines parameters for GetSummaries.
type GetSummariesParams struct {
	BookId QueryBookId `form:"book_id" json:"book_id"`
}

// PostSummariesJSONBody defines parameters for PostSummaries.
type PostSummariesJSONBody struct {
	// BookId book_id
	BookId *string `json:"book_id,omitempty"`

	// Content summary
	Content *string `json:"content,omitempty"`
}

// PutSummaryJSONBody defines parameters for PutSummary.
type PutSummaryJSONBody struct {
	// Content summary
	Content *string `json:"content,omitempty"`
}

// PostTagsJSONBody defines parameters for PostTags.
type PostTagsJSONBody struct {
	// Name tag
	Name *string `json:"name,omitempty"`
}

// PutTagJSONBody defines parameters for PutTag.
type PutTagJSONBody struct {
	// Name tag
	Name *string `json:"name,omitempty"`
}

// LoginJSONBody defines parameters for Login.
type LoginJSONBody struct {
	// Password password
	Password *string `json:"password,omitempty"`

	// Username username
	Username *string `json:"username,omitempty"`
}

// SignupJSONBody defines parameters for Signup.
type SignupJSONBody struct {
	// Password password
	Password *string `json:"password,omitempty"`

	// Username username
	Username *string `json:"username,omitempty"`
}

// PostBooksJSONRequestBody defines body for PostBooks for application/json ContentType.
type PostBooksJSONRequestBody PostBooksJSONBody

// PostSummariesJSONRequestBody defines body for PostSummaries for application/json ContentType.
type PostSummariesJSONRequestBody PostSummariesJSONBody

// PutSummaryJSONRequestBody defines body for PutSummary for application/json ContentType.
type PutSummaryJSONRequestBody PutSummaryJSONBody

// PostTagsJSONRequestBody defines body for PostTags for application/json ContentType.
type PostTagsJSONRequestBody PostTagsJSONBody

// PutTagJSONRequestBody defines body for PutTag for application/json ContentType.
type PutTagJSONRequestBody PutTagJSONBody

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody LoginJSONBody

// SignupJSONRequestBody defines body for Signup for application/json ContentType.
type SignupJSONRequestBody SignupJSONBody
