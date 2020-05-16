package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"jaytaylor.com/html2text"
	"log"
	"mvdan.cc/xurls"
	"net/http"
	"net/url"
	"os"
	"time"
)

type RegisterParams struct {
	Posts Posts `json:"posts"`
}

type Post struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	ScreenImageURL string `json:"screenImageUrl"`
	Body           string `json:"body"`
}

type Posts []*Post

const usage = `
Usage:
  indexer -[flags] [sitemap_url] [search_app_url]

Flags:
  -help            help for indexer
  -api_key string  auth token for indexing
`

func main() {
	var (
		help   = flag.Bool("help", false, "help flag")
		apiKey = flag.String("api_key", "", "api key to be used for indexing")
	)

	flag.Parse()
	if *help {
		fmt.Print(usage)
		os.Exit(0)
	}
	if *apiKey == "" {
		fmt.Println("--api_key flag must be set")
		fmt.Print(usage)
		os.Exit(0)
	}
	if len(flag.Args()) != 2 {
		fmt.Println("both of sitemap url and search app url must be set")
		fmt.Print(usage)
		os.Exit(1)
	}
	sitemapURL, err := url.Parse(flag.Arg(0))
	if err != nil {
		log.Fatal(errors.Wrap(err, "parse sitemap url"))
	}
	searchAppURL, err := url.Parse(flag.Arg(1))
	if err != nil {
		log.Fatal(errors.Wrap(err, "parse search app url"))
	}

	urls, err := getURLsFromSitemap(sitemapURL.String())
	if err != nil {
		log.Fatal(errors.Wrap(err, "get urls from sitemap"))
	}
	log.Printf("%d post urls are found in %s", len(urls), sitemapURL.String())

	posts, err := getPosts(urls)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("all post info are fetched")

	if err := indexPosts(posts, searchAppURL, *apiKey); err != nil {
		log.Fatal(err)
	}
	log.Printf("%d posts are registered successfully", len(posts))
}

func getURLsFromSitemap(sitemapURL string) ([]string, error) {
	res, err := http.Get(sitemapURL)
	if err != nil {
		return nil, errors.Wrapf(err, "make GET request to %s", sitemapURL)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.Wrapf(err, "status code: %d, error: %s", res.StatusCode, res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "read sitemap body: %v", err)
	}
	return xurls.Strict.FindAllString(string(body), -1)[1:], nil
}

func indexPosts(posts Posts, searchAppURL *url.URL, apiKey string) error {
	postsJson, err := json.Marshal(&RegisterParams{Posts: posts})
	if err != nil {
		log.Fatal(errors.Wrap(err, "convert posts to json"))
	}
	req, err := http.NewRequest("POST", searchAppURL.String(), bytes.NewReader(postsJson))
	if err != nil {
		return errors.Wrapf(err, "init post request to %s", searchAppURL)
	}
	req.Header = map[string][]string{"Authorization": {apiKey}}

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return errors.Wrapf(err, "make post request to %s", searchAppURL)
	}
	if res.StatusCode >= 300 {
		defer res.Body.Close()
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return errors.Wrap(err, "parse response body")
		}
		return errors.Errorf("register posts failed: %s", string(resBody))
	}
	return nil
}

func getPosts(urls []string) (Posts, error) {
	var posts Posts
	eg := errgroup.Group{}
	postChan := make(chan *Post, len(urls))

	for _, u := range urls {
		u := u
		eg.Go(func() error {
			time.Sleep(1 * time.Second)
			post, err := getPost(u)
			if err != nil {
				return errors.Wrapf(err, "get post info from %s", u)
			}
			postChan <- post
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	close(postChan)

	for {
		post, ok := <-postChan
		if !ok {
			break
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(url string) (*Post, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "make GET request to %s", url)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, errors.Wrapf(err, "make GET request to %s, got status: %d", url, res.StatusCode)
		}
		return nil, errors.Errorf("make GET request to %s, got status: %d, body: %s", url, res.StatusCode, string(body))
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "init goquery document")
	}
	title := doc.Find("title").Text()
	var description, screenImageURL string
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); name == "description" {
			description, _ = s.Attr("content")
		}
		if name, _ := s.Attr("property"); name == "og:image" {
			screenImageURL, _ = s.Attr("content")
		}
	})
	body := doc.Find("main").Text()
	bodyStr, err := html2text.FromString(body)
	if err != nil {
		return nil, errors.Wrap(err, "convert html to string")
	}
	return &Post{ID: url, Title: title, Description: description, ScreenImageURL: screenImageURL, Body: bodyStr}, nil
}
