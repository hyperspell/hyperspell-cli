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

var memoriesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates an existing document in the index. You can update the text, collection,\ntitle, and metadata. The document must already exist or a 404 will be returned.\nThis works for documents from any source (vault, slack, gmail, etc.).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "source",
			Usage:    `Allowed values: "reddit", "notion", "slack", "google_calendar", "google_mail", "box", "dropbox", "github", "google_drive", "vault", "web_crawler", "trace", "microsoft_teams", "gmail_actions".`,
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "resource-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "collection",
			Usage:    "The collection to move the document to — deprecated, set the collection using metadata instead.",
			Default:  map[string]any{},
			BodyPath: "collection",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Custom metadata for filtering. Keys must be alphanumeric with underscores, max 64 chars. Values must be string, number, boolean, or null. Will be merged with existing metadata.",
			Default:  map[string]any{},
			BodyPath: "metadata",
		},
		&requestflag.Flag[any]{
			Name:     "text",
			Usage:    "Full text of the document. If provided, the document will be re-indexed.",
			Default:  map[string]any{},
			BodyPath: "text",
		},
		&requestflag.Flag[any]{
			Name:     "title",
			Usage:    "Title of the document.",
			Default:  map[string]any{},
			BodyPath: "title",
		},
	},
	Action:          handleMemoriesUpdate,
	HideHelpCommand: true,
}

var memoriesList = cli.Command{
	Name:    "list",
	Usage:   "This endpoint allows you to paginate through all documents in the index. You can\nfilter the documents by title, date, metadata, etc.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "collection",
			Usage:     "Filter documents by collection.",
			QueryPath: "collection",
		},
		&requestflag.Flag[any]{
			Name:      "cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[any]{
			Name:      "filter",
			Usage:     `Filter documents by metadata using MongoDB-style operators. Example: {"department": "engineering", "priority": {"$gt": 3}}`,
			QueryPath: "filter",
		},
		&requestflag.Flag[int64]{
			Name:      "size",
			Default:   50,
			QueryPath: "size",
		},
		&requestflag.Flag[any]{
			Name:      "source",
			Usage:     "Filter documents by source.",
			QueryPath: "source",
		},
		&requestflag.Flag[any]{
			Name:      "status",
			Usage:     "Filter documents by status.",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleMemoriesList,
	HideHelpCommand: true,
}

var memoriesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a memory and its associated chunks from the index.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "source",
			Usage:    `Allowed values: "reddit", "notion", "slack", "google_calendar", "google_mail", "box", "dropbox", "github", "google_drive", "vault", "web_crawler", "trace", "microsoft_teams", "gmail_actions".`,
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "resource-id",
			Required: true,
		},
	},
	Action:          handleMemoriesDelete,
	HideHelpCommand: true,
}

var memoriesAdd = cli.Command{
	Name:    "add",
	Usage:   "Adds an arbitrary document to the index. This can be any text, email, call\ntranscript, etc. The document will be processed and made available for querying\nonce the processing is complete.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Full text of the document.",
			Required: true,
			BodyPath: "text",
		},
		&requestflag.Flag[any]{
			Name:     "collection",
			Usage:    "The collection to add the document to — deprecated, set the collection using metadata instead.",
			BodyPath: "collection",
		},
		&requestflag.Flag[any]{
			Name:     "date",
			Usage:    "Date of the document. Depending on the document, this could be the creation date or date the document was last updated (eg. for a chat transcript, this would be the date of the last message). This helps the ranking algorithm and allows you to filter by date range.",
			BodyPath: "date",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Custom metadata for filtering. Keys must be alphanumeric with underscores, max 64 chars. Values must be string, number, boolean, or null.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "resource-id",
			Usage:    "The resource ID to add the document to. If not provided, a new resource ID will be generated. If provided, the document will be updated if it already exists.",
			BodyPath: "resource_id",
		},
		&requestflag.Flag[any]{
			Name:     "title",
			Usage:    "Title of the document.",
			BodyPath: "title",
		},
	},
	Action:          handleMemoriesAdd,
	HideHelpCommand: true,
}

