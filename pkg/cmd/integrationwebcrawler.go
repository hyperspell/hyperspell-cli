// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/hyperspell/hyperspell-cli/internal/apiquery"
	"github.com/hyperspell/hyperspell-cli/internal/requestflag"
	"github.com/hyperspell/hyperspell-go"
	"github.com/hyperspell/hyperspell-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var integrationsWebCrawlerIndex = cli.Command{
	Name:    "index",
	Usage:   "Recursively crawl a website to make it available for indexed search.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "url",
			Usage:     "The base URL of the website to crawl",
			Required:  true,
			QueryPath: "url",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of pages to crawl in total",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "max-depth",
			Usage:     "Maximum depth of links to follow during crawling",
			Default:   2,
			QueryPath: "max_depth",
		},
	},
	Action:          handleIntegrationsWebCrawlerIndex,
	HideHelpCommand: true,
}

func handleIntegrationsWebCrawlerIndex(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := hyperspell.IntegrationWebCrawlerIndexParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Integrations.WebCrawler.Index(ctx, params, options...)
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
		Title:          "integrations:web-crawler index",
		Transform:      transform,
	})
}
