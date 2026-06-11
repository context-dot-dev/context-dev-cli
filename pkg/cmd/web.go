// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/context-dot-dev/context-dev-cli/internal/apiquery"
	"github.com/context-dot-dev/context-dev-cli/internal/requestflag"
	"github.com/context-dot-dev/context-go-sdk"
	"github.com/context-dot-dev/context-go-sdk/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var webExtract = requestflag.WithInnerFlags(cli.Command{
	Name:    "extract",
	Usage:   "Crawl a website, use the provided JSON Schema and instructions to prioritize\nrelevant internal links, and extract structured data from the selected pages.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "schema",
			Usage:    "JSON Schema for the returned data object. TypeScript Zod users can pass a JSON Schema generated from a Zod object; Python users can pass the equivalent JSON Schema object.",
			Required: true,
			BodyPath: "schema",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The starting website URL to crawl and extract from. Must include http:// or https://.",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[bool]{
			Name:     "fact-check",
			Usage:    "When true, every returned value must be grounded in facts stated on the page; fields that cannot be supported by the page are returned as null/empty. When false (default), the model may make reasonable inferences and derivations from the page content (e.g. ideal customer, competitor analysis, recommendations) while keeping verifiable specifics (names, quotes, URLs, dates, metrics) faithful to the source.",
			Default:  false,
			BodyPath: "factCheck",
		},
		&requestflag.Flag[bool]{
			Name:     "follow-subdomains",
			Usage:    "When true, follow links on subdomains of the starting URL's domain.",
			Default:  false,
			BodyPath: "followSubdomains",
		},
		&requestflag.Flag[bool]{
			Name:     "include-frames",
			Usage:    "When true, iframe contents are included in Markdown before extraction.",
			Default:  false,
			BodyPath: "includeFrames",
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			Usage:    "Optional extraction guidance, such as which facts to prioritize or how to interpret fields in the schema.",
			BodyPath: "instructions",
		},
		&requestflag.Flag[int64]{
			Name:     "max-age-ms",
			Usage:    "Return cached scrape results if a prior scrape for the same parameters is younger than this many milliseconds. Defaults to 7 days (604800000 ms).",
			Default:  604800000,
			BodyPath: "maxAgeMs",
		},
		&requestflag.Flag[int64]{
			Name:     "max-depth",
			Usage:    "Optional maximum link depth from the starting URL (0 = only the starting page). If omitted, there is no crawl depth limit.",
			BodyPath: "maxDepth",
		},
		&requestflag.Flag[int64]{
			Name:     "max-pages",
			Usage:    "Maximum number of pages to analyze for extraction. Hard cap: 50. Defaults to 5.",
			Default:  5,
			BodyPath: "maxPages",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "pdf",
			Default:  map[string]any{"shouldParse": true},
			BodyPath: "pdf",
		},
		&requestflag.Flag[int64]{
			Name:     "stop-after-ms",
			Usage:    "Soft time budget for the crawl in milliseconds. Min: 10000 (10s). Max: 110000 (110s). Default: 80000 (80s).",
			Default:  80000,
			BodyPath: "stopAfterMs",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-ms",
			Usage:    "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			BodyPath: "timeoutMS",
		},
		&requestflag.Flag[int64]{
			Name:     "wait-for-ms",
			Usage:    "Optional browser wait time in milliseconds after initial page load for each crawled page.",
			BodyPath: "waitForMs",
		},
	},
	Action:          handleWebExtract,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"pdf": {
		&requestflag.InnerFlag[int64]{
			Name:       "pdf.end",
			Usage:      "Last 1-based PDF page to parse. Must be greater than or equal to start when both are provided.",
			InnerField: "end",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "pdf.should-parse",
			Usage:      "When true, PDF pages are fetched and parsed. When false, PDF pages are skipped.",
			InnerField: "shouldParse",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "pdf.start",
			Usage:      "First 1-based PDF page to parse.",
			InnerField: "start",
		},
	},
})

var webExtractCompetitors = cli.Command{
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
		&requestflag.Flag[int64]{
			Name:      "timeout-ms",
			Usage:     "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			QueryPath: "timeoutMS",
		},
	},
	Action:          handleWebExtractCompetitors,
	HideHelpCommand: true,
}

