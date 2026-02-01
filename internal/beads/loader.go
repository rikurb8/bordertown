package beads

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	BeadsDir   = ".beads"
	IssuesFile = "issues.jsonl"
)

// Issue represents a beads issue/task.
type Issue struct {
	ID           string       `json:"id"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Status       string       `json:"status"`
	Priority     int          `json:"priority"`
	IssueType    string       `json:"issue_type"`
	Owner        string       `json:"owner"`
	CreatedAt    time.Time    `json:"created_at"`
	CreatedBy    string       `json:"created_by"`
	UpdatedAt    time.Time    `json:"updated_at"`
	ClosedAt     *time.Time   `json:"closed_at,omitempty"`
	CloseReason  string       `json:"close_reason,omitempty"`
	Dependencies []Dependency `json:"dependencies,omitempty"`
}

// Dependency represents a dependency relationship between issues.
type Dependency struct {
	IssueID     string    `json:"issue_id"`
	DependsOnID string    `json:"depends_on_id"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
}

// IsOpen returns true if the issue is not closed.
func (i *Issue) IsOpen() bool {
	return i.Status != "closed"
}

// IsEpic returns true if the issue is an epic.
func (i *Issue) IsEpic() bool {
	return i.IssueType == "epic"
}

// LoadIssues reads all issues from the .beads/issues.jsonl file.
func LoadIssues(rootDir string) ([]Issue, error) {
	path := filepath.Join(rootDir, BeadsDir, IssuesFile)
	return LoadIssuesFromFile(path)
}

// LoadIssuesFromFile reads issues from a specific JSONL file.
func LoadIssuesFromFile(path string) ([]Issue, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open issues file: %w", err)
	}
	defer file.Close()

	var issues []Issue
	scanner := bufio.NewScanner(file)

	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if line == "" {
			continue
		}

		var issue Issue
		if err := json.Unmarshal([]byte(line), &issue); err != nil {
			return nil, fmt.Errorf("parse issue at line %d: %w", lineNum, err)
		}
		issues = append(issues, issue)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read issues file: %w", err)
	}

	return issues, nil
}

// FindBeadsRoot walks up the directory tree to find a .beads directory.
func FindBeadsRoot(startDir string) (string, error) {
	dir := startDir
	for {
		beadsPath := filepath.Join(dir, BeadsDir)
		if info, err := os.Stat(beadsPath); err == nil && info.IsDir() {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("no %s directory found", BeadsDir)
		}
		dir = parent
	}
}
