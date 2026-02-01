package beads

// PlanningStatus indicates whether an epic needs more planning.
type PlanningStatus struct {
	NeedsPlanning bool
	Reasons       []string
}

// AnalyzeEpic determines if an epic needs more planning based on heuristics.
func AnalyzeEpic(group EpicGroup) PlanningStatus {
	status := PlanningStatus{}

	// Only analyze open epics
	if !group.Epic.IsOpen() {
		return status
	}

	// Check: Epic has no child tasks
	if len(group.Children) == 0 {
		status.NeedsPlanning = true
		status.Reasons = append(status.Reasons, "no tasks defined")
	}

	// Check: Epic has very few tasks (1-2)
	if len(group.Children) > 0 && len(group.Children) <= 2 {
		status.NeedsPlanning = true
		status.Reasons = append(status.Reasons, "only has 1-2 tasks")
	}

	// Check: Epic description is vague/short (less than 50 chars)
	if len(group.Epic.Description) < 50 {
		status.NeedsPlanning = true
		status.Reasons = append(status.Reasons, "description is brief")
	}

	// Check: All tasks are done but epic is still open
	if len(group.Children) > 0 {
		allDone := true
		for _, child := range group.Children {
			if child.IsOpen() {
				allDone = false
				break
			}
		}
		if allDone {
			status.NeedsPlanning = true
			status.Reasons = append(status.Reasons, "all tasks complete but epic still open")
		}
	}

	return status
}

// EpicSummary provides a summary of an epic's status.
type EpicSummary struct {
	Epic          Issue
	TotalTasks    int
	OpenTasks     int
	ClosedTasks   int
	Planning      PlanningStatus
}

// SummarizeEpic creates a summary for an epic group.
func SummarizeEpic(group EpicGroup) EpicSummary {
	summary := EpicSummary{
		Epic:       group.Epic,
		TotalTasks: len(group.Children),
		Planning:   AnalyzeEpic(group),
	}

	for _, child := range group.Children {
		if child.IsOpen() {
			summary.OpenTasks++
		} else {
			summary.ClosedTasks++
		}
	}

	return summary
}