var webExtractFonts = cli.Command{
	Name:    "extract-fonts",
	Usage:   "Scrape font information from a website including font families, usage\nstatistics, fallbacks, and element/word counts.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "direct-url",
			Usage:     "A specific URL to fetch fonts from directly, bypassing domain resolution (e.g., 'https://example.com/design-system'). When provided, fonts are extracted from this exact URL. You must provide either 'domain' or 'directUrl', but not both.",
			QueryPath: "directUrl",
		},
		&requestflag.Flag[string]{
			Name:      "domain",
			Usage:     "Domain name to extract fonts from (e.g., 'example.com', 'google.com'). The domain will be automatically normalized and validated. You must provide either 'domain' or 'directUrl', but not both.",
			QueryPath: "domain",
		},
		&requestflag.Flag[int64]{
			Name:      "max-age-ms",
			Usage:     "Maximum age in milliseconds for cached data before the API performs a hard refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms) are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1 year.",
			Default:   7776000000,
			QueryPath: "maxAgeMs",
		},
		&requestflag.Flag[int64]{
			Name:      "timeout-ms",
			Usage:     "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			QueryPath: "timeoutMS",
		},
	},
	Action:          handleWebExtractFonts,
	HideHelpCommand: true,
}

var webExtractStyleguide = cli.Command{
	Name:    "extract-styleguide",
	Usage:   "Extract a comprehensive design system from a website including colors,\ntypography, spacing, shadows, and UI components.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "direct-url",
			Usage:     "A specific URL to fetch the styleguide from directly, bypassing domain resolution (e.g., 'https://example.com/design-system'). When provided, the styleguide is extracted from this exact URL. You must provide either 'domain' or 'directUrl', but not both.",
			QueryPath: "directUrl",
		},
		&requestflag.Flag[string]{
			Name:      "domain",
			Usage:     "Domain name to extract styleguide from (e.g., 'example.com', 'google.com'). The domain will be automatically normalized and validated. You must provide either 'domain' or 'directUrl', but not both.",
			QueryPath: "domain",
		},
		&requestflag.Flag[int64]{
			Name:      "max-age-ms",
			Usage:     "Maximum age in milliseconds for cached data before the API performs a hard refresh. Defaults to 3 months (7776000000 ms). Values below 1 day (86400000 ms) are clamped to 1 day; values above 1 year (31536000000 ms) are clamped to 1 year.",
			Default:   7776000000,
			QueryPath: "maxAgeMs",
		},
		&requestflag.Flag[int64]{
			Name:      "timeout-ms",
			Usage:     "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			QueryPath: "timeoutMS",
		},
	},
	Action:          handleWebExtractStyleguide,
	HideHelpCommand: true,
}

