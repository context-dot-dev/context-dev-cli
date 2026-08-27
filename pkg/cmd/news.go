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
	Usage:   "Searches live and historical company news for one company, identified in\nsearchBy by name, domain, ticker (optionally disambiguated by exchange), or\nISIN. Results can be filtered by publisher domain, publisher country, article\nlanguage, article type, and published-at date, and include stable story IDs,\nsource metadata, verified entity relevance, and cursor pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "search-by",
			Usage:    "What to search for.",
			Required: true,
			BodyPath: "searchBy",
		},
		&requestflag.Flag[*string]{
			Name:     "cursor",
			Usage:    "Opaque next_cursor from the previous response, or null for the first page.",
			BodyPath: "cursor",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "filter-by",
			Usage:    "Optional result filters.",
			BodyPath: "filterBy",
		},
		&requestflag.Flag[int64]{
			Name:     "limit",
			Usage:    "Maximum results to return. Defaults to 10.",
			Default:  10,
			BodyPath: "limit",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sort-by",
			Usage:    "Result ordering. Defaults to newest.",
			Default:  map[string]any{"type": "newest"},
			BodyPath: "sortBy",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Optional tags for tracking usage. Up to 20 tags, each 1 to 50 characters.",
			BodyPath: "tags",
		},
	},
	Action:          handleNewsSearch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"search-by": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "search-by.entity",
			Usage:      "The company to search news for, identified by name, domain, ticker, or ISIN.",
			InnerField: "entity",
		},
		&requestflag.InnerFlag[string]{
			Name:       "search-by.type",
			Usage:      "How to search. Only entity search is supported.",
			InnerField: "type",
		},
	},
	"filter-by": {
		&requestflag.InnerFlag[[]string]{
			Name:       "filter-by.article-language",
			Usage:      "Article languages to include. Up to 3.",
			InnerField: "articleLanguage",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter-by.article-type",
			Usage:      "Article types to include. Up to 3.",
			InnerField: "articleType",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter-by.date",
			Usage:      "Published-at window in epoch milliseconds.",
			InnerField: "date",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter-by.source-country",
			Usage:      "Publisher countries to include, as lowercase ISO 3166-1 alpha-2 codes. Up to 3.",
			InnerField: "sourceCountry",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter-by.source-domain",
			Usage:      "Publisher domains to include. Up to 3.",
			InnerField: "sourceDomain",
		},
	},
	"sort-by": {
		&requestflag.InnerFlag[string]{
			Name:       "sort-by.type",
			Usage:      "Result ordering.",
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
