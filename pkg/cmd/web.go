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

var webAnswers = requestflag.WithInnerFlags(cli.Command{
	Name:    "answers",
	Usage:   "Research a task across the live web and Context.dev's company and people data,\nand get the answer in the JSON shape you provide, with the URLs it used. `fast`\n(10 credits) suits short lookups; `ultra` (100 credits, the default) reads more\nsources for deeper research.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "task",
			Usage:    "What to research, in plain language, up to 2000 characters. Include domains or URLs to focus the research.",
			Required: true,
			BodyPath: "task",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "controls",
			Usage:    "Limit which sites the research can search, read, and cite.",
			BodyPath: "controls",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "json-format",
			Usage:    "Example of the answer object to return, not a JSON Schema. Up to 8 levels deep, 500 values, and 16,000 characters. Values that cannot be found come back as `null`. Defaults to `{ \"result\": \"\" }`.",
			BodyPath: "json_format",
		},
		&requestflag.Flag[string]{
			Name:     "mode",
			Usage:    "`fast` prioritizes speed and adds extra verification for people and companies; `ultra` searches and reads more sources for deeper research. Defaults to `ultra`.",
			BodyPath: "mode",
		},
		&requestflag.Flag[[]string]{
			Name:     "policy",
			Usage:    "Access policy IDs to apply; each URL must pass all of them. Requires access policies enabled for your organization; otherwise returns `ACCESS_POLICIES_NOT_ENABLED`.",
			BodyPath: "policies",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Labels for this request. Filter usage and request logs by them in the dashboard or with `GET /org/logs`.",
			BodyPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "timeout-opts",
			Usage:    "Request deadline and what to return when it passes.",
			BodyPath: "timeoutOpts",
		},
		&requestflag.Flag[string]{
			Name:     "zdr",
			Usage:    "`enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.",
			BodyPath: "zdr",
		},
	},
	Action:          handleWebAnswers,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"controls": {
		&requestflag.InnerFlag[[]string]{
			Name:       "controls.exclude-sources",
			Usage:      "Never search, read, or cite these sites, sections, or pages. Same format as `sources`.",
			InnerField: "excludeSources",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "controls.sources",
			Usage:      "Only search, read, and cite these sites, sections, or pages. Each entry is a URL prefix: `example.com` covers the site and its subdomains, `example.com/docs` covers that section, and `example.com/docs/setup.html` is one page.",
			InnerField: "sources",
		},
	},
	"timeout-opts": {
		&requestflag.InnerFlag[int64]{
			Name:       "timeout-opts.milliseconds",
			Usage:      "Deadline in milliseconds.",
			InnerField: "milliseconds",
		},
		&requestflag.InnerFlag[string]{
			Name:       "timeout-opts.behavior",
			Usage:      "`fail` returns 408 at the deadline. `return-partial` returns the results collected so far and marks the response as partial.",
			InnerField: "behavior",
		},
	},
})

var webExtractCompetitors = requestflag.WithInnerFlags(cli.Command{
	Name:    "extract-competitors",
	Usage:   "Find a company's direct competitors from its website and web search results.\nEach competitor comes with a confidence level and the search results behind it.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "domain",
			Usage:     "Company domain, such as `stripe.com`. Full URLs are reduced to their domain.",
			Required:  true,
			QueryPath: "domain",
		},
		&requestflag.Flag[int64]{
			Name:      "num-competitors",
			Usage:     "Exact number of direct competitors to return.",
			Default:   5,
			QueryPath: "numCompetitors",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Comma-separated labels for this request, such as `production,team-alpha`. Filter usage and request logs by them.",
			QueryPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "timeout-opts",
			Usage:     "Request deadline and what to return when it passes.",
			QueryPath: "timeoutOpts",
		},
		&requestflag.Flag[string]{
			Name:      "zdr",
			Usage:     "`enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.",
			Default:   "disabled",
			QueryPath: "zdr",
		},
	},
	Action:          handleWebExtractCompetitors,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"timeout-opts": {
		&requestflag.InnerFlag[int64]{
			Name:       "timeout-opts.milliseconds",
			Usage:      "Deadline in milliseconds.",
			InnerField: "milliseconds",
		},
		&requestflag.InnerFlag[string]{
			Name:       "timeout-opts.behavior",
			Usage:      "`fail` returns 408 at the deadline. `return-partial` returns the results collected so far and marks the response as partial.",
			InnerField: "behavior",
		},
	},
})

