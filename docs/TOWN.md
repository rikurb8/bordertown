# Town Command

The `town` command manages your Bordertown workspace configuration.

## Quick Start

```bash
# Initialize a new workspace
bt town init

# Initialize with a custom name
bt town init --name my-project

# Initialize with description
bt town init --name my-project --description "My awesome project"

# Overwrite existing config
bt town init --force
```

## What is a Town?

A **Town** is your workspace directory - the root folder containing your projects (rigs), agents, and configuration. The `town.yml` file stores workspace-level settings.

## town.yml Configuration

Running `bt town init` creates a `town.yml` file:

```yaml
version: 1
name: my-workspace
description: "Optional description"
rigs_dir: ./rigs
mayor:
  model: claude-sonnet
defaults:
  agent_model: claude-haiku
```

### Fields

| Field | Description | Default |
|-------|-------------|---------|
| `version` | Config schema version | `1` |
| `name` | Workspace name | Directory name |
| `description` | What this workspace is for | (empty) |
| `rigs_dir` | Where project repos live | `./rigs` |
| `mayor.model` | Model for mayor commands | `claude-sonnet` |
| `defaults.agent_model` | Default model for agents | `claude-haiku` |

## Commands

### `town init`

Creates a new `town.yml` configuration file.

**Flags:**

- `--name` - Set the workspace name (defaults to current directory name)
- `--description` - Set a description for the workspace
- `--force` - Overwrite existing `town.yml` if it exists

**Examples:**

```bash
# Basic init (uses directory name)
bt town init

# With all options
bt town init --name mytown --description "Development workspace" --force
```
