package git

import (
	"errors"
	"os/exec"
	"strings"
)

// GetCommitsFrom returns a list of commit hashes
// from the base commit to the current HEAD.
func GetCommitsFrom(hash string) ([]string, error) {
	cmd := exec.Command(
		"git",
		"rev-list",
		hash+"..HEAD",
	)

	bb, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return strings.Split(
		strings.TrimSpace(string(bb)),
		"\n",
	), nil
}

// GetCommitMessage retrieves the commit message for a given commit hash.
func GetCommitMessage(commitHash string) (string, error) {
	cmd := exec.Command(
		"git",
		"log",
		"-1",
		"--pretty=%B",
		commitHash,
	)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

// GetBaseCommitFromBranch returns the base commit hash
// from the specified branch to the current branch.
func GetBaseCommitFromBranch(from string) (string, error) {
	if from == "" {
		return "", errors.New("branch name cannot be empty")
	}

	to, err := getCurrentBranchName()
	if err != nil {
		return "", err
	}

	cmd := exec.Command(
		"git",
		"merge-base",
		from,
		to,
	)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

// getCurrentBranch returns the name of the current git branch.
func getCurrentBranchName() (string, error) {
	cmd := exec.Command(
		"git",
		"rev-parse",
		"--abbrev-ref",
		"HEAD",
	)

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
