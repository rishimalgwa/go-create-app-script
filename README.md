# Go-App-Gen
A Golang script that creates a basic MVC project structure with common directories and files.
## Instructions to run 🛠️
```
$ git clone https://github.com/rishimalgwa/go-create-app-script.git
$ cd go-create-app-script
$ go build main.go
$ ./create-project
```

## Usage 💻️

1. Clone this repository to your local machine.
2. Run the `main.go` file using the `go run` command.
3. Enter the project name when prompted.
4. Enter the project path (optional) when prompted. Leave empty for the current directory.
5. Wait for the script to complete. The project structure will be created in the specified directory.

## Application Directory Structure 🏛️
```
📦my-app
 ┣ 📂controllers
 ┃ ┣ 📜handler.go
 ┃ ┗ 📜userController.go
 ┣ 📂database
 ┃ ┗ 📜db.go
 ┣ 📂helpers
 ┃ ┗ 📜env.go
 ┣ 📂middlewares
 ┃ ┗ 📜verify.go
 ┣ 📂models
 ┃ ┗ 📜userModel.go
 ┣ 📂nginx
 ┃ ┗ 📜nginx.conf
 ┣ 📂routes
 ┃ ┗ 📜userRoutes.go
 ┣ 📂utils
 ┃ ┣ 📜token.go
 ┃ ┗ 📜validate.go
 ┣ 📜.env
 ┣ 📜.gitignore
 ┣ 🐋Dockerfile
 ┣ 🐋docker-compose.yml
 ┗ 📜main.go

```


This is the directory structure for a Go web application. The structure follows the MVC architecture, with the **controllers/** directory containing the application's controllers, the **models/** directory containing the models, and the **middlewares/** directory containing middleware functions.

  

The **helpers/** directory contains helper functions that can be used across the application, and the **utils/** directory contains utility functions. The **nginx/** directory contains the nginx configuration file for running the application behind an nginx proxy.

  

The **routes/** directory contains the route definitions for the application, and the **database/** directory contains the database connection configuration.

  

The **.gitignore** file specifies the files and directories that should be ignored by version control, and the **.env** file contains environment variables for the application.

  

The **Dockerfile** and **docker-compose.yml** files are used for containerization of the application, and the **main.go** file contains the main function that starts the application.