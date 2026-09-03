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

var liveGetResource = cli.Command{
	Name:    "get-resource",
	Usage:   "Fetch one resource live by id. A single fetch may fan out into several resources\n(e.g. a thread → its messages); all are returned.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "source",
			Usage:     `Allowed values: "reddit", "notion", "slack", "google_calendar", "google_mail", "imap", "google_meet", "box", "dropbox", "github", "gitlab", "google_drive", "vault", "web_crawler", "trace", "microsoft_outlook", "microsoft_teams", "granola", "fathom", "fireflies", "figma", "linear", "hubspot", "salesforce", "coda", "confluence", "jira", "metabase", "gong", "clickup", "lightfield", "pylon", "fellow", "odoo", "external_mcp".`,
			Required:  true,
			PathParam: "source",
		},
		&requestflag.Flag[string]{
			Name:      "resource-id",
			Required:  true,
			PathParam: "resource_id",
		},
		&requestflag.Flag[*string]{
			Name:      "connection-id",
			Usage:     "Specific connection id.",
			QueryPath: "connection_id",
		},
		&requestflag.Flag[bool]{
			Name:      "index",
			Usage:     "Also queue this resource for indexing.",
			Default:   false,
			QueryPath: "index",
		},
	},
	Action:          handleLiveGetResource,
	HideHelpCommand: true,
}

var liveListResources = cli.Command{
	Name:    "list-resources",
	Usage:   "Page through a source's resources live (no indexing side effect).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "source",
			Usage:     `Allowed values: "reddit", "notion", "slack", "google_calendar", "google_mail", "imap", "google_meet", "box", "dropbox", "github", "gitlab", "google_drive", "vault", "web_crawler", "trace", "microsoft_outlook", "microsoft_teams", "granola", "fathom", "fireflies", "figma", "linear", "hubspot", "salesforce", "coda", "confluence", "jira", "metabase", "gong", "clickup", "lightfield", "pylon", "fellow", "odoo", "external_mcp".`,
			Required:  true,
			PathParam: "source",
		},
		&requestflag.Flag[*string]{
			Name:      "connection-id",
			Usage:     "Specific connection id.",
			QueryPath: "connection_id",
		},
		&requestflag.Flag[*string]{
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
	Action:          handleLiveListResources,
	HideHelpCommand: true,
}

var liveListSources = cli.Command{
	Name:            "list-sources",
	Usage:           "List the user's connected sources and the live capabilities each supports.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleLiveListSources,
	HideHelpCommand: true,
}

var liveSearch = cli.Command{
	Name:    "search",
	Usage:   "Search a source live for content that may not be indexed yet. With `index=true`,\neach hit is queued for indexing (no-op for live-only sources like Google\nCalendar — see `notes` in the response).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "source",
			Usage:     `Allowed values: "reddit", "notion", "slack", "google_calendar", "google_mail", "imap", "google_meet", "box", "dropbox", "github", "gitlab", "google_drive", "vault", "web_crawler", "trace", "microsoft_outlook", "microsoft_teams", "granola", "fathom", "fireflies", "figma", "linear", "hubspot", "salesforce", "coda", "confluence", "jira", "metabase", "gong", "clickup", "lightfield", "pylon", "fellow", "odoo", "external_mcp".`,
			Required:  true,
			PathParam: "source",
		},
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "Live search query.",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[*string]{
			Name:     "connection-id",
			Usage:    "Specific connection id when the user has multiple for this source.",
			BodyPath: "connection_id",
		},
		&requestflag.Flag[bool]{
			Name:     "index",
			Usage:    "If true, queue each hit for indexing so it's on-hand next time.",
			Default:  false,
			BodyPath: "index",
		},
	},
	Action:          handleLiveSearch,
	HideHelpCommand: true,
}

func handleLiveGetResource(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("resource-id") && len(unusedArgs) > 0 {
		cmd.Set("resource-id", unusedArgs[0])
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

	params := hyperspell.LiveGetResourceParams{
		Source: hyperspell.LiveGetResourceParamsSource(cmd.Value("source").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Live.GetResource(
		ctx,
		cmd.Value("resource-id").(string),
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
		Title:          "live get-resource",
		Transform:      transform,
	})
}

func handleLiveListResources(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("source") && len(unusedArgs) > 0 {
		cmd.Set("source", unusedArgs[0])
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

	params := hyperspell.LiveListResourcesParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Live.ListResources(
			ctx,
			hyperspell.LiveListResourcesParamsSource(cmd.Value("source").(string)),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "live list-resources",
			Transform:      transform,
		})
	} else {
		iter := client.Live.ListResourcesAutoPaging(
			ctx,
			hyperspell.LiveListResourcesParamsSource(cmd.Value("source").(string)),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "live list-resources",
			Transform:      transform,
		})
	}
}

func handleLiveListSources(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Live.ListSources(ctx, options...)
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
		Title:          "live list-sources",
		Transform:      transform,
	})
}

func handleLiveSearch(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("source") && len(unusedArgs) > 0 {
		cmd.Set("source", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := hyperspell.LiveSearchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Live.Search(
		ctx,
		hyperspell.LiveSearchParamsSource(cmd.Value("source").(string)),
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
		Title:          "live search",
		Transform:      transform,
	})
}
