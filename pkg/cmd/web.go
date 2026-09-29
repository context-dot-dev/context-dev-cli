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
	Usage:   "Research the web and return a sourced answer in your JSON shape. Choose `fast`\nfor a short task or `ultra` for deeper research.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "task",
			Usage:    "Research task. The agent selects company/profile lookups, web searches, or page reads. Include domains or URLs to focus the research.",
			Required: true,
			BodyPath: "task",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "json-format",
			Usage:    "Example answer object, not JSON Schema. Up to 8 levels, 500 values, and 16000 characters; unknowns may be null.",
			BodyPath: "json_format",
		},
		&requestflag.Flag[string]{
			Name:     "mode",
			Usage:    "`fast` prioritizes speed, with extra verification for people and companies; `ultra` supports deeper research (default).",
			BodyPath: "mode",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Labels for filtering usage in the dashboard.",
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
	"timeout-opts": {
		&requestflag.InnerFlag[int64]{
			Name:       "timeout-opts.milliseconds",
			Usage:      "Deadline in milliseconds.",
			InnerField: "milliseconds",
		},
		&requestflag.InnerFlag[string]{
			Name:       "timeout-opts.behavior",
			Usage:      `"fail" returns 408 at the deadline. "return-partial" returns available results; inspect the response’s partial flag.`,
			InnerField: "behavior",
		},
	},
})

var webExtractCompetitors = requestflag.WithInnerFlags(cli.Command{
	Name:    "extract-competitors",
	Usage:   "Analyze a company's landing page and web search evidence to return direct\ncompetitors for the same product or market.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "domain",
			Usage:     "Company domain to analyze, such as `stripe.com`. Full http(s) URLs are accepted and normalized to their domain.",
			Required:  true,
			QueryPath: "domain",
		},
		&requestflag.Flag[int64]{
			Name:      "num-competitors",
			Usage:     "Exact number of direct competitors to return. Defaults to 5.",
			Default:   5,
			QueryPath: "numCompetitors",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Comma-separated labels for filtering usage, e.g. `production,team-alpha`.",
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
			Usage:      `"fail" returns 408 at the deadline. "return-partial" returns available results; inspect the response’s partial flag.`,
			InnerField: "behavior",
		},
	},
})

var webExtractStyleguide = requestflag.WithInnerFlags(cli.Command{
	Name:    "extract-styleguide",
	Usage:   "Extract colors, typography, spacing, and component styles from a website.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "color-scheme",
			Usage:     "Optional browser color scheme to emulate for websites that respond to prefers-color-scheme. This value is part of the styleguide cache key.",
			QueryPath: "colorScheme",
		},
		&requestflag.Flag[string]{
			Name:      "direct-url",
			Usage:     "Exact URL to inspect. Provide either `domain` or `directUrl`, not both.",
			QueryPath: "directUrl",
		},
		&requestflag.Flag[string]{
			Name:      "domain",
			Usage:     "Domain name to extract styleguide from (e.g., 'example.com', 'google.com'). The domain will be automatically normalized and validated. You must provide either 'domain' or 'directUrl', but not both.",
			QueryPath: "domain",
		},
		&requestflag.Flag[*int64]{
			Name:      "max-age-ms",
			Usage:     "Maximum age of cached brand data in ms. Defaults to 3 months; clamped to 0–1 year. `0` refreshes.",
			Default:   requestflag.Ptr[int64](7776000000),
			QueryPath: "maxAgeMs",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Comma-separated labels for filtering usage, e.g. `production,team-alpha`.",
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
			Usage:      `"fail" returns 408 at the deadline. "return-partial" returns available results; inspect the response’s partial flag. "return-partial" requires at least 5000 ms.`,
			InnerField: "behavior",
		},
	},
})

