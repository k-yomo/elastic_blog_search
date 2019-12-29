package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/k-yomo/elastic_blog_search/src/gen/posts"
	"github.com/pkg/errors"
	"goa.design/goa/v3/security"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// posts service implementation.
type postssrvc struct {
	logger   *log.Logger
	esClient *elasticsearch.Client
}

// NewPosts returns the posts service implementation.
func NewPostsService(logger *log.Logger, esClient *elasticsearch.Client) posts.Service {
	return &postssrvc{logger, esClient}
}

const PostsIndex = "posts"

type indexParams struct {
	Index *index `json:"index"`
}

type index struct {
	Index string `json:"_index"`
	Type  string `json:"_type"`
	ID    string `json:"_id"`
}

type post struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

// registers blog posts to be searched
func (p *postssrvc) Register(ctx context.Context, payload *posts.RegisterPayload) (res int, err error) {
	var bulkIndexParamsByte []byte
	for _, p := range payload.Posts {
		post := &post{ID: *p.ID, Title: *p.Title, Description: *p.Description, Body: *p.Body}
		postIndexByte, err := json.Marshal(&indexParams{Index: &index{Index: PostsIndex, Type: PostsIndex, ID: post.ID}})
		if err != nil {
			return 500, posts.MakeInternal(err)
		}
		postByte, err := json.Marshal(post)
		if err != nil {
			return 500, posts.MakeInternal(err)
		}
		bulkIndexParamsByte = append(bulkIndexParamsByte, postIndexByte...)
		bulkIndexParamsByte = append(bulkIndexParamsByte, []byte("\n")...)
		bulkIndexParamsByte = append(bulkIndexParamsByte, postByte...)
		bulkIndexParamsByte = append(bulkIndexParamsByte, []byte("\n")...)
	}

	response, err := p.esClient.Bulk(bytes.NewReader(bulkIndexParamsByte))
	if err != nil {
		return 500, posts.MakeInternal(errors.Wrap(err, "bulk insert to elasticsearch failed"))
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 500, posts.MakeInternal(errors.Wrap(err, "read response body failed"))
	}
	if response.StatusCode >= 400 {
		return 500, posts.MakeInternal(
			errors.Errorf(
				"bulk insert to elasticsearch failed with status: %d, body: %s",
				response.StatusCode,
				string(body),
			),
		)
	}

	return 201, nil
}

type searchResult struct {
	Hits *hits `json:"hits"`
}

type hits struct {
	Total    int64       `json:"total"`
	MaxScore interface{} `json:"max_score"`
	Hits     []*hit      `json:"hits"`
}

type hit struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source *source `json:"_source"`
}

type source struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Title       string `json:"title"`
}

func (p *postssrvc) APIKeyAuth(ctx context.Context, key string, schema *security.APIKeyScheme) (context.Context, error) {
	if key != os.Getenv("API_SECRET_KEY") {
		return nil, posts.MakeUnauthenticated(errors.New("invalid api key"))
	}
	return ctx, nil
}

// search posts
func (p *postssrvc) Search(ctx context.Context, payload *posts.SearchPayload) (res *posts.SearchResult, err error) {
	response, err := p.esClient.Search(
		p.esClient.Search.WithIndex(PostsIndex),
		p.esClient.Search.WithBody(buildQuery(payload.Query, payload.Page, payload.PageSize)),
	)
	if err != nil {
		return nil, posts.MakeInternal(errors.Wrap(err, "posts.request to elasticposts.failed"))
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, posts.MakeInternal(errors.Wrapf(err, "read error response failed, response: %s", response.String()))
		}
		return nil, posts.MakeInternal(
			errors.Errorf("search request to elasticsearch failed with status %s, body: %s", response.Status(), body),
		)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, posts.MakeInternal(errors.Wrap(err, "read response body failed"))
	}
	sr := &searchResult{}
	if err := json.Unmarshal(body, sr); err != nil {
		return nil, posts.MakeInternal(errors.Wrapf(err, "unmarshal response body json failed, body: %s", body))
	}

	postList := make([]*posts.PostOutput, len(sr.Hits.Hits))
	for i, hit := range sr.Hits.Hits {
		postList[i] = &posts.PostOutput{
			ID:          hit.Source.ID,
			Title:       hit.Source.Title,
			Description: hit.Source.Description,
		}
	}

	return &posts.SearchResult{
		Posts:     postList,
		Page:      payload.Page,
		TotalPage: calcTotalPage(uint(sr.Hits.Total), payload.PageSize),
	}, nil
}

func buildQuery(query string, page, pageSize uint) io.Reader {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`{
	"_source": ["id", "title", "description"],
	"query": {
		"multi_match" : {
			"query": %q,
			"type": "most_fields",
			"fields": ["title^200", "description^20", "body"]
		}
	},
	"from": %d,
	"size": %d,
	"sort": [{ "_score" : "desc" }, { "_doc" : "asc" }]
}`, query, (page-1)*pageSize, pageSize))
	// page starts from 0 in elasticsearch

	return strings.NewReader(b.String())
}

func calcTotalPage(totalItems, pageSize uint) uint {
	return (totalItems + pageSize - 1) / pageSize
}
