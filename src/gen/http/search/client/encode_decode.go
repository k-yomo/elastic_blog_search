// Code generated by goa v3.0.7, DO NOT EDIT.
//
// search HTTP client encoders and decoders
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

	search "github.com/k-yomo/elastic_blog_search/src/gen/search"
	searchviews "github.com/k-yomo/elastic_blog_search/src/gen/search/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildSearchRequest instantiates a HTTP request object with method and path
// set to call the "search" service "search" endpoint
func (c *Client) BuildSearchRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SearchSearchPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("search", "search", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSearchRequest returns an encoder for requests sent to the search
// search server.
func EncodeSearchRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*search.SearchPayload)
		if !ok {
			return goahttp.ErrInvalidType("search", "search", "*search.SearchPayload", v)
		}
		values := req.URL.Query()
		values.Add("query", p.Query)
		if p.Page != nil {
			values.Add("page", fmt.Sprintf("%v", *p.Page))
		}
		if p.PageSize != nil {
			values.Add("pageSize", fmt.Sprintf("%v", *p.PageSize))
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeSearchResponse returns a decoder for responses returned by the search
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
				return nil, goahttp.ErrDecodingError("search", "search", err)
			}
			p := NewSearchPostCollectionOK(body)
			view := "default"
			vres := searchviews.PostCollection{p, view}
			if err = searchviews.ValidatePostCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("search", "search", err)
			}
			res := search.NewPostCollection(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("search", "search", resp.StatusCode, string(body))
		}
	}
}
