// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/hyperspell-cli/internal/apiquery"
	"github.com/stainless-sdks/hyperspell-cli/internal/requestflag"
	"github.com/stainless-sdks/hyperspell-go"
	"github.com/stainless-sdks/hyperspell-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var vaultsList = cli.Command{
	Name:    "list",
	Usage:   "This endpoint lists all collections, and how many documents are in each\ncollection. All documents that do not have a collection assigned are in the\n`null` collection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "size",
			Default:   50,
			QueryPath: "size",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleVaultsList,
	HideHelpCommand: true,
}

func handleVaultsList(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := hyperspell.VaultListParams{}

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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Vaults.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "vaults list", obj, format, transform)
	} else {
		iter := client.Vaults.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "vaults list", iter, format, transform, maxItems)
	}
}
