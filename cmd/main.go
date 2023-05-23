package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	r := gin.Default();
	r.POST("/webhook", handleWebhook)
	
	err := r.Run(":8080")

	if err != nil {
		log.Fatal(err)
	}
}

func handleWebhook(c *gin.Context)  {
	payload := struct {
		Ref string `json:"ref"`
		HeadCommit struct {
			ID string `json:"id"`
		}{}
	}{}

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imageName := "golang-webhook:" + payload.HeadCommit.ID
	buildDockerImage(payload.Ref, payload.HeadCommit.ID, imageName)

	runDockerContainer(imageName)

	c.JSON(http.StatusOK, gin.H{"message": "Docker image build and execution complete"})
}

func buildDockerImageV2(ref, commitID, imageName string) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatalf("Failed to create Docker client: %s", err.Error())
	}

	buildOptions := types.ImageBuildOptions{
		Tags:        []string{imageName},
		BuildArgs:   map[string]*string{"REF": &ref, "COMMIT_ID": &commitID},
		ContextPath: ".",
	}

	buildResponse, err := cli.ImageBuild(ctx, nil, buildOptions)
	if err != nil {
		log.Fatalf("Failed to build Docker image: %s", err.Error())
	}

	defer buildResponse.Body.Close()

	// 解析构建日志并进行适当处理
	// ...

	log.Printf("Docker image built successfully: %s", imageName)
}

func buildDockerImage(ref, commitID, imageName string)  {
	// TODO: 实现Docker镜像构建逻辑，可以使用Docker命令行工具或Docker SDK for Go
	log.Printf("Building Docker image for ref %s, commit ID %s, image name %s", ref, commitID, imageName)
	// 构建镜像的具体命令

		// 构建镜像的具体命令
		cmd := exec.Command("docker", "build", "-t", imageName, "--build-arg", "REF="+ref, "--build-arg", "COMMIT_ID="+commitID, ".")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		
		if err != nil {
			log.Printf("Docker image build failed: %s", err)
		}

		log.Printf("Docker image build complete")	
}

func runDockerImage(imageName string) {
	// TODO: 实现Docker镜像执行逻辑，可以使用Docker命令行工具或Docker SDK for Go
	log.Printf("Running Docker image %s", imageName)
	// 执行镜像的具体命令
}