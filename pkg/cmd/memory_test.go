// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"strings"
	"testing"

	"github.com/hyperspell/hyperspell-cli/internal/mocktest"
	"github.com/hyperspell/hyperspell-cli/internal/requestflag"
)

func TestMemoriesUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"memories", "update",
			"--source", "reddit",
			"--resource-id", "resource_id",
			"--collection", "string",
			"--date", "'2019-12-27T18:11:19.117Z'",
			"--metadata", "{foo: string}",
			"--text", "string",
			"--title", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"collection: string\n" +
			"date: '2019-12-27T18:11:19.117Z'\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"text: string\n" +
			"title: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--user-id", "string",
			"memories", "update",
			"--source", "reddit",
			"--resource-id", "resource_id",
		)
	})
}

func TestMemoriesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"memories", "list",
			"--max-items", "10",
			"--collection", "collection",
			"--cursor", "cursor",
			"--filter", "filter",
			"--size", "0",
			"--source", "reddit",
			"--status", "pending",
		)
	})
}

func TestMemoriesDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"memories", "delete",
			"--source", "reddit",
			"--resource-id", "resource_id",
		)
	})
}

func TestMemoriesAdd(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"memories", "add",
			"--text", "...",
			"--collection", "my-collection",
			"--date", "'2019-12-27T18:11:19.117Z'",
			"--metadata", "{author: John Doe, date: '2025-05-20T02:31:00Z', rating: 3}",
			"--resource-id", "resource_id",
			"--title", "My Document",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"text: ...\n" +
			"collection: my-collection\n" +
			"date: '2019-12-27T18:11:19.117Z'\n" +
			"metadata:\n" +
			"  author: John Doe\n" +
			"  date: '2025-05-20T02:31:00Z'\n" +
			"  rating: 3\n" +
			"resource_id: resource_id\n" +
			"title: My Document\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--user-id", "string",
			"memories", "add",
		)
	})
}

func TestMemoriesAddBulk(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"memories", "add-bulk",
			"--item", "{text: ..., collection: my-collection, date: '2019-12-27T18:11:19.117Z', metadata: {author: John Doe, date: '2025-05-20T02:31:00Z', rating: 3}, resource_id: resource_id, title: My Document}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(memoriesAddBulk)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"memories", "add-bulk",
			"--item.text", "...",
			"--item.collection", "my-collection",
			"--item.date", "2019-12-27T18:11:19.117Z",
			"--item.metadata", "{author: John Doe, date: '2025-05-20T02:31:00Z', rating: 3}",
			"--item.resource-id", "resource_id",
			"--item.title", "My Document",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"items:\n" +
			"  - text: ...\n" +
			"    collection: my-collection\n" +
			"    date: '2019-12-27T18:11:19.117Z'\n" +
			"    metadata:\n" +
			"      author: John Doe\n" +
			"      date: '2025-05-20T02:31:00Z'\n" +
			"      rating: 3\n" +
			"    resource_id: resource_id\n" +
			"    title: My Document\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--user-id", "string",
			"memories", "add-bulk",
		)
	})
}

func TestMemoriesGet(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"memories", "get",
			"--source", "reddit",
			"--resource-id", "resource_id",
		)
	})
}

