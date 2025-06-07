package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the configuration for linting commits.
type Config struct {
	// Branch - rules for git branches.
	Branch Branch `yaml:"branch"`
	// Type - rules for commit types.
	Type Type `yaml:"type"`
	// Scope - rules for commit scopes.
	Scope Scope `yaml:"scope"`
	// Subject - rules for commit subjects.
	Subject Subject `yaml:"subject"`
	// Body - rules for commit bodies.
	Body Body `yaml:"body"`
	// Task - rules for commit tasks.
	Task Task `yaml:"task"`
	// Forbidden - rules to forbid some information in commits.
	Forbidden Forbidden `yaml:"forbidden"`
}

// Branch - rules for git branches.
type Branch struct {
	// Default - the name of the default branch.
	Default string `yaml:"default"`
}

// Type - rules for commit types.
type Type struct {
	// List - a list of allowed commit types.
	List []string `yaml:"list"`
	// Required - whether the type is required in commit messages.
	Required bool `yaml:"required"`
}

// Scope - rules for commit scopes.
type Scope struct {
	// Pattern - a regex pattern that scopes must match.
	Pattern string `yaml:"pattern"`
	// Required - whether the scope is required in commit messages.
	Required bool `yaml:"required"`
}

// Subject - rules for commit subjects.
type Subject struct {
	// MinLength - minimum length of the subject.
	MinLength int `yaml:"min_length"`
	// MaxLength - maximum length of the subject.
	MaxLength int `yaml:"max_length"`
}

// Body - rules for commit bodies.
type Body struct {
	// RequiredForTypes - a list of commit types that require a body.
	RequiredForTypes []string `yaml:"required_for_types"`
	// MinLength - minimum length of the body.
	MinLength int `yaml:"min_length"`
	// MaxLength - maximum length of the body.
	MaxLength int `yaml:"max_length"`
	// RequiredForBreakingChange - whether a body is required for breaking changes.
	RequiredForBreakingChange bool `yaml:"required_for_breaking_change"`
	// RequireBlankLine - whether a blank line is required before the body.
	RequireBlankLine bool `yaml:"require_blank_line"`
}

// Task - rules for commit tasks.
type Task struct {
	// Location - a list of required locations where tasks must be found in commit messages.
	Location []string `yaml:"location"`
	// Pattern - a regex pattern that tasks must match.
	Pattern string `yaml:"pattern"`
	// BranchPattern - a regex pattern for branch names that contain tasks.
	BranchPattern string `yaml:"branch_pattern"`
	// Required - whether tasks are required in commit messages.
	Required bool `yaml:"required"`
}

// Forbidden - rules to forbid some information in commits.
type Forbidden struct {
	// Words - a list of words that are forbidden in commit messages.
	Words []string `yaml:"words"`
}

// LoadFrom loads the configuration from a YAML file at the specified path.
func LoadFrom(path string) (Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
