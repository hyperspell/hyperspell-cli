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

var entitiesList = cli.Command{
	Name:    "list",
	Usage:   "List entities available to the current app.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:      "cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   100,
			QueryPath: "limit",
		},
		&requestflag.Flag[*int64]{
			Name:      "min-supporting-documents",
			QueryPath: "min_supporting_documents",
		},
		&requestflag.Flag[*string]{
			Name:      "name-prefix",
			QueryPath: "name_prefix",
		},
		&requestflag.Flag[*string]{
			Name:      "search",
			QueryPath: "search",
		},
		&requestflag.Flag[string]{
			Name:      "sort-by",
			Usage:     `Allowed values: "id", "name", "type", "prominence".`,
			Default:   "id",
			QueryPath: "sort_by",
		},
		&requestflag.Flag[string]{
			Name:      "sort-dir",
			Usage:     `Allowed values: "asc", "desc".`,
			Default:   "asc",
			QueryPath: "sort_dir",
		},
		&requestflag.Flag[*string]{
			Name:      "status",
			Usage:     "How strongly the entity's current identity has been established.",
			QueryPath: "status",
		},
		&requestflag.Flag[*string]{
			Name:      "type",
			QueryPath: "type",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleEntitiesList,
	HideHelpCommand: true,
}

var entitiesGet = cli.Command{
	Name:    "get",
	Usage:   "Fetch a single entity belonging to the current app.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "entity-id",
			Required:  true,
			PathParam: "entity_id",
		},
	},
	Action:          handleEntitiesGet,
	HideHelpCommand: true,
}

var entitiesSearch = cli.Command{
	Name:    "search",
	Usage:   "Search the current app's entities by meaning.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[int64]{
			Name:     "limit",
			Default:  20,
			BodyPath: "limit",
		},
		&requestflag.Flag[*string]{
			Name:     "type",
			BodyPath: "type",
		},
	},
	Action:          handleEntitiesSearch,
	HideHelpCommand: true,
}

func handleEntitiesList(ctx context.Context, cmd *cli.Command) error {
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

	params := hyperspell.EntityListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Entities.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "entities list",
			Transform:      transform,
		})
	} else {
		iter := client.Entities.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "entities list",
			Transform:      transform,
		})
	}
}

func handleEntitiesGet(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("entity-id") && len(unusedArgs) > 0 {
		cmd.Set("entity-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Entities.Get(ctx, cmd.Value("entity-id").(string), options...)
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
		Title:          "entities get",
		Transform:      transform,
	})
}

func handleEntitiesSearch(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := hyperspell.EntitySearchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Entities.Search(ctx, params, options...)
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
		Title:          "entities search",
		Transform:      transform,
	})
}
