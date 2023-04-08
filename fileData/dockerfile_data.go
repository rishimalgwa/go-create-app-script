package filedata

func GetDockerFileData() string {
	return `FROM golang:alpine

RUN apk update && apk upgrade && \
	apk add --no-cache git

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE 8000

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main`

}
