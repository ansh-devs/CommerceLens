FROM golang:alpine as BuildStage

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main cmd/main.go

# -----Second Stage---------

FROM alpine:latest

COPY --from=BuildStage /app/main .

EXPOSE 8082

RUN ls
CMD ["./main"]