var memoriesAddBulk = requestflag.WithInnerFlags(cli.Command{
	Name:    "add-bulk",
	Usage:   "Adds multiple documents to the index in a single request.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "item",
			Usage:    "List of memories to ingest. Maximum 100 items.",
			Required: true,
			BodyPath: "items",
		},
	},
	Action:          handleMemoriesAddBulk,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"item": {
		&requestflag.InnerFlag[string]{
			Name:       "item.text",
			Usage:      "Full text of the document.",
			InnerField: "text",
		},
		&requestflag.InnerFlag[any]{
			Name:       "item.collection",
			Usage:      "The collection to add the document to — deprecated, set the collection using metadata instead.",
			InnerField: "collection",
		},
		&requestflag.InnerFlag[any]{
			Name:       "item.date",
			Usage:      "Date of the document. Depending on the document, this could be the creation date or date the document was last updated (eg. for a chat transcript, this would be the date of the last message). This helps the ranking algorithm and allows you to filter by date range.",
			InnerField: "date",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "item.metadata",
			Usage:      "Custom metadata for filtering. Keys must be alphanumeric with underscores, max 64 chars. Values must be string, number, boolean, or null.",
			InnerField: "metadata",
		},
		&requestflag.InnerFlag[string]{
			Name:       "item.resource-id",
			Usage:      "The resource ID to add the document to. If not provided, a new resource ID will be generated. If provided, the document will be updated if it already exists.",
			InnerField: "resource_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "item.title",
			Usage:      "Title of the document.",
			InnerField: "title",
		},
	},
})

var memoriesGet = cli.Command{
	Name:    "get",
	Usage:   "Retrieves a document by provider and resource_id.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "source",
			Usage:    `Allowed values: "reddit", "notion", "slack", "google_calendar", "google_mail", "box", "dropbox", "github", "google_drive", "vault", "web_crawler", "trace", "microsoft_teams", "gmail_actions".`,
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "resource-id",
			Required: true,
		},
	},
	Action:          handleMemoriesGet,
	HideHelpCommand: true,
}

var memoriesSearch = requestflag.WithInnerFlags(cli.Command{
	Name:    "search",
	Usage:   "Retrieves documents matching the query.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "Query to run.",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[bool]{
			Name:     "answer",
			Usage:    "If true, the query will be answered along with matching source documents.",
			Default:  false,
			BodyPath: "answer",
		},
		&requestflag.Flag[int64]{
			Name:     "effort",
			Usage:    "Effort level. 0 = pass query through verbatim. 1 = LLM rewrites the query for better retrieval and extracts date filters.",
			Default:  0,
			BodyPath: "effort",
		},
		&requestflag.Flag[int64]{
			Name:     "max-results",
			Usage:    "Maximum number of results to return.",
			Default:  10,
			BodyPath: "max_results",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "options",
			Usage:    "Search options for the query.",
			BodyPath: "options",
		},
		&requestflag.Flag[[]string]{
			Name:     "source",
			Usage:    "Only query documents from these sources.",
			BodyPath: "sources",
		},
	},
	Action:          handleMemoriesSearch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"options": {
		&requestflag.InnerFlag[any]{
			Name:       "options.after",
			Usage:      "Only query documents created on or after this date.",
			InnerField: "after",
		},
		&requestflag.InnerFlag[string]{
			Name:       "options.answer-model",
			Usage:      "Model to use for answer generation when answer=True",
			InnerField: "answer_model",
		},
		&requestflag.InnerFlag[any]{
			Name:       "options.before",
			Usage:      "Only query documents created before this date.",
			InnerField: "before",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "options.box",
			Usage:      "Search options for Box",
			InnerField: "box",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "options.filter",
			Usage:      "Metadata filters using MongoDB-style operators. Example: {'status': 'published', 'priority': {'$gt': 3}}",
			InnerField: "filter",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "options.google-calendar",
			Usage:      "Search options for Google Calendar",
			InnerField: "google_calendar",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "options.google-drive",
			Usage:      "Search options for Google Drive",
			InnerField: "google_drive",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "options.google-mail",
			Usage:      "Search options for Gmail",
			InnerField: "google_mail",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "options.max-results",
			Usage:      "Maximum number of results to return.",
			InnerField: "max_results",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "options.memory-types",
			Usage:      "Filter by memory type. Defaults to generic memories only. Pass multiple types to include procedures, etc.",
			InnerField: "memory_types",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "options.notion",
			Usage:      "Search options for Notion",
			InnerField: "notion",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "options.reddit",
			Usage:      "Search options for Reddit",
			InnerField: "reddit",
		},
		&requestflag.InnerFlag[any]{
			Name:       "options.resource-ids",
			Usage:      "Only return results from these specific resource IDs. Useful for scoping searches to specific documents (e.g., a specific email thread or uploaded file).",
			InnerField: "resource_ids",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "options.slack",
			Usage:      "Search options for Slack",
			InnerField: "slack",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "options.vault",
			Usage:      "Search options for vault",
			InnerField: "vault",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "options.web-crawler",
			Usage:      "Search options for Web Crawler",
			InnerField: "web_crawler",
		},
	},
})

