# Challenge Martian Robots
This project is the solution for the challenge https://github.com/Drafteame/code-challenges/tree/main/backend

## Description
This proyect using the hexagonal arquitecture, uses a txt with a template name as a source in the folder /files, to carry out use cases with it, modify the content of this

## Run
go run ./cmd/main.go

## Test
go test ./... -cover

## Docker
to build image
```
docker build -t my-app .
```

to run container
```
docker run --name my-app my-app
```