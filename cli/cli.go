package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"

	readConfig "goDeploy/utils"
)

func main() {
	config, err := readConfig.ReadConfigFile("config.json")

	if err != nil {
		log.Fatalf("Failed to read config file: %s", err.Error())
	}

	repoURL := config.RepoURL
	containerName := config.ContainerName
	dockerImageName := config.DockerImageName
	dockerImageTag := config.DockerImageTag
	remoteDir := config.RemoteDir
	githubToken := config.GithubToken

	// 拉取代码
	err = cloneRepo(repoURL, remoteDir, githubToken)
	if err != nil {
		log.Fatalf("Failed to clone repository: %s", err.Error())
	}

	// 停止和删除容器
	stopContainer(containerName)
	removeContainer(containerName)

	// 删除旧的Docker镜像
	removeDockerImage(dockerImageName, dockerImageTag)

	// 构建新的Docker镜像
	buildDockerImage(remoteDir, dockerImageName, dockerImageTag)

	// 删除临时文件
	// removeTempFiles(remoteDir, tarFile)

	log.Println("CLI program execution complete")
}

// 从GitHub上克隆代码
func cloneRepo(repoURL, remoteDir, githubToken string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	_, _, err := client.Repositories.DownloadContents(ctx, "", "", remoteDir, &github.RepositoryContentGetOptions{
		Ref: "master",
	})

	if err != nil {
		return fmt.Errorf("failed to clone repository: %s", err.Error())
	}

	return nil
}

// 停止容器
func stopContainer(containerName string) {
	cmd := exec.Command("docker", "stop", containerName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to stop container: %s", err.Error())
	}
}

// 删除容器
func removeContainer(containerName string) {
	cmd := exec.Command("docker", "rm", containerName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to remove container: %s", err.Error())
	}
}

// 删除Docker镜像
func removeDockerImage(dockerImageName, dockerImageTag string) {
	cmd := exec.Command("docker", "rmi", "-f", fmt.Sprintf("%s:%s", dockerImageName, dockerImageTag))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to remove Docker image: %s", err.Error())
	}
}

// 构建Docker镜像
func buildDockerImage(remoteDir, dockerImageName, dockerImageTag string) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatalf("Failed to create Docker client: %s", err.Error())
	}

	buildOptions := types.ImageBuildOptions{
		Tags:          []string{fmt.Sprintf("%s:%s", dockerImageName, dockerImageTag)},
		RemoteContext: remoteDir,
	}

	buildResponse, err := cli.ImageBuild(ctx, nil, buildOptions)
	if err != nil {
		log.Fatalf("Failed to build Docker image: %s", err.Error())
	}

	defer buildResponse.Body.Close()

	log.Printf("Docker image built successfully: %s:%s", dockerImageName, dockerImageTag)
}

// 删除临时文件
func removeTempFiles(remoteDir, tarFile string) {
	os.Remove(fmt.Sprintf("%s/%s", remoteDir, tarFile))
	os.Remove(fmt.Sprintf("%s/temp_key", remoteDir))
}
