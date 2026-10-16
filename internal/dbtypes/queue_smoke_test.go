package dbtypes

import "testing"

// Deliberate failure: merge-queue smoke test for #600 (PR-level demonstration
// that a failing required check keeps a PR out of the merge queue). Closed,
// never merged.
func TestMergeQueueSmokeFails(t *testing.T) {
	t.Fatal("deliberate failure: this PR must never be queued or merged")
}
