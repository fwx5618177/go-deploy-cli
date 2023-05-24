package readConfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	RepoURL         string `json:"REPO_URL"`
	ContainerName   string `json:"CONTAINER_NAME"`
	DockerImageName string `json:"DOCKER_IMAGE_NAME"`
	DockerImageTag  string `json:"DOCKER_IMAGE_TAG"`
	RemoteDir       string `json:"REMOTE_DIR"`
	GithubToken     string `json:"GITHUB_TOKEN"`
}

func ReadConfigFile(filepath string) (Config, error) {
	var config Config

	file, err := ioutil.ReadFile(filepath)

	if err != nil {
		return config, fmt.Errorf("Failed to read config file: %v", err)
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		return config, fmt.Errorf("Failed to parse config file: %v", err)
	}

	return config, nil
}