var webExtractStyleguide = requestflag.WithInnerFlags(cli.Command{
	Name:    "extract-styleguide",
	Usage:   "Extract a website's design system: colors, typography, spacing, shadows, button\nand card styles, and downloadable font files. Every component style includes\nready-to-use CSS.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "color-scheme",
			Usage:     "Emulate a light or dark color scheme. Each scheme is cached separately.",
			QueryPath: "colorScheme",
		},
		&requestflag.Flag[string]{
			Name:      "direct-url",
			Usage:     "Exact page URL to analyze. Provide either `domain` or `directUrl`, not both.",
			QueryPath: "directUrl",
		},
		&requestflag.Flag[string]{
			Name:      "domain",
			Usage:     "Domain to analyze, such as `stripe.com`. Provide either `domain` or `directUrl`, not both.",
			QueryPath: "domain",
		},
		&requestflag.Flag[*int64]{
			Name:      "max-age-ms",
			Usage:     "Maximum age of a cached result, in milliseconds. `0` always fetches fresh. Defaults to 90 days (7776000000 ms). Values above 1 year (31536000000 ms) are capped at 1 year.",
			Default:   requestflag.Ptr[int64](7776000000),
			QueryPath: "maxAgeMs",
		},
		&requestflag.Flag[[]string]{
			Name:      "policy",
			Usage:     "Comma-separated access policy IDs to apply; each URL must pass all of them. Requires access policies enabled for your organization; otherwise returns `ACCESS_POLICIES_NOT_ENABLED`.",
			QueryPath: "policies",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Comma-separated labels for this request, such as `production,team-alpha`. Filter usage and request logs by them.",
			QueryPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "timeout-opts",
			Usage:     "Request deadline and what to return when it passes.",
			QueryPath: "timeoutOpts",
		},
		&requestflag.Flag[string]{
			Name:      "zdr",
			Usage:     "`enabled` turns off request logs and traces and keeps request and response content out of usage logs. Brand data and hosted assets may still be stored. Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.",
			QueryPath: "zdr",
		},
	},
	Action:          handleWebExtractStyleguide,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"timeout-opts": {
		&requestflag.InnerFlag[int64]{
			Name:       "timeout-opts.milliseconds",
			Usage:      "Deadline in milliseconds.",
			InnerField: "milliseconds",
		},
		&requestflag.InnerFlag[string]{
			Name:       "timeout-opts.behavior",
			Usage:      "`fail` returns 408 at the deadline. `return-partial` returns the results collected so far and marks the response as partial. `return-partial` requires at least 5000 ms. On `POST /web/scrape`, an output that is not ready in time fails with `REQUEST_TIMEOUT` inside a 200 response under either behavior; 408 means the whole request could not respond by the deadline.",
			InnerField: "behavior",
		},
	},
})