var webScreenshot = requestflag.WithInnerFlags(cli.Command{
	Name:    "screenshot",
	Usage:   "Capture a screenshot of a website.",
	Suggest: true,
	Flags: []cli.Flag{
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
		&requestflag.Flag[string]{
			Name:      "handle-cookie-popup",
			Usage:     "Optional parameter to control cookie/consent popup handling. If 'true', we dismiss cookie banner before capture. If 'false' or not provided, captures the page without that step.",
			Default:   "false",
			QueryPath: "handleCookiePopup",
		},
		&requestflag.Flag[int64]{
			Name:      "max-age-ms",
			Usage:     "Return a cached screenshot if a prior screenshot for the same parameters exists and is younger than this many milliseconds. Defaults to 1 day (86400000 ms) when omitted. Max is 30 days (2592000000 ms). Set to 0 to always capture fresh.",
			Default:   86400000,
			QueryPath: "maxAgeMs",
		},
		&requestflag.Flag[string]{
			Name:      "page",
			Usage:     "Optional parameter to specify which page type to screenshot. If provided, the system will scrape the domain's links and use heuristics to find the most appropriate URL for the specified page type (30 supported languages). If not provided, screenshots the main domain landing page. Only applicable when using 'domain', not 'directUrl'.",
			QueryPath: "page",
		},
		&requestflag.Flag[int64]{
			Name:      "timeout-ms",
			Usage:     "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			QueryPath: "timeoutMS",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "viewport",
			Usage:     "Optional browser viewport dimensions for the screenshot. Defaults to 1920x1080.",
			Default:   map[string]any{"width": 1920, "height": 1080},
			QueryPath: "viewport",
		},
		&requestflag.Flag[int64]{
			Name:      "wait-for-ms",
			Usage:     "Optional browser wait time in milliseconds after initial page load before taking the screenshot. Min: 0. Max: 30000 (30 seconds).  Defaults to 3000 ms when omitted.",
			Default:   3000,
			QueryPath: "waitForMs",
		},
	},
	Action:          handleWebScreenshot,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
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
	Usage:   "Search the web and optionally scrape each result to Markdown in one round-trip.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "Natural-language search query.",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[[]string]{
			Name:     "exclude-domain",
			Usage:    `Blocklist — drop results from these domains. Example: ["pinterest.com", "reddit.com"].`,
			BodyPath: "excludeDomains",
		},
		&requestflag.Flag[string]{
			Name:     "freshness",
			Usage:    "Restrict results to content published within this window.",
			BodyPath: "freshness",
		},
		&requestflag.Flag[[]string]{
			Name:     "include-domain",
			Usage:    `Allowlist — only return results from these domains. Example: ["arxiv.org", "github.com"].`,
			BodyPath: "includeDomains",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "markdown-options",
			Usage:    "Inline Markdown scraping for each result. Set `enabled: true` to activate.",
			BodyPath: "markdownOptions",
		},
		&requestflag.Flag[bool]{
			Name:     "query-fanout",
			Usage:    "Expand the query into multiple parallel variants for broader recall.",
			BodyPath: "queryFanout",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-ms",
			Usage:    "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			BodyPath: "timeoutMS",
		},
	},
	Action:          handleWebSearch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"markdown-options": {
		&requestflag.InnerFlag[bool]{
			Name:       "markdown-options.enabled",
			Usage:      "Scrape each result to Markdown. Off by default to keep search cheap and fast.",
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
			Usage:      "Cache TTL in ms for scraped Markdown keyed by URL + options. Default 1 day, max 30 days. Set to 0 to force a fresh scrape.",
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
		&requestflag.InnerFlag[int64]{
			Name:       "markdown-options.timeout-ms",
			Usage:      "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			InnerField: "timeoutMS",
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
})

var webWebCrawlMd = requestflag.WithInnerFlags(cli.Command{
	Name:    "web-crawl-md",
	Usage:   "Performs a crawl starting from a given URL, extracts page content as Markdown,\nand returns results for all crawled pages.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The starting URL for the crawl (must include http:// or https:// protocol)",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[[]string]{
			Name:     "exclude-selector",
			Usage:    `CSS selectors to remove before each crawled page is converted to Markdown. Applied after includeSelectors. Exclusion takes precedence: an element matching both is removed. Examples: "nav", "footer", ".ad-banner", "[aria-hidden=true]".`,
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
			Usage:    `CSS selectors. When provided, only matching HTML subtrees (and their descendants) are kept before each crawled page is converted to Markdown. When omitted, the entire document is kept. Examples: "article.main", "#content", "[role=main]".`,
			BodyPath: "includeSelectors",
		},
		&requestflag.Flag[int64]{
			Name:     "max-age-ms",
			Usage:    "Return a cached result if a prior scrape for the same parameters exists and is younger than this many milliseconds. Defaults to 1 day (86400000 ms) when omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.",
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
			Usage:    "Maximum number of pages to crawl. Hard cap: 500.",
			Default:  100,
			BodyPath: "maxPages",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "pdf",
			Usage:    "PDF parsing controls. Use start/end to limit text extraction and OCR to an inclusive 1-based page range.",
			Default:  map[string]any{"shouldParse": true},
			BodyPath: "pdf",
		},
		&requestflag.Flag[bool]{
			Name:     "shorten-base64-images",
			Usage:    "Truncate base64-encoded image data in the Markdown output",
			Default:  true,
			BodyPath: "shortenBase64Images",
		},
		&requestflag.Flag[int64]{
			Name:     "stop-after-ms",
			Usage:    "Soft time budget for the crawl in milliseconds. After each scrape, the crawler checks the elapsed time and, if exceeded, returns the pages collected so far instead of continuing. Min: 10000 (10s). Max: 110000 (110s). Default: 80000 (80s).",
			Default:  80000,
			BodyPath: "stopAfterMs",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-ms",
			Usage:    "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			BodyPath: "timeoutMS",
		},
		&requestflag.Flag[string]{
			Name:     "url-regex",
			Usage:    "Regex pattern. Only URLs matching this pattern will be followed and scraped.",
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
			Usage:    "Optional browser wait time in milliseconds after initial page load for each crawled page. Min: 0. Max: 30000 (30 seconds). ",
			BodyPath: "waitForMs",
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
})

var webWebScrapeHTML = requestflag.WithInnerFlags(cli.Command{
	Name:    "web-scrape-html",
	Usage:   "Scrapes the given URL and returns the raw HTML content of the page.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "url",
			Usage:     "Full URL to scrape (must include http:// or https:// protocol)",
			Required:  true,
			QueryPath: "url",
		},
		&requestflag.Flag[[]string]{
			Name:      "exclude-selector",
			Usage:     `CSS selectors to remove from the result. Applied after includeSelectors. Exclusion takes precedence: an element matching both is removed. Examples: "nav", "footer", ".ad-banner", "[aria-hidden=true]".`,
			QueryPath: "excludeSelectors",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "headers",
			Usage:     "Optional outbound HTTP headers forwarded only to the target URL, sent as deep-object query params such as headers[X-Custom]=value. When provided, caching is bypassed: the result is neither read from nor written to cache.",
			QueryPath: "headers",
		},
		&requestflag.Flag[bool]{
			Name:      "include-frames",
			Usage:     "When true, iframes are rendered inline into the returned HTML.",
			Default:   false,
			QueryPath: "includeFrames",
		},
		&requestflag.Flag[[]string]{
			Name:      "include-selector",
			Usage:     `CSS selectors. When provided, only matching subtrees (and their descendants) are kept and everything else is dropped. When omitted, the entire document is kept. Examples: "article.main", "#content", "[role=main]".`,
			QueryPath: "includeSelectors",
		},
		&requestflag.Flag[int64]{
			Name:      "max-age-ms",
			Usage:     "Return a cached result if a prior scrape for the same parameters exists and is younger than this many milliseconds. Defaults to 1 day (86400000 ms) when omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.",
			Default:   86400000,
			QueryPath: "maxAgeMs",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "pdf",
			Usage:     "PDF parsing controls. Use start/end to limit text extraction and OCR to an inclusive 1-based page range.",
			Default:   map[string]any{"shouldParse": true},
			QueryPath: "pdf",
		},
		&requestflag.Flag[int64]{
			Name:      "timeout-ms",
			Usage:     "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			QueryPath: "timeoutMS",
		},
		&requestflag.Flag[bool]{
			Name:      "use-main-content-only",
			Usage:     "When true, return only the page's main content in the HTML response, excluding headers, footers, sidebars, and navigation when detectable.",
			Default:   false,
			QueryPath: "useMainContentOnly",
		},
		&requestflag.Flag[int64]{
			Name:      "wait-for-ms",
			Usage:     "Optional browser wait time in milliseconds after initial page load. Min: 0. Max: 30000 (30 seconds). ",
			QueryPath: "waitForMs",
		},
	},
	Action:          handleWebWebScrapeHTML,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"pdf": {
		&requestflag.InnerFlag[int64]{
			Name:       "pdf.end",
			Usage:      "Last 1-based PDF page to parse. When omitted, parsing ends at the last page. Must be greater than or equal to start when both are provided.",
			InnerField: "end",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "pdf.should-parse",
			Usage:      "When true, PDF URLs are fetched and parsed. When false, PDF URLs are skipped and a 400 WEBSITE_ACCESS_ERROR is returned.",
			InnerField: "shouldParse",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "pdf.start",
			Usage:      "First 1-based PDF page to parse. When omitted, parsing starts at the first page.",
			InnerField: "start",
		},
	},
})

