package session

import (
	"fmt"
	"os"
	"os/exec"
)

// Tool represents the CLI tool to use for sessions.
type Tool string

const (
	ToolClaude   Tool = "claude"
	ToolOpencode Tool = "opencode"
)

// ParseTool converts a string to a Tool, defaulting to Claude.
func ParseTool(s string) Tool {
	switch s {
	case "opencode":
		return ToolOpencode
	case "claude", "":
		return ToolClaude
	default:
		return ToolClaude
	}
}

// Options configures how a session is spawned.
type Options struct {
	Tool         Tool
	Prompt       string // Initial prompt/message to send
	SystemPrompt string // System prompt (for claude)
	WorkDir      string // Working directory
	Interactive  bool   // Attach stdin/stdout for interactive use
}

// Spawn starts a new CLI session with the given options.
// If Interactive is true, the session takes over stdin/stdout.
// Returns when the session exits.
func Spawn(opts Options) error {
	if opts.WorkDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("get working directory: %w", err)
		}
		opts.WorkDir = cwd
	}

	var cmd *exec.Cmd

	switch opts.Tool {
	case ToolOpencode:
		cmd = buildOpencodeCmd(opts)
	case ToolClaude:
		cmd = buildClaudeCmd(opts)
	default:
		return fmt.Errorf("unknown tool: %s", opts.Tool)
	}

	cmd.Dir = opts.WorkDir

	if opts.Interactive {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("session exited: %w", err)
	}

	return nil
}

// buildClaudeCmd constructs the command for Claude Code CLI.
func buildClaudeCmd(opts Options) *exec.Cmd {
	args := []string{}

	// Add system prompt if provided
	if opts.SystemPrompt != "" {
		args = append(args, "--system-prompt", opts.SystemPrompt)
	}

	// Add initial prompt if provided
	if opts.Prompt != "" {
		args = append(args, "--print", opts.Prompt)
	}

	return exec.Command("claude", args...)
}

// buildOpencodeCmd constructs the command for opencode CLI.
func buildOpencodeCmd(opts Options) *exec.Cmd {
	args := []string{}

	// Opencode uses -p for prompt
	if opts.Prompt != "" {
		args = append(args, "-p", opts.Prompt)
	}

	// Note: opencode doesn't have a --system-prompt flag in the same way,
	// but we can include context in the prompt itself

	return exec.Command("opencode", args...)
}

// Available checks if the specified tool is available on the system.
func Available(tool Tool) bool {
	var name string
	switch tool {
	case ToolClaude:
		name = "claude"
	case ToolOpencode:
		name = "opencode"
	default:
		return false
	}

	_, err := exec.LookPath(name)
	return err == nil
}
