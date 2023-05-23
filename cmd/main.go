package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
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

func buildDockerImage(ref, commitID, imageName string)  {
	// TODO: 实现Docker镜像构建逻辑，可以使用Docker命令行工具或Docker SDK for Go
	log.Printf("Building Docker image for ref %s, commit ID %s, image name %s", ref, commitID, imageName)
	// 构建镜像的具体命令
	
}

func runDockerImage(imageName string) {
	// TODO: 实现Docker镜像执行逻辑，可以使用Docker命令行工具或Docker SDK for Go
	log.Printf("Running Docker image %s", imageName)
	// 执行镜像的具体命令
}