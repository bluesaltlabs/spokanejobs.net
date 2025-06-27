package git

import (
	"encoding/json"
	"fmt"
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
		CommitMessage: fmt.Sprintf("Update job data for %s on %s", company, time.Now().Format("yyyy-mm-dd HH:mm:ss")),
	}
}

// SyncJobs writes jobs to the repo and commits/pushes
func (gs *GitSync) SyncJobs(jobs []types.ScrapedJob) error {
	if err := gs.ensureRepo(); err != nil {
		return err
	}
	if err := gs.pullLatest(); err != nil {
		return err
	}
	if err := gs.writeJobsFile(jobs); err != nil {
		return err
	}
	if err := gs.gitAddCommitPush(); err != nil {
		return err
	}
	return nil
}

func (gs *GitSync) ensureRepo() error {
	if _, err := os.Stat(gs.RepoPath); os.IsNotExist(err) {
		return fmt.Errorf("repo path does not exist: %s", gs.RepoPath)
	}
	return nil
}

func (gs *GitSync) pullLatest() error {
	cmd := exec.Command("git", "pull")
	cmd.Dir = gs.RepoPath
	return cmd.Run()
}

func (gs *GitSync) writeJobsFile(jobs []types.ScrapedJob) error {
	// Write to e.g. <repo>/api/companies/company.json or <repo>/api/jobs.json for consolidated
	jobsDir := filepath.Join(gs.RepoPath, gs.DataSubdir)
	if err := os.MkdirAll(jobsDir, 0755); err != nil {
		return err
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
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func (gs *GitSync) gitAddCommitPush() error {
	cmdAdd := exec.Command("git", "add", ".")
	cmdAdd.Dir = gs.RepoPath
	if err := cmdAdd.Run(); err != nil {
		return err
	}
	cmdCommit := exec.Command("git", "commit", "-m", gs.CommitMessage)
	cmdCommit.Dir = gs.RepoPath
	_ = cmdCommit.Run() // ignore error if nothing to commit
	cmdPush := exec.Command("git", "push")
	cmdPush.Dir = gs.RepoPath
	return cmdPush.Run()
}
