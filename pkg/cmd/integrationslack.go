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

var integrationsSlackList = cli.Command{
	Name:    "list",
	Usage:   "List Slack conversations accessible to the user via the live Nango connection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:      "channel",
			Usage:     "List of Slack channels to include (by id, name, or #name).",
			Default:   []string{},
			QueryPath: "channels",
		},
		&requestflag.Flag[*bool]{
			Name:      "exclude-archived",
			Usage:     "If set, pass 'exclude_archived' to Slack. If None, omit the param.",
			QueryPath: "exclude_archived",
		},
		&requestflag.Flag[bool]{
			Name:      "include-dms",
			Usage:     "Include direct messages (im) when listing conversations.",
			Default:   false,
			QueryPath: "include_dms",
		},
		&requestflag.Flag[bool]{
			Name:      "include-group-dms",
			Usage:     "Include group DMs (mpim) when listing conversations.",
			Default:   false,
			QueryPath: "include_group_dms",
		},
		&requestflag.Flag[bool]{
			Name:      "include-private",
			Usage:     "Include private channels when constructing Slack 'types'.",
			Default:   false,
			QueryPath: "include_private",
		},
	},
	Action:          handleIntegrationsSlackList,
	HideHelpCommand: true,
}

func handleIntegrationsSlackList(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := hyperspell.IntegrationSlackListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Integrations.Slack.List(ctx, params, options...)
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
		Title:          "integrations:slack list",
		Transform:      transform,
	})
}
