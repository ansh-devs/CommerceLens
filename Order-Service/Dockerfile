FROM golang:alpine as BuildStage

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o main cmd/main.go

# -----Second Stage---------

FROM alpine:latest

COPY --from=BuildStage /app/main .

EXPOSE 8080

RUN ls
CMD ["./main"]
