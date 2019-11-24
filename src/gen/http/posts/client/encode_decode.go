// Code generated by goa v3.0.7, DO NOT EDIT.
//
// posts HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/k-yomo/elastic_blog_search/src/design -o src/

package client

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	posts "github.com/k-yomo/elastic_blog_search/src/gen/posts"
	goahttp "goa.design/goa/v3/http"
)

// BuildRegisterRequest instantiates a HTTP request object with method and path
// set to call the "posts" service "register" endpoint
func (c *Client) BuildRegisterRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RegisterPostsPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("posts", "register", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRegisterRequest returns an encoder for requests sent to the posts
// register server.
func EncodeRegisterRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*posts.RegisterPayload)
		if !ok {
			return goahttp.ErrInvalidType("posts", "register", "*posts.RegisterPayload", v)
		}
		req.Header.Set("Authorization", p.Key)
		body := NewRegisterRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("posts", "register", err)
		}
		return nil
	}
}

// DecodeRegisterResponse returns a decoder for responses returned by the posts
// register endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeRegisterResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body int
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("posts", "register", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("posts", "register", resp.StatusCode, string(body))
		}
	}
}

// BuildSearchRequest instantiates a HTTP request object with method and path
// set to call the "posts" service "search" endpoint
func (c *Client) BuildSearchRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SearchPostsPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("posts", "search", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSearchRequest returns an encoder for requests sent to the posts search
// server.
func EncodeSearchRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*posts.SearchPayload)
		if !ok {
			return goahttp.ErrInvalidType("posts", "search", "*posts.SearchPayload", v)
		}
		values := req.URL.Query()
		values.Add("query", p.Query)
		values.Add("page", fmt.Sprintf("%v", p.Page))
		values.Add("pageSize", fmt.Sprintf("%v", p.PageSize))
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeSearchResponse returns a decoder for responses returned by the posts
// search endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeSearchResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body SearchResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("posts", "search", err)
			}
			err = ValidateSearchResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("posts", "search", err)
			}
			res := NewSearchResultOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("posts", "search", resp.StatusCode, string(body))
		}
	}
}

// marshalPostsPostParamsToPostParamsRequestBody builds a value of type
// *PostParamsRequestBody from a value of type *posts.PostParams.
func marshalPostsPostParamsToPostParamsRequestBody(v *posts.PostParams) *PostParamsRequestBody {
	res := &PostParamsRequestBody{
		ID:          v.ID,
		Title:       v.Title,
		Description: v.Description,
		Body:        v.Body,
	}

	return res
}

// marshalPostParamsRequestBodyToPostsPostParams builds a value of type
// *posts.PostParams from a value of type *PostParamsRequestBody.
func marshalPostParamsRequestBodyToPostsPostParams(v *PostParamsRequestBody) *posts.PostParams {
	res := &posts.PostParams{
		ID:          v.ID,
		Title:       v.Title,
		Description: v.Description,
		Body:        v.Body,
	}

	return res
}

// unmarshalPostOutputResponseBodyToPostsPostOutput builds a value of type
// *posts.PostOutput from a value of type *PostOutputResponseBody.
func unmarshalPostOutputResponseBodyToPostsPostOutput(v *PostOutputResponseBody) *posts.PostOutput {
	res := &posts.PostOutput{
		ID:          *v.ID,
		Title:       *v.Title,
		Description: *v.Description,
	}

	return res
}