var memoriesStatus = cli.Command{
	Name:            "status",
	Usage:           "This endpoint shows the indexing progress of documents, both by provider and\ntotal.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleMemoriesStatus,
	HideHelpCommand: true,
}

var memoriesUpload = cli.Command{
	Name:    "upload",
	Usage:   "This endpoint will upload a file to the index and return a resource_id. The file\nwill be processed in the background and the memory will be available for\nquerying once the processing is complete. You can use the `resource_id` to query\nthe memory later, and check the status of the memory.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "file",
			Usage:     "The file to ingest.",
			Required:  true,
			BodyPath:  "file",
			FileInput: true,
		},
		&requestflag.Flag[any]{
			Name:     "collection",
			Usage:    "The collection to add the document to — deprecated, set the collection using metadata instead.",
			BodyPath: "collection",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Custom metadata as JSON string for filtering. Keys must be alphanumeric with underscores, max 64 chars. Values must be string, number, or boolean.",
			BodyPath: "metadata",
		},
	},
	Action:          handleMemoriesUpload,
	HideHelpCommand: true,
}

func handleMemoriesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("resource-id") && len(unusedArgs) > 0 {
		cmd.Set("resource-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := hyperspell.MemoryUpdateParams{
		Source: hyperspell.MemoryUpdateParamsSource(cmd.Value("source").(string)),
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Memories.Update(
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
		Title:          "memories update",
		Transform:      transform,
	})
}

func handleMemoriesList(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := hyperspell.MemoryListParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Memories.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "memories list",
			Transform:      transform,
		})
	} else {
		iter := client.Memories.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "memories list",
			Transform:      transform,
		})
	}
}

func handleMemoriesDelete(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("resource-id") && len(unusedArgs) > 0 {
		cmd.Set("resource-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := hyperspell.MemoryDeleteParams{
		Source: hyperspell.MemoryDeleteParamsSource(cmd.Value("source").(string)),
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
	_, err = client.Memories.Delete(
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
		Title:          "memories delete",
		Transform:      transform,
	})
}

func handleMemoriesAdd(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := hyperspell.MemoryAddParams{}

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
	_, err = client.Memories.Add(ctx, params, options...)
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
		Title:          "memories add",
		Transform:      transform,
	})
}

func handleMemoriesAddBulk(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := hyperspell.MemoryAddBulkParams{}

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
	_, err = client.Memories.AddBulk(ctx, params, options...)
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
		Title:          "memories add-bulk",
		Transform:      transform,
	})
}

func handleMemoriesGet(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("resource-id") && len(unusedArgs) > 0 {
		cmd.Set("resource-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := hyperspell.MemoryGetParams{
		Source: hyperspell.MemoryGetParamsSource(cmd.Value("source").(string)),
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
	_, err = client.Memories.Get(
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
		Title:          "memories get",
		Transform:      transform,
	})
}

func handleMemoriesSearch(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := hyperspell.MemorySearchParams{}

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
	_, err = client.Memories.Search(ctx, params, options...)
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
		Title:          "memories search",
		Transform:      transform,
	})
}

func handleMemoriesStatus(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Memories.Status(ctx, options...)
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
		Title:          "memories status",
		Transform:      transform,
	})
}

func handleMemoriesUpload(ctx context.Context, cmd *cli.Command) error {
	client := hyperspell.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := hyperspell.MemoryUploadParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		MultipartFormEncoded,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Memories.Upload(ctx, params, options...)
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
		Title:          "memories upload",
		Transform:      transform,
	})
}
