package beads

import (
	"testing"
)

func TestAnalyzeEpic_NoTasks(t *testing.T) {
	group := EpicGroup{
		Epic:     Issue{ID: "epic-1", Status: "open", IssueType: "epic", Description: "A very detailed description of what this epic is about and its goals"},
		Children: nil,
	}

	status := AnalyzeEpic(group)

	if !status.NeedsPlanning {
		t.Error("expected needs planning for epic with no tasks")
	}
	if len(status.Reasons) == 0 {
		t.Error("expected at least one reason")
	}
}

func TestAnalyzeEpic_FewTasks(t *testing.T) {
	group := EpicGroup{
		Epic: Issue{ID: "epic-1", Status: "open", IssueType: "epic", Description: "A very detailed description of what this epic is about and its goals"},
		Children: []Issue{
			{ID: "task-1", Status: "open"},
		},
	}

	status := AnalyzeEpic(group)

	if !status.NeedsPlanning {
		t.Error("expected needs planning for epic with few tasks")
	}
}

func TestAnalyzeEpic_ShortDescription(t *testing.T) {
	group := EpicGroup{
		Epic: Issue{ID: "epic-1", Status: "open", IssueType: "epic", Description: "Short"},
		Children: []Issue{
			{ID: "task-1", Status: "open"},
			{ID: "task-2", Status: "open"},
			{ID: "task-3", Status: "open"},
		},
	}

	status := AnalyzeEpic(group)

	if !status.NeedsPlanning {
		t.Error("expected needs planning for epic with short description")
	}
}

func TestAnalyzeEpic_AllTasksComplete(t *testing.T) {
	group := EpicGroup{
		Epic: Issue{ID: "epic-1", Status: "open", IssueType: "epic", Description: "A very detailed description of what this epic is about and its goals"},
		Children: []Issue{
			{ID: "task-1", Status: "closed"},
			{ID: "task-2", Status: "closed"},
			{ID: "task-3", Status: "closed"},
		},
	}

	status := AnalyzeEpic(group)

	if !status.NeedsPlanning {
		t.Error("expected needs planning when all tasks complete but epic open")
	}
}

func TestAnalyzeEpic_WellPlanned(t *testing.T) {
	group := EpicGroup{
		Epic: Issue{ID: "epic-1", Status: "open", IssueType: "epic", Description: "A very detailed description of what this epic is about and all its requirements and acceptance criteria"},
		Children: []Issue{
			{ID: "task-1", Status: "open"},
			{ID: "task-2", Status: "open"},
			{ID: "task-3", Status: "closed"},
		},
	}

	status := AnalyzeEpic(group)

	if status.NeedsPlanning {
		t.Errorf("expected well-planned epic to not need planning, got reasons: %v", status.Reasons)
	}
}

func TestAnalyzeEpic_ClosedEpic(t *testing.T) {
	group := EpicGroup{
		Epic:     Issue{ID: "epic-1", Status: "closed", IssueType: "epic", Description: "Short"},
		Children: nil,
	}

	status := AnalyzeEpic(group)

	if status.NeedsPlanning {
		t.Error("closed epics should not need planning")
	}
}

func TestSummarizeEpic(t *testing.T) {
	group := EpicGroup{
		Epic: Issue{ID: "epic-1", Status: "open", IssueType: "epic", Description: "Detailed description here with enough content to pass the length check"},
		Children: []Issue{
			{ID: "task-1", Status: "open"},
			{ID: "task-2", Status: "closed"},
			{ID: "task-3", Status: "closed"},
		},
	}

	summary := SummarizeEpic(group)

	if summary.TotalTasks != 3 {
		t.Errorf("expected 3 total tasks, got %d", summary.TotalTasks)
	}
	if summary.OpenTasks != 1 {
		t.Errorf("expected 1 open task, got %d", summary.OpenTasks)
	}
	if summary.ClosedTasks != 2 {
		t.Errorf("expected 2 closed tasks, got %d", summary.ClosedTasks)
	}
}
