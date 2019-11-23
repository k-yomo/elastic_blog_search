// Code generated by goa v3.0.7, DO NOT EDIT.
//
// register service
//
// Command:
// $ goa gen github.com/k-yomo/elastic_blog_search/src/design -o src/

package register

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// register service registers blog posts to be searched
type Service interface {
	// Register implements register.
	Register(context.Context, []*Post) (res int, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "register"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"register"}

type Post struct {
	// Post's id
	ID *string
	// Post's title
	Title *string
	// Post's description
	Description *string
	// Post's body
	Body *string
}

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorized(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "Unauthorized",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "BadRequest",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
