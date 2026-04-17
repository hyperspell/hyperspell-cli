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

var evaluateGetQuery = cli.Command{
	Name:    "get-query",
	Usage:   "Retrieve the result of a previous query.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query-id",
			Required: true,
		},
	},
	Action:          handleEvaluateGetQuery,
	HideHelpCommand: true,
}

var evaluateScoreHighlight = cli.Command{
	Name:    "score-highlight",
	Usage:   "Score an individual highlight.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "highlight-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "comment",
			Usage:    "Comment on the chunk",
			BodyPath: "comment",
		},
		&requestflag.Flag[float64]{
			Name:     "score",
			Usage:    "Rating of the chunk from -1 (bad) to +1 (good).",
			Default:  0,
			BodyPath: "score",
		},
	},
	Action:          handleEvaluateScoreHighlight,
	HideHelpCommand: true,
}

var evaluateScoreQuery = cli.Command{
	Name:    "score-query",
	Usage:   "Score the result of a query.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query-id",
			Required: true,
		},
		&requestflag.Flag[float64]{
			Name:     "score",
			Usage:    "Rating of the query result from -1 (bad) to +1 (good).",
			Default:  0,
			BodyPath: "score",
		},
	},
	Action:          handleEvaluateScoreQuery,
	HideHelpCommand: true,
}

func handleEvaluateGetQuery(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("query-id") && len(unusedArgs) > 0 {
		cmd.Set("query-id", unusedArgs[0])
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
	_, err = client.Evaluate.GetQuery(ctx, cmd.Value("query-id").(string), options...)
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
		Title:          "evaluate get-query",
		Transform:      transform,
	})
}

func handleEvaluateScoreHighlight(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("highlight-id") && len(unusedArgs) > 0 {
		cmd.Set("highlight-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := hyperspell.EvaluateScoreHighlightParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Evaluate.ScoreHighlight(
		ctx,
		cmd.Value("highlight-id").(string),
		params,
		options...,
	)
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
		Title:          "evaluate score-highlight",
		Transform:      transform,
	})
}

func handleEvaluateScoreQuery(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("query-id") && len(unusedArgs) > 0 {
		cmd.Set("query-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := hyperspell.EvaluateScoreQueryParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Evaluate.ScoreQuery(
		ctx,
		cmd.Value("query-id").(string),
		params,
		options...,
	)
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
		Title:          "evaluate score-query",
		Transform:      transform,
	})
}
