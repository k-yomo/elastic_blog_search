// Code generated by goa v3.0.7, DO NOT EDIT.
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

	registerc "github.com/k-yomo/elastic_blog_search/src/gen/http/register/client"
	searchc "github.com/k-yomo/elastic_blog_search/src/gen/http/search/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `register register
search search
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` register register --body '[
      {
         "body": "Assumenda nesciunt nesciunt quasi voluptates perferendis.",
         "description": "Dolores alias incidunt sunt ut veniam.",
         "id": "Nihil quisquam.",
         "title": "Earum dolores qui."
      }
   ]'` + "\n" +
		os.Args[0] + ` search search --query "Quia ipsum omnis repellat nostrum autem." --page 6514126617171776835 --page-size 12386236855430162696` + "\n" +
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
		registerFlags = flag.NewFlagSet("register", flag.ContinueOnError)

		registerRegisterFlags    = flag.NewFlagSet("register", flag.ExitOnError)
		registerRegisterBodyFlag = registerRegisterFlags.String("body", "REQUIRED", "")

		searchFlags = flag.NewFlagSet("search", flag.ContinueOnError)

		searchSearchFlags        = flag.NewFlagSet("search", flag.ExitOnError)
		searchSearchQueryFlag    = searchSearchFlags.String("query", "REQUIRED", "")
		searchSearchPageFlag     = searchSearchFlags.String("page", "", "")
		searchSearchPageSizeFlag = searchSearchFlags.String("page-size", "", "")
	)
	registerFlags.Usage = registerUsage
	registerRegisterFlags.Usage = registerRegisterUsage

	searchFlags.Usage = searchUsage
	searchSearchFlags.Usage = searchSearchUsage

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
		case "register":
			svcf = registerFlags
		case "search":
			svcf = searchFlags
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
		case "register":
			switch epn {
			case "register":
				epf = registerRegisterFlags

			}

		case "search":
			switch epn {
			case "search":
				epf = searchSearchFlags

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
		case "register":
			c := registerc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "register":
				endpoint = c.Register()
				data, err = registerc.BuildRegisterPayload(*registerRegisterBodyFlag)
			}
		case "search":
			c := searchc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "search":
				endpoint = c.Search()
				data, err = searchc.BuildSearchPayload(*searchSearchQueryFlag, *searchSearchPageFlag, *searchSearchPageSizeFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// registerUsage displays the usage of the register command and its subcommands.
func registerUsage() {
	fmt.Fprintf(os.Stderr, `register service registers blog posts to be searched
Usage:
    %s [globalflags] register COMMAND [flags]

COMMAND:
    register: Register implements register.

Additional help:
    %s register COMMAND --help
`, os.Args[0], os.Args[0])
}
func registerRegisterUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] register register -body JSON

Register implements register.
    -body JSON: 

Example:
    `+os.Args[0]+` register register --body '[
      {
         "body": "Assumenda nesciunt nesciunt quasi voluptates perferendis.",
         "description": "Dolores alias incidunt sunt ut veniam.",
         "id": "Nihil quisquam.",
         "title": "Earum dolores qui."
      }
   ]'
`, os.Args[0])
}

// searchUsage displays the usage of the search command and its subcommands.
func searchUsage() {
	fmt.Fprintf(os.Stderr, `search service searches blog posts with given params
Usage:
    %s [globalflags] search COMMAND [flags]

COMMAND:
    search: Search implements search.

Additional help:
    %s search COMMAND --help
`, os.Args[0], os.Args[0])
}
func searchSearchUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] search search -query STRING -page UINT -page-size UINT

Search implements search.
    -query STRING: 
    -page UINT: 
    -page-size UINT: 

Example:
    `+os.Args[0]+` search search --query "Quia ipsum omnis repellat nostrum autem." --page 6514126617171776835 --page-size 12386236855430162696
`, os.Args[0])
}
