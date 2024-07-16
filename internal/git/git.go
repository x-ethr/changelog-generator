package git

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os/exec"
	"strings"
)

type Content struct {
	Subject string
	Body    string
}

type Commit struct {
	Hash string `json:"commit-hash"`

	Author struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Date  string `json:"date"`
	} `json:"author"`

	Committer struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Date  string `json:"date"`
	} `json:"committer"`
}

func (c Commit) Content(ctx context.Context) Content {
	var subject, body bytes.Buffer

	{
		cmd := exec.CommandContext(ctx, "git", "show", "--no-merges", "-s", c.Hash, "--format=%s")

		cmd.Stdout = &subject

		if e := cmd.Run(); e != nil {
			panic(e)
		}
	}

	{
		cmd := exec.CommandContext(ctx, "git", "show", "--no-merges", "-s", c.Hash, "--format=%b")

		cmd.Stdout = &body

		if e := cmd.Run(); e != nil {
			panic(e)
		}
	}

	return Content{
		Subject: strings.TrimSpace(subject.String()),
		Body:    strings.TrimSpace(body.String()),
	}
}

func Commits(ctx context.Context) []Commit {
	cmd := exec.CommandContext(ctx, "git", "log", "--no-merges", "--pretty=format:{ \"commit-hash\": \"%H\", \"author\": { \"name\": \"%an\", \"email\": \"%ae\", \"date\": \"%ad\" }, \"committer\": { \"name\": \"%cn\", \"email\": \"%ce\", \"date\": \"%cd\" } }")

	var buffer bytes.Buffer

	cmd.Stdout = &buffer

	if e := cmd.Run(); e != nil {
		panic(e)
	}

	scanner := bufio.NewScanner(&buffer)

	var commits []Commit
	for scanner.Scan() {
		line := scanner.Bytes()

		var commit Commit
		if e := json.Unmarshal(line, &commit); e != nil {
			panic(e)
		}

		commits = append(commits, commit)
	}

	return commits
}

func Tags(ctx context.Context) []string {
	cmd := exec.CommandContext(ctx, "git", "tag", "--list")

	var buffer bytes.Buffer

	cmd.Stdout = &buffer

	if e := cmd.Run(); e != nil {
		panic(e)
	}

	scanner := bufio.NewScanner(&buffer)

	var tags []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		tags = append(tags, line)
	}

	return tags
}

// Verify validated the executing system's git binary.
func Verify(ctx context.Context) error {
	path, e := exec.LookPath("git")
	if e != nil {
		return fmt.Errorf("git not found: %w", e)
	}

	slog.DebugContext(ctx, "Executable Found", slog.Group("git", slog.String("path", path)))

	return nil
}