var webScreenshot = requestflag.WithInnerFlags(cli.Command{
	Name:    "screenshot",
	Usage:   "Capture a screenshot of a website's home page, a page type such as pricing, or\nan exact URL. Returns the visible viewport, a scroll slice, or the full page.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:      "clear-popups",
			Usage:     "Dismiss cookie banners and close other popups and overlays before capture.",
			Default:   false,
			QueryPath: "clearPopups",
		},
		&requestflag.Flag[string]{
			Name:      "color-scheme",
			Usage:     "Emulate a light or dark color scheme.",
			QueryPath: "colorScheme",
		},
		&requestflag.Flag[string]{
			Name:      "country",
			Usage:     "Country to fetch from, as an ISO 3166-1 alpha-2 code.",
			QueryPath: "country",
		},
		&requestflag.Flag[string]{
			Name:      "direct-url",
			Usage:     "Exact URL to capture, such as `https://example.com/pricing`. Provide either `domain` or `directUrl`, not both.",
			QueryPath: "directUrl",
		},
		&requestflag.Flag[string]{
			Name:      "domain",
			Usage:     "Domain to capture, such as `stripe.com`. Provide either `domain` or `directUrl`, not both.",
			QueryPath: "domain",
		},
		&requestflag.Flag[string]{
			Name:      "full-screenshot",
			Usage:     "Capture the full scrollable page instead of only the viewport.",
			QueryPath: "fullScreenshot",
		},
		&requestflag.Flag[bool]{
			Name:      "handle-cookie-popup",
			Usage:     "Dismiss cookie consent banners before capture.",
			Default:   false,
			QueryPath: "handleCookiePopup",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "headers",
			Usage:     "HTTP headers to send to the target origin, such as `headers[Authorization]=Bearer token`. Requests with headers skip the cache and return the image as a data URL.",
			QueryPath: "headers",
		},
		&requestflag.Flag[*int64]{
			Name:      "max-age-ms",
			Usage:     "Maximum age of a cached result, in milliseconds. `0` always fetches fresh. Defaults to 1 day (86400000 ms). Maximum: 30 days (2592000000 ms).",
			Default:   requestflag.Ptr[int64](86400000),
			QueryPath: "maxAgeMs",
		},
		&requestflag.Flag[string]{
			Name:      "page",
			Usage:     "Capture this type of page instead of the home page. Context.dev finds it among the domain's links, in 30 languages. Only with `domain`.",
			QueryPath: "page",
		},
		&requestflag.Flag[*int64]{
			Name:      "scroll-offset",
			Usage:     "Return the viewport-sized slice of the full page that starts this many pixels from the top. Step by the viewport height, such as `0`, `1080`, `2160`, to walk a long page. The last slice can be shorter. Overrides `fullScreenshot`.",
			QueryPath: "scrollOffset",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Comma-separated labels for this request, such as `production,team-alpha`. Filter usage and request logs by them.",
			QueryPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "timeout-opts",
			Usage:     "Request deadline and what to return when it passes.",
			QueryPath: "timeoutOpts",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "viewport",
			Usage:     "Browser viewport size in pixels.",
			Default:   map[string]any{"width": 1920, "height": 1080},
			QueryPath: "viewport",
		},
		&requestflag.Flag[*int64]{
			Name:      "wait-for-ms",
			Usage:     "Time to wait after the page loads before capturing, in milliseconds. `0` waits 500 ms. `timeoutOpts.milliseconds` must be at least `waitForMs` plus 10000 ms, or the request fails with `TIMEOUT_TOO_SHORT_FOR_WAIT`.",
			Default:   requestflag.Ptr[int64](3000),
			QueryPath: "waitForMs",
		},
		&requestflag.Flag[string]{
			Name:      "zdr",
			Usage:     "`enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.",
			Default:   "disabled",
			QueryPath: "zdr",
		},
	},
	Action:          handleWebScreenshot,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"timeout-opts": {
		&requestflag.InnerFlag[int64]{
			Name:       "timeout-opts.milliseconds",
			Usage:      "Deadline in milliseconds.",
			InnerField: "milliseconds",
		},
		&requestflag.InnerFlag[string]{
			Name:       "timeout-opts.behavior",
			Usage:      "`fail` returns 408 at the deadline. `return-partial` returns the results collected so far and marks the response as partial. `return-partial` requires at least 5000 ms. On `POST /web/scrape`, an output that is not ready in time fails with `REQUEST_TIMEOUT` inside a 200 response under either behavior; 408 means the whole request could not respond by the deadline.",
			InnerField: "behavior",
		},
	},
	"viewport": {
		&requestflag.InnerFlag[int64]{
			Name:       "viewport.height",
			Usage:      "Viewport height in pixels.",
			InnerField: "height",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "viewport.width",
			Usage:      "Viewport width in pixels.",
			InnerField: "width",
		},
	},
})

