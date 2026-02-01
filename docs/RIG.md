# Rig Beads

Rig beads are the tasks, epics, and issues that belong to a specific project rig. Each rig is a git repository with its own beads database and workflow.

## What is a Rig?

A **Rig** is a project container that lives inside your town. It wraps a git repository and its local agents, and it has its own bead history.

## Rig Beads vs Town Beads

- **Town beads** are Bordertown's own work items, managed in this repo.
- **Rig beads** are work items for a specific rig, managed inside that rig's repo.

The dashboard shows town beads alongside a summary of each rig's beads.

## Managing Rig Beads

Use `bd` from inside the rig directory to create and track rig beads:

```bash
cd rigs/my-rig
bd ready
bd create --title="Feature work" --type=task --priority=2
```

## Rig Registry

Rigs are tracked in the Bordertown registry. Add rigs with:

```bash
bordertown rig add <name> <git-remote>
```
