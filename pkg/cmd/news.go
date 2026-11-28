// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/context-dot-dev/context-dev-cli/internal/apiquery"
	"github.com/context-dot-dev/context-dev-cli/internal/requestflag"
	"github.com/context-dot-dev/context-go-sdk/v2"
	"github.com/context-dot-dev/context-go-sdk/v2/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var newsSearch = requestflag.WithInnerFlags(cli.Command{
	Name:    "search",
	Usage:   "Find news about a company by name, domain, ticker, or ISIN, or search all news\nwith a natural-language query. Each article includes its publisher, type of\ncoverage, and, for company searches, whether the company is the main subject.\nNarrow results by publisher, country, language, article type, or publish date.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "search-by",
			Usage:    "What to search for: news about a specific company, or a natural-language query across all news.",
			Required: true,
			BodyPath: "searchBy",
		},
		&requestflag.Flag[*string]{
			Name:     "cursor",
			Usage:    "Cursor from the previous page's `next_cursor`. Omit for the first page.",
			BodyPath: "cursor",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "filter-by",
			Usage:    "Filters for the results. Use at most one of `sourceDomain`, `sourceCountry`, `articleLanguage`, or `articleType`, optionally combined with `date`.",
			BodyPath: "filterBy",
		},
		&requestflag.Flag[int64]{
			Name:     "limit",
			Usage:    "Maximum number of articles to return per page.",
			Default:  10,
			BodyPath: "limit",
		},
		&requestflag.Flag[[]string]{
			Name:     "policy",
			Usage:    "Access policy IDs to apply; each URL must pass all of them. Requires access policies enabled for your organization; otherwise returns `ACCESS_POLICIES_NOT_ENABLED`.",
			BodyPath: "policies",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sort-by",
			Usage:    "How to order results. Company searches default to `newest`; natural-language searches always use `relevance`.",
			BodyPath: "sortBy",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Labels for this request. Filter usage and request logs by them in the dashboard or with `GET /org/logs`.",
			BodyPath: "tags",
		},
	},
	Action:          handleNewsSearch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter-by": {
		&requestflag.InnerFlag[[]string]{
			Name:       "filter-by.article-language",
			Usage:      "Only return articles in these languages, as ISO 639 codes such as `en`.",
			InnerField: "articleLanguage",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter-by.article-type",
			Usage:      "Only return these types of coverage, such as `press_release`.",
			InnerField: "articleType",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter-by.date",
			Usage:      "Publish date range, in epoch milliseconds. `from` must not be later than `to`.",
			InnerField: "date",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter-by.source-country",
			Usage:      "Only return articles from publishers in these countries, as ISO 3166-1 alpha-2 codes such as `us`.",
			InnerField: "sourceCountry",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter-by.source-domain",
			Usage:      "Only return articles from these publisher domains, such as `reuters.com`.",
			InnerField: "sourceDomain",
		},
	},
	"sort-by": {
		&requestflag.InnerFlag[string]{
			Name:       "sort-by.type",
			Usage:      "`newest` puts the most recently published articles first; `relevance` puts the closest matches first.",
			InnerField: "type",
		},
	},
})

func handleNewsSearch(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := contextdev.NewsSearchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.News.Search(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "news search",
		Transform:      transform,
	})
}
