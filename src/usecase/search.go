package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/k-yomo/elastic_blog_search/src/gen/search"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

// search service implementation.
type searchrsrvc struct {
	logger   *log.Logger
	esClient *elasticsearch.Client
}

// NewSearch returns the search service implementation.
func NewSearch(logger *log.Logger, esClient *elasticsearch.Client) search.Service {
	return &searchrsrvc{logger, esClient}
}

const (
	defaultPage     = 1
	defaultPageSize = 50
)

type searchResult struct {
	Hits hits `json:"hits"`
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

func (s *searchrsrvc) Search(ctx context.Context, p *search.SearchPayload) (res *search.SearchResult, err error) {
	var page, pageSize uint
	if p.Page != nil {
		page = *p.Page
	} else {
		page = defaultPage
	}
	if p.PageSize != nil {
		pageSize = *p.PageSize
	} else {
		pageSize = defaultPageSize
	}

	response, err := s.esClient.Search(
		s.esClient.Search.WithIndex(PostsIndex),
		s.esClient.Search.WithBody(s.buildQuery(p.Query, page, pageSize)),
	)
	if err != nil {
		return nil, errors.Wrap(err, "search request to elasticsearch failed")
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, errors.Wrapf(err, "read error response failed")
		}
		return nil, errors.Errorf("search request to elasticsearch failed with status %s, body: %s", response.Status(), body)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body failed")
	}
	sr := &searchResult{}
	if err := json.Unmarshal(body, sr); err != nil {
		return nil, errors.Wrap(err, "unmarshal response body json failed")
	}

	posts := make([]*search.Post, len(sr.Hits.Hits))
	for i, hit := range sr.Hits.Hits {
		posts[i] = &search.Post{
			ID:          hit.Source.ID,
			Title:       hit.Source.Title,
			Description: hit.Source.Description,
		}
	}

	return &search.SearchResult{
		Posts:     posts,
		Page:      page,
		TotalPage: calcTotalPage(uint(sr.Hits.Total), pageSize),
	}, nil
}

func (s *searchrsrvc) buildQuery(query string, page, pageSize uint) io.Reader {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`{
	"_source": ["id", "title", "description"],
	"query" : {
		"multi_match" : {
			"query" : %q,
			"type": "most_fields",
			"fields" : ["title^100", "description^20", "body"]
		}
	},
	"from" : %d,
	"size" : %d,
	"sort" : [ { "_score" : "desc" }, { "_doc" : "asc" } ]
}`, query, (page-1)*pageSize, pageSize))
	// page starts from 0 in elasticsearch

	return strings.NewReader(b.String())
}

func calcTotalPage(totalItems, pageSize uint) uint {
	totalPage := totalItems / pageSize
	if totalItems%pageSize != 0 {
		totalPage++
	}
	return totalPage
}
