package git

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func HeadTag(repoPath string) (string, error) {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return "", fmt.Errorf("unable to open repo: %w", err)
	}
	ref, err := r.Head()
	if err != nil {
		return "", fmt.Errorf("unable fetch head: %w", err)
	}
	tagRefs, _ := r.Tags()
	var tagName string

	tagRefs.ForEach(func(t *plumbing.Reference) error {
		if t.Hash().String() == ref.Hash().String() {
			tagName = t.Name().Short()
			return fmt.Errorf("found")
		}
		return nil
	})

	return tagName, nil
}