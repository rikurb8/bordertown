package prompts

import (
	"strings"
	"testing"
)

func TestEpicPlanningPromptContainsKeyElements(t *testing.T) {
	// Verify the prompt contains essential guidance
	requiredPhrases := []string{
		"epic",
		"task",
		"bd create",
		"bd dep add",
		"priority",
		"dependencies",
		"P0", "P1", "P2", "P3",
	}

	for _, phrase := range requiredPhrases {
		if !strings.Contains(EpicPlanningPrompt, phrase) {
			t.Errorf("EpicPlanningPrompt should contain %q", phrase)
		}
	}
}

func TestEpicPlanningInitialPrompt(t *testing.T) {
	tests := []struct {
		name    string
		title   string
		wantSub string
	}{
		{
			name:    "empty title",
			title:   "",
			wantSub: "plan a new epic",
		},
		{
			name:    "with title",
			title:   "User authentication",
			wantSub: "User authentication",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EpicPlanningInitialPrompt(tt.title)
			if !strings.Contains(got, tt.wantSub) {
				t.Errorf("EpicPlanningInitialPrompt(%q) = %q, want to contain %q",
					tt.title, got, tt.wantSub)
			}
		})
	}
}