var webWebScrapeImages = requestflag.WithInnerFlags(cli.Command{
	Name:    "web-scrape-images",
	Usage:   "Extract image assets from a web page, including standard URLs, inline SVGs, data\nURIs, responsive image sources, metadata, CSS backgrounds, video posters, and\nembeds. The base request costs 1 credit. When enrichment is enabled, the entire\ncall costs 5 credits.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "url",
			Usage:     "Page URL to inspect. Must include http:// or https://.",
			Required:  true,
			QueryPath: "url",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "enrichment",
			Usage:     "Optional per-image processing, sent as deep-object query params such as enrichment[resolution]=true.",
			QueryPath: "enrichment",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "headers",
			Usage:     "Optional outbound HTTP headers forwarded only to the target URL, sent as deep-object query params such as headers[X-Custom]=value. When provided, caching is bypassed: the result is neither read from nor written to cache.",
			QueryPath: "headers",
		},
		&requestflag.Flag[int64]{
			Name:      "max-age-ms",
			Usage:     "Reuse a cached result this many milliseconds old or newer. Default: 86400000 (1 day). Set to 0 to bypass cache. Maximum: 2592000000 (30 days).",
			Default:   86400000,
			QueryPath: "maxAgeMs",
		},
		&requestflag.Flag[int64]{
			Name:      "timeout-ms",
			Usage:     "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			QueryPath: "timeoutMS",
		},
		&requestflag.Flag[int64]{
			Name:      "wait-for-ms",
			Usage:     "Optional browser wait time in milliseconds after initial page load before collecting images. Min: 0. Max: 30000 (30 seconds). ",
			QueryPath: "waitForMs",
		},
	},
	Action:          handleWebWebScrapeImages,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"enrichment": {
		&requestflag.InnerFlag[bool]{
			Name:       "enrichment.classification",
			Usage:      "Classify each image by visual asset type.",
			InnerField: "classification",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "enrichment.hosted-url",
			Usage:      "Host materializable images on the Brand.dev CDN and return their URL and MIME type.",
			InnerField: "hostedUrl",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "enrichment.max-time-per-ms",
			Usage:      "Per-image enrichment timeout in milliseconds. Default: 30000. Maximum: 60000.",
			InnerField: "maxTimePerMs",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "enrichment.resolution",
			Usage:      "Measure image width and height when possible.",
			InnerField: "resolution",
		},
	},
})

