package main

import (
	"bufio"
	filedata "create-project/fileData"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter project name: ")
	projectName, _ := reader.ReadString('\n')
	projectName = strings.TrimSpace(projectName)

	fmt.Print("Enter project path (leave empty for current directory): ")
	projectPath, _ := reader.ReadString('\n')
	projectPath = strings.TrimSpace(projectPath)

	if projectPath == "" {
		projectPath = "./" + projectName
	} else {
		projectPath = projectPath + "/" + projectName
	}

	fmt.Println("Creating project in directory:", projectPath)

	os.MkdirAll(projectPath+"/controllers", 0755)
	os.MkdirAll(projectPath+"/models", 0755)
	os.MkdirAll(projectPath+"/middlewares", 0755)
	os.MkdirAll(projectPath+"/helpers", 0755)
	os.MkdirAll(projectPath+"/utils", 0755)
	os.MkdirAll(projectPath+"/nginx", 0755)
	os.MkdirAll(projectPath+"/routes", 0755)
	os.MkdirAll(projectPath+"/database", 0755)

	nginxFile, _ := os.Create(projectPath + "/nginx/nginx.conf")
	nginxFile.Close()

	controllerFile, _ := os.Create(projectPath + "/controllers/handler.go")
	controllerFile.Close()
	controllerFile, _ = os.Create(projectPath + "/controllers/userController.go")
	controllerFile.Close()

	modelFile, _ := os.Create(projectPath + "/models/userModel.go")
	modelFile.Close()

	middlewareFile, _ := os.Create(projectPath + "/middlewares/verify.go")
	middlewareFile.Close()

	helperFile, _ := os.Create(projectPath + "/helpers/env.go")
	helperFile.Close()

	utilFile, _ := os.Create(projectPath + "/utils/token.go")
	utilFile.WriteString(filedata.GetTokenData())
	utilFile.Close()
	utilFile, _ = os.Create(projectPath + "/utils/validate.go")
	utilFile.WriteString(filedata.GetValidateData())
	utilFile.Close()

	routeFile, _ := os.Create(projectPath + "/routes/userRoutes.go")
	routeFile.Close()

	gitignoreFile, _ := os.Create(projectPath + "/.gitignore")
	gitignoreFile.Close()

	envFile, _ := os.Create(projectPath + "/.env")
	envFile.Close()

	dockerfile, _ := os.Create(projectPath + "/Dockerfile")
	dockerfile.WriteString(filedata.GetDockerFileData())
	dockerfile.Close()

	dockerComposeFile, _ := os.Create(projectPath + "/docker-compose.yml")
	dockerComposeFile.Close()

	databaseFile, _ := os.Create(projectPath + "/database/db.go")
	databaseFile.WriteString(filedata.GetDBFileData(projectName))
	databaseFile.Close()

	mainFile, _ := os.Create(projectPath + "/main.go")
	mainFile.WriteString(filedata.GetMainFileData())
	mainFile.Close()

	fmt.Println("Project structure created successfully!")
}