var webSearch = requestflag.WithInnerFlags(cli.Command{
	Name:    "search",
	Usage:   "Search the web and get ranked results, with optional page Markdown or\nquery-relevant passages for each one. Supports search operators and filters for\ndomain, freshness, country, and language.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "Search query. Accepts natural language and search operators such as `site:`, `-site:`, `inurl:`, `intitle:`, quoted phrases, and `OR`.",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			Usage:    "Localize results to this country, as an ISO 3166-1 alpha-2 code such as `us` or `de`.",
			BodyPath: "country",
		},
		&requestflag.Flag[*int64]{
			Name:     "description-max-characters",
			Usage:    "Maximum length of each result's `description`, in characters.",
			Default:  nil,
			BodyPath: "descriptionMaxCharacters",
		},
		&requestflag.Flag[[]string]{
			Name:     "exclude-domain",
			Usage:    "Leave out results from these domains, such as `pinterest.com`.",
			BodyPath: "excludeDomains",
		},
		&requestflag.Flag[string]{
			Name:     "freshness",
			Usage:    "Restrict results to content published within this window.",
			BodyPath: "freshness",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "highlights-options",
			Usage:    "Return passages from each result page that are relevant to the query. Pages are read with the `markdownOptions` settings.",
			BodyPath: "highlightsOptions",
		},
		&requestflag.Flag[[]string]{
			Name:     "include-domain",
			Usage:    "Only return results from these domains, such as `arxiv.org`.",
			BodyPath: "includeDomains",
		},
		&requestflag.Flag[any]{
			Name:     "languages",
			Usage:    "Languages to return results in, as BCP 47 codes such as `en` or `de`, or `all` for any language.",
			Default:  "all",
			BodyPath: "languages",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "markdown-options",
			Usage:    "Scrape each result page to Markdown. Set `enabled` to `true` to turn it on.",
			BodyPath: "markdownOptions",
		},
		&requestflag.Flag[int64]{
			Name:     "num-results",
			Usage:    "Number of results to return.",
			Default:  10,
			BodyPath: "numResults",
		},
		&requestflag.Flag[[]string]{
			Name:     "policy",
			Usage:    "Access policy IDs to apply; each URL must pass all of them. Requires access policies enabled for your organization; otherwise returns `ACCESS_POLICIES_NOT_ENABLED`.",
			BodyPath: "policies",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Labels for this request. Filter usage and request logs by them in the dashboard or with `GET /org/logs`.",
			BodyPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "timeout-opts",
			Usage:    "Request deadline and what to return when it passes.",
			BodyPath: "timeoutOpts",
		},
		&requestflag.Flag[string]{
			Name:     "zdr",
			Usage:    "`enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.",
			Default:  "disabled",
			BodyPath: "zdr",
		},
	},
	Action:          handleWebSearch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"highlights-options": {
		&requestflag.InnerFlag[bool]{
			Name:       "highlights-options.enabled",
			Usage:      "Return relevant passages for each result. Adds 1 credit per 10 results.",
			InnerField: "enabled",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "highlights-options.max-characters",
			Usage:      "Maximum combined length of passages per result, in characters.",
			InnerField: "maxCharacters",
		},
	},
	"markdown-options": {
		&requestflag.InnerFlag[bool]{
			Name:       "markdown-options.enabled",
			Usage:      "Scrape each result page to Markdown. Adds 1 credit per result that returns Markdown.",
			InnerField: "enabled",
		},
		&requestflag.InnerFlag[any]{
			Name:       "markdown-options.include-frames",
			Usage:      "Which iframe contents to include in the Markdown. `auto` inlines only iframes that hold page content, such as same-site documents, framesets, job boards, forms, and embedded documents, and skips ads and widgets. `true` inlines every iframe; `false` inlines none.",
			InnerField: "includeFrames",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "markdown-options.include-images",
			Usage:      "Include image references in the Markdown.",
			InnerField: "includeImages",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "markdown-options.include-links",
			Usage:      "Keep link URLs in the Markdown.",
			InnerField: "includeLinks",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "markdown-options.max-age-ms",
			Usage:      "Maximum age of a cached result, in milliseconds. `0` always fetches fresh. Defaults to 180 days (15552000000 ms) with Markdown, or 365 days (31536000000 ms) with highlights only.",
			InnerField: "maxAgeMs",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "markdown-options.pdf",
			Usage:      "Whether to parse PDFs and which pages to read.",
			InnerField: "pdf",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "markdown-options.shorten-base64-images",
			Usage:      "Shorten inline base64 image data.",
			InnerField: "shortenBase64Images",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "markdown-options.timeout-opts",
			Usage:      "Request deadline and what to return when it passes.",
			InnerField: "timeoutOpts",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "markdown-options.use-main-content-only",
			Usage:      "Keep only the main content, without navigation, headers, footers, or sidebars.",
			InnerField: "useMainContentOnly",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "markdown-options.wait-for-ms",
			Usage:      "Time to wait after each result page loads before reading it, in milliseconds. `0` waits 500 ms. Helps with JavaScript-heavy pages.",
			InnerField: "waitForMs",
		},
	},
	"timeout-opts": {
		&requestflag.InnerFlag[int64]{
			Name:       "timeout-opts.milliseconds",
			Usage:      "Deadline in milliseconds.",
			InnerField: "milliseconds",
		},
		&requestflag.InnerFlag[string]{
			Name:       "timeout-opts.behavior",
			Usage:      "`fail` returns 408 at the deadline. `return-partial` returns the results collected so far and marks the response as partial.",
			InnerField: "behavior",
		},
	},
})

