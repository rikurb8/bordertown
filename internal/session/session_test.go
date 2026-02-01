package session

import (
	"testing"
)

func TestParseTool(t *testing.T) {
	tests := []struct {
		input    string
		expected Tool
	}{
		{"claude", ToolClaude},
		{"opencode", ToolOpencode},
		{"", ToolClaude},
		{"unknown", ToolClaude},
		{"CLAUDE", ToolClaude}, // case sensitive, falls back to default
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ParseTool(tt.input)
			if got != tt.expected {
				t.Errorf("ParseTool(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestBuildClaudeCmd(t *testing.T) {
	tests := []struct {
		name     string
		opts     Options
		wantArgs []string
	}{
		{
			name:     "no options",
			opts:     Options{Tool: ToolClaude},
			wantArgs: []string{},
		},
		{
			name: "with prompt",
			opts: Options{
				Tool:   ToolClaude,
				Prompt: "hello world",
			},
			wantArgs: []string{"--print", "hello world"},
		},
		{
			name: "with system prompt",
			opts: Options{
				Tool:         ToolClaude,
				SystemPrompt: "You are a helpful assistant",
			},
			wantArgs: []string{"--system-prompt", "You are a helpful assistant"},
		},
		{
			name: "with both prompts",
			opts: Options{
				Tool:         ToolClaude,
				SystemPrompt: "system",
				Prompt:       "user",
			},
			wantArgs: []string{"--system-prompt", "system", "--print", "user"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := buildClaudeCmd(tt.opts)
			if cmd.Path == "" && len(cmd.Args) > 0 {
				// exec.Command sets Args[0] to the command name
				t.Log("Command built successfully")
			}
			// Check args (skip Args[0] which is the command name)
			gotArgs := cmd.Args[1:]
			if len(gotArgs) != len(tt.wantArgs) {
				t.Errorf("got %d args, want %d", len(gotArgs), len(tt.wantArgs))
				return
			}
			for i, arg := range gotArgs {
				if arg != tt.wantArgs[i] {
					t.Errorf("arg[%d] = %q, want %q", i, arg, tt.wantArgs[i])
				}
			}
		})
	}
}

func TestBuildOpencodeCmd(t *testing.T) {
	tests := []struct {
		name     string
		opts     Options
		wantArgs []string
	}{
		{
			name:     "no options",
			opts:     Options{Tool: ToolOpencode},
			wantArgs: []string{},
		},
		{
			name: "with prompt",
			opts: Options{
				Tool:   ToolOpencode,
				Prompt: "hello world",
			},
			wantArgs: []string{"-p", "hello world"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := buildOpencodeCmd(tt.opts)
			gotArgs := cmd.Args[1:]
			if len(gotArgs) != len(tt.wantArgs) {
				t.Errorf("got %d args, want %d", len(gotArgs), len(tt.wantArgs))
				return
			}
			for i, arg := range gotArgs {
				if arg != tt.wantArgs[i] {
					t.Errorf("arg[%d] = %q, want %q", i, arg, tt.wantArgs[i])
				}
			}
		})
	}
}
