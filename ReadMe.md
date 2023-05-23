- main.go                // 应用程序的入口文件
- server/                // HTTP服务器相关代码
  - handler.go           // 路由处理程序
  - webhook.go           // GitHub Webhook事件处理逻辑
- docker/                // Docker相关代码
  - builder.go           // Docker镜像构建逻辑
  - executor.go          // Docker镜像执行逻辑
- utils/                 // 工具函数或公共代码
  - signature.go         // 签名验证相关函数
- config/                // 配置文件或配置相关代码
  - config.go            // 应用程序配置结构体及加载函数
  - config.json          // 配置文件（可选）


这是一个基本的目录结构示例，你可以根据你的项目需求进行调整和扩展。下面是对各个目录的功能说明：

main.go：作为应用程序的入口文件，初始化服务器和其他必要的组件。
server/：包含与HTTP服务器相关的代码，例如路由处理程序和GitHub Webhook事件处理逻辑。
docker/：包含与Docker相关的代码，例如Docker镜像构建逻辑和Docker镜像执行逻辑。
utils/：包含通用的工具函数或公共代码，例如签名验证相关的函数。
config/：包含应用程序的配置文件或配置相关的代码，例如加载应用程序配置的函数和配置文件。
