package beads

import (
	"testing"
)

func TestGroupByEpic(t *testing.T) {
	issues := []Issue{
		{ID: "epic-1", Title: "Epic 1", IssueType: "epic", Status: "open",
			Dependencies: []Dependency{
				{IssueID: "epic-1", DependsOnID: "task-1", Type: "blocks"},
				{IssueID: "epic-1", DependsOnID: "task-2", Type: "blocks"},
			}},
		{ID: "task-1", Title: "Task 1", IssueType: "task", Status: "open"},
		{ID: "task-2", Title: "Task 2", IssueType: "task", Status: "closed"},
		{ID: "orphan-1", Title: "Orphan", IssueType: "task", Status: "open"},
	}

	grouped := GroupByEpic(issues)

	// Check epics
	if len(grouped.Epics) != 1 {
		t.Fatalf("expected 1 epic, got %d", len(grouped.Epics))
	}
	if grouped.Epics[0].Epic.ID != "epic-1" {
		t.Errorf("expected epic-1, got %s", grouped.Epics[0].Epic.ID)
	}
	if len(grouped.Epics[0].Children) != 2 {
		t.Errorf("expected 2 children, got %d", len(grouped.Epics[0].Children))
	}

	// Check orphans
	if len(grouped.Orphans) != 1 {
		t.Fatalf("expected 1 orphan, got %d", len(grouped.Orphans))
	}
	if grouped.Orphans[0].ID != "orphan-1" {
		t.Errorf("expected orphan-1, got %s", grouped.Orphans[0].ID)
	}
}

func TestGroupByEpic_MultipleEpics(t *testing.T) {
	issues := []Issue{
		{ID: "epic-a", IssueType: "epic", Status: "open",
			Dependencies: []Dependency{
				{IssueID: "epic-a", DependsOnID: "task-a1", Type: "blocks"},
			}},
		{ID: "epic-b", IssueType: "epic", Status: "closed",
			Dependencies: []Dependency{
				{IssueID: "epic-b", DependsOnID: "task-b1", Type: "blocks"},
			}},
		{ID: "task-a1", IssueType: "task", Status: "open"},
		{ID: "task-b1", IssueType: "task", Status: "closed"},
	}

	grouped := GroupByEpic(issues)

	if len(grouped.Epics) != 2 {
		t.Fatalf("expected 2 epics, got %d", len(grouped.Epics))
	}

	openEpics := grouped.OpenEpics()
	if len(openEpics) != 1 {
		t.Errorf("expected 1 open epic, got %d", len(openEpics))
	}
}

func TestGroupByEpic_NoEpics(t *testing.T) {
	issues := []Issue{
		{ID: "task-1", IssueType: "task", Status: "open"},
		{ID: "task-2", IssueType: "task", Status: "open"},
	}

	grouped := GroupByEpic(issues)

	if len(grouped.Epics) != 0 {
		t.Errorf("expected 0 epics, got %d", len(grouped.Epics))
	}
	if len(grouped.Orphans) != 2 {
		t.Errorf("expected 2 orphans, got %d", len(grouped.Orphans))
	}
}