var webWebScrapeMd = requestflag.WithInnerFlags(cli.Command{
	Name:    "web-scrape-md",
	Usage:   "Scrapes the given URL into LLM usable Markdown.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "url",
			Usage:     "Full URL to scrape into LLM usable Markdown (must include http:// or https:// protocol)",
			Required:  true,
			QueryPath: "url",
		},
		&requestflag.Flag[[]string]{
			Name:      "exclude-selector",
			Usage:     `CSS selectors to remove before conversion to Markdown. Applied after includeSelectors. Exclusion takes precedence: an element matching both is removed. Examples: "nav", "footer", ".ad-banner", "[aria-hidden=true]".`,
			QueryPath: "excludeSelectors",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "headers",
			Usage:     "Optional outbound HTTP headers forwarded only to the target URL, sent as deep-object query params such as headers[X-Custom]=value. When provided, caching is bypassed: the result is neither read from nor written to cache.",
			QueryPath: "headers",
		},
		&requestflag.Flag[bool]{
			Name:      "include-frames",
			Usage:     "When true, the contents of iframes are rendered to Markdown.",
			Default:   false,
			QueryPath: "includeFrames",
		},
		&requestflag.Flag[bool]{
			Name:      "include-images",
			Usage:     "Include image references in Markdown output",
			Default:   false,
			QueryPath: "includeImages",
		},
		&requestflag.Flag[bool]{
			Name:      "include-links",
			Usage:     "Preserve hyperlinks in Markdown output",
			Default:   true,
			QueryPath: "includeLinks",
		},
		&requestflag.Flag[[]string]{
			Name:      "include-selector",
			Usage:     `CSS selectors. When provided, only matching HTML subtrees (and their descendants) are kept before conversion to Markdown. When omitted, the entire document is kept. Examples: "article.main", "#content", "[role=main]".`,
			QueryPath: "includeSelectors",
		},
		&requestflag.Flag[int64]{
			Name:      "max-age-ms",
			Usage:     "Return a cached result if a prior scrape for the same parameters exists and is younger than this many milliseconds. Defaults to 1 day (86400000 ms) when omitted. Max is 30 days (2592000000 ms). Set to 0 to always scrape fresh.",
			Default:   86400000,
			QueryPath: "maxAgeMs",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "pdf",
			Usage:     "PDF parsing controls. Use start/end to limit text extraction and OCR to an inclusive 1-based page range.",
			Default:   map[string]any{"shouldParse": true},
			QueryPath: "pdf",
		},
		&requestflag.Flag[bool]{
			Name:      "shorten-base64-images",
			Usage:     "Shorten base64-encoded image data in the Markdown output",
			Default:   true,
			QueryPath: "shortenBase64Images",
		},
		&requestflag.Flag[int64]{
			Name:      "timeout-ms",
			Usage:     "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			QueryPath: "timeoutMS",
		},
		&requestflag.Flag[bool]{
			Name:      "use-main-content-only",
			Usage:     "Extract only the main content of the page, excluding headers, footers, sidebars, and navigation",
			Default:   false,
			QueryPath: "useMainContentOnly",
		},
		&requestflag.Flag[int64]{
			Name:      "wait-for-ms",
			Usage:     "Optional browser wait time in milliseconds after initial page load before converting the page to Markdown. Min: 0. Max: 30000 (30 seconds). ",
			QueryPath: "waitForMs",
		},
	},
	Action:          handleWebWebScrapeMd,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"pdf": {
		&requestflag.InnerFlag[int64]{
			Name:       "pdf.end",
			Usage:      "Last 1-based PDF page to parse. When omitted, parsing ends at the last page. Must be greater than or equal to start when both are provided.",
			InnerField: "end",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "pdf.should-parse",
			Usage:      "When true, PDF URLs are fetched and parsed. When false, PDF URLs are skipped and a 400 WEBSITE_ACCESS_ERROR is returned.",
			InnerField: "shouldParse",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "pdf.start",
			Usage:      "First 1-based PDF page to parse. When omitted, parsing starts at the first page.",
			InnerField: "start",
		},
	},
})

