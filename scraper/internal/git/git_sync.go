package git

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
)

// GitSync handles syncing job data to a git repo
// Example usage:
//
//	gs := NewGitSync("/path/to/repo_data", "company_slug")
//	err := gs.SyncJobs(jobs)
type GitSync struct {
	RepoPath      string
	Company       string
	DataSubdir    string // e.g. "jobs" or "companies" if needed
	CommitMessage string
}

func NewGitSync(repoPath, company, dataSubdir string) *GitSync {
	return &GitSync{
		RepoPath:      repoPath,
		Company:       company,
		DataSubdir:    dataSubdir,
		CommitMessage: fmt.Sprintf("Update job data for %s on %s", company, time.Now().Format("2006-01-02 15:04:05")),
	}
}

// SyncJobs writes jobs to the repo and commits/pushes
func (gs *GitSync) SyncJobs(jobs []types.ScrapedJob) error {
	log.Printf("[GitSync] Starting sync for %s with %d jobs", gs.Company, len(jobs))

	if err := gs.ensureRepo(); err != nil {
		return fmt.Errorf("ensure repo failed: %w", err)
	}

	if err := gs.pullLatest(); err != nil {
		return fmt.Errorf("pull latest failed: %w", err)
	}

	if err := gs.writeJobsFile(jobs); err != nil {
		return fmt.Errorf("write jobs file failed: %w", err)
	}

	if err := gs.gitAddCommitPush(); err != nil {
		return fmt.Errorf("git add/commit/push failed: %w", err)
	}

	log.Printf("[GitSync] Successfully completed sync for %s", gs.Company)
	return nil
}

func (gs *GitSync) ensureRepo() error {
	if _, err := os.Stat(gs.RepoPath); os.IsNotExist(err) {
		return fmt.Errorf("repo path does not exist: %s", gs.RepoPath)
	}
	log.Printf("[GitSync] Repo path verified: %s", gs.RepoPath)
	return nil
}

func (gs *GitSync) pullLatest() error {
	log.Printf("[GitSync] Pulling latest changes from remote")
	cmd := exec.Command("git", "pull")
	cmd.Dir = gs.RepoPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git pull failed: %w", err)
	}
	log.Printf("[GitSync] Successfully pulled latest changes")
	return nil
}

func (gs *GitSync) writeJobsFile(jobs []types.ScrapedJob) error {
	// Write to e.g. <repo>/api/companies/company.json or <repo>/api/jobs.json for consolidated
	jobsDir := filepath.Join(gs.RepoPath, gs.DataSubdir)
	if err := os.MkdirAll(jobsDir, 0755); err != nil {
		return fmt.Errorf("failed to create jobs directory: %w", err)
	}

	var filePath string
	if gs.Company == "consolidated" {
		// For consolidated jobs, write directly as jobs.json
		filePath = filepath.Join(jobsDir, "jobs.json")
	} else {
		// For individual companies, write as company.json
		filePath = filepath.Join(jobsDir, gs.Company+".json")
	}

	data, err := json.MarshalIndent(jobs, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal jobs to JSON: %w", err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	log.Printf("[GitSync] Successfully wrote %d jobs to %s", len(jobs), filePath)
	return nil
}

func (gs *GitSync) gitAddCommitPush() error {
	log.Printf("[GitSync] Starting git operations")

	// Check if there are any changes to commit
	cmdStatus := exec.Command("git", "status", "--porcelain")
	cmdStatus.Dir = gs.RepoPath
	output, err := cmdStatus.Output()
	if err != nil {
		return fmt.Errorf("git status failed: %w", err)
	}

	if len(output) == 0 {
		log.Printf("[GitSync] No changes to commit, skipping git operations")
		return nil
	}

	log.Printf("[GitSync] Changes detected: %s", string(output))

	// Add all changes
	cmdAdd := exec.Command("git", "add", ".")
	cmdAdd.Dir = gs.RepoPath
	if err := cmdAdd.Run(); err != nil {
		return fmt.Errorf("git add failed: %w", err)
	}
	log.Printf("[GitSync] Successfully added changes")

	// Commit changes
	cmdCommit := exec.Command("git", "commit", "-m", gs.CommitMessage)
	cmdCommit.Dir = gs.RepoPath
	if err := cmdCommit.Run(); err != nil {
		return fmt.Errorf("git commit failed: %w", err)
	}
	log.Printf("[GitSync] Successfully committed with message: %s", gs.CommitMessage)

	// Push changes
	cmdPush := exec.Command("git", "push")
	cmdPush.Dir = gs.RepoPath
	if err := cmdPush.Run(); err != nil {
		return fmt.Errorf("git push failed: %w", err)
	}
	log.Printf("[GitSync] Successfully pushed to remote")

	return nil
}
