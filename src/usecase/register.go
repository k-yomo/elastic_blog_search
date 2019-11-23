package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/pkg/errors"
	"goa.design/goa/v3/security"
	"io/ioutil"
	"log"
	"os"

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

func (s *registersrvc) APIKeyAuth(ctx context.Context, key string, schema *security.APIKeyScheme) (context.Context, error) {
	if key != os.Getenv("API_SECRET_KEY") {
		return nil, register.MakeUnauthenticated(errors.New("invalid api key"))
	}
	return ctx, nil
}

// Register implements register.
func (s *registersrvc) Register(ctx context.Context, payload *register.RegisterPayload) (res int, err error) {
	var bulkIndexParamsByte []byte
	for _, p := range payload.Posts {
		post := mapFromPostPayload(p)
		postIndexByte, err := json.Marshal(&IndexParams{Index: &Index{Index: PostsIndex, Type: PostsIndex, ID: post.ID}})
		if err != nil {
			return 500, register.MakeInternal(err)
		}
		postByte, err := json.Marshal(post)
		if err != nil {
			return 500, register.MakeInternal(err)
		}
		bulkIndexParamsByte = append(bulkIndexParamsByte, postIndexByte...)
		bulkIndexParamsByte = append(bulkIndexParamsByte, []byte("\n")...)
		bulkIndexParamsByte = append(bulkIndexParamsByte, postByte...)
		bulkIndexParamsByte = append(bulkIndexParamsByte, []byte("\n")...)
	}

	response, err := s.esClient.Bulk(bytes.NewReader(bulkIndexParamsByte))
	if err != nil {
		return 500, register.MakeInternal(errors.Wrap(err, "bulk insert to elasticsearch failed"))
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 500, register.MakeInternal(errors.Wrap(err, "read response body failed"))
	}
	if response.StatusCode >= 400 {
		return 500, register.MakeInternal(
			errors.Errorf(
				"bulk insert to elasticsearch failed with status: %d, body: %s",
				response.StatusCode,
				string(body),
			),
		)
	}

	return 201, nil
}

func mapFromPostPayload(p *register.Post) *Post {
	return &Post{
		ID:          *p.ID,
		Title:       *p.Title,
		Description: *p.Description,
		Body:        *p.Body,
	}
}
