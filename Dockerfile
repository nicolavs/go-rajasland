FROM golang:1.17-alpine as builder

WORKDIR /project/go-docker/

# COPY go.mod, go.sum and download the dependencies
COPY go.* ./
RUN go mod download

# COPY All things inside the project and build
COPY . .
RUN go build -o /project/go-docker/build/myapp .

FROM scratch
COPY --from=builder /project/go-docker/build/myapp /project/go-docker/build/myapp
COPY --from=builder /platform/migrations /db/migrations

EXPOSE 8080
ENTRYPOINT [ "/project/go-docker/build/myapp" ]