var webWebScrapeSitemap = cli.Command{
	Name:    "web-scrape-sitemap",
	Usage:   "Crawl an entire website's sitemap and return all discovered page URLs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "domain",
			Usage:     "Domain to build a sitemap for",
			Required:  true,
			QueryPath: "domain",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "headers",
			Usage:     "Optional outbound HTTP headers forwarded only to the target URL, sent as deep-object query params such as headers[X-Custom]=value. When provided, caching is bypassed: the result is neither read from nor written to cache.",
			QueryPath: "headers",
		},
		&requestflag.Flag[int64]{
			Name:      "max-links",
			Usage:     "Maximum number of links to return from the sitemap crawl. Defaults to 10,000. Minimum is 1, maximum is 100,000.",
			Default:   10000,
			QueryPath: "maxLinks",
		},
		&requestflag.Flag[int64]{
			Name:      "timeout-ms",
			Usage:     "Optional timeout in milliseconds for the request. If the request takes longer than this value, it will be aborted with a 408 status code. Maximum allowed value is 300000ms (5 minutes).",
			QueryPath: "timeoutMS",
		},
		&requestflag.Flag[string]{
			Name:      "url-regex",
			Usage:     "Optional RE2-compatible regex pattern. Only URLs matching this pattern are returned and counted against maxLinks.",
			QueryPath: "urlRegex",
		},
	},
	Action:          handleWebWebScrapeSitemap,
	HideHelpCommand: true,
}

func handleWebExtract(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.WebExtractParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Web.Extract(ctx, params, options...)
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
		Title:          "web extract",
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

func handleWebExtractFonts(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.WebExtractFontsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Web.ExtractFonts(ctx, params, options...)
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
		Title:          "web extract-fonts",
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

func handleWebWebScrapeHTML(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.WebWebScrapeHTMLParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Web.WebScrapeHTML(ctx, params, options...)
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
		Title:          "web web-scrape-html",
		Transform:      transform,
	})
}

func handleWebWebScrapeImages(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.WebWebScrapeImagesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Web.WebScrapeImages(ctx, params, options...)
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
		Title:          "web web-scrape-images",
		Transform:      transform,
	})
}

func handleWebWebScrapeMd(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.WebWebScrapeMdParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Web.WebScrapeMd(ctx, params, options...)
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
		Title:          "web web-scrape-md",
		Transform:      transform,
	})
}

func handleWebWebScrapeSitemap(ctx context.Context, cmd *cli.Command) error {
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

	params := contextdev.WebWebScrapeSitemapParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Web.WebScrapeSitemap(ctx, params, options...)
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
		Title:          "web web-scrape-sitemap",
		Transform:      transform,
	})
}