var webScreenshot = requestflag.WithInnerFlags(cli.Command{
	Name:    "screenshot",
	Usage:   "Capture a screenshot of a website.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:      "clear-popups",
			Usage:     "Optional parameter for comprehensive popup cleanup. If 'true', the browser dismisses detected cookie/consent UI and clears other detected obstructive popups and overlays before capture. If 'false' or not provided, this parameter requests no cleanup; handleCookiePopup can still request cookie/consent handling independently.",
			Default:   false,
			QueryPath: "clearPopups",
		},
		&requestflag.Flag[string]{
			Name:      "color-scheme",
			Usage:     "Optional parameter to choose the site's visual theme in the screenshot. Use 'light' or 'dark' when the site offers both appearances.",
			QueryPath: "colorScheme",
		},
		&requestflag.Flag[string]{
			Name:      "country",
			Usage:     "Fetch from this country (ISO 3166-1 alpha-2).",
			QueryPath: "country",
		},
		&requestflag.Flag[string]{
			Name:      "direct-url",
			Usage:     "A specific URL to screenshot directly, bypassing domain resolution (e.g., 'https://example.com/pricing'). When provided, the screenshot is taken of this exact URL. You must provide either 'domain' or 'directUrl', but not both.",
			QueryPath: "directUrl",
		},
		&requestflag.Flag[string]{
			Name:      "domain",
			Usage:     "Domain name to take screenshot of (e.g., 'example.com', 'google.com'). The domain will be automatically normalized and validated. You must provide either 'domain' or 'directUrl', but not both.",
			QueryPath: "domain",
		},
		&requestflag.Flag[string]{
			Name:      "full-screenshot",
			Usage:     "Optional parameter to determine screenshot type. If 'true', takes a full page screenshot capturing all content. If 'false' or not provided, takes a viewport screenshot (standard browser view).",
			QueryPath: "fullScreenshot",
		},
		&requestflag.Flag[bool]{
			Name:      "handle-cookie-popup",
			Usage:     "Optional parameter to control cookie/consent popup handling. If 'true', we dismiss cookie banner before capture. If 'false' or not provided, captures the page without that step.",
			Default:   false,
			QueryPath: "handleCookiePopup",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "headers",
			Usage:     "Optional outbound HTTP headers, using the same JSON object or deep-object query format as other scrape endpoints (for example headers[Authorization]=Bearer token). Headers are scoped to the target origin during capture. For domain/page requests, discovery receives no custom headers and only pages on the resolved origin are eligible. Non-empty headers bypass screenshot caching and return an in-memory data URL; no screenshot is uploaded. Empty objects behave like omitted headers.",
			QueryPath: "headers",
		},
		&requestflag.Flag[*int64]{
			Name:      "max-age-ms",
			Usage:     "Return a cached screenshot if a prior screenshot for the same parameters exists and is younger than this many milliseconds. Defaults to 1 day (86400000 ms) when omitted. Max is 30 days (2592000000 ms). Set to 0 to always capture fresh.",
			Default:   requestflag.Ptr[int64](86400000),
			QueryPath: "maxAgeMs",
		},
		&requestflag.Flag[string]{
			Name:      "page",
			Usage:     "Optional parameter to specify which page type to screenshot. If provided, the system will scrape the domain's links and use heuristics to find the most appropriate URL for the specified page type (30 supported languages). If not provided, screenshots the main domain landing page. Only applicable when using 'domain', not 'directUrl'.",
			QueryPath: "page",
		},
		&requestflag.Flag[*int64]{
			Name:      "scroll-offset",
			Usage:     "Optional vertical scroll offset in pixels for capturing a long page in viewport-sized chunks. When provided, the full page is captured once and the returned image is the viewport-sized slice that begins at this Y offset (e.g. request scrollOffset=0, then 1080, then 2160 to walk a 1920x1080 landing page top to bottom). The final slice may be shorter than the viewport height. Takes precedence over fullScreenshot. Max: 100000.",
			QueryPath: "scrollOffset",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Comma-separated labels for filtering usage, e.g. `production,team-alpha`.",
			QueryPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "timeout-opts",
			Usage:     "Request deadline and what to return when it passes.",
			QueryPath: "timeoutOpts",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "viewport",
			Usage:     "Optional browser viewport dimensions for the screenshot. Defaults to 1920x1080.",
			Default:   map[string]any{"width": 1920, "height": 1080},
			QueryPath: "viewport",
		},
		&requestflag.Flag[*int64]{
			Name:      "wait-for-ms",
			Usage:     "Optional browser wait time in milliseconds after initial page load before taking the screenshot. Min: 0. Max: 30000 (30 seconds). Defaults to 3000 ms when omitted. When combined with timeoutOpts, timeoutOpts.milliseconds must be at least waitForMs + 10000 ms; a shorter deadline is rejected with 400 TIMEOUT_TOO_SHORT_FOR_WAIT.",
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
			Usage:      `"fail" returns 408 at the deadline. "return-partial" returns available results; inspect the response’s partial flag. "return-partial" requires at least 5000 ms.`,
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
	Usage:   "Search the web and optionally return page content or relevant passages with each\nresult.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "Search query. Accepts natural language as well as Google-style search operators such as `site:`, `-site:`, `inurl:`, `intitle:`, quoted phrases, and `OR`.",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			Usage:    "Two-letter ISO 3166-1 alpha-2 country code to localize results to a specific country (maps to Google's `gl` parameter). Example: \"us\", \"gb\", \"de\".",
			BodyPath: "country",
		},
		&requestflag.Flag[[]string]{
			Name:     "exclude-domain",
			Usage:    `Blocklist — drop results from these domains. Up to 100 domains. Example: ["pinterest.com", "reddit.com"].`,
			BodyPath: "excludeDomains",
		},
		&requestflag.Flag[string]{
			Name:     "freshness",
			Usage:    "Restrict results to content published within this window.",
			BodyPath: "freshness",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "highlights-options",
			Usage:    "Passages from each result page that are relevant to the query. Pages are read with the `markdownOptions` settings.",
			BodyPath: "highlightsOptions",
		},
		&requestflag.Flag[[]string]{
			Name:     "include-domain",
			Usage:    `Allowlist — only return results from these domains. Up to 100 domains. Example: ["arxiv.org", "github.com"].`,
			BodyPath: "includeDomains",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "markdown-options",
			Usage:    "Inline Markdown scraping for each result. Set `enabled: true` to activate.",
			BodyPath: "markdownOptions",
		},
		&requestflag.Flag[int64]{
			Name:     "num-results",
			Usage:    "Number of results to request and return (10–100). Defaults to 10.",
			Default:  10,
			BodyPath: "numResults",
		},
		&requestflag.Flag[bool]{
			Name:     "query-fanout",
			Usage:    "Currently has no effect.",
			BodyPath: "queryFanout",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Labels for filtering usage in the dashboard.",
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
			Usage:      "Maximum combined length of passages per result.",
			InnerField: "maxCharacters",
		},
	},
	"markdown-options": {
		&requestflag.InnerFlag[bool]{
			Name:       "markdown-options.enabled",
			Usage:      "Scrape each result to Markdown. Adds 1 credit per 10 results.",
			InnerField: "enabled",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "markdown-options.include-frames",
			Usage:      "Render iframe contents into the Markdown.",
			InnerField: "includeFrames",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "markdown-options.include-images",
			Usage:      "Emit image references in the Markdown.",
			InnerField: "includeImages",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "markdown-options.include-links",
			Usage:      "Keep hyperlinks in the Markdown.",
			InnerField: "includeLinks",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "markdown-options.max-age-ms",
			Usage:      "Cache TTL in ms for scraped Markdown keyed by URL + options. Default 15 days, max 30 days. Set to 0 to force a fresh scrape.",
			InnerField: "maxAgeMs",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "markdown-options.pdf",
			Usage:      "PDF handling. Use start/end to bound text extraction and OCR to a page range.",
			InnerField: "pdf",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "markdown-options.shorten-base64-images",
			Usage:      "Truncate inline base64 image payloads to keep responses small.",
			InnerField: "shortenBase64Images",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "markdown-options.timeout-opts",
			Usage:      "Request deadline and what to return when it passes.",
			InnerField: "timeoutOpts",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "markdown-options.use-main-content-only",
			Usage:      "Strip nav, header, footer, and sidebar — keep only the primary article content.",
			InnerField: "useMainContentOnly",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "markdown-options.wait-for-ms",
			Usage:      "Extra wait after page load before rendering, in ms (0–30000). Useful for JS-heavy pages.",
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
			Usage:      `"fail" returns 408 at the deadline. "return-partial" returns available results; inspect the response’s partial flag.`,
			InnerField: "behavior",
		},
	},
})

