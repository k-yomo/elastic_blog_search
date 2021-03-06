// Code generated by goa v3.1.2, DO NOT EDIT.
//
// server HTTP client CLI support package
//
// Command:
// $ goa gen github.com/k-yomo/elastic_blog_search/src/design -o src/

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	postsc "github.com/k-yomo/elastic_blog_search/src/gen/http/posts/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `posts (register|search|related-posts)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` posts register --body '{
      "posts": [
         {
            "body": "Harum molestiae nemo quo.",
            "description": "Tenetur autem libero sint voluptatem.",
            "id": "Est ipsa laboriosam assumenda veritatis sapiente ullam.",
            "screenImageUrl": "Possimus illo voluptatibus corrupti.",
            "title": "Debitis asperiores quasi."
         },
         {
            "body": "Harum molestiae nemo quo.",
            "description": "Tenetur autem libero sint voluptatem.",
            "id": "Est ipsa laboriosam assumenda veritatis sapiente ullam.",
            "screenImageUrl": "Possimus illo voluptatibus corrupti.",
            "title": "Debitis asperiores quasi."
         }
      ]
   }' --key "Corrupti quasi ut consequatur sit quis repellendus."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		postsFlags = flag.NewFlagSet("posts", flag.ContinueOnError)

		postsRegisterFlags    = flag.NewFlagSet("register", flag.ExitOnError)
		postsRegisterBodyFlag = postsRegisterFlags.String("body", "REQUIRED", "")
		postsRegisterKeyFlag  = postsRegisterFlags.String("key", "REQUIRED", "")

		postsSearchFlags        = flag.NewFlagSet("search", flag.ExitOnError)
		postsSearchQueryFlag    = postsSearchFlags.String("query", "REQUIRED", "")
		postsSearchPageFlag     = postsSearchFlags.String("page", "1", "")
		postsSearchPageSizeFlag = postsSearchFlags.String("page-size", "50", "")

		postsRelatedPostsFlags     = flag.NewFlagSet("related-posts", flag.ExitOnError)
		postsRelatedPostsURLFlag   = postsRelatedPostsFlags.String("url", "REQUIRED", "")
		postsRelatedPostsCountFlag = postsRelatedPostsFlags.String("count", "5", "")
	)
	postsFlags.Usage = postsUsage
	postsRegisterFlags.Usage = postsRegisterUsage
	postsSearchFlags.Usage = postsSearchUsage
	postsRelatedPostsFlags.Usage = postsRelatedPostsUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "posts":
			svcf = postsFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "posts":
			switch epn {
			case "register":
				epf = postsRegisterFlags

			case "search":
				epf = postsSearchFlags

			case "related-posts":
				epf = postsRelatedPostsFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "posts":
			c := postsc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "register":
				endpoint = c.Register()
				data, err = postsc.BuildRegisterPayload(*postsRegisterBodyFlag, *postsRegisterKeyFlag)
			case "search":
				endpoint = c.Search()
				data, err = postsc.BuildSearchPayload(*postsSearchQueryFlag, *postsSearchPageFlag, *postsSearchPageSizeFlag)
			case "related-posts":
				endpoint = c.RelatedPosts()
				data, err = postsc.BuildRelatedPostsPayload(*postsRelatedPostsURLFlag, *postsRelatedPostsCountFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// postsUsage displays the usage of the posts command and its subcommands.
func postsUsage() {
	fmt.Fprintf(os.Stderr, `Service is the posts service interface.
Usage:
    %s [globalflags] posts COMMAND [flags]

COMMAND:
    register: registers blog posts to be searched
    search: search blog posts
    related-posts: get related blog posts

Additional help:
    %s posts COMMAND --help
`, os.Args[0], os.Args[0])
}
func postsRegisterUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] posts register -body JSON -key STRING

registers blog posts to be searched
    -body JSON: 
    -key STRING: 

Example:
    `+os.Args[0]+` posts register --body '{
      "posts": [
         {
            "body": "Harum molestiae nemo quo.",
            "description": "Tenetur autem libero sint voluptatem.",
            "id": "Est ipsa laboriosam assumenda veritatis sapiente ullam.",
            "screenImageUrl": "Possimus illo voluptatibus corrupti.",
            "title": "Debitis asperiores quasi."
         },
         {
            "body": "Harum molestiae nemo quo.",
            "description": "Tenetur autem libero sint voluptatem.",
            "id": "Est ipsa laboriosam assumenda veritatis sapiente ullam.",
            "screenImageUrl": "Possimus illo voluptatibus corrupti.",
            "title": "Debitis asperiores quasi."
         }
      ]
   }' --key "Corrupti quasi ut consequatur sit quis repellendus."
`, os.Args[0])
}

func postsSearchUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] posts search -query STRING -page UINT -page-size UINT

search blog posts
    -query STRING: 
    -page UINT: 
    -page-size UINT: 

Example:
    `+os.Args[0]+` posts search --query "A blanditiis sed et animi est." --page 17474747091011644362 --page-size 1374544153326036729
`, os.Args[0])
}

func postsRelatedPostsUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] posts related-posts -url STRING -count UINT

get related blog posts
    -url STRING: 
    -count UINT: 

Example:
    `+os.Args[0]+` posts related-posts --url "Esse corrupti molestiae sequi a hic delectus." --count 2959245027571265226
`, os.Args[0])
}