var webWebCrawlMd = requestflag.WithInnerFlags(cli.Command{
	Name:    "web-crawl-md",
	Usage:   "Crawl up to 500 pages from a start URL and get each page back as Markdown with\nmetadata. Control scope with depth limits, regex patterns, and subdomain rules.\nFor larger crawls, use a batch.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "Start URL, including `http://` or `https://`.",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			Usage:    "Country to fetch from, as an ISO 3166-1 alpha-2 code.",
			BodyPath: "country",
		},
		&requestflag.Flag[any]{
			Name:     "exclude-selector",
			Usage:    "Remove elements matching these CSS selectors. Overrides `includeSelectors`.",
			BodyPath: "excludeSelectors",
		},
		&requestflag.Flag[string]{
			Name:     "exclude-url-regex",
			Usage:    "Never follow or scrape URLs matching this regex, even when they match `urlRegex`. The start URL is always scraped.",
			BodyPath: "excludeUrlRegex",
		},
		&requestflag.Flag[bool]{
			Name:     "follow-subdomains",
			Usage:    "Follow links to subdomains of the start URL's domain, such as `docs.example.com` from `example.com`. `www` and the apex domain always count as the same site.",
			Default:  false,
			BodyPath: "followSubdomains",
		},
		&requestflag.Flag[any]{
			Name:     "include-frames",
			Usage:    "Which iframe contents to include in each page's Markdown. `auto` inlines only iframes that hold page content, such as same-site documents, framesets, job boards, forms, and embedded documents, and skips ads and widgets. `true` inlines every iframe; `false` inlines none.",
			Default:  "auto",
			BodyPath: "includeFrames",
		},
		&requestflag.Flag[bool]{
			Name:     "include-images",
			Usage:    "Include image references in the Markdown (`html-conversion` only).",
			Default:  false,
			BodyPath: "includeImages",
		},
		&requestflag.Flag[bool]{
			Name:     "include-links",
			Usage:    "Keep link URLs in the Markdown (`html-conversion` only).",
			Default:  true,
			BodyPath: "includeLinks",
		},
		&requestflag.Flag[any]{
			Name:     "include-selector",
			Usage:    "Keep only elements matching these CSS selectors.",
			BodyPath: "includeSelectors",
		},
		&requestflag.Flag[*int64]{
			Name:     "max-age-ms",
			Usage:    "Maximum age of a cached result, in milliseconds. `0` always fetches fresh. Defaults to 1 day (86400000 ms). Maximum: 30 days (2592000000 ms).",
			Default:  requestflag.Ptr[int64](86400000),
			BodyPath: "maxAgeMs",
		},
		&requestflag.Flag[int64]{
			Name:     "max-depth",
			Usage:    "Maximum link depth from the start URL. `0` crawls only the start URL. Omit for no depth limit.",
			BodyPath: "maxDepth",
		},
		&requestflag.Flag[int64]{
			Name:     "max-pages",
			Usage:    "Maximum number of pages to crawl.",
			Default:  100,
			BodyPath: "maxPages",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "pdf",
			Usage:    "Whether to parse PDFs and which pages to read.",
			Default:  map[string]any{"shouldParse": true, "ocr": false},
			BodyPath: "pdf",
		},
		&requestflag.Flag[[]string]{
			Name:     "policy",
			Usage:    "Access policy IDs to apply; each URL must pass all of them. Requires access policies enabled for your organization; otherwise returns `ACCESS_POLICIES_NOT_ENABLED`.",
			BodyPath: "policies",
		},
		&requestflag.Flag[bool]{
			Name:     "settle-animations",
			Usage:    "Wait briefly for CSS animations and transitions to settle before reading each page.",
			Default:  false,
			BodyPath: "settleAnimations",
		},
		&requestflag.Flag[bool]{
			Name:     "shorten-base64-images",
			Usage:    "Shorten inline base64 image data (`html-conversion` only).",
			Default:  true,
			BodyPath: "shortenBase64Images",
		},
		&requestflag.Flag[[]string]{
			Name:     "source-preference",
			Usage:    "Markdown sources in order of preference: `html-conversion` converts the page HTML, `markdown-direct` uses Markdown the site publishes.",
			Default:  []string{"html-conversion", "markdown-direct"},
			BodyPath: "sourcePreference",
		},
		&requestflag.Flag[int64]{
			Name:     "stop-after-ms",
			Usage:    "Soft time budget for the crawl, in milliseconds. When it runs out, the crawl returns the pages collected so far. Larger values are capped at 110 seconds.",
			Default:  80000,
			BodyPath: "stopAfterMs",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Labels for this request. Filter usage and request logs by them in the dashboard or with `GET /org/logs`.",
			BodyPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "timeout-opts",
			Usage:    "Request deadline and what to return when it passes.",
			BodyPath: "timeoutOpts",
		},
		&requestflag.Flag[string]{
			Name:     "url-regex",
			Usage:    "Only follow and scrape URLs matching this regex. A prefix pattern of the form `^<start URL>` also follows a redirect of the start page.",
			BodyPath: "urlRegex",
		},
		&requestflag.Flag[bool]{
			Name:     "use-main-content-only",
			Usage:    "Keep only the main content, without navigation, headers, footers, or sidebars.",
			Default:  false,
			BodyPath: "useMainContentOnly",
		},
		&requestflag.Flag[*int64]{
			Name:     "wait-for-ms",
			Usage:    "Time to wait after each page loads before reading it, in milliseconds. `0` waits 500 ms. `timeoutOpts.milliseconds` must be at least `waitForMs` plus 10000 ms, or the request fails with `TIMEOUT_TOO_SHORT_FOR_WAIT`.",
			Default:  requestflag.Ptr[int64](3500),
			BodyPath: "waitForMs",
		},
		&requestflag.Flag[string]{
			Name:     "zdr",
			Usage:    "`enabled` turns on zero data retention. Returns 403 `ZDR_NOT_ENABLED` unless your organization has ZDR.",
			Default:  "disabled",
			BodyPath: "zdr",
		},
	},
	Action:          handleWebWebCrawlMd,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"pdf": {
		&requestflag.InnerFlag[int64]{
			Name:       "pdf.end",
			Usage:      "Last page to parse. Defaults to the last page. Must be at least `start`.",
			InnerField: "end",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "pdf.ocr",
			Usage:      "Read scanned PDF pages with OCR. Pages that already have text are kept as is.",
			InnerField: "ocr",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "pdf.should-parse",
			Usage:      "Parse PDF pages. When `false`, PDFs are left out of the results and not counted as failures.",
			InnerField: "shouldParse",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "pdf.start",
			Usage:      "First page to parse, starting at 1.",
			InnerField: "start",
		},
	},
	"timeout-opts": {
		&requestflag.InnerFlag[int64]{
			Name:       "timeout-opts.milliseconds",
			Usage:      "Deadline in milliseconds.",
			InnerField: "milliseconds",
		},
		&requestflag.InnerFlag[string]{
			Name:       "timeout-opts.behavior",
			Usage:      "`fail` returns 408 at the deadline. `return-partial` returns the results collected so far and marks the response as partial.",
			InnerField: "behavior",
		},
	},
})

func handleWebAnswers(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.WebAnswersParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Web.Answers(ctx, params, options...)
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
		Title:          "web answers",
		Transform:      transform,
	})
}

func handleWebExtractCompetitors(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := contextdev.WebExtractCompetitorsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Web.ExtractCompetitors(ctx, params, options...)
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
		Title:          "web extract-competitors",
		Transform:      transform,
	})
}

func handleWebExtractStyleguide(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := contextdev.WebExtractStyleguideParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Web.ExtractStyleguide(ctx, params, options...)
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
		Title:          "web extract-styleguide",
		Transform:      transform,
	})
}

func handleWebScreenshot(ctx context.Context, cmd *cli.Command) error {
	client := contextdev.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := contextdev.WebScreenshotParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Web.Screenshot(ctx, params, options...)
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
		Title:          "web screenshot",
		Transform:      transform,
	})
}

func handleWebSearch(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.WebSearchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Web.Search(ctx, params, options...)
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
		Title:          "web search",
		Transform:      transform,
	})
}

func handleWebWebCrawlMd(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.WebWebCrawlMdParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Web.WebCrawlMd(ctx, params, options...)
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
		Title:          "web web-crawl-md",
		Transform:      transform,
	})
}