var webWebCrawlMd = requestflag.WithInnerFlags(cli.Command{
	Name:    "web-crawl-md",
	Usage:   "Crawl a website and return page content as Markdown. Use a batch for crawls\nbeyond 500 pages.",
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
			Usage:    "Fetch from this country (ISO 3166-1 alpha-2).",
			BodyPath: "country",
		},
		&requestflag.Flag[[]string]{
			Name:     "exclude-selector",
			Usage:    "Remove matching elements after inclusions. Exclusions take precedence.",
			BodyPath: "excludeSelectors",
		},
		&requestflag.Flag[bool]{
			Name:     "follow-subdomains",
			Usage:    "When true, follow links on subdomains of the starting URL's domain (e.g. docs.example.com when starting from example.com). www and apex are always treated as equivalent.",
			Default:  false,
			BodyPath: "followSubdomains",
		},
		&requestflag.Flag[bool]{
			Name:     "include-frames",
			Usage:    "When true, the contents of iframes are rendered to Markdown for each crawled page.",
			Default:  false,
			BodyPath: "includeFrames",
		},
		&requestflag.Flag[bool]{
			Name:     "include-images",
			Usage:    "Include image references in the Markdown output",
			Default:  false,
			BodyPath: "includeImages",
		},
		&requestflag.Flag[bool]{
			Name:     "include-links",
			Usage:    "Preserve hyperlinks in the Markdown output",
			Default:  true,
			BodyPath: "includeLinks",
		},
		&requestflag.Flag[[]string]{
			Name:     "include-selector",
			Usage:    "Keep matching HTML subtrees before converting each page to Markdown.",
			BodyPath: "includeSelectors",
		},
		&requestflag.Flag[int64]{
			Name:     "max-age-ms",
			Usage:    "Maximum cache age in milliseconds. Defaults to 1 day; `0` fetches fresh.",
			Default:  86400000,
			BodyPath: "maxAgeMs",
		},
		&requestflag.Flag[int64]{
			Name:     "max-depth",
			Usage:    "Maximum link depth from the starting URL (0 = only the starting page)",
			BodyPath: "maxDepth",
		},
		&requestflag.Flag[int64]{
			Name:     "max-pages",
			Usage:    "Maximum pages to crawl.",
			Default:  100,
			BodyPath: "maxPages",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "pdf",
			Usage:    "PDF handling. `start`/`end` limit parsing to an inclusive, 1-based page range.",
			Default:  map[string]any{"shouldParse": true, "ocr": false},
			BodyPath: "pdf",
		},
		&requestflag.Flag[bool]{
			Name:     "settle-animations",
			Usage:    "Wait briefly for CSS animations and transitions to settle before reading each page.",
			Default:  false,
			BodyPath: "settleAnimations",
		},
		&requestflag.Flag[bool]{
			Name:     "shorten-base64-images",
			Usage:    "Truncate base64-encoded image data in the Markdown output",
			Default:  true,
			BodyPath: "shortenBase64Images",
		},
		&requestflag.Flag[int64]{
			Name:     "stop-after-ms",
			Usage:    "Soft crawl deadline in milliseconds. Returns pages collected before the next deadline check.",
			Default:  80000,
			BodyPath: "stopAfterMs",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Labels for filtering usage in the dashboard.",
			BodyPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "timeout-opts",
			Usage:    "Request deadline and what to return when it passes.",
			BodyPath: "timeoutOpts",
		},
		&requestflag.Flag[string]{
			Name:     "url-regex",
			Usage:    "Regex pattern. Only URLs matching this pattern will be followed and scraped. An automatic prefix scope in the form ^<starting URL> follows a redirect of the starting page.",
			BodyPath: "urlRegex",
		},
		&requestflag.Flag[bool]{
			Name:     "use-main-content-only",
			Usage:    "Extract only the main content, stripping headers, footers, sidebars, and navigation",
			Default:  false,
			BodyPath: "useMainContentOnly",
		},
		&requestflag.Flag[int64]{
			Name:     "wait-for-ms",
			Usage:    "Browser wait time in milliseconds after initial page load for each crawled page. Defaults to 3500 (3.5 seconds). Min: 0. Max: 30000 (30 seconds).",
			Default:  3500,
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
			Usage:      "Last 1-based PDF page to parse. When omitted, parsing ends at the last page. Must be greater than or equal to start when both are provided.",
			InnerField: "end",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "pdf.ocr",
			Usage:      "Read scanned PDF pages with OCR; preserve pages that already contain text.",
			InnerField: "ocr",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "pdf.should-parse",
			Usage:      "When true, PDF pages are fetched and parsed. When false, PDF pages are skipped entirely (not included in results and not counted as failures).",
			InnerField: "shouldParse",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "pdf.start",
			Usage:      "First 1-based PDF page to parse. When omitted, parsing starts at the first page.",
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
			Usage:      `"fail" returns 408 at the deadline. "return-partial" returns available results; inspect the response’s partial flag.`,
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