func TestMemoriesSearch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"memories", "search",
			"--query", "What does Hyperspell do?",
			"--answer=true",
			"--effort", "0",
			"--max-results", "0",
			"--options", "{after: '2019-12-27T18:11:19.117Z', answer_model: llama-3.1, before: '2019-12-27T18:11:19.117Z', box: {weight: 0}, filter: {}, google_calendar: {calendar_id: calendar_id, weight: 0}, google_drive: {weight: 0}, google_mail: {label_ids: [string], weight: 0}, max_results: 200, memory_types: [procedure], notion: {notion_page_ids: [string], weight: 0}, reddit: {period: hour, sort: relevance, subreddit: subreddit, weight: 0}, resource_ids: [string], slack: {channels: [string], exclude_archived: true, include_dms: true, include_group_dms: true, include_private: true, weight: 0}, vault: {weight: 0}, web_crawler: {max_depth: 0, url: url, weight: 0}}",
			"--source", "vault",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(memoriesSearch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"memories", "search",
			"--query", "What does Hyperspell do?",
			"--answer=true",
			"--effort", "0",
			"--max-results", "0",
			"--options.after", "2019-12-27T18:11:19.117Z",
			"--options.answer-model", "llama-3.1",
			"--options.before", "2019-12-27T18:11:19.117Z",
			"--options.box", "{weight: 0}",
			"--options.filter", "{}",
			"--options.google-calendar", "{calendar_id: calendar_id, weight: 0}",
			"--options.google-drive", "{weight: 0}",
			"--options.google-mail", "{label_ids: [string], weight: 0}",
			"--options.max-results", "200",
			"--options.memory-types", "[procedure]",
			"--options.notion", "{notion_page_ids: [string], weight: 0}",
			"--options.reddit", "{period: hour, sort: relevance, subreddit: subreddit, weight: 0}",
			"--options.resource-ids", "[string]",
			"--options.slack", "{channels: [string], exclude_archived: true, include_dms: true, include_group_dms: true, include_private: true, weight: 0}",
			"--options.vault", "{weight: 0}",
			"--options.web-crawler", "{max_depth: 0, url: url, weight: 0}",
			"--source", "vault",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: What does Hyperspell do?\n" +
			"answer: true\n" +
			"effort: 0\n" +
			"max_results: 0\n" +
			"options:\n" +
			"  after: '2019-12-27T18:11:19.117Z'\n" +
			"  answer_model: llama-3.1\n" +
			"  before: '2019-12-27T18:11:19.117Z'\n" +
			"  box:\n" +
			"    weight: 0\n" +
			"  filter: {}\n" +
			"  google_calendar:\n" +
			"    calendar_id: calendar_id\n" +
			"    weight: 0\n" +
			"  google_drive:\n" +
			"    weight: 0\n" +
			"  google_mail:\n" +
			"    label_ids:\n" +
			"      - string\n" +
			"    weight: 0\n" +
			"  max_results: 200\n" +
			"  memory_types:\n" +
			"    - procedure\n" +
			"  notion:\n" +
			"    notion_page_ids:\n" +
			"      - string\n" +
			"    weight: 0\n" +
			"  reddit:\n" +
			"    period: hour\n" +
			"    sort: relevance\n" +
			"    subreddit: subreddit\n" +
			"    weight: 0\n" +
			"  resource_ids:\n" +
			"    - string\n" +
			"  slack:\n" +
			"    channels:\n" +
			"      - string\n" +
			"    exclude_archived: true\n" +
			"    include_dms: true\n" +
			"    include_group_dms: true\n" +
			"    include_private: true\n" +
			"    weight: 0\n" +
			"  vault:\n" +
			"    weight: 0\n" +
			"  web_crawler:\n" +
			"    max_depth: 0\n" +
			"    url: url\n" +
			"    weight: 0\n" +
			"sources:\n" +
			"  - vault\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--user-id", "string",
			"memories", "search",
		)
	})
}

func TestMemoriesStatus(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"memories", "status",
		)
	})
}

func TestMemoriesUpload(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"memories", "upload",
			"--file", mocktest.TestFile(t, "Example data"),
			"--collection", "collection",
			"--metadata", "metadata",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		testFile := mocktest.TestFile(t, "Example data")
		// Test piping YAML data over stdin
		pipeDataStr := "" +
			"file: Example data\n" +
			"collection: collection\n" +
			"metadata: metadata\n"
		pipeDataStr = strings.ReplaceAll(pipeDataStr, "Example data", testFile)
		pipeData := []byte(pipeDataStr)
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--user-id", "string",
			"memories", "upload",
		)
	})
}
