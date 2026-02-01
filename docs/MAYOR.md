# Mayor Command

The `mayor` command provides project oversight and planning tools for your Bordertown workspace.

## Quick Start

```bash
# Review epics and planning status
bt mayor review

# Include closed epics and issues
bt mayor review --all
```

## Commands

### `mayor review`

Analyzes your beads issues grouped by epic and indicates which epics need more planning.

**Output includes:**

- All open epics with their child tasks
- Task counts (open/closed) per epic
- Planning warnings for epics that need attention
- Orphan issues not linked to any epic

**Flags:**

- `--all` - Include closed epics and issues in the output

**Example output:**

```
Mayor Review

○ bt-abc - Feature Implementation
  Tasks: 2 open, 1 closed
    ○ bt-def - Implement core logic
    ○ bt-ghi - Add tests
    ● bt-jkl - Design API

○ bt-xyz - Another Epic
  Tasks: 0 open, 0 closed
  ⚠ Needs planning: no tasks defined, description is brief

Orphan Issues (no epic)
  ○ bt-123 [task] - Standalone task
```

## How Epic Grouping Works

Issues are grouped by epic using dependencies:

1. An issue belongs to an epic if **the epic depends on it** (the issue blocks the epic)
2. Issues not linked to any epic appear in the "Orphan Issues" section

To link a task to an epic:
```bash
bd dep add <epic-id> <task-id>
```

## Planning Indicators

The `review` command flags epics that may need more planning:

| Warning | Meaning |
|---------|---------|
| "no tasks defined" | Epic has no linked tasks |
| "only has 1-2 tasks" | Epic may be under-planned |
| "description is brief" | Epic description is less than 50 characters |
| "all tasks complete but epic still open" | Consider closing the epic |

## Symbols

- `○` Open issue
- `●` Closed issue
- `⚠` Warning/needs attention
