package main

import (
	"reflect"
	"sort"
	"testing"
)

// // TestMergeContacts tests the mergeContacts function.
func TestMergeContacts(t *testing.T) {
	setA := []Contact{
		{Phone: "2135550100", Email: "", StateID: "", Platform: "CRM 1"},
		{Phone: "4155550199", Email: "test@alpineiq.com", StateID: "", Platform: "POS1"},
		{Phone: "2135550100", Email: "", StateID: "", Platform: "CRM 1"},
	}

	expectedMerged := []Contact{
		{Phone: "2135550100", Email: "", StateID: "", Platform: "CRM 1"},
		{Phone: "4155550199", Email: "test@alpineiq.com", StateID: "", Platform: "POS1"},
	}

	merged := mergeContacts(setA)

	// Sort expected merged contacts by platform for comparison
	sort.Slice(expectedMerged, func(i, j int) bool {
		return expectedMerged[i].Platform < expectedMerged[j].Platform
	})

	// Sort actual merged contacts by platform for comparison
	sort.Slice(merged, func(i, j int) bool {
		return merged[i].Platform < merged[j].Platform
	})

	// Comparison of merged contacts vs. expected
	if !reflect.DeepEqual(merged, expectedMerged) {
		t.Errorf("merged contacts not as expected: %v, Got: %v", expectedMerged, merged)
	}
}
