# STAGE 1 : build
FROM golang:1.19 as build
# set working directory
WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

COPY ./ ./

RUN GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /app/build/linux/interview-service

# STAGE 2 : run dependencies
FROM gcr.io/distroless/base-debian11

COPY --from=build /app/build/linux/interview-service /interview-service