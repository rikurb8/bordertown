package beads

// EpicGroup represents an epic with its child issues.
type EpicGroup struct {
	Epic     Issue
	Children []Issue
}

// GroupedIssues contains issues organized by epic.
type GroupedIssues struct {
	Epics   []EpicGroup
	Orphans []Issue // Issues not associated with any epic
}

// GroupByEpic organizes issues by their parent epic.
// An issue belongs to an epic if the epic depends on it (the issue blocks the epic).
func GroupByEpic(issues []Issue) GroupedIssues {
	result := GroupedIssues{}

	// Build lookup maps
	issueByID := make(map[string]Issue)
	for _, issue := range issues {
		issueByID[issue.ID] = issue
	}

	// Find which issues block which epics
	// If epic.dependencies contains {depends_on_id: X}, then X blocks the epic
	epicChildren := make(map[string][]string) // epicID -> []childID
	childToEpic := make(map[string]string)    // childID -> epicID

	var epics []Issue
	for _, issue := range issues {
		if issue.IsEpic() {
			epics = append(epics, issue)
			for _, dep := range issue.Dependencies {
				if dep.Type == "blocks" {
					// This epic depends on dep.DependsOnID
					childID := dep.DependsOnID
					epicChildren[issue.ID] = append(epicChildren[issue.ID], childID)
					childToEpic[childID] = issue.ID
				}
			}
		}
	}

	// Build epic groups
	for _, epic := range epics {
		group := EpicGroup{Epic: epic}
		for _, childID := range epicChildren[epic.ID] {
			if child, ok := issueByID[childID]; ok {
				group.Children = append(group.Children, child)
			}
		}
		result.Epics = append(result.Epics, group)
	}

	// Find orphans (non-epic issues not linked to any epic)
	for _, issue := range issues {
		if !issue.IsEpic() {
			if _, hasParent := childToEpic[issue.ID]; !hasParent {
				result.Orphans = append(result.Orphans, issue)
			}
		}
	}

	return result
}

// OpenEpics returns only epic groups where the epic is still open.
func (g GroupedIssues) OpenEpics() []EpicGroup {
	var open []EpicGroup
	for _, eg := range g.Epics {
		if eg.Epic.IsOpen() {
			open = append(open, eg)
		}
	}
	return open
}

// OpenOrphans returns only orphan issues that are still open.
func (g GroupedIssues) OpenOrphans() []Issue {
	var open []Issue
	for _, issue := range g.Orphans {
		if issue.IsOpen() {
			open = append(open, issue)
		}
	}
	return open
}
