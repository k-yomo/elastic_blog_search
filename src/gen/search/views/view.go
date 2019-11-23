// Code generated by goa v3.0.7, DO NOT EDIT.
//
// search views
//
// Command:
// $ goa gen github.com/k-yomo/elastic_blog_search/src/design -o src/

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// PostCollection is the viewed result type that is projected based on a view.
type PostCollection struct {
	// Type to project
	Projected PostCollectionView
	// View to render
	View string
}

// PostCollectionView is a type that runs validations on a projected type.
type PostCollectionView []*PostView

// PostView is a type that runs validations on a projected type.
type PostView struct {
	// Post's id
	ID *string
	// Post's title
	Title *string
	// Post's description
	Description *string
}

var (
	// PostCollectionMap is a map of attribute names in result type PostCollection
	// indexed by view name.
	PostCollectionMap = map[string][]string{
		"default": []string{
			"id",
			"title",
			"description",
		},
	}
	// PostMap is a map of attribute names in result type Post indexed by view name.
	PostMap = map[string][]string{
		"default": []string{
			"id",
			"title",
			"description",
		},
	}
)

// ValidatePostCollection runs the validations defined on the viewed result
// type PostCollection.
func ValidatePostCollection(result PostCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidatePostCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidatePostCollectionView runs the validations defined on
// PostCollectionView using the "default" view.
func ValidatePostCollectionView(result PostCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidatePostView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidatePostView runs the validations defined on PostView using the
// "default" view.
func ValidatePostView(result *PostView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	if result.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "result"))
	}
	return
}