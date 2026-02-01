package prompts

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadEpicPlanningPrompt_Builtin(t *testing.T) {
	// Non-existent directory should fall back to built-in
	result := LoadEpicPlanningPrompt("/nonexistent", "")
	if result != EpicPlanningPrompt {
		t.Error("expected built-in prompt when no custom file exists")
	}
}

func TestLoadEpicPlanningPrompt_ConfiguredPath(t *testing.T) {
	tmpDir := t.TempDir()
	customPrompt := "Custom epic planning prompt content"

	// Write custom prompt to configured path
	customPath := filepath.Join(tmpDir, "my-prompt.md")
	if err := os.WriteFile(customPath, []byte(customPrompt), 0644); err != nil {
		t.Fatal(err)
	}

	result := LoadEpicPlanningPrompt(tmpDir, "my-prompt.md")
	if result != customPrompt {
		t.Errorf("expected custom prompt, got %q", result)
	}
}

func TestLoadEpicPlanningPrompt_ConfiguredAbsolutePath(t *testing.T) {
	tmpDir := t.TempDir()
	customPrompt := "Absolute path prompt"

	customPath := filepath.Join(tmpDir, "absolute-prompt.md")
	if err := os.WriteFile(customPath, []byte(customPrompt), 0644); err != nil {
		t.Fatal(err)
	}

	// Use absolute path
	result := LoadEpicPlanningPrompt("/other/dir", customPath)
	if result != customPrompt {
		t.Errorf("expected custom prompt from absolute path, got %q", result)
	}
}

func TestLoadEpicPlanningPrompt_DefaultLocation(t *testing.T) {
	tmpDir := t.TempDir()
	customPrompt := "Default location prompt"

	// Create default directory structure
	defaultDir := filepath.Join(tmpDir, DefaultPromptDir)
	if err := os.MkdirAll(defaultDir, 0755); err != nil {
		t.Fatal(err)
	}

	defaultPath := filepath.Join(defaultDir, DefaultEpicPlanningFile)
	if err := os.WriteFile(defaultPath, []byte(customPrompt), 0644); err != nil {
		t.Fatal(err)
	}

	// No configured path, should find default
	result := LoadEpicPlanningPrompt(tmpDir, "")
	if result != customPrompt {
		t.Errorf("expected prompt from default location, got %q", result)
	}
}

func TestLoadEpicPlanningPrompt_ConfiguredPathTakesPrecedence(t *testing.T) {
	tmpDir := t.TempDir()
	configuredPrompt := "Configured prompt"
	defaultPrompt := "Default prompt"

	// Create both configured and default
	configuredPath := filepath.Join(tmpDir, "custom.md")
	if err := os.WriteFile(configuredPath, []byte(configuredPrompt), 0644); err != nil {
		t.Fatal(err)
	}

	defaultDir := filepath.Join(tmpDir, DefaultPromptDir)
	if err := os.MkdirAll(defaultDir, 0755); err != nil {
		t.Fatal(err)
	}
	defaultPath := filepath.Join(defaultDir, DefaultEpicPlanningFile)
	if err := os.WriteFile(defaultPath, []byte(defaultPrompt), 0644); err != nil {
		t.Fatal(err)
	}

	// Configured path should take precedence
	result := LoadEpicPlanningPrompt(tmpDir, "custom.md")
	if result != configuredPrompt {
		t.Errorf("expected configured prompt to take precedence, got %q", result)
	}
}

func TestPromptFileExists(t *testing.T) {
	tmpDir := t.TempDir()

	// No files exist
	if PromptFileExists(tmpDir, "") {
		t.Error("expected false when no files exist")
	}

	// Create default file
	defaultDir := filepath.Join(tmpDir, DefaultPromptDir)
	if err := os.MkdirAll(defaultDir, 0755); err != nil {
		t.Fatal(err)
	}
	defaultPath := filepath.Join(defaultDir, DefaultEpicPlanningFile)
	if err := os.WriteFile(defaultPath, []byte("test"), 0644); err != nil {
		t.Fatal(err)
	}

	if !PromptFileExists(tmpDir, "") {
		t.Error("expected true when default file exists")
	}

	// Create configured file
	configuredPath := filepath.Join(tmpDir, "custom.md")
	if err := os.WriteFile(configuredPath, []byte("test"), 0644); err != nil {
		t.Fatal(err)
	}

	if !PromptFileExists(tmpDir, "custom.md") {
		t.Error("expected true when configured file exists")
	}
}
