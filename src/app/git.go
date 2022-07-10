package app

import (
	"bufio"
	"errors"
	"os"
	"path"
	"strings"

	"github.com/cubismod/fly-image-updater/src/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
)

func parseDockerfile(dir string) ([]string, error) {
	dockerfile := path.Join(dir, "Dockerfile")

	content, err := os.Open(dockerfile)
	
	if err != nil {
		return nil, err
	}

	// read Dockerfile to find first FROM line
	scanner := bufio.NewScanner(content)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "FROM") {
			split := strings.Split(text, " ")
			if len(split) == 2 {
				// expecting this for example
				// FROM node:18.4.0-alpine
				return split, nil
			}
		}
	}
	return nil, errors.New("your Dockerfile is missing the FROM line or is malformed")
}

// Clones a Git repository into a temporary
// directory which is removed after usage
// with Fly
func RepoEnv(config config.Config, app config.App) {
	dir, err := os.MkdirTemp(os.TempDir(), "fly-git")

	if err != nil {
		logrus.Error(err)
		return
	}

	defer os.RemoveAll(dir)

	_, err2 := git.PlainClone(dir, false, &git.CloneOptions{
		URL: app.GitRepo,
	})

	if err2 != nil {
		logrus.Error(err)
		return
	}

	parseDockerfile(dir)

	// re
}
