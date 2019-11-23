package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v7"
	"io/ioutil"
	"log"

	register "github.com/k-yomo/elastic_blog_search/src/gen/register"
)

const PostsIndex = "posts"

// register service implementation.
type registersrvc struct {
	logger   *log.Logger
	esClient *elasticsearch.Client
}

// NewRegister returns the register service implementation.
func NewRegister(logger *log.Logger, esClient *elasticsearch.Client) register.Service {
	return &registersrvc{logger, esClient}
}

type IndexParams struct {
	Index *Index `json:"index"`
}

type Index struct {
	Index string `json:"_index"`
	Type  string `json:"_type"`
	ID    string `json:"_id"`
}

type Post struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

// Register implements register.
func (s *registersrvc) Register(ctx context.Context, postPayloads []*register.Post) (res int, err error) {
	var bulkIndexParamsByte []byte
	for _, payload := range postPayloads {
		post := mapFromPayload(payload)
		postIndexByte, err := json.Marshal(&IndexParams{Index: &Index{Index: PostsIndex, Type: "items", ID: post.ID}})
		if err != nil {
			return 500, err
		}
		postByte, err := json.Marshal(post)
		if err != nil {
			return 500, err
		}
		bulkIndexParamsByte = append(bulkIndexParamsByte, postIndexByte...)
		bulkIndexParamsByte = append(bulkIndexParamsByte, []byte("\n")...)
		bulkIndexParamsByte = append(bulkIndexParamsByte, postByte...)
		bulkIndexParamsByte = append(bulkIndexParamsByte, []byte("\n")...)
	}

	response, err := s.esClient.Bulk(bytes.NewReader(bulkIndexParamsByte))
	if err != nil {
		return 500, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 500, err
	}
	if response.StatusCode >= 400 {
		return 500, errors.New(string(body))
	}

	return 201, nil
}

func mapFromPayload(p *register.Post) *Post {
	return &Post{
		ID:          *p.ID,
		Title:       *p.Title,
		Description: *p.Description,
		Body:        *p.Body,
	}
}
